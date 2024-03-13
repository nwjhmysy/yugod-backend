## 1.`Http`协议详解

### 请求报文

```
结构：---------------------------
请求行----->方法+空格+url+http版本
请求头部--->首部：域值（...多对）
空行
请求体--->请求携带的参数（浏览器发送的GET请求一般没有请求体）
```

- 请求行

  ```
  请求行由请求方法字段、URL字段和HTTP协议版本字段3个字段组成，它们用空格分隔。比如 GET /data/info.html HTTP/1.1
  
  方法字段就是HTTP使用的请求方法，比如常见的GET/POST
  
  其中HTTP协议版本有两种：HTTP1.0/HTTP1.1 可以这样区别：
  
  HTTP1.0对于每个连接都只能传送一个请求和响应，请求就会关闭，HTTP1.0没有Host字段;
  而HTTP1.1在同一个连接中可以传送多个请求和响应，多个请求可以重叠和同时进行，HTTP1.1必须有Host字段。
  ```

- 请求头

  ```
  HTTP客户程序(例如浏览器)，向服务器发送请求的时候必须指明请求类型(一般是GET或者 POST)。如有必要，客户程序还可以选择发送其他的请求头。大多数请求头并不是必需的，但Content-Length除外。对于POST请求来说 Content-Length必须出现。
  ```

  常见的请求头字段含义：

  **Accept**： 浏览器可接受的MIME类型。

  **Accept-Charset**：浏览器可接受的字符集。

  **Accept-Encoding**：浏览器能够进行解码的数据编码方式，比如gzip。Servlet能够向支持gzip的浏览器返回经gzip编码的HTML页面。许多情形下这可以减少5到10倍的下载时间。

  **Accept-Language**：浏览器所希望的语言种类，当服务器能够提供一种以上的语言版本时要用到。

  **Authorization**：授权信息，通常出现在对服务器发送的WWW-Authenticate头的应答中。

  **Content-Length**：表示请求消息正文的长度。

  **Host**： 客户机通过这个头告诉服务器，想访问的主机名。Host头域指定请求资源的Intenet主机和端口号，必须表示请求url的原始服务器或网关的位置。HTTP/1.1请求必须包含主机头域，否则系统会以400状态码返回。

  **If-Modified-Since**：客户机通过这个头告诉服务器，资源的缓存时间。只有当所请求的内容在指定的时间后又经过修改才返回它，否则返回304“NotModified”应答。

  **Referer**：客户机通过这个头告诉服务器，它是从哪个资源来访问服务器的(防盗链)。包含一个URL，用户从该URL代表的页面出发访问当前请求的页面。

  **User-Agent**：User-Agent头域的内容包含发出请求的用户信息。浏览器类型，如果Servlet返回的内容与浏览器类型有关则该值非常有用。

  **Cookie**：客户机通过这个头可以向服务器带数据，这是最重要的请求头信息之一。

  **Pragma**：指定“no-cache”值表示服务器必须返回一个刷新后的文档，即使它是代理服务器而且已经有了页面的本地拷贝。

  **From**：请求发送者的email地址，由一些特殊的Web客户程序使用，浏览器不会用到它。

  **Connection**：处理完这次请求后是否断开连接还是继续保持连接。如果Servlet看到这里的值为“Keep-Alive”，或者看到请求使用的是HTTP 1.1(HTTP1.1默认进行持久连接)，它就可以利用持久连接的优点，当页面包含多个元素时(例如Applet，图片)，显著地减少下载所需要的时间。要实现这一点，Servlet需要在应答中发送一个Content-Length头，最简单的实现方法是：先把内容写入ByteArrayOutputStream，然后在正式写出内容之前计算它的大小。

  **Range**：Range头域可以请求实体的一个或者多个子范围。例如，

  表示头500个字节：bytes=0-499

  表示第二个500字节：bytes=500-999

  表示最后500个字节：bytes=-500

  表示500字节以后的范围：bytes=500-

  第一个和最后一个字节：bytes=0-0,-1

  同时指定几个范围：bytes=500-600,601-999

  但是服务器可以忽略此请求头，如果无条件GET包含Range请求头，响应会以状态码206(PartialContent)返回而不是以200 (OK)。

