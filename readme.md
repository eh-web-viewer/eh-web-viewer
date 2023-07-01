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

后端读一下配置。

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
