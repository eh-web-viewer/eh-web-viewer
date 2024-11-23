# 前端爆炸了

不想修，不想写vue。

人多的话用react写一个。

或者自己想写网页端的话可以直接引用镜像站的内容。已经做了跨域，都允许的。

----

# eh-front

一个前端练手项目。以适合移动端的操作方式看eh漫画。
## how to run 
[`back/main.go`](back/main.go)选择调试（或者运行exe）
front目录 `npm run dev`

## intro

突然发现ehentai不要配额了，于是就试试看写个前端，反正流量大头的hathnetwork又不走前端流量。
大概只是模仿记忆里面的ehviewer吧。

### 首页
翻页，瀑布流
显示漫画用的卡片。
喜欢的tag高亮


## API

return in

request path
request queries
request key params.
return type
payload

### 首页

search params

payload

gallarys: Arrays[]
next page
prev page


## notes for project

屎山

### TODO

what's next:

search:

- [] but how?

index:

- [] show tags and other info
- [] cover image should load the first /s/
- [] async tags.

gallery:

- [] show tags and other info
- [] 2 mode (waterfall or pages)
- [] 要不直接加载大图得了

ui/ux:

- [] 后端读一下配置。
- 

----

# 自用

## 怎么快速上传部署

后端：
```sh
./download.sh  back/back.bin  eh-web-viewer
```

前端

```sh 
cd /var/www/ehwv
nc -l 1324 -v| tar -xzvf - # server
```

```sh
tar -czvf - assets index.html vite.svg | /c/bin/Nmap/ncat.exe [localhost] 1324 # local
```

## 更新记录

# 今まで
EhwvGalleryInfo
是啥
EhwvImageBox.vue
是啥
都不记得了

不想改了。下次从view这边改。
# 08/30

终于想到可以做个跟新纪录以防忘记怎么做了。

现在是啥，好像是改版面的问题，现在好麻烦。
目标是：
- [ ] gallery页面能原地加载下一页
- [ ] index页面的bug修复。并不是跟着router走而是 按钮->过程->更新页面和url。
- [ ] image似乎有点毛病。放在之后说吧
- [ ] trival or misc
  - [ ] 标签可点击，样式
  - [ ] index的筛选

# ----

~~curl加上proxy。~~
~~去用fiber里的~~
没proxy。

怎么做前端来着。

1MDmyTSmGxgtdzXUSFu3WiZs1tSY7dMmse


在unmounted之后数据会掉
那么想个办法



```ts

myObject[Symbol.iterator] = function* () {
    const keys = Object.keys(this);
  
    for (const key of keys) {
      yield [key, this[key]];
    }
  };

```

我是不是存一个ref就行了。。

keep alive 不顶用所以存在另一个文件里面。



TODO

给api加上特定ip访问ex

界面

搜索要怎么做

预览的卡片


详细页面

store.ts中存储两个页面的meta和list

```js
const queryString = 'param1=value1&param2=value2&param3=value3';
const params = new URLSearchParams(queryString);

// Access individual query parameters
console.log(params.get('param1')); // "value1"
console.log(params.get('param2')); // "value2"
console.log(params.get('param3')); // "value3"

// Get all query parameters as an object
const paramObject = {};
for (const [key, value] of params) {
  paramObject[key] = value;
}
console.log(paramObject); // { param1: "value1", param2: "value2", param3: "value3" }

```
```js
const code = 'console.log("Hello, World!");';
eval(code); // Output: Hello, World!

```
```vue
<template>
  <router-link :to="{ path: '/?next=1' }">
</template>
```
will trigger update

```vue
<template>
  <div v-show="prevURLString !== route.fullPath">
    <router-link :to="prevLocation">
      <button id="index-top-hook" style="width: 100%;">上一页</button>
    </router-link>
  </div>
</template>
```

template not work   


```ts
  const newImageList:IImage[] = gallery.value.urls.map(async (url) => {
    const image = await fetchImage(url)
    return image
  })
```

(method) Array<string>.map<Promise<IImage>>(callbackfn: (value: string, index: number, array: string[]) => Promise<IImage>, thisArg?: any): Promise<IImage>[]

Type 'Promise<IImage>[]' is not assignable to type 'IImage[]'.
  Type 'Promise<IImage>' is missing the following properties from type 'IImage': query, galleryQuery, nextPageQuery, prevPageQuery, and 3 more.ts(2322)


tag的顺序


搜搜

tag的过滤


```vue
<template>
  <div class="scroll-container" @scroll="handleScroll">
    <div class="content">
      <!-- 内容 -->
    </div>
    <div class="element" :class="{ 'show': showElement }">
      <!-- 要显示或隐藏的元素 -->
    </div>
  </div>
</template>

<script lang="js">
export default {
  data() {
    return {
      showElement: false,
      prevScrollTop: 0
    };
  },
  methods: {
    handleScroll(event) {
      console.log("handleScroll")
      const scrollTop = event.target.scrollTop;
      
      if (scrollTop > this.prevScrollTop) {
        // 向下滚动，隐藏元素
        this.showElement = false;
      } else {
        // 向上滚动，显示元素
        this.showElement = true;
      }
      
      this.prevScrollTop = scrollTop;
    }
  }
}
</script>

<style>
.scroll-container {
  height: 500px;
  overflow-y: scroll;
  position: relative;
  background-color: aqua;
}

.content {
  height: 2000px; /* 超过容器高度以触发滚动 */
}

.element {
  position: fixed;
  bottom: 10px;
  left: 50%;
  transform: translateX(-50%);
  background-color: #f2f2f2;
  padding: 10px;
  transition: opacity 0.3s ease-in-out;
  opacity: 0;
}

.element.show {
  opacity: 1;
}
</style>

```

scroll


迷了，之前query怎么跑起来的
可能是用了string？


```ts
function myFunction() {
  return new Promise((resolve, reject) => {
    // Perform asynchronous operation or any logic here

    if (/* condition for success */) {
      resolve("Operation completed successfully");
    } else {
      reject("Error occurred");
    }
  });
}
```
promise 接受一个 函数作为输入初始化
为其传入两个callback
运行传入函数,当callback被call时promise完成

```sh
tsc soup.ts ; node soup.js ; rm soup.js
```


搜索框:


doujin
#9E2720

manga
#DB6C24

artist cg
#D38F1D

game cg
#6A936D

westen
#AB9F60

non-h
#5FA9CF

image set
#325CA2

cosplay
#6A32A2

asian porn
#A23282

misc
#777777

TODO:

docker
