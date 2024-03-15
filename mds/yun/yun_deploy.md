### 1.主机记录

云服务器上域名解析时的主机记录（Host Record）是指将域名映射到特定的 IP 地址或其他网络资源的记录。主机记录是域名系统（DNS）中的一部分，它告诉 DNS 查询请求如何定位和解析特定的主机或服务。

主机记录包括以下几种常见类型：

1. **A 记录（Address Record）：** 将域名映射到一个 IPv4 地址。这通常用于将域名指向一个具体的服务器。
2. **CNAME 记录（Canonical Name）：** 将域名指向另一个域名，实现别名的效果。主要用于将一个域名指向另一个域名，常用于负载均衡、网站托管等。
3. **MX 记录（Mail Exchange）：** 用于指定邮件服务器的地址，确保电子邮件能够正确地路由到目标邮件服务器。
4. **TXT 记录（Text Record）：** 用于存储文本信息，通常用于验证域名所有权或提供其他相关的信息。
5. **NS 记录（Name Server）：** 指定域名服务器的地址，告诉 DNS 查询请求去哪里查找该域名的解析信息。
6. **PTR 记录（Pointer Record）：** 用于反向 DNS 查询，将 IP 地址解析为相应的域名。

这些主机记录在域名解析过程中扮演关键角色。当用户在浏览器中输入一个域名时，系统会发起 DNS 查询，根据域名的主机记录找到相应的 IP 地址或其他资源，然后建立连接。主机记录因此对于确保域名能够正确映射到相应的服务器或服务非常重要。

### 2.域名解析

域名解析后会映射到一个`IP`地址。

下面是一般的步骤：

- **登录域名注册商控制台：** 使用你在注册商注册域名时创建的账户登录。
- **找到域名管理界面：** 在注册商的控制面板中找到与你想要设置的域名相关的管理界面。
- **添加主机记录：** 在 DNS 设置或域名解析页面，添加相应的主机记录。例如，如果你想将域名指向云服务器的 IP 地址，你需要添加一个 A 记录，将域名映射到该 IP 地址。
- **保存更改：** 完成设置后，保存更改。DNS 记录的生效时间可能需要一些时间，这是因为 DNS 记录需要在全球的 DNS 服务器中传播。
- **验证设置：** 确保域名解析设置生效后，可以通过浏览器访问域名来验证是否正确地指向了云服务器。

设置主机类型A

@ ：一般表示空值（即解析后的域名可以没有前缀）

### 3.通过域名访问主机

解析后的域名指向一个`IP`地址

例如，通过以下方式，你可以直接访问主机的某个端口：

```
http://example.com:8080
```

**访问`IP`时有如果后面不加端口号会默认访问`80`端口或`443`端口。**

- `80`

  HTTP 服务的默认端口

- `443`

  HTTPS 服务的默认端口

使用80端口，如果没有配置 SSL 证书，浏览器会认为是使用了 HTTP 服务，而不是 HTTPS。在常见的 Web 浏览器中，如果你访问一个使用80端口的网站，浏览器默认会将连接视为普通的 HTTP 连接，而不会尝试使用 SSL/TLS 进行加密通信。

HTTPS 使用的默认端口是443，而 HTTP 使用的默认端口是80。当浏览器访问一个网站时，如果URL没有明确指定协议（http:// 或 https://），浏览器会默认使用 HTTP，并且连接至网站的80端口。

因此，在没有配置 SSL 证书的情况下，通过80端口访问的网站会被认定为使用了 HTTP 服务。要使连接通过 HTTPS 进行加密通信，你需要在服务器上配置 SSL 证书，并确保浏览器可以建立安全的连接。

### **4.在 Nginx 服务器上：**

- 在 Nginx 配置文件中找到相关的 `server` 部分，添加 SSL 配置。
- 示例如下：

```
server {
    listen 443 ssl;
    server_name your_domain.com;

    ssl_certificate /path/to/your/certificate.crt;
    ssl_certificate_key /path/to/your/private.key;

    # 其他 SSL 配置...

    location / {
        # 其他 Nginx 配置...
    }
}
```

- 保存配置文件，并使用 `nginx -s reload` 命令重新加载 Nginx 配置。

请注意，SSH 证书和 SSL/TLS 证书用于不同的目的，因此在配置过程中需区分使用场景。 SSH 证书用于身份验证和远程登录，而 SSL/TLS 证书用于加密传输，特别是在 Web 服务器上提供 HTTPS 服务。

### 5.配置`SSL`证书

#### 1.在云服务器控制台申请SSL证书

可以申请一年的证书

#### 2.下载

申请成功后下载所需密钥和证书

**文件夹名称**：`xxx.xxx_nginx`

**文件夹内容**：

`xxx.xxx_bundle.crt` 证书文件（需要）

`xxx.xxx_bundle.pem` 证书文件

`xxx.xxx.key` 私钥文件（需要）

`xxx.xxx.csr` CSR 文件

**说明**

