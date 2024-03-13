### 1.`openapitools/openapi-generator-cli`

1. 创建`docker-compose.yml` `generate`命令中的各参数官网更详细

   ```yaml
   # docker-compose.yml
   version: '3'
   services:
     openapi-generator-cli:
       image: openapitools/openapi-generator-cli:v7.0.1
       command:
         [
           'generate',
           '-i',
           './tmp/src/openapi.v3.yaml',# 卷中yaml文件映射到容器中的位置
           '-o',
           'tmp/dist',# 在容器中生成代码的输出路径
           '-g',
           'typescript-axios',# 生成的前端代码类型
           '--additional-properties=withSeparateModelsAndApi=true',
           '--model-package=model',
           '--api-package=api',
           '--type-mappings=Date=string',
         ]
       volumes:
         - ./submodules/performance-spec:/tmp/src # 定义卷
         - ./src/services:/tmp/dist # 定义卷
   ```

   用于定义和运行容器的配置，command中的是启动openapitools容器后的配置。

2. 准备API文档的yaml文件

   openAPI的yaml文件可以通过`apicurito`编辑器生成，前后端公用一套。

3. 启动docker。

4. 通过`docker compose`启动容器

   ```sh
   docker-compose run --rm openapi-generator-cli
   ```

   `openapi-generator-cli`:在`docker-compose.yaml`中配置过了。

   `--rm`:在容器停止运行后自动删除容器。

### 2.`mysql:latest`

1. 创建`Dockerfile`文件

   ```yaml
   # 使用官方 MySQL 镜像作为基础镜像
   FROM mysql:latest
   
   # 设置 MySQL root 用户的密码
   ENV MYSQL_ROOT_PASSWORD=root123
   
   # 创建一个新的数据库
   ENV MYSQL_DATABASE=performance_mysql
   
   # 创建一个新用户，并设置用户密码
   ENV MYSQL_USER=yinsiyu
   ENV MYSQL_PASSWORD=ysy123
   
   # 将 MySQL 的配置文件复制到容器中
   # COPY my.cnf /etc/mysql/my.cnf
   
   # 暴露 MySQL 服务端口
   EXPOSE 3306
   
   ```

   创建新用户的目的：

   在 MySQL 中，通常建议不要使用 `root` 用户进行正常应用程序的连接和操作，这是因为 `root` 用户拥有很高的权限，可能导致潜在的安全风险。在生产环境中，最好为应用程序创建一个专用的、有限权限的用户，并使用该用户进行数据库操作。

   

2. 创建脚本文件`Makefile`

   在根目录创建脚本文件`Makefile`，用于存放并执行创建镜像和启动容器的命令

   ```makefile
   build-mysql-image:
   	docker build -t performance_mysql_image .
   
   run-performance-mysql:
   	docker-compose up -d
   ```

   

3. 创建容器

   执行命令：

   ```sh
   make build-mysql-image
   ```

   

4. 创建`mysql-data`文件夹

   在根目录创建`mysql-data`，用于将数据存放在本地。

5. 创建`docker-compose.yml`

   用于配置和启动容器。

   ```yaml
   version: "3"
   
   services:
     mysql:
       image: performance_mysql_image
       container_name: performance_mysql
       environment:
         MYSQL_ROOT_PASSWORD: root123
         MYSQL_DATABASE: performance_mysql
         MYSQL_USER: yinsiyu
         MYSQL_PASSWORD: ysy123
       ports:
         - "3306:3306"
       volumes:
         #   - ./my.cnf:/etc/mysql/my.cnf
         - ./mysql-data:/var/lib/mysql # 将mysql-data文件夹映射到容器的/var/lib/mysql路径下
   ```

6. 创建并启动容器

   ```sh
   make run-performance-mysql
   ```



### 3.使用`docker`启动前端项目（以vue3为例）

1. 创建项目

2. 创建`.dockerignore`

   在copy文件时忽略某些文件，如：node_modules

   ```
   node_modules
   ```

3. 创建`Dockerfile`文件

   `Dockerfile`文件的作用在下面有说明

   ```dockerfile
   # 基于 Node 镜像构建 Vite 项目
   FROM node:18.18.2-alpine AS builder
   
   WORKDIR /app
   
   COPY package.json .
   COPY package-lock.json .
   
   RUN npm config set registry https://registry.npmjs.org
   RUN npm install
   
   COPY . .
   RUN npm run build-only
   
   # 构建 Nginx 服务器并拷贝 Vite 项目的构建文件
   FROM nginx:alpine
   
   COPY --from=builder /app/dist /usr/share/nginx/html
   COPY nginx.conf /etc/nginx/nginx.conf
   
   
   EXPOSE 80
   
   CMD ["nginx", "-g", "daemon off;"]
   
   ```

   ```
   docker build --platform=linux/amd64 -t <hub_path/image_name> .
   ```

   使用`node:18.18.2-alpine`将项目压缩并打包，然后再拷贝到`Nginx`服务器中。

