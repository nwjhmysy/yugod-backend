### 1.本地`mac`连接`ECS`云服务器

命令一：`sudo -i`

目的：使用`root`用户权限

命令二：`ssh -i <.pem密钥文件> <用户@云服务器IP地址>` 

目的：连接云服务器

例如：

```shell
ssh -i /Users/yinsiyu/Desktop/aliyun_ysy.pem root@47.108.217.27
```

### 2.创建文件夹

```shell
mkdir folder_name # folder_name是文件夹名字
```

以下是一些常用的`mkdir`命令选项：

- `-p`：递归地创建文件夹。如果父文件夹不存在，该选项将同时创建父文件夹。
- `-m`：设置文件夹的权限模式。可以使用八进制数值或符号模式。
- `-v`：显示详细的输出，以便查看创建过程。

### 3.删除文件/文件夹

```shell
#删除文件
rm example.txt

#删除文件夹
rm -r /path/to/directory
```

`-r` 选项表示递归删除，将连同目录中的所有文件和子目录一起删除。如果目录中有写保护的文件或子目录，系统将提示你进行确认。

请注意，删除目录及其内容是一个不可逆操作，请谨慎操作，确保你真正想要删除的是正确的目录。

如果一次性删除文件夹及里面的所有文件时：

```shell
rm -rf /path/to/directory
```

### 4.查看所有的运行端口

```shell
sudo netstat -tuln
```

这个命令使用`netstat`工具，并结合一些选项来列出所有处于监听状态的端口。

解释一下命令中的选项：

- `-t`：显示TCP协议相关的端口。
- `-u`：显示UDP协议相关的端口。
- `-l`：仅显示监听状态的端口。
- `-n`：以数字形式显示端口号，而不是解析为服务名。

运行该命令后，会显示一个端口列表，其中包括本地地址、远程地址、协议和状态等信息。你可以查找对应的端口号来确定正在运行的服务。

### 5.查看有一端口

请注意，使用`netstat`命令可能需要管理员权限（通过`sudo`命令），并且结果可能会包含大量信息。你可以使用其他的过滤工具（如`grep`）来进一步筛选和查找特定的端口或服务。例如，使用`grep`过滤特定的端口号：

```shell
sudo netstat -tuln | grep <port_number>
```

将`<port_number>`替换为你要查找的具体端口号。这将过滤显示仅包含指定端口号的行。

注意，根据系统配置和安全策略，某些端口可能会被隐藏或限制访问，因此可能不会在列表中显示出来。

### 6.查看占用某一端口的进程的`PID`

`lsof` 可以列出当前系统打开的文件和进程相关的信息，包括网络连接。

以下是使用 `lsof` 命令查看占用特定端口的进程名称的示例：

```shell
sudo lsof -i :<port_number>
```

### 7.根据进程的`PID`结束该进程

```shell
sudo kill <PID>
```

### 8.服务

开启某一服务

```shell
sudo systemctl start <service_name>
```

查看某一服务状态

```shell
sudo systemctl status <service_name>
```

关闭某一服务

```shell
sudo systemctl stop <service_name>
```

