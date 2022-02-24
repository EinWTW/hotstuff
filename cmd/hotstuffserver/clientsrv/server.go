package clientsrv

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	hotstuffgorums "github.com/EinWTW/hotstuff/backend/gorums"
	pbv "github.com/EinWTW/hotstuff/cmd/hotstuffserver/proto/veritashs"
	"github.com/EinWTW/hotstuff/consensus/chainedhotstuff"
	"github.com/EinWTW/hotstuff/internal/cli"
	"github.com/EinWTW/hotstuff/internal/logging"
	"github.com/EinWTW/hotstuff/storage/redisdb"
	"github.com/EinWTW/hotstuff/synchronizer"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/relab/gorums"
	"github.com/relab/hotstuff"
	"github.com/relab/hotstuff/client"
	"github.com/relab/hotstuff/config"
	"github.com/relab/hotstuff/crypto"
	"github.com/relab/hotstuff/leaderrotation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/proto"
)

type options struct {
	RootCAs         []string `mapstructure:"root-cas"`
	Privkey         string   `mapstructure:"privkey"`
	Cert            string
	SelfID          hotstuff.ID `mapstructure:"self-id"`
	PmType          string      `mapstructure:"pacemaker"`
	LeaderID        hotstuff.ID `mapstructure:"leader-id"`
	ViewTimeout     int         `mapstructure:"view-timeout"`
	BatchSize       int         `mapstructure:"batch-size"`
	PrintThroughput bool        `mapstructure:"print-throughput"`
	PrintCommands   bool        `mapstructure:"print-commands"`
	ClientAddr      string      `mapstructure:"client-listen"`
	PeerAddr        string      `mapstructure:"peer-listen"`
	RedisAddr       string      `mapstructure:"redis-address"`
	TLS             bool
	Interval        int
	Output          string
	Replicas        []struct {
		ID         hotstuff.ID
		PeerAddr   string `mapstructure:"peer-address"`
		ClientAddr string `mapstructure:"client-address"`
		RedisAddr  string `mapstructure:"redis-address"`
		Pubkey     string `mapstructure:"pubkey"`
		Cert       string `mapstructure:"cert"`
	}
}

func InitHotstuffServer(configFile string) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	var conf options
	err := cli.ReadConfig(&conf, configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read config: %v\n", err)
		os.Exit(1)
	}

	// TODO: replace with go 1.16 signal.NotifyContext
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-signals
		fmt.Fprintf(os.Stderr, "Exiting...")
		cancel()
	}()

	start(ctx, &conf)
}

func start(ctx context.Context, conf *options) {
	privkey, err := crypto.ReadPrivateKeyFile(conf.Privkey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read private key file: %v\n", err)
		os.Exit(1)
	}

	var creds credentials.TransportCredentials
	var tlsCert tls.Certificate
	if conf.TLS {
		creds, tlsCert = loadCreds(conf)
	}

	var clientAddress string
	replicaConfig := config.NewConfig(conf.SelfID, privkey, creds)
	for _, r := range conf.Replicas {
		key, err := crypto.ReadPublicKeyFile(r.Pubkey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read public key file '%s': %v\n", r.Pubkey, err)
			os.Exit(1)
		}

		info := &config.ReplicaInfo{
			ID:      r.ID,
			Address: r.PeerAddr,
			PubKey:  key,
		}

		if r.ID == conf.SelfID {
			// override own addresses if set
			if conf.ClientAddr != "" {
				clientAddress = conf.ClientAddr
			} else {
				clientAddress = r.ClientAddr
			}

			if conf.PeerAddr != "" {
				info.Address = conf.PeerAddr
			}

		}

		replicaConfig.Replicas[r.ID] = info
	}

	logging.NameLogger(fmt.Sprintf("hs%d", conf.SelfID))

	srv := newClientServer(conf, replicaConfig, &tlsCert)
	err = srv.Start(clientAddress)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start HotStuff: %v\n", err)
		os.Exit(1)
	}

	<-ctx.Done()
	srv.Stop()
}

func loadCreds(conf *options) (credentials.TransportCredentials, tls.Certificate) {
	if conf.Cert == "" {
		for _, replica := range conf.Replicas {
			if replica.ID == conf.SelfID {
				conf.Cert = replica.Cert
			}
		}
	}

	tlsCert, err := tls.LoadX509KeyPair(conf.Cert, conf.Privkey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse certificate: %v\n", err)
		os.Exit(1)
	}

	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get system cert pool: %v\n", err)
		os.Exit(1)
	}

	for _, ca := range conf.RootCAs {
		cert, err := ioutil.ReadFile(ca)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read CA file: %v\n", err)
			os.Exit(1)
		}
		if !rootCAs.AppendCertsFromPEM(cert) {
			fmt.Fprintf(os.Stderr, "Failed to add CA to cert pool.\n")
			os.Exit(1)
		}
	}
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		RootCAs:      rootCAs,
		ClientCAs:    rootCAs,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	})

	return creds, tlsCert
}

// cmdID is a unique identifier for a command
type cmdID struct {
	clientID    uint32
	sequenceNum uint64
}

type clientSrv struct {
	ctx          context.Context
	cancel       context.CancelFunc
	conf         *options
	gorumsSrv    *gorums.Server
	hsSrv        *hotstuffgorums.Server
	cfg          *hotstuffgorums.Config
	hs           hotstuff.Consensus
	pm           hotstuff.ViewSynchronizer
	cmdCache     *cmdCache
	rediskv      *redisdb.RedisKV
	mut          sync.Mutex
	finishedCmds map[cmdID]chan struct{}

	lastExecTime int64
}