4. `nginx.conf`配置文件

   ```
   # nginx.conf
   
   user nginx;
   worker_processes auto;
   error_log /var/log/nginx/error.log;
   pid /var/run/nginx.pid;
   
   events {
       worker_connections 1024;
   }
   
   http {
       include /etc/nginx/mime.types;
       default_type application/octet-stream;
   
       server {
           listen 80 default_server;
           listen [::]:80 default_server;
   
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
   }
   ```

   

5. 启动容器

   ```
   version: '3'
   
   services:
     nginx:
       image: <hub_path/image_name>
       container_name: <container_name>
       ports:
         - "8080:80"
       restart: always
   ```

   ```
   docker-compose up -d
   ```

   

7. 注意

   运行项目时可能会遇到这样的问题：容器中的项目正常运行，但是宿主机无法访问项目。

   原因是：项目的服务器配置只允许本机访问，不对其他IP地址暴露。

   解决：修改`package.json`文件的`scripts`部分，添加`--host 0.0.0.0`：

   ```json
   "scripts": {
     "serve": "vue-cli-service serve --host 0.0.0.0"
   }
   // 本项目是用的vite，所以是
   "scripts": {
     "dev": "vite --host 0.0.0.0"
   }
   ```


### 4.如何保证go项目的docker容器和mysql的docker容器在同一个网络下运行并且两者可以连接？

要确保 Go 项目的 Docker 容器和 MySQL 的 Docker 容器在同一个网络下运行并能够相互连接，可以遵循以下步骤：

1. **创建 Docker 网络：** 使用以下命令创建一个自定义的 Docker 网络。这将使得两个容器可以使用服务名称进行通信。

   ```
   docker network create mynetwork
   ```

   这里的 `mynetwork` 是自定义网络的名称，你可以根据需要选择一个适当的名称。

2. **启动 MySQL 容器：** 启动 MySQL 容器，并将其连接到刚刚创建的网络。确保在启动 MySQL 容器时设置了正确的数据库用户、密码和其他配置。

   ```
   docker run -d --network=mynetwork --name=mysql-container -e MYSQL_ROOT_PASSWORD=root_password -e MYSQL_DATABASE=my_database -e MYSQL_USER=db_user -e MYSQL_PASSWORD=db_password mysql:latest
   ```

   这里，`--network=mynetwork` 选项将 MySQL 容器连接到刚刚创建的网络。

3. **启动 Go 项目容器：** 启动 Go 项目容器，并将其连接到相同的网络。

   ```
   docker run -d --network=mynetwork --name=go-container your-go-image
   ```

   这里，`--network=mynetwork` 选项将 Go 项目容器连接到相同的网络。

4. **在 Go 项目中使用 MySQL 服务名称：** 在 Go 项目中，确保使用 MySQL 容器的服务名称而不是具体的 IP 地址。服务名称即 MySQL 容器的名称，在上面的例子中是 `mysql-container`。在 Go 项目的数据库连接配置中，使用这个服务名称来连接 MySQL。

   ```
   db, err := sql.Open("mysql", "db_user:db_password@tcp(mysql-container:3306)/my_database")
   ```

   这样，Go 项目容器就可以通过服务名称 `mysql-container` 找到 MySQL 容器，并与之通信。

通过这些步骤，你可以确保 Go 项目的 Docker 容器和 MySQL 的 Docker 容器在同一个网络下运行，并能够通过服务名称进行连接。

### 5.`Docker`网络相关

要查看 Docker 中的网络，你可以使用 `docker network` 命令。以下是一些常用的命令：

1. **列出所有网络：**

   ```
   docker network ls
   ```

   这会列出系统中所有的 Docker 网络，包括默认网络和用户创建的网络。

2. **查看特定网络的详细信息：**

   ```
   docker network inspect <network-name>
   ```

   这个命令会提供有关特定网络的详细信息，包括网络的配置、连接到网络的容器等。

   例如：

   ```
   docker network inspect mynetwork
   ```

   这里的 `mynetwork` 是网络的名称，你可以替换成你感兴趣的网络名称。

