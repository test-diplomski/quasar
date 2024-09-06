module github.com/jtomic1/config-schema-service

go 1.22.3

require (
	github.com/c12s/meridian v1.0.0
	github.com/c12s/oort v1.0.0
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/xeipuuv/gojsonschema v1.2.0
	go.etcd.io/etcd/client/v3 v3.5.11
	golang.org/x/mod v0.14.0
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.1
	sigs.k8s.io/yaml v1.4.0
)

require (
	github.com/c12s/magnetar v1.0.0 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/nats-io/nats.go v1.31.0 // indirect
	github.com/nats-io/nkeys v0.4.5 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	go.etcd.io/etcd/api/v3 v3.5.11 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.11 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.17.0 // indirect
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
)

replace github.com/c12s/oort => ../oort

replace github.com/c12s/meridian => ../meridian

replace github.com/c12s/magnetar => ../magnetar
