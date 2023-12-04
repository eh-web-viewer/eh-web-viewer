
# 再做啥

- [EhwvImage](#ehwvimage)
  - 显示图片，好像好了。
- ~~api缓存。~~
  
- ~~图片预加载功能？~~

别做了，把tag的watcher和高亮做出来就行了
还有gallery和tag显示的分离
几乎改不动了。。

草，乱七八糟的，我自己也找不到，不想写了。
要么以后重写。

## [EhwvImage](./front/src/components/EhwvImage.vue)

## [EhwvImageView](./front/src/views/EhwvImageView.vue)

# 笔记

## 组件之间的传参

```vue
const props = defineProps<{
  src?: string
  timeout?: number
}>()
```

```vue
:src="[src]" :timeout="[timeout]"
```

TODO：默认值怎么弄。

## emit事件

```vue
const emits = defineEmits([
  'update:error', 
])
```

```vue
@error="onError"
```