CSR 文件是申请证书时由您上传或系统在线生成的，提供给 CA 机构。**安装时可忽略该文件**。

#### 3.配置`nginx.conf`

```nginx
server {
        listen 80;
        server_name yugod.top;

        # Redirect all HTTP requests to HTTPS
        return 301 https://www.yugod.top;
    }

server {
        #SSL 默认访问端口号为 443
        listen 443 ssl; 
        #请填写绑定证书的域名
        server_name yugod.top; 
        #请填写证书文件的相对路径或绝对路径
        ssl_certificate /ssl/yugod.top_bundle.crt; 
        #请填写私钥文件的相对路径或绝对路径
        ssl_certificate_key /ssl/yugod.top.key; 
        ssl_session_timeout 5m;
        #请按照以下协议配置
        ssl_protocols TLSv1.2 TLSv1.3; 
        #请按照以下套件配置，配置加密套件，写法遵循 openssl 标准。
        ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:HIGH:!aNULL:!MD5:!RC4:!DHE; 
        ssl_prefer_server_ciphers on;

        root /usr/share/nginx/html;
        index index.html index.htm;

        location / {
            try_files $uri $uri/ /index.html;
        }

        error_page 500 502 503 504 /50x.html;
        location = /50x.html {
            root /usr/share/nginx/html;
        }
    }
```

##### 重定向`http`服务

在`80`端口的配置中

添加重定向配置

```
# Redirect all HTTP requests to HTTPS
return 301 https://www.yugod.top;
```



#### 4.将文件放在配置文件中对应的位置

根据上面的`nginx.conf`

将`xxx.xxx_bundle.crt` 和`xxx.xxx.key`放到：

云服务器中的`/ssl`目录中

或

docker容器中的`/ssl`目录中

#### 5.注意

在控制台——>'我的证书'中

将证书绑定域名

### 6.在云服务器启动`docker`容器

```
docker run --name <容器名字>  -d -p 80:80 -p 443:443 --restart=always <镜像>
```

### 7.会用到的一些`docker`命令

```
# 登陆
docker login

# 查看容器
docker ps

# 根据容器id关闭容器
docker stop <contain_id>

#启动容器
docker run --name <容器名字>  -d -p 80:80 -p 443:443 --restart=always <镜像>
```

### 8.`HTTPS`无法加载`HTTP`资源问题

问题描述：当HTTPS的页面尝试发出一个HTTP请求时，浏览器会默认阻止这种请求，因为这会导致所谓的“混合内容”（Mixed Content）问题，可能会导致安全漏洞。

出现场景：

前后端共同部署在一台云服务器上，前端开通了443端口。

解决：

利用`nginx`对前段发出的请求进行转发。

例如：

```nginx
server {
	# https://yugod.top
	# 443 服务
	
	location /api {
            proxy_pass http://yugod.top:8080/api;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
        }
}
```

以上面的为例：

在前端代码中，使用`https://yugod.top/api`作为`baseURL`请求。会被转发为`http://yugod.top:8080/api`

### 9.后端部署如何开启`https`服务

基本原理同上

使用nginx服务器，对443端口下的路由进行重定向（前提是安装了SSH证书）。



### 云端启动容器命令

后端

```
sudo docker run --name yugod-backend-v1.1 -d -p 8080:8080 --restart=always yinsiyu/yugod-backend:v1.1
```

前端

```
sudo docker run --name yugod-frontend-v1.1 -d -p 80:80 -p 443:443 --restart=always yinsiyu/yugod-frontend:v1.1
```

### 10.创建并启动`mysql`

将`docker-compose.db.yml`上传到云服务器

```yaml
version: '3.7'

services:
  mysql:
    image: mysql:8.0
    restart: always
    container_name: mysql-yugod
    environment:
      MYSQL_ROOT_PASSWORD: ysy123
      MYSQL_DATABASE: yugod_db
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql

volumes:
  mysql_data:

```

使用`docker-compose`构建mysql的容器并启动（有坑）

在控制台开启`3306`端口

| **应用类型** | **来源**  | **协议** | **端口** | **策略** | **备注**         |
| ------------ | --------- | -------- | -------- | -------- | ---------------- |
| MySQL (3306) | 0.0.0.0/0 | TCP      | 3306     | 允许     | MySQL服务 (3306) |

坑：

云服务器安装了`docker-compose`

但是权限不足

报错：

Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock

原因：

```
docker进程使用 Unix Socket 而不是 TCP 端口。而默认情况下，Unix socket 属于 root 用户，因此需要 root权限 才能访问。
```

解决：

查看 /var/run/docker.sock所在用户组，将用户重新加入docker组中

```
sudo groupadd docker #添加docker用户组

sudo gpasswd -a $XXX docker #检测当前用户是否已经在docker用户组中，其中XXX为用户名，例如我的，liangll

sudo gpasswd -a $USER docker #将当前用户添加至docker用户组

newgrp docker #更新docker用户组
```

