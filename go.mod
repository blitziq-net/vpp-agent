module go.ligato.io/vpp-agent/v3

go 1.23.8

replace go.ligato.io/cn-infra/v2 => github.com/joeybole/cn-infra/v2 v2.6.0-alpha

replace github.com/armon/go-metrics => github.com/hashicorp/go-metrics v0.5.3

require (
	github.com/alecthomas/jsonschema v0.0.0-20220216202328-9eeeec9d044b
	github.com/common-nighthawk/go-figure v0.0.0-20210622060536-734e95fb86be
	github.com/coreos/go-iptables v0.8.0
	github.com/docker/cli v27.2.1+incompatible
	github.com/docker/docker v24.0.5+incompatible
	github.com/docker/go-connections v0.5.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-errors/errors v1.5.1
	github.com/goccy/go-graphviz v0.0.6
	github.com/goccy/go-yaml v1.17.1
	github.com/google/go-cmp v0.7.0
	github.com/gorilla/mux v1.8.1
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/iancoleman/orderedmap v0.3.0
	github.com/jhump/protoreflect v1.17.0
	github.com/lunixbochs/struc v0.0.0-20241101090106-8d528fa2c543
	github.com/mitchellh/go-ps v1.0.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/moby/term v0.5.2
	github.com/namsral/flag v1.7.4-pre
	github.com/olekukonko/tablewriter v0.0.5
	github.com/onsi/gomega v1.37.0
	github.com/pkg/errors v0.9.1
	github.com/pkg/profile v1.7.0
	github.com/prometheus/client_golang v1.21.1
	github.com/segmentio/textio v1.2.0
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/cobra v1.9.1
	github.com/spf13/pflag v1.0.6
	github.com/spf13/viper v1.20.1
	github.com/unrolled/render v1.7.0
	github.com/vishvananda/netlink v1.3.0
	github.com/vishvananda/netns v0.0.5
	github.com/xeipuuv/gojsonschema v1.2.0
	go.etcd.io/etcd/client/v3 v3.5.21
	go.fd.io/govpp v0.11.0
	go.ligato.io/cn-infra/v2 v2.5.0-alpha.0.20230824082901-356dce1f1754
	golang.org/x/sys v0.31.0
	google.golang.org/grpc v1.71.1
	google.golang.org/protobuf v1.36.6
)

require (
	cloud.google.com/go/compute/metadata v0.6.0 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20250102033503-faa5f7b0171c // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/Shopify/sarama v1.38.1 // indirect
	github.com/armon/go-metrics v0.5.4 // indirect
	github.com/bennyscetbun/jsongo v1.1.2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/bshuster-repo/logrus-logstash-hook v1.1.0 // indirect
	github.com/bsm/sarama-cluster v2.1.15+incompatible // indirect
	github.com/bufbuild/protocompile v0.14.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/distribution/reference v0.6.0 // indirect
	github.com/docker/distribution v2.8.3+incompatible // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/eapache/go-resiliency v1.7.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/evalphobia/logrus_fluent v0.5.4 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/felixge/fgprof v0.9.5 // indirect
	github.com/fluent/fluent-logger-golang v1.9.0 // indirect
	github.com/fogleman/gg v1.3.0 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/ftrvxmtrx/fd v0.0.0-20150925145434-c6d800382fff // indirect
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v1.0.0 // indirect
	github.com/google/pprof v0.0.0-20250403155104-27863c87afa6 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/hashicorp/consul/api v1.32.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-metrics v0.5.4 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/serf v0.10.2 // indirect
	github.com/howeyc/crc16 v0.0.0-20171223171357-2b2a61e366a6 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.1 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/philhofer/fwd v1.1.3-0.20240916144458-20a13a1f6b7c // indirect
	github.com/pierrec/lz4/v4 v4.1.22 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.63.0 // indirect
	github.com/prometheus/procfs v0.16.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20250401214520-65e299d6c5c9 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/sagikazarmark/locafero v0.9.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.14.0 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/tinylib/msgp v1.2.5 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	go.etcd.io/etcd/api/v3 v3.5.21 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.21 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.35.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/exp v0.0.0-20250305212735-054e65f0b394 // indirect
	golang.org/x/image v0.25.0 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	golang.org/x/time v0.11.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250324211829-b45e905df463 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gotest.tools/v3 v3.5.0 // indirect
)
