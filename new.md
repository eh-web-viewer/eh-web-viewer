
# 再做啥

- EhevImage
  - 显示图片，好像好了。

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
