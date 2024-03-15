### 知识1

mac上关闭某个端口的方法

1. 查找端口的进程`ID`——`PID`

   ```
   sudo lsof -i :<端口号>
   ```

   出现类似以下信息

   ```
   COMMAND  PID    USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
   main    6015 yinsiyu    5u  IPv6 0xb9e2185adf72afff      0t0  TCP *:8849 (LISTEN)
   ```

2. 使用`kill`命令关闭它

   ```
   sudo kill <PID>
   ```

   

### 1.初始化项目

```
go mod init <golang_project_name>
```

### 2.安装`gin`

```
go get -u github.com/gin-gonic/gin
```

### 3.安装`gorm`(可选)

```
go get -u gorm.io/gorm
```

### 4.创建`main.go`

```go
func main() {

	WaitExit()
}
// 启动程序时打印消息
func WaitExit() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server.")
}
```

### 5.配置程序

创建`/config`目录存放`yml`配置文件

创建`/app/config`目录存放配置程序

#### 方式一.使用环境变量

提前在`docker`环境中设置好需要的环境变量

#### 方式二.使用`yml`文件

在根目录创建`config`用于存放配置的`yml`文件。

安装包`gopkg.in/yaml.v2`,用于处理`yaml`文件

```
go get gopkg.in/yaml.v2
```

`app.yml`

```
Port: 8849
Debug: true
Swagger: true
ReadTimeout: 120
WriteTimeout: 120
MaximumUploadFileSize: 20971520
```



在`/app/config/app.go`中进行配置

首先根据环境变量配置，如果未设置再根据`yml`文件进行配置

```go
type AppConfig struct {
	// System define
	Port                  int  `yaml:"Port"`
	Debug                 bool `yaml:"Debug"`
	Mode                  string
	MaximumUploadFileSize int64  `yaml:"MaximumUploadFileSize"`
	FrontendURL           string `yaml:"FrontendURL"`
	BackendURL            string `yaml:"BackendURL"`
}

func (app *AppConfig) setAppMode() {
	if app.Debug {
		app.Mode = "debug"
	} else {
		app.Mode = "release"
	}
}

func (app *AppConfig) overwritePortIfNeeded(key string) error {
	port := os.Getenv(key)
	portNumber, err := strconv.Atoi(port)
	if err == nil && portNumber > 0 && portNumber < 65536 {
		app.Port = portNumber
	}
	return err
}

var App AppConfig

func init() {
	var setting AppConfig

	if util.GetEnvBooleanValue("APP_USE_ENV") {
		// 使用环境变量配置项目
		setting = AppConfig{
			Debug:       util.GetEnvBooleanValue("APP_DEBUG"),
			FrontendURL: os.Getenv("APP_FRONTEND_URL"),
			BackendURL:  os.Getenv("APP_BACKEND_URL"),
		}
		if err := setting.overwritePortIfNeeded("APP_PORT"); err != nil {
			setting.Port = 8080
		}
		maxUploadFileSize, _ := strconv.Atoi(os.Getenv("APP_MAXIMUM_UPLOAD_FILE_SIZE"))
		setting.MaximumUploadFileSize = int64(maxUploadFileSize)
	} else {
		// 使用yml配置项目
		config, err := os.ReadFile("config/app.yml")
		if err != nil {
			log.Fatal("App config not set.")
		}
		yamlErr := yaml.Unmarshal(config, &setting)
		if yamlErr != nil {
			log.Fatal("App config read error.")
		}
	}
  App = setting

	App.setAppMode()
  App.overwritePortIfNeeded("PORT")
}
```

### 6.启动`gin server`

创建`/app/boot`目录，在`boot`中存放启动程序

`/app/boot/gin.go`

```go
// to launch gin server
func GinServer() {
	// 创建gin实例
	engine := gin.New()

	engine.MaxMultipartMemory = config.App.MaximumUploadFileSize

	// CORS

	// Routers

	// 配置server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.App.Port),
		Handler:      engine,
		ReadTimeout:  time.Duration(60) * time.Second,
		WriteTimeout: time.Duration(60) * time.Second,
	}

	// Start server
	go func() {
		log.Println("Server started.")
		log.Println("Port" + server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server listen: %s\n", err)
		}
	}()
}

func init() {
	gin.SetMode(config.App.Mode)
}
```