3. **查看容器的网络信息：**

   如果你想查看特定容器连接到的网络，可以使用以下命令：

   ```
   docker network inspect <container-name>
   ```

   例如：

   ```
   docker network inspect my-container
   ```

   这会提供与容器相关联的网络信息，包括 IP 地址、网关等。

4. **显示特定网络上的容器列表：**

   ```
   docker network ls
   ```

   这会列出连接到指定网络的所有容器。

   例如：

   ```
   docker network ls
   ```

   这里的 `mynetwork` 是网络的名称。

这些命令可以帮助你了解 Docker 中的网络配置，并查看容器之间以及容器与主机之间的连接关系。

### 6.如果不使用Docker网络，两个容器之间可以相互访问吗？

在 Docker 中，即使不使用特定的用户创建网络，容器之间也可以通过它们的 IP 地址进行相互访问，前提是这些容器在同一个网络命名空间中。Docker 默认提供了一个默认的桥接网络，容器被连接到这个网络中，使得它们可以直接使用对方的 IP 地址进行通信。

例如，假设你有两个容器，它们分别是 `container1` 和 `container2`。即使你没有显式地将它们连接到自定义网络，它们依然可以通过各自的 IP 地址相互访问。

```
# 启动容器1
docker run -d --name container1 your-image1

# 启动容器2
docker run -d --name container2 your-image2
```

然后，你可以在 `container1` 中使用 `ping` 或其他网络工具访问 `container2`：

```
docker exec -it container1 ping <container2-ip-address>
```

或者在 `container2` 中访问 `container1`：

```
docker exec -it container2 ping <container1-ip-address>
```

然而，这种方法具有一些缺点，因为容器的 IP 地址可能会发生变化，而且直接使用 IP 地址不够灵活。因此，使用 Docker 的网络功能，特别是用户创建的自定义网络，是一种更好的方式，它提供了更灵活、可维护和可扩展的网络配置选项。

### 7.卷相关

要查看 Docker 中的卷，你可以使用 `docker volume` 命令。以下是一些常用的命令：

1. **列出所有卷：**

   ```
   docker volume ls
   ```

   这将列出系统中所有的 Docker 卷。

1. **删除某个卷**

   ```
   docker volume remove <卷名>
   ```

   

2. **查看特定卷的详细信息：**

   ```
   docker volume inspect <volume-name>
   ```

   这个命令提供有关特定卷的详细信息，包括卷的名称、驱动、挂载点等。

   例如：

   ```
   docker volume inspect my_volume
   ```

   这里的 `my_volume` 是卷的名称，你可以替换成你感兴趣的卷的名称。

3. **显示特定卷上的容器列表：**

   ```
   docker volume ls -qf "name=<volume-name>"
   ```

   这将列出连接到指定卷的所有容器。

   例如：

   ```
   docker volume ls -qf "name=my_volume"
   ```

   这里的 `my_volume` 是卷的名称。

4. **删除不再使用的卷：**

   ```
   docker volume prune
   ```

   这将删除所有没有与容器关联的卷。请谨慎使用，确保你不需要这些卷的数据。

这些命令可以帮助你查看和管理 Docker 中的卷。如果你发现卷没有按预期工作，你可以检查卷的详细信息以获取更多的信息。

使用 Docker 卷时，不需要提前手动创建卷。Docker 提供了命令和选项，它会在需要时自动创建卷。

当你在 `docker-compose.yml` 文件中使用卷时，例如：

```yml
version: '3'
services:
  myapp:
    image: your-image
    volumes:
      - my_volume:/container_path
```

在这个例子中，`my_volume` 是一个卷的名称，如果它不存在，Docker 会在启动服务时自动创建这个卷。

如果你在命令行中使用 `docker run`，也可以使用 `-v` 选项来指定卷。如果指定的卷不存在，Docker 也会在运行容器时创建它。例如：

```bash
docker run -v my_volume:/container_path your-image
```

在这里，`my_volume` 是卷的名称，`/container_path` 是容器内部的路径。如果 `my_volume` 不存在，Docker 会创建这个卷。

总之，Docker 会在需要时自动创建卷，你无需提前手动创建它们。这使得卷的管理更加方便，因为你可以在 Docker Compose 文件或 `docker run` 命令中定义卷，而无需事先手动操作。

### 8.三个基础命令

```
docker build --platform=linux/amd64 -t <image_name> .
```

```
docker-compose -f <docker-compose.yml_name> up -d
```

```
// 设置容器名称，端口映射，后台运行
docker run -d --name <container-name> -p <host-port>:<container-port> <image-name>
```

