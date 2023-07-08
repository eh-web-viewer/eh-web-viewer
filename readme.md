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
  <router-link :to="{ path: '/?next=1' }">
```
will trigger update

```vue

  <div v-show="prevURLString !== route.fullPath">
    <router-link :to="prevLocation">
      <button id="index-top-hook" style="width: 100%;">上一页</button>
    </router-link>
  </div>
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