### 7.注册路由

创建`app/router`目录存放路由

根目录分为`/`和`/api`两组,目的是为了让`api`结构更加清晰

#### `root`

```go
func SetupRootRouter(engine *gin.Engine) {
	rootRouter := engine.Group("/")
	rootRouter.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is running.")
	})
}
```

#### `api`

```go
func SetupApiRouter(engine *gin.Engine) {
	apiRouter := engine.Group("api")

	// test
	testApiRouter := apiRouter.Group("/test")
	testApiRouter.GET("/hello", func(ctx *gin.Context) {
		// ...
	})

	// auth

	// user

	// ...
}
```

#### 使用路由

在`/app/boot/gin.go`中加入上面注册的路由代码

```go
	// Routers
	router.SetupRootRouter(engine)
	router.SetupApiRouter(engine)
```

### 8.使用`MVC`架构

创建目录`/app/controller`

创建目录`/app/service`

创建目录`/app/dao`

### 9.封装`response`

创建目录`/app/lib/response/`,存放相应类型

```go
func (r *Gin) SendJSON(code int, obj interface{}) {
	r.Ctx.Header("Content-Type", "application/json; charset=utf-8")
	r.Ctx.Header("Cache-Control", "no-cache")
	r.Ctx.Header("Pragma", "no-cache")
	r.Ctx.Header("Expires", "0")
	r.Ctx.Header("X-Content-Type-Options", "nosniff")
	r.Ctx.JSON(code, obj)
}

// 200
func (r *Gin) Success(response interface{}) {
	r.SendJSON(http.StatusOK, response)
}

// 400
	// ValidationError
	// ClientError
	
// 404

// 500

// ...
```

创建目录`/app/openapi/`，存放结构体类型（包括相应体类型）

```go
type CommonResponse struct {
	Message string         `json:"message"`
	Status  ResponseStatus `json:"status"`
}
```

### 10.配置并使用中间件

创建目录`/app/middleware`,存放中间件

#### `CORS`（全局中间件）

导入相关的包

```
go get github.com/gin-contrib/cors
```

基础配置

`/app/middleware/cors.go`

```go
var CORS gin.HandlerFunc

func init() {
	CORS = cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	})
}
```

**使用**

`/app/boot/gin.go`

```go
	// CORS
	engine.Use(middleware.CORS)
```

#### 局部中间件

演示

`/app/middleware/auth.go`

```go
var ParamFilter gin.HandlerFunc

func init() {
	// 演示
	// 需求：只有携带了参数value，且 value >= 10 的GET请求才能通过

	ParamFilter = func(ctx *gin.Context) {
		resp := response.Gin{Ctx: ctx}

		// 只允许通过GET请求
		if resp.Ctx.Request.Method != "GET" {
      // 终止请求
			ctx.Abort()
			resp.ClientError("不是GET请求")
			return
		}

		param, err := strconv.Atoi(ctx.Query("value"))

		if err != nil {
			ctx.Abort()
			resp.ClientError("参数获取失败")
			return
		}

		if param < 10 {
			ctx.Abort()
			resp.ClientError("value小于10")
			return
		}
		// 通过请求
		ctx.Next()
	}
}
```

**使用**

`/app/router/api.go`

```go
	testApiRouter := apiRouter.Group("/test", middleware.ParamFilter)
```

所有的`/api/test`请求都会经过该中间件

### 冷知识2

查看`docker`容器的日志

```
docker logs container_name_or_id
```

查看环境变量

```
echo $VARIABLE_NAME
或
echo ｜ grep VARIABLE_NAME
查看全部
env
```

### 11.创建`docker`镜像

使用`golang:1.21-alpine`镜像对程序进行打包

然后copy到`alpine:latest`镜像中

`Dockerfile`

