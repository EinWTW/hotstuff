module github.com/wtwinlab/hotstuff

go 1.13

require (
	github.com/felixge/fgprof v0.9.1
	github.com/go-redis/redis/v8 v8.11.4
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.5.2
	github.com/mattn/go-isatty v0.0.12
	github.com/relab/gorums v0.7.0
	github.com/relab/hotstuff v0.2.2
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
)

replace github.com/relab/hotstuff v0.2.2 => github.com/wtwinlab/hotstuff v0.2.3-0.20220429064733-fff49f0b79e0
