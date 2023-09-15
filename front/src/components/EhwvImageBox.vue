EhwvImageBox:
show the image on exhentai.org/s/:key/:id 
params:
  url?: the /s/:key/:id
  image?: object that get on exhentai.org/s/:key/:id, will show the image and request the altImage when failed
  src?: if image is not set, show this.
tests:
  - [ ] on error
  - [ ] on retry


<template>
  <img :src="imgSrc" :alt="imgSrc" @error="imgError" @load="imgLoaded" @loadstart="imgLoadStart" >
</template>

<script setup lang="ts">
import { 
  ref,
  onMounted,
  onUpdated,
} from 'vue';
// not used
// import { useRoute, useRouter } from 'vue-router'
// const route = useRoute()
// const router = useRouter()
import { fetchImage, IImage } from '@/functions/api'
import { findPath } from '@/functions/utils';

const props = defineProps<{
  src?: string, // default
  image?: IImage, // object
  url?: string, // something like /s/:key/:id
}>()

const imgSrc = ref("/error.webp")
// let status = "init"
let imgObj = props.image

const timeout = 5000
let timeoutId : NodeJS.Timeout

async function imgError() {
  if (typeof imgObj === 'undefined') return
  fetchImage(findPath(imgObj.query) + "?nl=" + imgObj.altQuery)
  .then(obj => {
    imgObj = obj
    imgSrc.value = imgObj.image
  })
  clearTimeout(timeoutId);
  timeoutId = setTimeout(() => {
    imgError()
  }, timeout);
}


function onMountedOrUpdated(){
  // 
  if (typeof props.src !== 'undefined') {
    imgSrc.value = props.src
  }
  if (typeof props.image !== 'undefined') {
    imgObj = props.image
    imgSrc.value = imgObj.image
  } else if (typeof props.url !== 'undefined') {
    fetchImage(props.url)
    .then(obj => {
      imgObj = obj
      imgSrc.value = imgObj.image
    })
  }
}


function imgLoadStart() {
  // set timeout if  
  clearTimeout(timeoutId); // what will happen if not initalized
  timeoutId = setTimeout(() => {
    imgError()
  }, timeout);
}

function imgLoaded() {
  // cancel the timer when image is loaded successfully
  clearTimeout(timeoutId);
}

onMounted(() => {
  onMountedOrUpdated()
})

onUpdated(() => {
  onMountedOrUpdated()
})


</script>