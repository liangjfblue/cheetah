module github.com/liangjfblue/gdiscovery

go 1.13

require (
	github.com/coreos/bbolt v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/etcd v3.3.19+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/btree v1.0.0 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.14.3 // indirect
	github.com/jonboulle/clockwork v0.1.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.5.1 // indirect
	github.com/soheilhy/cmux v0.1.4 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200122045848-3419fae592fc // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	go.etcd.io/bbolt v1.3.4 // indirect
	go.etcd.io/etcd v3.3.19+incompatible
	go.uber.org/zap v1.14.1 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	google.golang.org/grpc v1.26.0 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

//because bbold update module name at this commit
//https://github.com/etcd-io/bbolt/commit/e65d4d5d27e82efe8e16175a11ec2f6e8067c990#diff-37aff102a57d3d7b797f152915a6dc16
replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.3
