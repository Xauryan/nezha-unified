module github.com/nezhahq/nezha

go 1.26.6

require (
	github.com/VictoriaMetrics/VictoriaMetrics v1.151.0
	github.com/appleboy/gin-jwt/v2 v2.10.3
	github.com/dustinkirkland/golang-petname v0.0.0-20260215035315-f0c533e9ce9b
	github.com/gin-contrib/pprof v1.5.4
	github.com/gin-gonic/gin v1.12.0
	github.com/go-viper/mapstructure/v2 v2.5.0
	github.com/goccy/go-json v0.10.6
	github.com/golang-jwt/jwt/v4 v4.5.2
	github.com/gorilla/websocket v1.5.4-0.20250319132907-e064f32e3674
	github.com/hashicorp/go-uuid v1.0.3
	github.com/jinzhu/copier v0.4.0
	github.com/knadh/koanf/maps v0.1.3
	github.com/knadh/koanf/providers/env v1.1.0
	github.com/knadh/koanf/providers/file v1.2.1
	github.com/knadh/koanf/v2 v2.3.6
	github.com/leonelquinteros/gotext v1.7.2
	github.com/libdns/cloudflare v0.2.2
	github.com/libdns/he v1.2.2
	github.com/libdns/libdns v1.1.1
	github.com/mattn/go-sqlite3 v1.14.50
	github.com/miekg/dns v1.1.73
	github.com/modelcontextprotocol/go-sdk v1.7.0
	github.com/nezhahq/libdns-tencentcloud v0.0.0-20260628095405-2ab294ec675b
	github.com/oschwald/maxminddb-golang v1.13.1
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/robfig/cron/v3 v3.0.1
	github.com/sqids/sqids-go v0.4.1
	github.com/stretchr/testify v1.12.1
	github.com/swaggo/files v1.0.1
	github.com/swaggo/gin-swagger v1.6.1
	github.com/swaggo/swag v1.16.6
	github.com/tidwall/gjson v1.19.0
	golang.org/x/crypto v0.55.0
	golang.org/x/exp v0.0.0-20260824195058-e88cd73687aa
	golang.org/x/mod v0.40.0
	golang.org/x/net v0.58.0
	golang.org/x/oauth2 v0.36.0
	golang.org/x/sync v0.22.0
	golang.org/x/sys v0.47.0
	google.golang.org/grpc v1.83.2
	google.golang.org/protobuf v1.36.12
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/sqlite v1.6.0
	gorm.io/gorm v1.31.2
	sigs.k8s.io/yaml v1.6.0
)