```dockerfile
FROM golang:1.21-alpine as BUILD

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/linux_amd64/main


FROM alpine:latest
# 复制资源
COPY --from=BUILD /app/config /config
COPY --from=BUILD /app/bin/linux_amd64/main app/main

# Env设置 - App
ENV APP_USE_ENV="true"
ENV APP_PORT=8080
ENV APP_DEBUG="false"
ENV APP_FRONTEND_URL="http://www.example.com"
ENV APP_BACKEND_URL="http://api.example.com"

# 暴露端口
EXPOSE 8080

ENTRYPOINT ["app/main"]
```

使用`docker-compose.yml`运行容器

`docker-compose.yml`

```yaml
version: "3.1"

services:
  backend:
    platform: linux/amd64
    image: yinsiyu/gin-project
    container_name: gin-project
    restart: always
    ports:
      - "8849:8080"
```

将命令集合到`Makefile`中

`Makefile`

```makefile
build-linux-amd64:
	docker build --platform=linux/amd64 -t yinsiyu/gin-project .

launch-app:
	docker-compose up -d

image-push:
	docker push yinsiyu/gin-project
```

### 12.`open-api`

使用`api.yaml`规范`API`

使用`docker`镜像 ——`openapitools/openapi-generator-cli:Tag`

`docker-compose.api.yml`

```yml
version: '3'
services:
  openapi-generator-cli:
    image: openapitools/openapi-generator-cli:v7.2.0
    command:
      [
        'generate',
        '-i',
        './tmp/src/openapi.v3.yaml',
        '-o',
        'tmp/dist',
        '-g',
        'go-gin-server',
        '--additional-properties=packageName=openapi,withGoCodegenComment=true,apiPath=openapi,enumClassPrefix=true',
      ]
    volumes:
      - ./gin-api:/tmp/src
      - ./app:/tmp/dist

```

创建`openapi.v3.yaml`文件，通过`openapi-generator`容器生成所需的结构体，路由等等

创建`/app/gin-api/openapi.v3.yaml`

使用软件——`APICurito`编辑yaml文件

创建`.openapi-generator-ignore`忽略`openapi-generator`对其他文件的修改

`/app/.openapi-generator-ignore`

```
api/*
Dockerfile
go.mod
main.go
/openapi/*.md
```

### 13.整理`Makefile`

重新命名文件`docker-compose.yml`

`docker-compose.yml`——>`docker-compose.app.yml`

`Makefile`

```makefile
generate-api:
	docker-compose -f docker-compose.api.yml run --rm openapi-generator-cli

build-linux-amd64:
	docker build --platform=linux/amd64 -t yinsiyu/gin-project .

launch-app:
	docker-compose -f docker-compose.app.yml up -d

image-push:
	docker push yinsiyu/gin-project
```

### 14.配置并使用`mysql`驱动

#### 安装所需的库

安装`gorm`

```
go get -u gorm.io/gorm
```

安装`mysql`驱动

```
go get -u gorm.io/driver/mysql
```

#### 配置连接

- 创建`db.yml`

  在`/config`目录下创建`db.yml`文件（**有小坑**）

  ```yaml
  Username: root
  Password: ysy123
  Connection: tcp(www.yugod.top:3306)
  DatabaseName: yugod_db
  TablePrefix: ysy_
  ```

- 初始化配置信息结构体

  在`/app/config`目录下创建`db.go`

  注意：没有使用环境变量，可以配置一下。

  ```go
  type DBConfig struct {
  	Username     string `yaml:"Username"`
  	Password     string `yaml:"Password"`
  	Connection   string `yaml:"Connection"`
  	DatabaseName string `yaml:"DatabaseName"`
  }
  
  var DB DBConfig
  
  func (d DBConfig) GetDSN() string {
  	return fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True", d.Username, d.Password, d.Connection, d.DatabaseName)
  }
  
  func init() {
  	// Read config from YAML
  	var settings DBConfig
  
  	config, err := os.ReadFile("config/db.yml")
  
  	if err != nil {
  		log.Fatal("DB config not set.")
  	}
  	yamlErr := yaml.Unmarshal(config, &settings)
  	if yamlErr != nil {
  		log.Fatal("DB config read error.")
  	}
  
  	DB = settings
  }
  
  ```

