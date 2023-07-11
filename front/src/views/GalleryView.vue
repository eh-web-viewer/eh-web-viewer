<template>  
  <div>
    <gallery-info :gallery="gallery" v-if="(typeof gallery !== 'undefined')"></gallery-info>
    
    <div v-show="(curPage > 0)" v-if="(typeof gallery !== 'undefined')">
      <router-link :to="{path: gallery.query, query: { p: curPage-1 }, replace: true}">
        <button id="gallery-bottom-hook" style="width: 100%;">上一页</button>
      </router-link>
    </div>

    <template v-for="imageURL in imageURLList" :key="imageURL">
      <router-link :to="chopString(imageURL, HOST)">
        <image-box-url :url="imageURL"></image-box-url>
      </router-link>
      {{ findIndexFromImageUrl(imageURL) }}
    </template>

    <div v-show="(curPage < maxPage)" v-if="(typeof gallery !== 'undefined')">
      <router-link :to="{path: gallery.query, query: { p: curPage+1 }, replace: true}">
        <button id="gallery-bottom-hook" style="width: 100%;">下一页</button>
      </router-link>
    </div>

  </div>
</template>

<script setup lang="ts">
const HOST = "https://exhentai.org"
// components
import ImageBoxUrl from "@/components/ImageBoxUrl.vue";
// vue-core
import { onMounted, onUpdated, ref } from "vue";
// vue-route
import { useRoute, useRouter } from "vue-router";
const route = useRoute();
const router = useRouter();

import { fetchGallery, IGallery } from "@/functions/api";
import { chopString, getNumberFromString, getParam, findIndexFromImageUrl } from "@/functions/utils";
import GalleryInfo from "@/components/GalleryInfo.vue";

const gallery = ref<IGallery>()
const imageURLList = ref<string[]>([])
const curPage = ref(0)
const maxPage = ref(0)
// let fetching = false
// let curPage = 0
// let maxPage = 0
let lastQuery = ""
// not used
// function nextPage(curPage:number) {
//   if (typeof gallery.value === 'undefined') return "/"
//   return findPath(gallery.value.query)+"?p=" + (curPage+1)
// }
// function prevPage(curPage:number) {
//   if (typeof gallery.value === 'undefined') return "/"
//   return findPath(gallery.value.query)+"?p=" + (curPage-1)
// }
// not used
// function pageTo(page:number) {
//   if (typeof gallery.value === 'undefined') return "/"
//   return findPath(gallery.value.query)+"?p=" + (page) 
// }
// not used
// async function loadNextPage() {  
//   console.log("call loadNextPage", fetching)
//   fetching = true
//   curPage++
//   console.log("loadNextPage",curPage)
//   const path = findPath(gallery.value!.query)+"?p="+(curPage)
//   const nextGallery = await fetchGallery(path)
//   imageURLList.value = [...imageURLList.value, ...nextGallery.urls]
//   fetching = false
//   console.log("end of loadNextPage", curPage, fetching)
// }

// not used
// onBeforeMount(async  () => {
//   console.log("Gallery: onBeforeMount");
//   await router.isReady(); // use this or will get '/' only
// })

onMounted(async () => {
  console.log("Gallery: onMount");
  await router.isReady(); // use this or will get '/' only
  lastQuery = route.fullPath
  const query = route.fullPath
  gallery.value = await fetchGallery(query)

  console.log(route.query)
  curPage.value = parseInt(getParam(route.query, 'p')||'0')
  maxPage.value = Math.floor((getNumberFromString(gallery.value.pages)-0.5)/20)
  console.log(maxPage.value)
  
  // imageURLList.value = [...imageURLList.value, ...gallery.value.urls]
  imageURLList.value = gallery.value.urls

});

// not used
onUpdated(async () => {
  console.log("Gallery: onUpdated")
  await router.isReady()
  if (lastQuery === route.fullPath) return
  lastQuery = route.fullPath

  imageURLList.value = []

  const query = route.fullPath
  const newGallery = await fetchGallery(query)

  console.log(route.query)
  curPage.value = parseInt(getParam(route.query, 'p')||'0')
  imageURLList.value = newGallery.urls

});
</script>
