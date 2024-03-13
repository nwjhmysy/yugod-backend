### 1.创建应用

创建

```
npm create vue@latest
```

按照流程进行安装

### 2.安装`tailwind`

#### `Install Tailwind CSS`

安装并初始化配置文件

```
npm install -D tailwindcss postcss autoprefixer
npx tailwindcss init
```

#### `PostCSS configuration`

`postcss.config.js`

```js
module.exports = {
  plugins: {
    tailwindcss: {},
    autoprefixer: {},
  }
}
```

#### `Configure your template paths`

`tailwind.config.js`

```js
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{html,js,vue}"],
  theme: {
    extend: {},
  },
  plugins: [],
}
```

#### `Add the Tailwind to your CSS`

`style.css`

```
@tailwind base;
@tailwind components;
@tailwind utilities;
```

### 3.安装`prepare`

格式化代码工具

安装

```
npm i prettier
```

格式化规则

`.prettierrc`

```
{
  "semi": true,
  "singleQuote": true,
  "tabWidth": 2,
  "printWidth": 120
}
```

忽略规则

`.prettierignore`

```
// 例如
src/services/*
src/stories/*
```

命令

```
"format": "prettier -w src/*.{ts,vue} && prettier -w src/**/*.{ts,vue}",
```

### 4.使用`docker`

使用`docker`镜像——`node:18.18.2-alpine`打包项目

再将项目`copy`到`nginx:alpine`镜像

`Dockerfile`

```dockerfile
# 基于 Node 镜像构建 Vite 项目
FROM node:18.18.2-alpine AS builder

WORKDIR /app

COPY package.json .
COPY package-lock.json .

RUN npm config set registry https://registry.npmjs.org
RUN npm install

COPY . .
RUN npm run build

# 构建 Nginx 服务器并拷贝 Vite 项目的构建文件
FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf


EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
```

`docker-compose.yml`

```yml
version: '3'

services:
  nginx:
    image: yinsiyu/frontend-vue3
    container_name: myvue3
    ports:
      - "8080:80"
    restart: always
```

`Makefile`

```makefile
build-dev-image:
	docker build --platform=linux/amd64 -t yinsiyu/frontend-vue3 .

docker-run:
	docker-compose up -d

docker-down:
	docker-compose down

docker-push:
	docker push yinsiyu/frontend-vue3:latest
```

`nginx.conf`

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

### 5.静态资源实施思路（双语言版）

#### 整体思路

每个页面都可以获取语言类型的索引

改变语言种类时路由改变

改变路由中的语言类型参数时页面获取的索引也改变

根据语言类型索引拿到所需的静态资源

#### 静态资源结构

```
model
	- XXX
		- index.ts
		- XXX_zh.ts
		- XXX.ja.ts
		- type.ts
```

`index.ts`

```tsx
const test:Record<LOCALES,xxx> = {
  zh: xxx_zh,
  ja: xxx_ja,
}
```

通过语言索引获取对应的静态资源

#### 语言切换（SPA）

使用了全局属性

```tsx
import { onMounted, ref, watch } from 'vue'
import { defineStore } from 'pinia'
import { LOCALES } from '@/constants'
import { browserLocale } from '@/hooks/useLocale'
import { useRoute, useRouter } from 'vue-router'

export const useLangStore = defineStore('lang', () => {
  const router = useRouter()
  const route = useRoute()
  const lang = ref<LOCALES>(browserLocale())

  onMounted(() => {
    switch (route.params.lang) {
      case 'zh':
        lang.value = LOCALES.ZH
        break
      case 'ja':
        lang.value = LOCALES.JA
        break
      default:
        lang.value = browserLocale()
        break
    }
  })

  // 改变语言 lang 时,改变路由
  watch(lang, (val) => {
    router.push({ params: { lang: val } })
  })

  const setLangJa = () => {
    lang.value = LOCALES.JA
  }

  const setLangZh = () => {
    lang.value = LOCALES.ZH
  }
	// 切换路由时使用 moveTo 方法
  const moveTo = (path?: string) => {
    if (path === route.path) return
    router.push({ path: '/' + lang.value + path })
  }

  return { lang, setLangJa, setLangZh, moveTo }
})

```