require (
	cloud.google.com/go/auth v0.23.2 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.8 // indirect
	cloud.google.com/go/compute/metadata v0.9.0 // indirect
	filippo.io/edwards25519 v1.2.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.23.1 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.14.1 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.12.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns v1.2.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns v1.3.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph v0.10.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azcertificates v1.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/internal v1.2.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.8.0 // indirect
	github.com/BurntSushi/toml v1.6.0 // indirect
	github.com/G-Core/gcorelabscdn-go v1.0.37 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/VictoriaMetrics/easyproto v1.2.0 // indirect
	github.com/VictoriaMetrics/fastcache v1.13.3 // indirect
	github.com/VictoriaMetrics/metrics v1.44.0 // indirect
	github.com/VictoriaMetrics/metricsql v0.87.3 // indirect
	github.com/akamai/AkamaiOPEN-edgegrid-golang/v13 v13.4.0 // indirect
	github.com/alexbrainman/sspi v0.0.0-20250919150558-7d374ff0d59e // indirect
	github.com/alibabacloud-go/alb-20200616/v2 v2.3.2 // indirect
	github.com/alibabacloud-go/alibabacloud-gateway-fc-util v0.0.7 // indirect
	github.com/alibabacloud-go/alibabacloud-gateway-spi v0.0.5 // indirect
	github.com/alibabacloud-go/apig-20240327/v11 v11.2.3 // indirect
	github.com/alibabacloud-go/cas-20200407/v4 v4.7.1 // indirect
	github.com/alibabacloud-go/cdn-20180510/v10 v10.2.0 // indirect
	github.com/alibabacloud-go/cloudapi-20160714/v5 v5.8.0 // indirect
	github.com/alibabacloud-go/darabonba-openapi/v2 v2.2.4 // indirect
	github.com/alibabacloud-go/dcdn-20180115/v4 v4.1.0 // indirect
	github.com/alibabacloud-go/ddoscoo-20200101/v5 v5.0.2 // indirect
	github.com/alibabacloud-go/debug v1.0.1 // indirect
	github.com/alibabacloud-go/endpoint-util v1.1.1 // indirect
	github.com/alibabacloud-go/esa-20240910/v3 v3.13.0 // indirect
	github.com/alibabacloud-go/fc-20230330/v4 v4.8.2 // indirect
	github.com/alibabacloud-go/fc-open-20210406/v2 v2.0.12 // indirect
	github.com/alibabacloud-go/ga-20191120/v4 v4.0.2 // indirect
	github.com/alibabacloud-go/live-20161101/v3 v3.0.1 // indirect
	github.com/alibabacloud-go/nlb-20220430/v4 v4.1.3 // indirect
	github.com/alibabacloud-go/openapi-util v0.1.2 // indirect
	github.com/alibabacloud-go/slb-20140515/v4 v4.0.14 // indirect
	github.com/alibabacloud-go/tea v1.5.3 // indirect
	github.com/alibabacloud-go/tea-utils/v2 v2.0.9 // indirect
	github.com/alibabacloud-go/vod-20170321/v4 v4.12.1 // indirect
	github.com/alibabacloud-go/waf-openapi-20211001/v7 v7.8.5 // indirect
	github.com/aliyun/credentials-go v1.4.12 // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/avast/retry-go v3.0.0+incompatible // indirect
	github.com/aws/aws-sdk-go-v2 v1.47.0 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.33.5 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.20.5 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.20.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.5.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.8.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.5.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/acm v1.50.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/amplify v1.48.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.42.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.73.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.41.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.63.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/iam v1.64.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.14.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.66.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53 v1.64.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/signin v1.10.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.38.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.43.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.51.0 // indirect
	github.com/aws/smithy-go v1.28.1 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/baidubce/bce-sdk-go v0.9.276 // indirect
	github.com/benbjohnson/clock v1.3.5 // indirect
	github.com/bodgit/gssapi v0.0.4 // indirect
	github.com/bodgit/tsig v1.3.1 // indirect
	github.com/boombuler/barcode v1.1.0 // indirect
	github.com/bytedance/gopkg v0.1.4 // indirect
	github.com/bytedance/sonic v1.15.1 // indirect
	github.com/bytedance/sonic/loader v0.5.1 // indirect
	github.com/byteplus-sdk/byteplus-go-sdk-v2 v1.0.78 // indirect
	github.com/byteplus-sdk/byteplus-sdk-golang v1.0.72 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cenkalti/backoff/v5 v5.0.3 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/clbanning/mxj/v2 v2.7.0 // indirect
	github.com/cloudwego/base64x v0.1.7 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/disintegration/imaging v1.6.2 // indirect
	github.com/domodwyer/mailyak/v3 v3.6.2 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/fatih/color v1.19.0 // indirect
	github.com/felixge/httpsnoop v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.10.1 // indirect
	github.com/fxamacker/cbor/v2 v2.9.2 // indirect
	github.com/gabriel-vasile/mimetype v1.4.13 // indirect
	github.com/ganigeorgiev/fexpr v0.6.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gin-contrib/sse v1.1.1 // indirect
	github.com/go-acme/alidns-20150109/v5 v5.5.0 // indirect
	github.com/go-acme/esa-20240910/v3 v3.4.0 // indirect
	github.com/go-acme/jdcloud-sdk-go v1.64.0 // indirect
	github.com/go-acme/lego/v5 v5.3.1 // indirect
	github.com/go-acme/tencentclouddnspod v1.3.24 // indirect
	github.com/go-acme/tencentedgdeone v1.3.38 // indirect
	github.com/go-cmd/cmd v1.4.3 // indirect
	github.com/go-jose/go-jose/v4 v4.1.4 // indirect
	github.com/go-logr/logr v1.4.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/jsonpointer v1.0.0 // indirect
	github.com/go-openapi/jsonreference v1.0.0 // indirect
	github.com/go-openapi/spec v0.22.4 // indirect
	github.com/go-openapi/swag/conv v0.27.1 // indirect
	github.com/go-openapi/swag/jsonname v0.26.0 // indirect
	github.com/go-openapi/swag/jsonutils v0.27.1 // indirect
	github.com/go-openapi/swag/loading v0.27.1 // indirect
	github.com/go-openapi/swag/pools v0.27.1 // indirect
	github.com/go-openapi/swag/stringutils v0.27.1 // indirect
	github.com/go-openapi/swag/typeutils v0.27.1 // indirect
	github.com/go-openapi/swag/yamlutils v0.27.1 // indirect
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.30.3 // indirect
	github.com/go-resty/resty/v2 v2.17.2 // indirect
	github.com/goccy/go-yaml v1.19.2 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v1.0.0 // indirect
	github.com/google/go-querystring v1.2.0 // indirect
	github.com/google/jsonschema-go v0.4.3 // indirect
	github.com/google/s2a-go v0.1.9 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.20 // indirect
	github.com/googleapis/gax-go/v2 v2.24.0 // indirect
	github.com/gorilla/css v1.0.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.8 // indirect
	github.com/huaweicloud/huaweicloud-sdk-go-v3 v0.1.215 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/goidentity/v6 v6.0.1 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jdcloud-api/jdcloud-sdk-go v1.67.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jlaffaye/ftp v0.2.4 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.13-0.20220915233716-71ac16282d12 // indirect
	github.com/klauspost/compress v1.19.2 // indirect
	github.com/klauspost/cpuid/v2 v2.4.0 // indirect
	github.com/klauspost/crc32 v1.3.0 // indirect
	github.com/kr/fs v0.1.0 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.15 // indirect
	github.com/mattn/go-isatty v0.0.23 // indirect
	github.com/microcosm-cc/bluemonday v1.0.27 // indirect
	github.com/minio/crc64nvme v1.1.1 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/minio-go/v7 v7.3.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.3-0.20250322232337-35a7c28c31ee // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/namedotcom/go/v4 v4.0.2 // indirect
	github.com/ncruces/go-strftime v1.0.0 // indirect
	github.com/nrdcg/bunny-go v0.1.0 // indirect
	github.com/nrdcg/desec v0.11.1 // indirect
	github.com/nrdcg/goacmedns v0.2.0 // indirect
	github.com/nrdcg/namesilo v0.5.0 // indirect
	github.com/nrdcg/oci-go-sdk/certificatesmanagement/v1065 v1065.124.2 // indirect
	github.com/nrdcg/oci-go-sdk/common/v1065 v1065.124.2 // indirect
	github.com/nrdcg/oci-go-sdk/dns/v1065 v1065.120.0 // indirect
	github.com/nrdcg/porkbun v0.4.0 // indirect
	github.com/openshift/gssapi v0.0.0-20161010215902-5fb4217df13b // indirect
	github.com/ovh/go-ovh v1.9.0 // indirect
	github.com/pavlo-v-chernykh/keystore-go/v4 v4.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.3.1 // indirect
	github.com/peterhellberg/link v1.2.0 // indirect
	github.com/philhofer/fwd v1.2.0 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pkg/sftp v1.13.11 // indirect
	github.com/pocketbase/dbx v1.12.0 // indirect
	github.com/pocketbase/ozzo-validation/v4 v4.3.0 // indirect
	github.com/pocketbase/pocketbase v0.39.10 // indirect
	github.com/povsister/scp v0.0.0-20250701154629-777cf82de5df // indirect
	github.com/pquerna/otp v1.5.0 // indirect
	github.com/qiniu/dyn v1.3.0 // indirect
	github.com/qiniu/go-sdk/v7 v7.28.1 // indirect
	github.com/qiniu/x v1.17.0 // indirect
	github.com/quic-go/qpack v0.6.0 // indirect
	github.com/quic-go/quic-go v0.59.1 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/rs/xid v1.6.0 // indirect
	github.com/samber/lo v1.53.0 // indirect
	github.com/segmentio/asm v1.1.3 // indirect
	github.com/segmentio/encoding v0.5.4 // indirect
	github.com/sirupsen/logrus v1.9.4 // indirect
	github.com/sony/gobreaker/v2 v2.4.0 // indirect
	github.com/spf13/afero v1.15.0 // indirect
	github.com/spf13/cast v1.10.0 // indirect
	github.com/spf13/cobra v1.10.2 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/stretchr/objx v0.5.3 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/antiddos v1.3.131 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn v1.3.173 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/clb v1.3.176 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.3.180 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ga2 v1.3.180 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gaap v1.3.34 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live v1.3.175 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/scf v1.3.101 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssl v1.3.147 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/teo v1.3.180 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tse v1.3.168 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod v1.3.179 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/waf v1.3.171 // indirect
	github.com/tidwall/match v1.2.0 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
	github.com/tinylib/msgp v1.6.4 // indirect
	github.com/tjfoc/gmsm v1.4.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ucloud/ucloud-sdk-go v0.22.124 // indirect
	github.com/ugorji/go/codec v1.3.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fastrand v1.1.0 // indirect
	github.com/valyala/gozstd v1.25.0 // indirect
	github.com/valyala/histogram v1.2.0 // indirect
	github.com/valyala/quicktemplate v1.8.0 // indirect
	github.com/volcengine/volc-sdk-golang v1.0.256 // indirect
	github.com/volcengine/volcengine-go-sdk v1.2.52 // indirect
	github.com/vultr/govultr/v3 v3.31.2 // indirect
	github.com/wneessen/go-mail v0.8.1 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/xhit/go-str2duration/v2 v2.2.0 // indirect
	github.com/yandex-cloud/go-genproto v0.121.0 // indirect
	github.com/yandex-cloud/go-sdk/services/certificatemanager v0.0.97 // indirect
	github.com/yandex-cloud/go-sdk/services/dns v0.0.65 // indirect
	github.com/yandex-cloud/go-sdk/v2 v2.176.0 // indirect
	github.com/yosida95/uritemplate/v3 v3.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	github.com/yunify/qingcloud-sdk-go v2.0.0-alpha.38+incompatible // indirect
	github.com/zeebo/xxh3 v1.1.0 // indirect
	github.com/zenlayer/zenlayercloud-sdk-go v0.2.53 // indirect
	gitlab.ecloud.com/ecloud/ecloudsdkcloudcore v1.0.0 // indirect
	gitlab.ecloud.com/ecloud/ecloudsdkclouddns v1.0.1 // indirect
	gitlab.ecloud.com/ecloud/ecloudsdkcmcdn v1.0.0 // indirect
	gitlab.ecloud.com/ecloud/ecloudsdkcore v1.0.6 // indirect
	gitlab.ecloud.com/ecloud/ecloudsdkvlb v1.0.7 // indirect
	go.mongodb.org/mongo-driver v1.17.9 // indirect
	go.mongodb.org/mongo-driver/v2 v2.6.0 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.69.0 // indirect
	go.opentelemetry.io/otel v1.44.0 // indirect
	go.opentelemetry.io/otel/metric v1.44.0 // indirect
	go.opentelemetry.io/otel/trace v1.44.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/ratelimit v0.3.1 // indirect
	go.uber.org/zap v1.28.0 // indirect
	go.yaml.in/yaml/v2 v2.4.4 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/arch v0.27.0 // indirect
	golang.org/x/image v0.44.0 // indirect
	golang.org/x/term v0.45.0 // indirect
	golang.org/x/text v0.41.0 // indirect
	golang.org/x/time v0.15.0 // indirect
	golang.org/x/tools v0.49.0 // indirect
	google.golang.org/api v0.295.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20260720211330-0afa2a65878a // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260819154853-08b0e4226688 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.3 // indirect
	gopkg.in/ns1/ns1-go.v2 v2.18.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/api v0.36.3 // indirect
	k8s.io/apimachinery v0.36.3 // indirect
	k8s.io/client-go v0.36.3 // indirect
	k8s.io/klog/v2 v2.140.0 // indirect
	k8s.io/kube-openapi v0.0.0-20260721132016-d427ff9ee9ad // indirect
	k8s.io/utils v0.0.0-20260707023825-cf1189d6abe3 // indirect
	modernc.org/libc v1.74.1 // indirect
	modernc.org/mathutil v1.7.1 // indirect
	modernc.org/memory v1.11.0 // indirect
	modernc.org/sqlite v1.55.0 // indirect
	sigs.k8s.io/json v0.0.0-20250730193827-2d320260d730 // indirect
	sigs.k8s.io/randfill v1.0.0 // indirect
	sigs.k8s.io/structured-merge-diff/v6 v6.4.2 // indirect
	software.sslmate.com/src/go-pkcs12 v0.7.3 // indirect
)