- 空行

  它的作用是通过一个空行，告诉服务器请求头部到此为止。

- 请求数据

  若方法字段是GET，则此项为空，没有数据

  若方法字段是POST,则通常来说此处放置的就是要提交的数据

### 响应报文

```
结构：---------------------------
状态行----->版本+空格+状态码+原因短语
响应头部--->首部：域值（...多对）
空行
响应体--->请求携带的参数（浏览器发送的GET请求一般没有请求体）
```

- 响应行

  ```
  响应行一般由协议版本、状态码及其描述组成 比如 HTTP/1.1 200 OK
  
  其中协议版本HTTP/1.1或者HTTP/1.0，200就是它的状态码，OK则为它的描述。
  ```

  常见状态码

  | 状态码   | 表示含义                                                     |
  | -------- | ------------------------------------------------------------ |
  | 100～199 | 表示成功接收请求，要求客户端继续提交下一次请求才能完成整个处理过程。 |
  | 200～299 | 表示成功接收请求并已完成整个处理过程。常用200                |
  | 300～399 | 为完成请求，客户需进一步细化请求。例如：请求的资源已经移动一个新地址、常用302(意味着你请求我，我让你去找别人),307和304(我不给你这个资源，自己拿缓存) |
  | 400～499 | 客户端的请求有错误，常用404(意味着你请求的资源在web服务器中没有)403(服务器拒绝访问，权限不够) |
  | 500～599 | 服务器端出现错误，常用500                                    |

