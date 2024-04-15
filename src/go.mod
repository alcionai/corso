module github.com/alcionai/corso/src

go 1.21

replace (
	github.com/kopia/kopia => github.com/alcionai/kopia v0.12.2-0.20240322180947-41471159a0a4

	// Alcion fork removes the validation of email addresses as we might get incomplete email addresses
	github.com/xhit/go-simple-mail/v2 v2.16.0 => github.com/alcionai/go-simple-mail/v2 v2.0.0-20231220071811-c70ebcd9a41a
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.5.1
	github.com/alcionai/clues v0.0.0-20240125221452-9fc7746dd20c
	github.com/armon/go-metrics v0.4.1
	github.com/aws/aws-xray-sdk-go v1.8.3
	github.com/cenkalti/backoff/v4 v4.2.1
	github.com/fatih/color v1.16.0
	github.com/golang-jwt/jwt/v5 v5.2.0
	github.com/google/uuid v1.6.0
	github.com/h2non/gock v1.2.0
	github.com/jaytaylor/html2text v0.0.0-20230321000545-74c2419ad056
	github.com/jhillyerd/enmime v1.1.0
	github.com/kopia/kopia v0.15.0
	github.com/microsoft/kiota-abstractions-go v1.6.0
	github.com/microsoft/kiota-authentication-azure-go v1.0.2
	github.com/microsoft/kiota-http-go v1.3.1
	github.com/microsoft/kiota-serialization-form-go v1.0.0
	github.com/microsoft/kiota-serialization-json-go v1.0.7
	github.com/microsoftgraph/msgraph-sdk-go v1.37.0
	github.com/microsoftgraph/msgraph-sdk-go-core v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/puzpuzpuz/xsync/v3 v3.0.2
	github.com/rudderlabs/analytics-go v3.3.3+incompatible
	github.com/spatialcurrent/go-lazy v0.0.0-20211115014721-47315cc003d1
	github.com/spf13/cast v1.6.0
	github.com/spf13/cobra v1.8.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.18.2
	github.com/stretchr/testify v1.9.0
	github.com/tidwall/pretty v1.2.1
	github.com/tomlazar/table v0.1.2
	github.com/vbauerster/mpb/v8 v8.1.6 // do not update; keep at v8.1.6
	github.com/xhit/go-simple-mail/v2 v2.16.0
	go.uber.org/goleak v1.3.0
	go.uber.org/zap v1.26.0
	golang.org/x/exp v0.0.0-20231127185646-65229373498e
	golang.org/x/time v0.5.0
	golang.org/x/tools v0.17.0
	gotest.tools/v3 v3.5.1
)

require (
	github.com/arran4/golang-ical v0.2.4
	github.com/emersion/go-vcard v0.0.0-20230815062825-8fda7d206ec9
	jaytaylor.com/html2text v0.0.0-20230321000545-74c2419ad056
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/storage/azblob v1.2.1 // indirect
	github.com/VividCortex/ewma v1.2.0 // indirect
	github.com/acarl005/stripansi v0.0.0-20180116102854-5a71ef0e047d // indirect
	github.com/andybalholm/brotli v1.0.6 // indirect
	github.com/aws/aws-sdk-go v1.48.6 // indirect
	github.com/cention-sany/utf7 v0.0.0-20170124080048-26cad61bd60a // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gofrs/flock v0.8.1 // indirect
	github.com/gogs/chardet v0.0.0-20211120154057-b7413eaefb8f // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/h2non/parth v0.0.0-20190131123155-b4df798d6542 // indirect
	github.com/hashicorp/cronexpr v1.1.2 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/microsoft/kiota-serialization-multipart-go v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/ssor/bom v0.0.0-20170718123548-6386211fdfcf // indirect
	github.com/std-uritemplate/std-uritemplate/go v0.0.55 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/toorop/go-dkim v0.0.0-20201103131630-e1cd1a0a5208 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	go.opentelemetry.io/otel/metric v1.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231212172506-995d672761c0 // indirect
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.10.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.5.2 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.2.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chmduquesne/rollinghash v4.0.0+incompatible // indirect
	github.com/cjlapao/common-go v0.0.39 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dustin/go-humanize v1.0.1
	github.com/edsrzf/mmap-go v1.1.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.4 // indirect
	github.com/klauspost/cpuid/v2 v2.2.6 // indirect
	github.com/klauspost/pgzip v1.2.6 // indirect
	github.com/klauspost/reedsolomon v1.12.0 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/microsoft/kiota-serialization-text-go v1.0.0
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/minio-go/v7 v7.0.67
	github.com/minio/sha256-simd v1.0.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/natefinch/atomic v1.0.1 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.18.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/rs/xid v1.5.0 // indirect
	github.com/segmentio/backo-go v1.0.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/tidwall/gjson v1.17.0
	github.com/tidwall/match v1.1.1 // indirect
	github.com/xtgo/uuid v0.0.0-20140804021211-a0b114877d4c // indirect
	github.com/zeebo/blake3 v0.2.3 // indirect
	go.opentelemetry.io/otel v1.24.0 // indirect
	go.opentelemetry.io/otel/trace v1.24.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.19.0 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.21.0
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/grpc v1.60.1 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
