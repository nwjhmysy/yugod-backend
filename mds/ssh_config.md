在GitHub上添加SSH密钥是一个常见的步骤，以便你可以通过SSH协议与GitHub仓库进行安全的通信。以下是在GitHub上添加SSH密钥的步骤：

### **步骤 1: 检查SSH密钥**

首先，你需要检查是否已经有SSH密钥。在终端或命令提示符中输入以下命令：

```
ls -al ~/.ssh
```

或者在Windows系统上，你可以使用以下命令：

```
dir %userprofile%\.ssh
```

如果已经存在SSH密钥，你会看到类似以下的文件：

- `id_rsa`（私钥）
- `id_rsa.pub`（公钥）

如果不存在，或者你想要添加新的密钥，可以继续下面的步骤。

### **步骤 2: 生成SSH密钥**

如果你没有SSH密钥，可以通过以下命令生成：

```
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
```

如果生成的密钥不在`.ssh`文件夹下，执行下面命令：

```
ssh-keygen -t rsa -b 4096 -C "your_email@example.com" -f ~/.ssh/ysy_public
```

在这个命令中，替换 `your_email@example.com` 为你在GitHub 上注册时使用的电子邮件地址。你也可以使用默认设置，只需按回车键。

### **步骤 3: 添加SSH密钥到SSH代理**

在一些操作系统上，SSH代理可能已经在运行，但在一些情况下，你可能需要手动启动。使用以下命令来启动SSH代理：

```
eval "$(ssh-agent -s)"
```

然后，将生成的私钥添加到SSH代理：

```
ssh-add ~/.ssh/id_rsa
```

### **步骤 4: 添加SSH密钥到GitHub**

现在，你需要将公钥添加到GitHub。使用以下命令来复制公钥到剪贴板：

```
pbcopy < ~/.ssh/id_rsa.pub   # macOS
clip < ~/.ssh/id_rsa.pub     # Windows
```

然后，打开你的GitHub帐户，转到 Settings > SSH and GPG keys，点击 "New SSH key" 或 "Add SSH key"，将公钥粘贴到 "Key" 文本框中，并给密钥一个描述性的标题。

完成后，你的SSH密钥就添加到了GitHub。现在，你可以使用SSH协议克隆或推送GitHub仓库而不需要输入用户名和密码。

### SSH大致原理：

每个账号都可以设置公钥，计算机访问账号时会从通过本地的`SSH`代理中找对对应的私钥访问账号。

公钥和私钥的生成是根据账号的邮箱。

生成公钥和私钥的命令：

生成

```
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
```

生成的钥保存在`~/.ssh`

查看

```
ls -al ~/.ssh
```

还要确保计算机开启了SSH服务：

查看开启与否

```
eval "$(ssh-agent -s)"
```

还要确保将私钥添加到代理中：

添加

```
ssh-add -K ~/.ssh/id_rsa
```

### 删除

1. **转到 SSH 文件夹：** 输入以下命令并按 Enter，以转到存储 SSH 密钥的文件夹：

   ```
   cd ~/.ssh
   ```

2. **列出当前文件夹中的文件：** 使用以下命令列出当前文件夹中的文件：

   ```
   ls
   ```

   查看所有

   ```
   ls -al
   ```

3. **找到并删除要删除的私钥文件：** 根据你的需要找到要删除的私钥文件，并使用 `rm` 命令删除它。例如，如果要删除名为 `id_rsa_example` 的私钥文件，可以

   如果你的私钥文件有关联的公钥文件（通常是 `.pub` 结尾），你也可以删除它，如下所示：

   ```
   rm id_rsa_example id_rsa_example.pub
   ```

4. **验证删除：** 使用 `ls` 命令确保文件已被删除：

   ```
   ls -al
   ```

   如果你不再看到被删除的私钥文件，那么删除成功。

### 测试链接

1. **验证SSH连接：**

   - 运行以下命令验证是否可以通过SSH与GitHub建立连接：

     ```
     ssh -T git@github.com
     ```

   - 如果您在执行此命令时遇到权限问题，请检查是否使用了正确的SSH密钥。

   如果链接成功，会出现以下信息，**并且还会发送邮件**

   ```
   Hi nwjhmysy! You've successfully authenticated, but GitHub does not provide shell access.
   ```

2. **查看SSH连接错误消息：**

   - 如果在推送时遇到问题，运行以下命令以获取更详细的错误消息：

     ```
     ssh -vT git@github.com
     ```

- 这会显示详细的SSH连接信息，可以帮助您诊断问题。