func newClientServer(conf *options, replicaConfig *config.ReplicaConfig, tlsCert *tls.Certificate) *clientSrv {
	ctx, cancel := context.WithCancel(context.Background())

	serverOpts := []gorums.ServerOption{}
	grpcServerOpts := []grpc.ServerOption{}

	if conf.TLS {
		grpcServerOpts = append(grpcServerOpts, grpc.Creds(credentials.NewServerTLSFromCert(tlsCert)))
	}

	serverOpts = append(serverOpts, gorums.WithGRPCServerOptions(grpcServerOpts...))
	rdb, err := redisdb.NewRedisKV(conf.RedisAddr, "", 1)
	if err != nil {
		log.Println("New server redis db fail: " + conf.RedisAddr)
	}

	srv := &clientSrv{
		ctx:          ctx,
		cancel:       cancel,
		conf:         conf,
		gorumsSrv:    gorums.NewServer(serverOpts...),
		cmdCache:     newCmdCache(conf.BatchSize),
		finishedCmds: make(map[cmdID]chan struct{}),
		lastExecTime: time.Now().UnixNano(),
		rediskv:      rdb,
	}

	srv.cfg = hotstuffgorums.NewConfig(*replicaConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to init gorums backend: %s\n", err)
		os.Exit(1)
	}

	srv.hsSrv = hotstuffgorums.NewServer(*replicaConfig)

	var leaderRotation hotstuff.LeaderRotation
	switch conf.PmType {
	case "fixed":
		leaderRotation = leaderrotation.NewFixed(conf.LeaderID)
	case "round-robin":
		leaderRotation = leaderrotation.NewRoundRobin(srv.cfg)
	default:
		fmt.Fprintf(os.Stderr, "Invalid pacemaker type: '%s'\n", conf.PmType)
		os.Exit(1)
	}
	srv.pm = synchronizer.New(leaderRotation, time.Duration(conf.ViewTimeout)*time.Millisecond)
	srv.hs = chainedhotstuff.Builder{
		Config:       srv.cfg,
		Acceptor:     srv.cmdCache,
		Executor:     srv,
		Synchronizer: srv.pm,
		CommandQueue: srv.cmdCache,
	}.Build()
	// Use a custom server instead of the gorums one
	client.RegisterClientServer(srv.gorumsSrv, srv)
	return srv
}

func (srv *clientSrv) Start(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	err = srv.hsSrv.Start(srv.hs)
	if err != nil {
		return err
	}

	err = srv.cfg.Connect(10 * time.Second)
	if err != nil {
		return err
	}

	// sleep so that all replicas can be ready before we start
	time.Sleep(time.Duration(srv.conf.ViewTimeout) * time.Millisecond)

	srv.pm.Start()

	go func() {
		err := srv.gorumsSrv.Serve(lis)
		if err != nil {
			log.Println(err)
		}
	}()

	return nil
}

func (srv *clientSrv) Stop() {
	srv.pm.Stop()
	srv.cfg.Close()
	srv.rediskv.Close()
	srv.hsSrv.Stop()
	srv.gorumsSrv.Stop()
	srv.cancel()
}

func (srv *clientSrv) ExecCommand(_ context.Context, cmd *client.Command, out func(*empty.Empty, error)) {
	finished := make(chan struct{})
	id := cmdID{cmd.ClientID, cmd.SequenceNumber}
	srv.mut.Lock()
	srv.finishedCmds[id] = finished
	srv.mut.Unlock()

	srv.cmdCache.addCommand(cmd)

	go func(id cmdID, finished chan struct{}) {
		<-finished

		srv.mut.Lock()
		delete(srv.finishedCmds, id)
		srv.mut.Unlock()

		// send response
		out(&empty.Empty{}, nil)
	}(id, finished)
}

func (srv *clientSrv) Exec(cmd hotstuff.Command) {
	batch := new(client.Batch)
	err := proto.UnmarshalOptions{AllowPartial: true}.Unmarshal([]byte(cmd), batch)
	if err != nil {
		return
	}

	if len(batch.GetCommands()) > 0 && srv.conf.PrintThroughput {
		now := time.Now().UnixNano()
		prev := atomic.SwapInt64(&srv.lastExecTime, now)
		fmt.Printf("%d, %d\n", now-prev, len(batch.GetCommands()))
	}
	req := new(pbv.SetRequest)
	for _, cmd := range batch.GetCommands() {
		if err != nil {
			log.Printf("Failed to unmarshal command: %v\n", err)
		}
		if srv.conf.PrintCommands {
			log.Printf("Debug20220222-%s", cmd.Data)
		}

		srv.mut.Lock()
		if c, ok := srv.finishedCmds[cmdID{cmd.ClientID, cmd.SequenceNumber}]; ok {
			c <- struct{}{}
		}
		srv.mut.Unlock()
		if len(cmd.Data) > 0 {
			err = proto.Unmarshal(cmd.Data, req)
			if err != nil {
				log.Printf("Failed to unmarshal command: %v\n", err)
				return
			}
			err = srv.Set([]byte(req.Key), []byte(req.Value))
			if err != nil {
				log.Printf("Error set redis data: %v\n", err)
			}
		} else {
			log.Printf("######Debug20220219 clientSrv SetCommands empty: " + strconv.Itoa(int(cmd.SequenceNumber)))
		}
	}
}

func (srv *clientSrv) Get(key []byte) ([]byte, error) {
	value, err := srv.rediskv.Get(key)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (srv *clientSrv) Set(key, value []byte) error {
	srv.mut.Lock()
	err := srv.rediskv.Set(key, value)
	srv.mut.Unlock()
	// val, err := srv.rediskv.Get(key)
	// if err != nil {
	// 	log.Println("Debug20220222-gorumsConfig"+string(key), err)
	// }

	return err
}
