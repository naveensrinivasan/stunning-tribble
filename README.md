# stunning-tribble

Scan your `go` dependecies for [OSV](https://osv.dev).

### What is OSV?
>Database for open source vulnerabilities
OSV is a vulnerability database and triage infrastructure for open source projects aimed at helping both open source maintainers and consumers of open source.


## How do I use this tool?
1. `go install github.com/naveensrinivasan/stunning-tribble@latest`
1. Navigate to your `go.mod` folder
1. `go list -m -f '{{if not (or  .Main)}}{{.Path}}@{{.Version}}_{{.Replace}}{{end}}' all  | stunning-tribble`

- If there aren't issues, it would `exit` without an error. 
- If it finds any Vulnerability, it would print the vulnerability and `exit` with 1.

### Why build this?

Every time a PR comes in for updates to `go.mod`/`go.su` will help with any known OSV issues.

### I can't fix all of them. Can I ignore existing ones?

Yes, you can ignore existing ones by passing the ID via the command line as comma-separated. Here the tool will ignore `GO-2020-0018,GO-2020-0016`

Example 
`go list -m -f '{{if not (or  .Main)}}{{.Path}}@{{.Version}}_{{.Replace}}{{end}}' all   | stunning-tribble GO-2020-0018,GO-2020-0016`


#### Does it handle `replace` directive?

Yes, `go list -m -f '{{if not (or  .Main)}}{{.Path}}@{{.Version}}_{{.Replace}}{{end}}' all`


#### What is the input for this?
`go list -m -f '{{if not (or  .Main)}}{{.Path}}@{{.Version}}_{{.Replace}}{{end}}' all `
<details>
<summary>Here is an example of input that can be passed as stdin</summary>

```
cloud.google.com/go@v0.65.0_<nil>
cloud.google.com/go/bigquery@v1.8.0_<nil>
cloud.google.com/go/datastore@v1.1.0_<nil>
cloud.google.com/go/firestore@v1.1.0_<nil>
cloud.google.com/go/pubsub@v1.3.1_<nil>
cloud.google.com/go/storage@v1.10.0_<nil>
dmitri.shuralyov.com/gpu/mtl@v0.0.0-20190408044501-666a987793e9_<nil>
git.schwanenlied.me/yawning/bsaes.git@v0.0.0-20180720073208-c0276d75487e_github.com/Yawning/bsaes v0.0.0-20180720073208-c0276d75487e
github.com/BurntSushi/toml@v0.3.1_<nil>
github.com/BurntSushi/xgb@v0.0.0-20160522181843-27f122750802_<nil>
github.com/NebulousLabs/fastrand@v0.0.0-20181203155948-6fb6489aac4e_<nil>
github.com/NebulousLabs/go-upnp@v0.0.0-20180202185039-29b680b06c82_<nil>
github.com/OneOfOne/xxhash@v1.2.2_<nil>
github.com/Yawning/aez@v0.0.0-20180114000226-4dad034d9db2_<nil>
github.com/aead/chacha20@v0.0.0-20180709150244-8b13a72661da_<nil>
github.com/aead/siphash@v1.0.1_<nil>
github.com/alecthomas/template@v0.0.0-20190718012654-fb15b899a751_<nil>
github.com/alecthomas/units@v0.0.0-20190924025748-f65c72e2690d_<nil>
github.com/antihax/optional@v1.0.0_<nil>
github.com/armon/circbuf@v0.0.0-20150827004946-bbbad097214e_<nil>
github.com/armon/go-metrics@v0.0.0-20180917152333-f0300d1749da_<nil>
github.com/armon/go-radix@v0.0.0-20180808171621-7fddfc383310_<nil>
github.com/asaskevich/govalidator@v0.0.0-20190424111038-f61b66f89f4a_<nil>
github.com/benbjohnson/clock@v1.0.3_<nil>
github.com/beorn7/perks@v1.0.1_<nil>
github.com/bgentry/speakeasy@v0.1.0_<nil>
github.com/bketelsen/crypt@v0.0.3-0.20200106085610-5cbc8cc4026c_<nil>
github.com/btcsuite/btcd@v0.22.0-beta.0.20210803133449-f5a1fb9965e4_<nil>
github.com/btcsuite/btclog@v0.0.0-20170628155309-84c8d2346e9f_<nil>
github.com/btcsuite/btcutil@v1.0.3-0.20210527170813-e2ba6805a890_<nil>
github.com/btcsuite/btcutil/psbt@v1.0.3-0.20210527170813-e2ba6805a890_<nil>
github.com/btcsuite/btcwallet@v0.12.1-0.20210826004415-4ef582f76b02_<nil>
github.com/btcsuite/btcwallet/wallet/txauthor@v1.0.2-0.20210803004036-eebed51155ec_<nil>
github.com/btcsuite/btcwallet/wallet/txrules@v1.0.0_<nil>
github.com/btcsuite/btcwallet/wallet/txsizes@v1.0.1-0.20210519225359-6ab9b615576f_<nil>
github.com/btcsuite/btcwallet/walletdb@v1.3.6-0.20210803004036-eebed51155ec_<nil>
github.com/btcsuite/btcwallet/wtxmgr@v1.3.1-0.20210822222949-9b5a201c344c_<nil>
github.com/btcsuite/go-socks@v0.0.0-20170105172521-4720035b7bfd_<nil>
github.com/btcsuite/golangcrypto@v0.0.0-20150304025918-53f62d9b43e8_<nil>
github.com/btcsuite/goleveldb@v1.0.0_<nil>
github.com/btcsuite/snappy-go@v1.0.0_<nil>
github.com/btcsuite/websocket@v0.0.0-20150119174127-31079b680792_<nil>
github.com/btcsuite/winsvc@v1.0.0_<nil>
github.com/census-instrumentation/opencensus-proto@v0.2.1_<nil>
github.com/certifi/gocertifi@v0.0.0-20200922220541-2c3bb06c6054_<nil>
github.com/cespare/xxhash@v1.1.0_<nil>
github.com/cespare/xxhash/v2@v2.1.1_<nil>
github.com/chzyer/logex@v1.1.10_<nil>
github.com/chzyer/readline@v0.0.0-20180603132655-2972be24d48e_<nil>
github.com/chzyer/test@v0.0.0-20180213035817-a1ea475d72b1_<nil>
github.com/client9/misspell@v0.3.4_<nil>
github.com/cncf/udpa/go@v0.0.0-20201120205902-5459f2c99403_<nil>
github.com/cockroachdb/datadriven@v0.0.0-20200714090401-bf6692d28da5_<nil>
github.com/cockroachdb/errors@v1.2.4_<nil>
github.com/cockroachdb/logtags@v0.0.0-20190617123548-eb05cc24525f_<nil>
github.com/coreos/bbolt@v1.3.2_<nil>
github.com/coreos/etcd@v3.3.13+incompatible_<nil>
github.com/coreos/go-semver@v0.3.0_<nil>
github.com/coreos/go-systemd@v0.0.0-20190321100706-95778dfbb74e_<nil>
github.com/coreos/go-systemd/v22@v22.3.2_<nil>
github.com/coreos/pkg@v0.0.0-20180928190104-399ea9e2e55f_<nil>
github.com/cpuguy83/go-md2man/v2@v2.0.0_<nil>
github.com/creack/pty@v1.1.11_<nil>
github.com/davecgh/go-spew@v1.1.1_<nil>
github.com/decred/dcrd/lru@v1.0.0_<nil>
github.com/dgrijalva/jwt-go@v3.2.0+incompatible_github.com/golang-jwt/jwt v3.2.1+incompatible
github.com/dgryski/go-sip13@v0.0.0-20181026042036-e10d5fee7954_<nil>
github.com/dustin/go-humanize@v1.0.0_<nil>
github.com/envoyproxy/go-control-plane@v0.9.9-0.20210217033140-668b12f5399d_<nil>
github.com/envoyproxy/protoc-gen-validate@v0.1.0_<nil>
github.com/fatih/color@v1.7.0_<nil>
github.com/form3tech-oss/jwt-go@v3.2.3+incompatible_<nil>
github.com/frankban/quicktest@v1.2.2_<nil>
github.com/fsnotify/fsnotify@v1.4.9_<nil>
github.com/getsentry/raven-go@v0.2.0_<nil>
github.com/ghodss/yaml@v1.0.0_<nil>
github.com/go-errors/errors@v1.0.1_<nil>
github.com/go-gl/glfw@v0.0.0-20190409004039-e6da0acd62b1_<nil>
github.com/go-gl/glfw/v3.3/glfw@v0.0.0-20200222043503-6f7a984d4dc4_<nil>
github.com/go-kit/kit@v0.9.0_<nil>
github.com/go-kit/log@v0.1.0_<nil>
github.com/go-logfmt/logfmt@v0.5.0_<nil>
github.com/go-openapi/errors@v0.19.2_<nil>
github.com/go-openapi/strfmt@v0.19.5_<nil>
github.com/go-stack/stack@v1.8.0_<nil>
github.com/godbus/dbus/v5@v5.0.4_<nil>
github.com/gogo/protobuf@v1.3.2_<nil>
github.com/golang/glog@v0.0.0-20210429001901-424d2337a529_<nil>
github.com/golang/groupcache@v0.0.0-20210331224755-41bb18bfe9da_<nil>
github.com/golang/lint@v0.0.0-20180702182130-06c8688daad7_<nil>
github.com/golang/mock@v1.4.4_<nil>
github.com/golang/protobuf@v1.5.2_<nil>
github.com/google/btree@v1.0.1_<nil>
github.com/google/go-cmp@v0.5.6_<nil>
github.com/google/gofuzz@v1.0.0_<nil>
github.com/google/martian@v2.1.0+incompatible_<nil>
github.com/google/martian/v3@v3.0.0_<nil>
github.com/google/pprof@v0.0.0-20200708004538-1a94d8640e99_<nil>
github.com/google/renameio@v0.1.0_<nil>
github.com/google/uuid@v1.1.2_<nil>
github.com/googleapis/gax-go/v2@v2.0.5_<nil>
github.com/gopherjs/gopherjs@v0.0.0-20181017120253-0766667cb4d1_<nil>
github.com/gorilla/websocket@v1.4.2_<nil>
github.com/grpc-ecosystem/go-grpc-middleware@v1.3.0_<nil>
github.com/grpc-ecosystem/go-grpc-prometheus@v1.2.0_<nil>
github.com/grpc-ecosystem/grpc-gateway@v1.16.0_<nil>
github.com/grpc-ecosystem/grpc-gateway/v2@v2.5.0_<nil>
github.com/hashicorp/consul/api@v1.1.0_<nil>
github.com/hashicorp/consul/sdk@v0.1.1_<nil>
github.com/hashicorp/errwrap@v1.0.0_<nil>
github.com/hashicorp/go-cleanhttp@v0.5.1_<nil>
github.com/hashicorp/go-immutable-radix@v1.0.0_<nil>
github.com/hashicorp/go-msgpack@v0.5.3_<nil>
github.com/hashicorp/go-multierror@v1.0.0_<nil>
github.com/hashicorp/go-rootcerts@v1.0.0_<nil>
github.com/hashicorp/go-sockaddr@v1.0.0_<nil>
github.com/hashicorp/go-syslog@v1.0.0_<nil>
github.com/hashicorp/go-uuid@v1.0.1_<nil>
github.com/hashicorp/go.net@v0.0.1_<nil>
github.com/hashicorp/golang-lru@v0.5.1_<nil>
github.com/hashicorp/hcl@v1.0.0_<nil>
github.com/hashicorp/logutils@v1.0.0_<nil>
github.com/hashicorp/mdns@v1.0.0_<nil>
github.com/hashicorp/memberlist@v0.1.3_<nil>
github.com/hashicorp/serf@v0.8.2_<nil>
github.com/hpcloud/tail@v1.0.0_<nil>
github.com/ianlancetaylor/demangle@v0.0.0-20181102032728-5e5cf60278f6_<nil>
github.com/inconshreveable/mousetrap@v1.0.0_<nil>
github.com/jackpal/gateway@v1.0.5_<nil>
github.com/jackpal/go-nat-pmp@v0.0.0-20170405195558-28a68d0c24ad_<nil>
github.com/jedib0t/go-pretty@v4.3.0+incompatible_<nil>
github.com/jessevdk/go-flags@v1.4.0_<nil>
github.com/jonboulle/clockwork@v0.2.2_<nil>
github.com/jpillora/backoff@v1.0.0_<nil>
github.com/jrick/logrotate@v1.0.0_<nil>
github.com/json-iterator/go@v1.1.11_<nil>
github.com/jstemmer/go-junit-report@v0.9.1_<nil>
github.com/jtolds/gls@v4.20.0+incompatible_<nil>
github.com/juju/clock@v0.0.0-20190205081909-9c5c9712527c_<nil>
github.com/juju/errors@v0.0.0-20190806202954-0232dcc7464d_<nil>
github.com/juju/loggo@v0.0.0-20190526231331-6e530bcce5d8_<nil>
github.com/juju/retry@v0.0.0-20180821225755-9058e192b216_<nil>
github.com/juju/testing@v0.0.0-20190723135506-ce30eb24acd2_<nil>
github.com/juju/utils@v0.0.0-20180820210520-bf9cc5bdd62d_<nil>
github.com/juju/version@v0.0.0-20180108022336-b64dbd566305_<nil>
github.com/julienschmidt/httprouter@v1.3.0_<nil>
github.com/kisielk/errcheck@v1.5.0_<nil>
github.com/kisielk/gotool@v1.0.0_<nil>
github.com/kkdai/bstream@v0.0.0-20181106074824-b3251f7901ec_<nil>
github.com/konsorten/go-windows-terminal-sequences@v1.0.3_<nil>
github.com/kr/logfmt@v0.0.0-20140226030751-b84e30acd515_<nil>
github.com/kr/pretty@v0.1.0_<nil>
github.com/kr/pty@v1.1.1_<nil>
github.com/kr/text@v0.2.0_<nil>
github.com/lightninglabs/gozmq@v0.0.0-20191113021534-d20a764486bf_<nil>
github.com/lightninglabs/neutrino@v0.12.1_<nil>
github.com/lightninglabs/protobuf-hex-display@v1.4.3-hex-display_<nil>
github.com/lightningnetwork/lightning-onion@v1.0.2-0.20210520211913-522b799e65b1_<nil>
github.com/lightningnetwork/lnd/cert@v1.0.3_./cert
github.com/lightningnetwork/lnd/clock@v1.0.1_./clock
github.com/lightningnetwork/lnd/healthcheck@v1.0.2_./healthcheck
github.com/lightningnetwork/lnd/kvdb@v1.0.2_./kvdb
github.com/lightningnetwork/lnd/queue@v1.0.4_./queue
github.com/lightningnetwork/lnd/ticker@v1.0.0_./ticker
github.com/ltcsuite/ltcd@v0.0.0-20190101042124-f37f8bf35796_<nil>
github.com/ltcsuite/ltcutil@v0.0.0-20181217130922-17f3b04680b6_<nil>
github.com/magiconair/properties@v1.8.1_<nil>
github.com/mattn/go-colorable@v0.0.9_<nil>
github.com/mattn/go-isatty@v0.0.3_<nil>
github.com/mattn/go-runewidth@v0.0.9_<nil>
github.com/matttproud/golang_protobuf_extensions@v1.0.1_<nil>
github.com/miekg/dns@v1.1.43_<nil>
github.com/mitchellh/cli@v1.0.0_<nil>
github.com/mitchellh/go-homedir@v1.1.0_<nil>
github.com/mitchellh/go-testing-interface@v1.0.0_<nil>
github.com/mitchellh/gox@v0.4.0_<nil>
github.com/mitchellh/iochan@v1.0.0_<nil>
github.com/mitchellh/mapstructure@v1.1.2_<nil>
github.com/modern-go/concurrent@v0.0.0-20180306012644-bacd9c7ef1dd_<nil>
github.com/modern-go/reflect2@v1.0.1_<nil>
github.com/mwitkow/go-conntrack@v0.0.0-20190716064945-2f068394615f_<nil>
github.com/oklog/ulid@v1.3.1_<nil>
github.com/onsi/ginkgo@v1.7.0_<nil>
github.com/onsi/gomega@v1.4.3_<nil>
github.com/opentracing/opentracing-go@v1.1.0_<nil>
github.com/pascaldekloe/goe@v0.0.0-20180627143212-57f6aae5913c_<nil>
github.com/pelletier/go-toml@v1.2.0_<nil>
github.com/pkg/errors@v0.9.1_<nil>
github.com/pmezard/go-difflib@v1.0.0_<nil>
github.com/posener/complete@v1.1.1_<nil>
github.com/prometheus/client_golang@v1.11.0_<nil>
github.com/prometheus/client_model@v0.2.0_<nil>
github.com/prometheus/common@v0.26.0_<nil>
github.com/prometheus/procfs@v0.6.0_<nil>
github.com/prometheus/tsdb@v0.7.1_<nil>
github.com/rogpeppe/fastuuid@v1.2.0_<nil>
github.com/rogpeppe/go-internal@v1.3.0_<nil>
github.com/russross/blackfriday/v2@v2.0.1_<nil>
github.com/ryanuber/columnize@v0.0.0-20160712163229-9b3edd62028f_<nil>
github.com/sean-/seed@v0.0.0-20170313163322-e2103e2c3529_<nil>
github.com/shurcooL/sanitized_anchor_name@v1.0.0_<nil>
github.com/sirupsen/logrus@v1.7.0_<nil>
github.com/smartystreets/assertions@v0.0.0-20180927180507-b2de0cb4f26d_<nil>
github.com/smartystreets/goconvey@v1.6.4_<nil>
github.com/soheilhy/cmux@v0.1.5_<nil>
github.com/spaolacci/murmur3@v0.0.0-20180118202830-f09979ecbc72_<nil>
github.com/spf13/afero@v1.1.2_<nil>
github.com/spf13/cast@v1.3.0_<nil>
github.com/spf13/cobra@v1.1.3_<nil>
github.com/spf13/jwalterweatherman@v1.0.0_<nil>
github.com/spf13/pflag@v1.0.5_<nil>
github.com/spf13/viper@v1.7.0_<nil>
github.com/stretchr/objx@v0.2.0_<nil>
github.com/stretchr/testify@v1.7.0_<nil>
github.com/subosito/gotenv@v1.2.0_<nil>
github.com/tidwall/pretty@v1.0.0_<nil>
github.com/tmc/grpc-websocket-proxy@v0.0.0-20201229170055-e5319fda7802_<nil>
github.com/tv42/zbase32@v0.0.0-20160707012821-501572607d02_<nil>
github.com/urfave/cli@v1.20.0_<nil>
github.com/xiang90/probing@v0.0.0-20190116061207-43a291ad63a2_<nil>
github.com/yuin/goldmark@v1.3.5_<nil>
go.etcd.io/bbolt@v1.3.6_<nil>
go.etcd.io/etcd/api/v3@v3.5.0_<nil>
go.etcd.io/etcd/client/pkg/v3@v3.5.0_<nil>
go.etcd.io/etcd/client/v2@v2.305.0_<nil>
go.etcd.io/etcd/client/v3@v3.5.0_<nil>
go.etcd.io/etcd/pkg/v3@v3.5.0_<nil>
go.etcd.io/etcd/raft/v3@v3.5.0_<nil>
go.etcd.io/etcd/server/v3@v3.5.0_<nil>
go.mongodb.org/mongo-driver@v1.0.3_<nil>
go.opencensus.io@v0.22.4_<nil>
go.opentelemetry.io/contrib@v0.20.0_<nil>
go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc@v0.20.0_<nil>
go.opentelemetry.io/otel@v0.20.0_<nil>
go.opentelemetry.io/otel/exporters/otlp@v0.20.0_<nil>
go.opentelemetry.io/otel/metric@v0.20.0_<nil>
go.opentelemetry.io/otel/oteltest@v0.20.0_<nil>
go.opentelemetry.io/otel/sdk@v0.20.0_<nil>
go.opentelemetry.io/otel/sdk/export/metric@v0.20.0_<nil>
go.opentelemetry.io/otel/sdk/metric@v0.20.0_<nil>
go.opentelemetry.io/otel/trace@v0.20.0_<nil>
go.opentelemetry.io/proto/otlp@v0.7.0_<nil>
go.uber.org/atomic@v1.7.0_<nil>
go.uber.org/goleak@v1.1.10_<nil>
go.uber.org/multierr@v1.6.0_<nil>
go.uber.org/zap@v1.17.0_<nil>
golang.org/x/crypto@v0.0.0-20201002170205-7f63de1d35b0_<nil>
golang.org/x/exp@v0.0.0-20200224162631-6cc2880d07d6_<nil>
golang.org/x/image@v0.0.0-20190802002840-cff245a6509b_<nil>
golang.org/x/lint@v0.0.0-20210508222113-6edffad5e616_<nil>
golang.org/x/mobile@v0.0.0-20190719004257-d2bd2a29d028_<nil>
golang.org/x/mod@v0.4.2_<nil>
golang.org/x/net@v0.0.0-20210913180222-943fd674d43e_<nil>
golang.org/x/oauth2@v0.0.0-20210615190721-d04028783cf1_<nil>
golang.org/x/sync@v0.0.0-20210220032951-036812b2e83c_<nil>
golang.org/x/sys@v0.0.0-20210915083310-ed5796bab164_<nil>
golang.org/x/term@v0.0.0-20201126162022-7de9c90e9dd1_<nil>
golang.org/x/text@v0.3.6_<nil>
golang.org/x/time@v0.0.0-20210220033141-f8bda1e9f3ba_<nil>
golang.org/x/tools@v0.1.3_<nil>
golang.org/x/xerrors@v0.0.0-20200804184101-5ec99f83aff1_<nil>
google.golang.org/api@v0.30.0_<nil>
google.golang.org/appengine@v1.6.6_<nil>
google.golang.org/genproto@v0.0.0-20210617175327-b9e0b3197ced_<nil>
google.golang.org/grpc@v1.38.0_<nil>
google.golang.org/protobuf@v1.26.0_<nil>
gopkg.in/alecthomas/kingpin.v2@v2.2.6_<nil>
gopkg.in/check.v1@v1.0.0-20190902080502-41f04d3bba15_<nil>
gopkg.in/errgo.v1@v1.0.1_<nil>
gopkg.in/errgo.v2@v2.1.0_<nil>
gopkg.in/fsnotify.v1@v1.4.7_<nil>
gopkg.in/ini.v1@v1.51.0_<nil>
gopkg.in/macaroon-bakery.v2@v2.0.1_<nil>
gopkg.in/macaroon.v2@v2.0.0_<nil>
gopkg.in/mgo.v2@v2.0.0-20190816093944-a6b53ec6cb22_<nil>
gopkg.in/natefinch/lumberjack.v2@v2.0.0_<nil>
gopkg.in/resty.v1@v1.12.0_<nil>
gopkg.in/tomb.v1@v1.0.0-20141024135613-dd632973f1e7_<nil>
gopkg.in/yaml.v2@v2.4.0_<nil>
gopkg.in/yaml.v3@v3.0.0-20210107192922-496545a6307b_<nil>
honnef.co/go/tools@v0.0.1-2020.1.4_<nil>
rsc.io/binaryregexp@v0.2.0_<nil>
rsc.io/quote/v3@v3.1.0_<nil>
rsc.io/sampler@v1.3.0_<nil>
sigs.k8s.io/yaml@v1.2.0_<nil>
```
</details>  

#### What is the output when it fails?
It dumps the `osv` `json` result.
<details>
<summary>Here is an example of output</summary>

``` json
{
  "osv": [
    {
      "vulns": [
        {
          "id": "GO-2021-0089",
          "package": {
            "name": "github.com/buger/jsonparser",
            "ecosystem": "Go"
          },
          "details": "Parsing malformed JSON which contain opening brackets, but not closing brackes,\nleads to an infinite loop. If operating on untrusted user input this can be\nused as a denial of service vector.\n",
          "affects": {
            "ranges": [
              {
                "type": "SEMVER",
                "fixed": "0.0.0-20200321185410-91ac96899e49"
              }
            ]
          },
          "aliases": [
            "CVE-2020-10675"
          ],
          "modified": "2021-04-14T12:00:00Z",
          "published": "2021-04-14T12:00:00Z",
          "ecosystem_specific": {
            "symbols": [
              "findKeyStart"
            ]
          },
          "database_specific": {
            "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2021-0089.yaml",
            "source": "https://storage.googleapis.com/go-vulndb/github.com/buger/jsonparser.json"
          },
          "references": [
            {
              "type": "FIX",
              "url": "https://github.com/buger/jsonparser/pull/192"
            },
            {
              "type": "FIX",
              "url": "https://github.com/buger/jsonparser/commit/91ac96899e492584984ded0c8f9a08f10b473717"
            },
            {
              "type": "WEB",
              "url": "https://github.com/buger/jsonparser/issues/188"
            }
          ],
          "affected": [
            {
              "package": {
                "name": "github.com/buger/jsonparser",
                "ecosystem": "Go"
              },
              "ranges": [
                {
                  "type": "SEMVER",
                  "events": [
                    {
                      "introduced": "0"
                    },
                    {
                      "fixed": "0.0.0-20200321185410-91ac96899e49"
                    }
                  ]
                }
              ],
              "ecosystem_specific": {
                "symbols": [
                  "findKeyStart"
                ]
              },
              "database_specific": {
                "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2021-0089.yaml",
                "source": "https://storage.googleapis.com/go-vulndb/github.com/buger/jsonparser.json"
              }
            }
          ]
        },
        {
          "id": "GO-2021-0057",
          "package": {
            "name": "github.com/buger/jsonparser",
            "ecosystem": "Go"
          },
          "details": "Due to improper bounds checking, maliciously crafted JSON objects\ncan cause an out-of-bounds panic. If parsing user input, this may\nbe used as a denial of service vector.\n",
          "affects": {
            "ranges": [
              {
                "type": "SEMVER",
                "fixed": "1.1.1"
              }
            ]
          },
          "aliases": [
            "CVE-2020-35381"
          ],
          "modified": "2021-04-14T12:00:00Z",
          "published": "2021-04-14T12:00:00Z",
          "ecosystem_specific": {
            "symbols": [
              "searchKeys"
            ]
          },
          "database_specific": {
            "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2021-0057.yaml",
            "source": "https://storage.googleapis.com/go-vulndb/github.com/buger/jsonparser.json"
          },
          "references": [
            {
              "type": "FIX",
              "url": "https://github.com/buger/jsonparser/pull/221"
            },
            {
              "type": "FIX",
              "url": "https://github.com/buger/jsonparser/commit/df3ea76ece10095374fd1c9a22a4fb85a44efc42"
            },
            {
              "type": "WEB",
              "url": "https://github.com/buger/jsonparser/issues/219"
            }
          ],
          "affected": [
            {
              "package": {
                "name": "github.com/buger/jsonparser",
                "ecosystem": "Go"
              },
              "ranges": [
                {
                  "type": "SEMVER",
                  "events": [
                    {
                      "introduced": "0"
                    },
                    {
                      "fixed": "1.1.1"
                    }
                  ]
                }
              ],
              "ecosystem_specific": {
                "symbols": [
                  "searchKeys"
                ]
              },
              "database_specific": {
                "source": "https://storage.googleapis.com/go-vulndb/github.com/buger/jsonparser.json",
                "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2021-0057.yaml"
              }
            }
          ]
        }
      ]
    },
    {
      "vulns": [
        {
          "id": "GO-2020-0017",
          "package": {
            "name": "github.com/dgrijalva/jwt-go",
            "ecosystem": "Go"
          },
          "details": "If a JWT contains an audience claim with an array of strings, rather\nthan a single string, and `MapClaims.VerifyAudience` is called with\n`req` set to `false`, then audience verification will be bypassed,\nallowing an invalid set of audiences to be provided.\n",
          "affects": {
            "ranges": [
              {
                "type": "SEMVER",
                "introduced": "0.0.0-20150717181359-44718f8a89b0"
              }
            ]
          },
          "aliases": [
            "CVE-2020-26160"
          ],
          "modified": "2021-04-14T12:00:00Z",
          "published": "2021-04-14T12:00:00Z",
          "ecosystem_specific": {
            "symbols": [
              "MapClaims.VerifyAudience"
            ]
          },
          "database_specific": {
            "source": "https://storage.googleapis.com/go-vulndb/github.com/dgrijalva/jwt-go/v4.json",
            "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2020-0017.yaml"
          },
          "references": [
            {
              "type": "FIX",
              "url": "https://github.com/dgrijalva/jwt-go/commit/ec0a89a131e3e8567adcb21254a5cd20a70ea4ab"
            },
            {
              "type": "WEB",
              "url": "https://github.com/dgrijalva/jwt-go/issues/422"
            }
          ],
          "affected": [
            {
              "package": {
                "name": "github.com/dgrijalva/jwt-go",
                "ecosystem": "Go"
              },
              "ranges": [
                {
                  "type": "SEMVER",
                  "events": [
                    {
                      "introduced": "0.0.0-20150717181359-44718f8a89b0"
                    }
                  ]
                }
              ],
              "ecosystem_specific": {
                "symbols": [
                  "MapClaims.VerifyAudience"
                ]
              },
              "database_specific": {
                "source": "https://storage.googleapis.com/go-vulndb/github.com/dgrijalva/jwt-go/v4.json",
                "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2020-0017.yaml"
              }
            },
            {
              "package": {
                "name": "github.com/dgrijalva/jwt-go/v4",
                "ecosystem": "Go"
              },
              "ranges": [
                {
                  "type": "SEMVER",
                  "events": [
                    {
                      "introduced": "0"
                    },
                    {
                      "fixed": "4.0.0-preview1"
                    }
                  ]
                }
              ],
              "ecosystem_specific": {
                "symbols": [
                  "MapClaims.VerifyAudience"
                ]
              },
              "database_specific": {
                "source": "https://storage.googleapis.com/go-vulndb/github.com/dgrijalva/jwt-go/v4.json",
                "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2020-0017.yaml"
              }
            }
          ]
        }
      ]
    },
    {
      "vulns": [
        {
          "id": "GO-2020-0020",
          "package": {
            "name": "github.com/gorilla/handlers",
            "ecosystem": "Go"
          },
          "details": "Usage of the [`CORS`] handler may apply improper CORS headers, allowing\nthe requester to explicitly control the value of the Access-Control-Allow-Origin\nheader, which bypasses the expected behavior of the Same Origin Policy.\n",
          "affects": {
            "ranges": [
              {
                "type": "SEMVER",
                "fixed": "1.3.0"
              }
            ]
          },
          "modified": "2021-04-14T12:00:00Z",
          "published": "2021-04-14T12:00:00Z",
          "ecosystem_specific": {
            "symbols": [
              "cors.ServeHTTP"
            ]
          },
          "database_specific": {
            "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2020-0020.yaml",
            "source": "https://storage.googleapis.com/go-vulndb/github.com/gorilla/handlers.json"
          },
          "references": [
            {
              "type": "FIX",
              "url": "https://github.com/gorilla/handlers/pull/116"
            },
            {
              "type": "FIX",
              "url": "https://github.com/gorilla/handlers/commit/90663712d74cb411cbef281bc1e08c19d1a76145"
            }
          ],
          "affected": [
            {
              "package": {
                "name": "github.com/gorilla/handlers",
                "ecosystem": "Go"
              },
              "ranges": [
                {
                  "type": "SEMVER",
                  "events": [
                    {
                      "introduced": "0"
                    },
                    {
                      "fixed": "1.3.0"
                    }
                  ]
                }
              ],
              "ecosystem_specific": {
                "symbols": [
                  "cors.ServeHTTP"
                ]
              },
              "database_specific": {
                "source": "https://storage.googleapis.com/go-vulndb/github.com/gorilla/handlers.json",
                "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2020-0020.yaml"
              }
            }
          ]
        }
      ]
    },
    {
      "vulns": [
        {
          "id": "GO-2020-0008",
          "package": {
            "name": "github.com/miekg/dns",
            "ecosystem": "Go"
          },
          "details": "DNS message transaction IDs are generated using [`math/rand`] which\nmakes them relatively predictable. This reduces the complexity\nof response spoofing attacks against DNS clients.\n",
          "affects": {
            "ranges": [
              {
                "type": "SEMVER",
                "fixed": "1.1.25-0.20191211073109-8ebf2e419df7"
              }
            ]
          },
          "aliases": [
            "CVE-2019-19794"
          ],
          "modified": "2021-04-14T12:00:00Z",
          "published": "2021-04-14T12:00:00Z",
          "ecosystem_specific": {
            "symbols": [
              "id"
            ]
          },
          "database_specific": {
            "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2020-0008.yaml",
            "source": "https://storage.googleapis.com/go-vulndb/github.com/miekg/dns.json"
          },
          "references": [
            {
              "type": "FIX",
              "url": "https://github.com/miekg/dns/pull/1044"
            },
            {
              "type": "FIX",
              "url": "https://github.com/miekg/dns/commit/8ebf2e419df7857ac8919baa05248789a8ffbf33"
            },
            {
              "type": "WEB",
              "url": "https://github.com/miekg/dns/issues/1037"
            },
            {
              "type": "WEB",
              "url": "https://github.com/miekg/dns/issues/1043"
            }
          ],
          "affected": [
            {
              "package": {
                "name": "github.com/miekg/dns",
                "ecosystem": "Go"
              },
              "ranges": [
                {
                  "type": "SEMVER",
                  "events": [
                    {
                      "introduced": "0"
                    },
                    {
                      "fixed": "1.1.25-0.20191211073109-8ebf2e419df7"
                    }
                  ]
                }
              ],
              "ecosystem_specific": {
                "symbols": [
                  "id"
                ]
              },
              "database_specific": {
                "source": "https://storage.googleapis.com/go-vulndb/github.com/miekg/dns.json",
                "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2020-0008.yaml"
              }
            }
          ]
        }
      ]
    },
    {
      "vulns": [
        {
          "id": "GO-2020-0018",
          "package": {
            "name": "github.com/satori/go.uuid",
            "ecosystem": "Go"
          },
          "details": "UUIDs generated using [`NewV1`] and [`NewV4`] may not read the expected\nnumber of random bytes. These UUIDs may contain a significantly smaller\namount of entropy than expected, possibly leading to collisions.\n",
          "affects": {
            "ranges": [
              {
                "type": "SEMVER",
                "fixed": "1.2.1-0.20181016170032-d91630c85102"
              }
            ]
          },
          "modified": "2021-04-14T12:00:00Z",
          "published": "2021-04-14T12:00:00Z",
          "ecosystem_specific": {
            "symbols": [
              "NewV4",
              "rfc4122Generator.getClockSequence",
              "rfc4122Generator.getHardwareAddr"
            ]
          },
          "database_specific": {
            "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2020-0018.yaml",
            "source": "https://storage.googleapis.com/go-vulndb/github.com/satori/go.uuid.json"
          },
          "references": [
            {
              "type": "FIX",
              "url": "https://github.com/satori/go.uuid/pull/75"
            },
            {
              "type": "FIX",
              "url": "https://github.com/satori/go.uuid/commit/d91630c8510268e75203009fe7daf2b8e1d60c45"
            },
            {
              "type": "WEB",
              "url": "https://github.com/satori/go.uuid/issues/73"
            }
          ],
          "affected": [
            {
              "package": {
                "name": "github.com/satori/go.uuid",
                "ecosystem": "Go"
              },
              "ranges": [
                {
                  "type": "SEMVER",
                  "events": [
                    {
                      "introduced": "0"
                    },
                    {
                      "fixed": "1.2.1-0.20181016170032-d91630c85102"
                    }
                  ]
                }
              ],
              "ecosystem_specific": {
                "symbols": [
                  "NewV4",
                  "rfc4122Generator.getClockSequence",
                  "rfc4122Generator.getHardwareAddr"
                ]
              },
              "database_specific": {
                "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2020-0018.yaml",
                "source": "https://storage.googleapis.com/go-vulndb/github.com/satori/go.uuid.json"
              }
            }
          ]
        }
      ]
    },
    {
      "vulns": [
        {
          "id": "GO-2020-0016",
          "package": {
            "name": "github.com/ulikunitz/xz",
            "ecosystem": "Go"
          },
          "details": "An attacker can construct a series of bytes such that calling\n[`Reader.Read`] on the bytes could cause an infinite loop. If\nparsing user supplied input, this may be used as a denial of\nservice vector.\n",
          "affects": {
            "ranges": [
              {
                "type": "SEMVER",
                "fixed": "0.5.8"
              }
            ]
          },
          "aliases": [
            "CVE-2021-29482"
          ],
          "modified": "2021-04-14T12:00:00Z",
          "published": "2021-04-14T12:00:00Z",
          "ecosystem_specific": {
            "symbols": [
              "readUvarint"
            ]
          },
          "database_specific": {
            "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2020-0016.yaml",
            "source": "https://storage.googleapis.com/go-vulndb/github.com/ulikunitz/xz.json"
          },
          "references": [
            {
              "type": "FIX",
              "url": "https://github.com/ulikunitz/xz/commit/69c6093c7b2397b923acf82cb378f55ab2652b9b"
            },
            {
              "type": "WEB",
              "url": "https://github.com/ulikunitz/xz/issues/35"
            },
            {
              "type": "WEB",
              "url": "https://github.com/ulikunitz/xz/security/advisories/GHSA-25xm-hr59-7c27"
            }
          ],
          "affected": [
            {
              "package": {
                "name": "github.com/ulikunitz/xz",
                "ecosystem": "Go"
              },
              "ranges": [
                {
                  "type": "SEMVER",
                  "events": [
                    {
                      "introduced": "0"
                    },
                    {
                      "fixed": "0.5.8"
                    }
                  ]
                }
              ],
              "ecosystem_specific": {
                "symbols": [
                  "readUvarint"
                ]
              },
              "database_specific": {
                "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2020-0016.yaml",
                "source": "https://storage.googleapis.com/go-vulndb/github.com/ulikunitz/xz.json"
              }
            }
          ]
        }
      ]
    },
    {
      "vulns": [
        {
          "id": "GO-2020-0036",
          "package": {
            "name": "gopkg.in/yaml.v2",
            "ecosystem": "Go"
          },
          "details": "Due to unbounded aliasing, a crafted YAML file can cause consumption\nof significant system resources. If parsing user supplied input, this\nmay be used as a denial of service vector.\n",
          "affects": {
            "ranges": [
              {
                "type": "SEMVER",
                "fixed": "2.2.8"
              }
            ]
          },
          "aliases": [
            "CVE-2019-11254"
          ],
          "modified": "2021-04-14T12:00:00Z",
          "published": "2021-04-14T12:00:00Z",
          "ecosystem_specific": {
            "symbols": [
              "yaml_parser_fetch_more_tokens"
            ]
          },
          "database_specific": {
            "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2020-0036.yaml",
            "source": "https://storage.googleapis.com/go-vulndb/gopkg.in/yaml.v2.json"
          },
          "references": [
            {
              "type": "FIX",
              "url": "https://github.com/go-yaml/yaml/pull/555"
            },
            {
              "type": "FIX",
              "url": "https://github.com/go-yaml/yaml/commit/53403b58ad1b561927d19068c655246f2db79d48"
            },
            {
              "type": "WEB",
              "url": "https://bugs.chromium.org/p/oss-fuzz/issues/detail?id=18496"
            }
          ],
          "affected": [
            {
              "package": {
                "name": "gopkg.in/yaml.v2",
                "ecosystem": "Go"
              },
              "ranges": [
                {
                  "type": "SEMVER",
                  "events": [
                    {
                      "introduced": "0"
                    },
                    {
                      "fixed": "2.2.8"
                    }
                  ]
                }
              ],
              "ecosystem_specific": {
                "symbols": [
                  "yaml_parser_fetch_more_tokens"
                ]
              },
              "database_specific": {
                "source": "https://storage.googleapis.com/go-vulndb/gopkg.in/yaml.v2.json",
                "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2020-0036.yaml"
              }
            },
            {
              "package": {
                "name": "github.com/go-yaml/yaml",
                "ecosystem": "Go"
              },
              "ranges": [
                {
                  "type": "SEMVER",
                  "events": [
                    {
                      "introduced": "0"
                    }
                  ]
                }
              ],
              "ecosystem_specific": {
                "symbols": [
                  "yaml_parser_fetch_more_tokens"
                ]
              },
              "database_specific": {
                "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2020-0036.yaml",
                "source": "https://storage.googleapis.com/go-vulndb/gopkg.in/yaml.v2.json"
              }
            }
          ]
        },
        {
          "id": "GO-2021-0061",
          "package": {
            "name": "gopkg.in/yaml.v2",
            "ecosystem": "Go"
          },
          "details": "Due to unbounded alias chasing, a maliciously crafted YAML file\ncan cause the system to consume significant system resources. If\nparsing user input, this may be used as a denial of service vector.\n",
          "affects": {
            "ranges": [
              {
                "type": "SEMVER",
                "fixed": "2.2.3"
              }
            ]
          },
          "modified": "2021-04-14T12:00:00Z",
          "published": "2021-04-14T12:00:00Z",
          "ecosystem_specific": {
            "symbols": [
              "decoder.unmarshal"
            ]
          },
          "database_specific": {
            "source": "https://storage.googleapis.com/go-vulndb/gopkg.in/yaml.v2.json",
            "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2021-0061.yaml"
          },
          "references": [
            {
              "type": "FIX",
              "url": "https://github.com/go-yaml/yaml/pull/375"
            },
            {
              "type": "FIX",
              "url": "https://github.com/go-yaml/yaml/commit/bb4e33bf68bf89cad44d386192cbed201f35b241"
            }
          ],
          "affected": [
            {
              "package": {
                "name": "gopkg.in/yaml.v2",
                "ecosystem": "Go"
              },
              "ranges": [
                {
                  "type": "SEMVER",
                  "events": [
                    {
                      "introduced": "0"
                    },
                    {
                      "fixed": "2.2.3"
                    }
                  ]
                }
              ],
              "ecosystem_specific": {
                "symbols": [
                  "decoder.unmarshal"
                ]
              },
              "database_specific": {
                "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2021-0061.yaml",
                "source": "https://storage.googleapis.com/go-vulndb/gopkg.in/yaml.v2.json"
              }
            },
            {
              "package": {
                "name": "github.com/go-yaml/yaml",
                "ecosystem": "Go"
              },
              "ranges": [
                {
                  "type": "SEMVER",
                  "events": [
                    {
                      "introduced": "0"
                    }
                  ]
                }
              ],
              "ecosystem_specific": {
                "symbols": [
                  "decoder.unmarshal"
                ]
              },
              "database_specific": {
                "source": "https://storage.googleapis.com/go-vulndb/gopkg.in/yaml.v2.json",
                "url": "https://go.googlesource.com/vulndb/+/refs/heads/master/reports/GO-2021-0061.yaml"
              }
            }
          ]
        }
      ]
    }
  ]
}
```
</details>  

#### Why not print the output in table format?

This project aims to have the least amount of dependency to not worry about `osv` on dependencies.

#### Why name this as stunning-tribble mean?

GitHub generated the repository name. I am not good at naming things.