- 响应头

  ```
  响应头用于描述服务器的基本信息，以及数据的描述，服务器通过这些数据的描述信息，可以通知客户端如何处理等一会儿它回送的数据。
  
  设置HTTP响应头往往和状态码结合起来。
  例如，有好几个表示“文档位置已经改变”的状态代码都伴随着一个Location头，而401(Unauthorized)状态代码则必须伴随一个WWW-Authenticate头。
  然而，即使在没有设置特殊含义的状态代码时，指定应答头也是很有用的。
  应答头可以用来完成：设置Cookie，指定修改日期，指示浏览器按照指定的间隔刷新页面，声明文档的长度以便利用持久HTTP连接，……等等许多其他任务。
  ```

  常见的响应头字段含义：

  **Allow**：服务器支持哪些请求方法(如GET、POST等)。

  **Content-Encoding**：文档的编码(Encode)方法。只有在解码之后才可以得到Content-Type头指定的内容类型。利用gzip压缩文档能够显著地减少HTML文档的下载时间。Java的GZIPOutputStream可以很方便地进行gzip压缩，但只有Unix上的Netscape和Windows上的IE4、IE5才支持它。因此，Servlet应该通过查看Accept-Encoding头(即request.getHeader(“Accept- Encoding”))检查浏览器是否支持gzip，为支持gzip的浏览器返回经gzip压缩的HTML页面，为其他浏览器返回普通页面。

  **Content-Length**：表示内容长度。只有当浏览器使用持久HTTP连接时才需要这个数据。如果你想要利用持久连接的优势，可以把输出文档写入

  ByteArrayOutputStram，完成后查看其大小，然后把该值放入Content-Length头，最后通过byteArrayStream.writeTo(response.getOutputStream()发送内容。

  **Content-Type**：表示后面的文档属于什么MIME类型。Servlet默认为text/plain，但通常需要显式地指定为text/html。由于经常要设置

  Content-Type，因此HttpServletResponse提供了一个专用的方法setContentType。

  **Date**：当前的GMT时间，例如，Date:Mon,31Dec200104:25:57GMT。Date描述的时间表示世界标准时，换算成本地时间，需要知道用户所在的时区。你可以用setDateHeader来设置这个头以避免转换时间格式的麻烦。

  **Expires**：告诉浏览器把回送的资源缓存多长时间，-1或0则是不缓存。

  **Last-Modified**：文档的最后改动时间。客户可以通过If-Modified-Since请求头提供一个日期，该请求将被视为一个条件GET，只有改动时间迟于指定时间的文档才会返回，否则返回一个304(NotModified)状态。Last-Modified也可用setDateHeader方法来设置。

  **Location**：这个头配合302状态码使用，用于重定向接收者到一个新URI地址。表示客户应当到哪里去提取文档。Location通常不是直接设置的，而是通过HttpServletResponse的sendRedirect方法，该方法同时设置状态代码为302。

  **Refresh**：告诉浏览器隔多久刷新一次，以秒计。

  **Server**：服务器通过这个头告诉浏览器服务器的类型。Server响应头包含处理请求的原始服务器的软件信息。此域能包含多个产品标识和注释，产品标识一般按照重要性排序。Servlet一般不设置这个值，而是由Web服务器自己设置。

  **Set-Cookie**：设置和页面关联的Cookie。Servlet不应使用response.setHeader(“Set-Cookie”, …)，而是应使用HttpServletResponse提供的专用方法addCookie。

  **Transfer-Encoding**：告诉浏览器数据的传送格式。

  **WWW-Authenticate**：客户应该在Authorization头中提供什么类型的授权信息?在包含401(Unauthorized)状态行的应答中这个头是必需的。例如，response.setHeader(“WWW-Authenticate”,“BASICrealm=\”executives\”“)。注意Servlet一般不进行这方面的处理，而是让Web服务器的专门机制来控制受密码保护页面的访问。

  ```
  注：设置应答头最常用的方法是HttpServletResponse的setHeader，该方法有两个参数，分别表示应答头的名字和值。和设置状态代码相似，设置应答头应该在发送任何文档内容之前进行。
  
  setDateHeader方法和setIntHeadr方法专门用来设置包含日期和整数值的应答头，前者避免了把Java时间转换为GMT时间字符串的麻烦，后者则避免了把整数转换为字符串的麻烦。
  
  HttpServletResponse还提供了许多设置
  
  setContentType：设置Content-Type头。大多数Servlet都要用到这个方法。
  
  setContentLength：设置Content-Length头。对于支持持久HTTP连接的浏览器来说，这个函数是很有用的。
  
  addCookie：设置一个Cookie(Servlet API中没有setCookie方法，因为应答往往包含多个Set-Cookie头)。
  ```

- 空行

  代表响应头结束

- 响应体

  响应体就是响应的消息体，如果是纯数据就是返回纯数据，如果请求的是HTML页面，那么返回的就是HTML代码，如果是JS就是JS代码，如此之类。

## 2.常见`Content-Type`类型

`Content-type`定义了`http`请求的数据类型

如果设置在请求头中，则定义的是请求体的数据类型；

如果设置在响应头中，则定义的是响应体的数据类型；

- 请求头
  1. application/json：JSON 数据格式；
  2. application/x-www-form-urlencoded：表单默认的提数据格式；
  3. multipart/form-data：一般用于文件上传；
- 响应头
  1. text开头
     text/html： HTML格式
     text/plain：纯文本格式
     text/xml： XML格式
  2. 图片格式
     image/gif ：gif 图片格式
     image/jpeg ：jpg 图片格式
     image/png：png 图片格式
  3. application开头
     application/xhtml+xml：XHTML 格式
     application/xml：XML 数据格式
     application/atom+xml：Atom XML 聚合格式
     application/json：JSON 数据格式
     application/pdf：pdf 格式
     application/msword：Word 文档格式
     application/octet-stream：二进制流数据（如常见的文件下载）
     application/x-www-form-urlencoded：表单发送默认格式
  4. 媒体文件
     audio/x-wav：wav文件
     audio/x-ms-wma：w文件
     audio/mp3：mp3文件
     video/x-ms-wmv：wmv文件
     video/mpeg4：mp4文件
     video/avi：avi文件