- 连接数据库的方法

  在路径`app/db`下创建`db.go`

  通过上面的配置信息和方法在`InitDB`方法中尝试连接

  创建`app/dao/base.go`

  ```go
  var (
  	// DB reference to database
  	DB *gorm.DB
  )
  ```

  `app/db/db.go`

  ```go
  var DB *gorm.DB
  
  func InitDB() {
  	gormConfig := &gorm.Config{
  		// 表名配置：添加前缀
  		NamingStrategy: schema.NamingStrategy{
  			TablePrefix:   "ysy_",
  			SingularTable: true,
  		},
  	}
  	// mysql Debug
  	if config.App.Debug {
  		gormLogger := logger.New(
  			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
  			logger.Config{
  				SlowThreshold: time.Second, // Slow SQL threshold
  				LogLevel:      logger.Info, // Log level
  				Colorful:      false,       // Disable color
  			},
  		)
  		gormConfig.Logger = gormLogger
  	}
  
  	db, err := gorm.Open(mysql.Open(config.DB.GetDSN()), gormConfig)
  	if err != nil {
  		log.Fatalf("Got error when connect database, the error is '%v'", err)
  	}
  
  	// 传入dao层
  	dao.DB = db
  	err = db.AutoMigrate(&model.ClickVolume{})
  
  	if err != nil {
  		log.Fatalf("Got error when migrate database, the error is '%v'", err)
  	}
  }
  ```

- 启动连接

  创建`app/boot/db.go`

  ```go
  // InitDB 创建数据表
  func InitDB() {
  	db.InitDB()
  }
  ```

  在`main.go`中启动连接

  ```go
  boot.InitDB()
  ```

#### 定义模型

在使用 GORM 进行数据库操作时，通常需要定义模型（Model）来表示数据库中的表结构。当你修改了模型的定义，例如添加了新的字段或者修改了字段的类型，你需要手动执行相应的数据库迁移操作来使数据库结构与模型定义保持一致。

`AutoMigrate`方法

`AutoMigrate` 方法的作用就是自动检测模型的变化，并将数据库结构更新为最新定义。它会创建缺失的表、添加缺失的字段，或者修改现有字段的类型等。

注意：

`AutoMigrate` 方法一般不会直接导致已有数据的丢失，但它会修改数据库表的结构，可能会导致数据的丢失或不完整，具体取决于你的模型定义以及数据库中已有的数据。

在执行 `AutoMigrate` 方法时，GORM 会尝试按照你的模型定义来更新数据库表的结构。它会添加缺失的表、添加缺失的字段，或者修改字段的类型等。如果你的模型定义与数据库中的表结构不一致，那么 `AutoMigrate` 将会尝试使它们保持一致。

但是，如果模型定义的变化导致了数据丢失或不完整的情况，例如你删除了一个字段，那么相关的数据可能会丢失。同样，如果你修改了字段的类型，可能会导致数据丢失或截断。

在`/model`目录下创建模型

例如：

app/model/click-volume.go

```go
type ClickVolume struct {
	gorm.Model
	Name  string `gorm:"not null;type:varchar(255);comment:统计名称"`
	Count uint   `gorm:"not null;type:bigint;comment:点击次数"`
}
```

用`AutoMigrate`定义模型

```go
err = db.AutoMigrate(&model.ClickVolume{})
if err != nil {
  	// 迁移错误
		log.Fatalf("Got error when migrate database, the error is '%v'", err)
	}
```

用于定义模型的结构图最好继承`gorm.Model`(如上)

继承后的表会增加字段：

```go
type Model struct {
  	// 主键
    ID        uint `gorm:"primarykey"`
  	// 创建时间
    CreatedAt time.Time
  	// 更新时间
    UpdatedAt time.Time
  	// 删除时间
    DeletedAt DeletedAt `gorm:"index"`
}
```

如果通过gorm框架删除数据，在数据库中数据不会消失，只是改变`DeletedAt`字段的值

防止数据被误删

#### 补坑

连接数据库时

通常连接地址应该包含协议

www.yugod.top:3306❌

tcp（www.yugod.top:3306）✅

最后拼接的`dsn`类似于

```
dsn := "user:password@tcp(xxx.xxx.xxx.xxx:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
```

