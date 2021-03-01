# golang 语言常见第三方包

> 此处为个人常用
> 
> 更广泛的可参考 [awesome-go](https://github.com/avelino/awesome-go)

---

| 功能简要 | 分类 | 仓库地址 | 类别 | 备注 |
| --- | --- | --- | --- | --- | 
| 格式化输出 | 格式化 | [fmt](https://github.com/golang/go/tree/master/src/fmt) | 系统包 | 
| 时间处理 | 时间 | [time](https://github.com/golang/go/tree/master/src/time) | 系统包 |
| 字符串处理 | 字符串 | [strings](https://github.com/golang/go/tree/master/src/strings) | 系统包 | 
| 字符处理 | 字符 | [strconv](https://github.com/golang/go/tree/master/src/strconv) | 系统包 | 
| 排序 | 排序 | [sort](https://golang.org/pkg/sort/) | 系统包
| MD5/SHA256等 | Hash | [hash](https://github.com/golang/go/tree/master/src/hash) | 系统包 | 
| md5 | md5 | [md5](https://golang.org/pkg/crypto/md5/) | 系统包 | 
| 基本数学计算  | 数学计算 | [math](https://github.com/golang/go/tree/master/src/math) | 系统包 |
| 随机数 | 数学计算 | [rand](https://golang.org/pkg/math/rand/) | 系统包 | 
| 正则表达式 | 正则表达式 | [regexp](https://golang.org/pkg/regexp/) | 系统包 | 
| 操作系统 | OS | [os](https://github.com/golang/go/tree/master/src/os) | 系统包 | 
| 文件 | 文件 | [ioutil](https://github.com/golang/go/tree/master/src/io/ioutil) | 系统包 | 
| context | context | [context](https://golang.org/pkg/context/) | 系统包 | 
| 并发 | 并发 | [sync](https://golang.org/pkg/sync/) | 系统包 | 
| 类型解析| 反射 | [reflect](https://golang.org/pkg/reflect/) | 系统包| 
| 序列化 |序列化| [encoding/json](https://github.com/golang/go/tree/master/src/encoding/json) | 系统包 | 序列化 |
| 序列化 | 序列化|[vmihailenco/msgpack](https://github.com/vmihailenco/msgpack) | 第三方 | 序列化 | [msgpack 官方](https://github.com/msgpack/msgpack-go) | 
| MySQL ORM | ORM | [go-gorm/gorm](https://github.com/go-gorm/gorm) |  第三方 | 此为V2, V1 见 [jinzhu/gorm](github.com/jinzhu/gorm) |
| XORM ORM | ORM | [XORM](xorm.io/xorm) | 第三方 | 
| Mongo ORM | ORM| [mgo](https://github.com/globalsign/mgo) | 第三方 | 
| Snowflake | ID |[bwmarrin/snowflake](github.com/bwmarrin/snowflake) | 第三方 | twitter 官方 分布式ID的go语言实现 | ID |
| UUID | UUID |[satori/go.uuid](github.com/satori/go.uuid) | 官方 | ID | 
| Web 框架 | Web |[gin-gonic/gin](github.com/gin-gonic/gin) | 第三方 |
| Web 框架 | Web | [beego](github.com/astaxie/beego) | 第三方 | 
| Web 参数校验 | Web |[go-playground/validator](github.com/go-playground/validator/v10) |  第三方 | 
| Web API 文档 | Web API |[go-swagger/go-swagger](https://github.com/go-swagger/go-swagger) | 第三方 |
| Web API 文档 | Web API |[swaggo/swag](https://github.com/swaggo/swag) | 第三方 | |
| Web JWT 认证 | Web 认证 | b[dgrijalva/jwt-go](github.com/dgrijalva/jwt-go) | 第三方 | Web 安全认证 |
| Redis 驱动 | Redis | [go-redis/redis](github.com/go-redis/redis/v7) | 第三方 | 
| decimal | - | [shopspring/decimal](github.com/shopspring/decimal) | 第三方 |
| 科学计算 | - | [gonum](gonum.org/v1/gonum/mat) | 第三方 | 如概率论/线性代数等 | 
| 日志输出 |log| [sirupsen/logrus](github.com/sirupsen/logrus) | 第三方 | 
| 命令行解析 | command line | [alecthomas/kingpin.v2](gopkg.in/alecthomas/kingpin.v2) |  第三方 |
| 配置文件解析 ini|配置文件解析| [ini](gopkg.in/ini.v1) |  第三方 | 
| 配置文件解析 yaml | 配置文件解析| [yaml](gopkg.in/yaml.v2) | 第三方
| 配置文件解析 toml | 配置文件解析 | [toml](https://github.com/BurntSushi/toml) | 第三方 | 
| 火币 | 币商|[huobirdcenter/huobi_golang](github.com/huobirdcenter/huobi_golang) |  第三方 | 币商
| gateio |币商|  [gateio/gateapi-go](github.com/gateio/gateapi-go/v5) | 第三方 | 币商
| 支付-支付宝 |支付| [smartwalle/alipay/v3](github.com/smartwalle/alipay/v3) | 第三方 |  支付 
| 支付-微信支付 | 支付|[]() | 第三方 | 支付
| 微信公众号 | 微信|[silenceper/wechat](github.com/silenceper/wechat/v2) | 第三方
| 阿里云OSS | OSS|[aliyun/aliyun-oss-go-sdk](github.com/aliyun/aliyun-oss-go-sdk) | 第三方 | OSS
| 腾讯与OSS | OSS|[tencentyun/cos-go-sdk-v5](github.com/tencentyun/cos-go-sdk-v5) | 第三方 | OSS
| websocket | websocket | [websocket](github.com/gorilla/websocket) | 第三方 | 
| RPIO | RPIO| [go-rpio](github.com/stianeikeland/go-rpio) |  第三方 | 控制树莓派 | 
