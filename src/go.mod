module github.com/alcionai/corso/src

go 1.19

replace github.com/kopia/kopia => github.com/alcionai/kopia v0.12.2-0.20230502235504-2509b1d72a79

require (
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.2.0
	github.com/alcionai/clues v0.0.0-20230406223931-f48777f4773c
	github.com/armon/go-metrics v0.4.1
	github.com/aws/aws-sdk-go v1.44.275
	github.com/aws/aws-xray-sdk-go v1.8.1
	github.com/cenkalti/backoff/v4 v4.2.1
	github.com/google/uuid v1.3.0
	github.com/h2non/gock v1.2.0
	github.com/kopia/kopia v0.12.2-0.20230327171220-747baeebdab1
	github.com/microsoft/kiota-abstractions-go v1.0.0
	github.com/microsoft/kiota-authentication-azure-go v1.0.0
	github.com/microsoft/kiota-http-go v1.0.0
	github.com/microsoft/kiota-serialization-form-go v1.0.0
	github.com/microsoft/kiota-serialization-json-go v1.0.1
	github.com/microsoftgraph/msgraph-sdk-go v1.4.0
	github.com/microsoftgraph/msgraph-sdk-go-core v1.0.0
	github.com/pkg/errors v0.9.1
	github.com/rudderlabs/analytics-go v3.3.3+incompatible
	github.com/spatialcurrent/go-lazy v0.0.0-20211115014721-47315cc003d1
	github.com/spf13/cobra v1.7.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.16.0
	github.com/stretchr/testify v1.8.4
	github.com/tidwall/pretty v1.2.1
	github.com/tomlazar/table v0.1.2
	github.com/vbauerster/mpb/v8 v8.1.6
	go.uber.org/zap v1.24.0
	golang.org/x/exp v0.0.0-20230213192124-5e25df0256eb
	golang.org/x/time v0.3.0
	golang.org/x/tools v0.9.3
)

require (
	github.com/VividCortex/ewma v1.2.0 // indirect
	github.com/acarl005/stripansi v0.0.0-20180116102854-5a71ef0e047d // indirect
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/dnaeon/go-vcr v1.2.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gofrs/flock v0.8.1 // indirect
	github.com/h2non/parth v0.0.0-20190131123155-b4df798d6542 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.34.0 // indirect
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.6.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.3.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v0.7.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chmduquesne/rollinghash v4.0.0+incompatible // indirect
	github.com/cjlapao/common-go v0.0.39 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dustin/go-humanize v1.0.1
	github.com/edsrzf/mmap-go v1.1.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.16.5 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/klauspost/pgzip v1.2.5 // indirect
	github.com/klauspost/reedsolomon v1.11.7 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/microsoft/kiota-serialization-text-go v1.0.0
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/minio-go/v7 v7.0.52 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/natefinch/atomic v1.0.1 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.15.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.42.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/rivo/uniseg v0.4.3 // indirect
	github.com/rs/xid v1.4.0 // indirect
	github.com/segmentio/backo-go v1.0.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/tidwall/gjson v1.14.3 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/xtgo/uuid v0.0.0-20140804021211-a0b114877d4c // indirect
	github.com/yosida95/uritemplate/v3 v3.0.2 // indirect
	github.com/zeebo/blake3 v0.2.3 // indirect
	go.opentelemetry.io/otel v1.15.1 // indirect
	go.opentelemetry.io/otel/trace v1.15.1 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/mod v0.10.0 // indirect
	golang.org/x/net v0.10.0
	golang.org/x/sync v0.2.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/grpc v1.55.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