#### 在页面中的使用（例）

```tsx
const {lang} = storeToRefs(useLangStore());

const value = computed(() => {
  const key = lang?.value || LOCALES.ZH;
  return TEST_VALUE[key];
});
```

#### 切换`meta`

创建`meta`静态资源

在`/router/index.ts`中获取`meta`资源

`getMeta(route)`

```tsx
const getMeta = (route: RouteLocationNormalizedLoaded) => {
  const lang = route.params.lang as LOCALES || LOCALES.ZH;
  const meta = META_VALUE[lang];
  const metaKey = route.name ? route.name?.toString() : 'common';

  return meta[metaKey] || meta['common'];
};
```

在`router.afterEach()`方法中修改标签

使用`nextTick()`方法

`nextTick()`: DOM 更新循环结束后执行回调函数

```tsx
router.afterEach((to) => {
  // nextTick(): DOM 更新循环结束后执行回调函数
  nextTick(() => {
    // 为 HTML 添加 meta
    const meta = getMeta(to);

    // 添加 tittle：
    document.title = meta.title;
    document.querySelector('meta[name=description]')?.remove();

    // 添加 meta：
    // 查找已存在的 meta[name=viewport] 元素
    const viewportMeta = document.querySelector('meta[name="viewport"]');
    // 创建新的 meta 元素
    const descriptionMeta = document.createElement('meta');
    descriptionMeta.setAttribute('name', 'description');
    descriptionMeta.setAttribute('content', meta.description);
    // 插入到 viewportMeta 元素之后
    viewportMeta?.parentNode?.insertBefore(descriptionMeta, viewportMeta.nextSibling);
  });
});
```

### 6.获取环境变量

创建`.env`文件

在代码中获取

例如：

```tsx
import.meta.env.VITE_API_BATH_PATH
```

注意：

在Vite中，环境变量有两种类型：以 `VITE_` 开头的变量会被暴露到客户端代码中，而其他的则不会。

### 7.无法触发浏览器下载的问题

后端接口

功能：根据前端传来的文件路径获取本地的文件，然后将文件响应给前端

代码：

```go
func (md *mdAPI) DownloadMDByCode(ctx *gin.Context) {
	resp := response.Gin{Ctx: ctx}
	param := DownloadMdAPIParam{}

	// 获取参数 md_path 和 download_code
	if err := ctx.ShouldBindQuery(&param); err != nil {
		resp.ClientError("参数获取失败！")
		return
	}
	// 判断 code 码
	if param.DownloadCode != config.App.DownloadCode {
		resp.ClientError("下载码错误！")
		return
	}

	// 打开本地的 Markdown 文件
	filePath := "mds/" + param.MdPath + ".md"
	_, err := os.Stat(filePath)
	if err != nil {
		resp.ClientError("没找到文件！")
		return
	}

	// 设置响应头，告诉浏览器这是一个文件下载
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+strings.Replace(filePath, "/", "_", -1))
	ctx.Header("Expires", "0")
	ctx.Header("Cache-Control", "must-revalidate")
	ctx.Header("Pragma", "public")
	ctx.File(filePath)
}
```

问题：前端请求该接口并不能触发浏览器下载行为。

原因：

使用了`axiox`进行请求，虽然已经设置了响应头来指示浏览器下载文件，但有时浏览器可能会忽略这些头部信息，导致文件不会自动下载。这可能是由于浏览器的安全策略或用户的浏览器设置所致。

解决：

在页面设置一个不可见的`a`标签链接指向下载接口的URL。

如果依然需要请求的状态等，可以使用axios请求。

等于是请求了两次，一次是获取请求状态等参数，一次是触发浏览器的下载。

### 8.状态管理：`pinia`

设置状态和`active`

例如：

设置-

```tsx
import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useCounterStore = defineStore('counter', () => {
  const count = ref(0)
  const doubleCount = computed(() => count.value * 2)
  function increment() {
    count.value++
  }

  return { count, doubleCount, increment }
})
```

使用-

```tsx
const { count } = storeToRefs(useCounterStore())
const { doubleCount, increment } = useCounterStore()
```