require (
	github.com/certimate-go/certimate v0.0.0
	github.com/gofrs/flock v0.13.1
)

replace github.com/certimate-go/certimate => ./third_party/certimate

replace gitlab.ecloud.com/ecloud/ecloudsdkcloudcore v1.0.0 => ./third_party/certimate/pkg/sdk3rd-forked/gitlab.ecloud.com/ecloud/ecloudsdkcloudcore@v1.0.0+mod

replace gitlab.ecloud.com/ecloud/ecloudsdkclouddns v1.0.1 => ./third_party/certimate/pkg/sdk3rd-forked/gitlab.ecloud.com/ecloud/ecloudsdkclouddns@v1.0.1+mod

replace gitlab.ecloud.com/ecloud/ecloudsdkcmcdn v1.0.0 => ./third_party/certimate/pkg/sdk3rd-forked/gitlab.ecloud.com/ecloud/ecloudsdkcmcdn@v1.0.0

replace gitlab.ecloud.com/ecloud/ecloudsdkcore v1.0.2 => ./third_party/certimate/pkg/sdk3rd-forked/gitlab.ecloud.com/ecloud/ecloudsdkcore@v1.0.2

replace gitlab.ecloud.com/ecloud/ecloudsdkcore v1.0.6 => ./third_party/certimate/pkg/sdk3rd-forked/gitlab.ecloud.com/ecloud/ecloudsdkcore@v1.0.6

replace gitlab.ecloud.com/ecloud/ecloudsdkvlb v1.0.7 => ./third_party/certimate/pkg/sdk3rd-forked/gitlab.ecloud.com/ecloud/ecloudsdkvlb@v1.0.7
