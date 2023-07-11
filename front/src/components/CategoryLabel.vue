<template>
  <span 
    class="serch-box-catagory-content" 
    :style="{
      backgroundColor : color,
      opacity: selected ? '1' : '0.2'
    }"
  >
    {{ category }}
  </span>
</template>

<script setup lang="ts">
const table = {
  'doujinshi'     : '#9E2720',
  'manga'      : '#DB6C24',
  'artist cg'  : '#D38F1D',
  'game cg'    : '#6A936D',
  'western'     : '#AB9F60',
  'non-h'      : '#5FA9CF',
  'image set'  : '#325CA2',
  'cosplay'    : '#6A32A2',
  'asian porn' : '#A23282',
  'misc'       : '#777777',
} as Record<string, string> 

import { onBeforeMount, onUpdated, ref } from 'vue';

const props = defineProps<{
  category : string
  selected : boolean
}>()

// const category = ref(props.category)
// const selected = ref(props.selected)
const color = ref(table[props.category.toLocaleLowerCase()] ?? '#000000') // this will not change


onBeforeMount(() => {
  // console.log(category.value)
})
onUpdated(() => {
  console.log("C-Label: Updated", props.category)
})

</script>

