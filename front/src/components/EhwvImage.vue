NEW!

show the image, if on error, pass on error event to parent.

params:
  :img="[src]"
    the src of image.
  @error="[func]"
    what will call on error


<template>
  <img :src="computedSrc" :alt="computedSrc" @error="imgError" @load="imgLoaded" @loadstart="imgLoadStart" >
</template>


<script setup lang="ts">

import { 
  onMounted,
  onUpdated,
  computed,
} from 'vue';

const props = defineProps<{
  src?: string
  timeout?: number
}>()

const emits = defineEmits([
  'update:error', 
])


const computedSrc = computed(() => {
  return props.src ?? "/favicon.ico"
})
const computedTimeout = computed(() => {
  return props.timeout ?? 5000
})

let timeoutId : NodeJS.Timeout

async function imgError() {
  emits('update:error')
}


function imgLoadStart() {
  // set timeout if  
  // clear timer.
  clearTimeout(timeoutId); // TODO: what will happen if not initalized?
  timeoutId = setTimeout(() => {
    imgError()
  }, computedTimeout.value);
}

function imgLoaded() {
  // cancel the timer when image is loaded successfully
  clearTimeout(timeoutId);
}


function onMountedOrUpdated(){
  // do nothing?  
}

onMounted(() => {
  onMountedOrUpdated()
})

onUpdated(() => {
  onMountedOrUpdated()
})


</script>