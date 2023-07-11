<template>
  <!-- <div class="scroll-hook" @scroll="handleScroll"> -->
  <search-bar :show-search-box="showSearchBox" @submit="search"></search-bar> 

  <div v-show="prevPageButtonShow">
    <!-- <router-link :to="prevURLString"> -->
      <button id="index-top-hook" style="width: 100%;" @click="prevPageHandler">上一页</button>
    <!-- </router-link> -->
  </div>

  <template v-for="summary in previewList" :key="summary">
    <div @click="clicked(summary)" class="gallery-preview-wrapper" style="align-items: center; width: 100%; margin: 30px 0px; background-color: RGBA(255,255,255,0.1); height: 280px;">
      <router-link :to="{ path: summary.url.substring('https://exhentai.org'.length) }">
        <gallery-info :gallery="summary" />
      </router-link>
    </div>
  </template>

  <div v-show="nextPageButtonShow">
    <!-- <router-link :to="nextURLString"> -->
      <button id="index-bottom-hook" style="width: 100%;" @click=nextPageHandler>下一页</button>
    <!-- </router-link> -->
  </div>
  <!-- </div> -->
</template>

<script setup lang="ts">
// consts
const HOST = "https://exhentai.org"
// components
import GalleryInfo from "@/components/GalleryInfo.vue";
import SearchBar from "@/components/SearchBar.vue";
// vue-core
import { 
  ref,
  onBeforeMount,
  onMounted, 
  onBeforeUpdate, 
  onBeforeUnmount,
} from "vue"
// vue-route
import { 
  RouteLocationRaw,
  useRoute, useRouter,
} from 'vue-router'
const route = useRoute()
const router = useRouter()
// api
import {
  IGallery, 
  fetchIndex,
} from "@/functions/api";
// store
import {
  reactives,
} from '@/functions/store'
// const metaData = reactives.indexMetaData
const previewList = ref<IGallery[]>([])
const prevURLString = ref<string>("")
const nextURLString = ref<string>("")
const prevLocation = ref<RouteLocationRaw>(route.fullPath)
const nextLocation = ref<RouteLocationRaw>(route.fullPath)
const showSearchBox = ref(true)
const prevPageButtonShow = ref(false)
const nextPageButtonShow = ref(false)

let lastPath : string
let lastScrollTop = 0

function clicked(summary:IGallery) {
  const galleryData = reactives.galleryMetaData
  galleryData.preview = summary.preview
  galleryData.category = summary.category
  galleryData.title = summary.title
  galleryData.pages = summary.pages
  galleryData.tags = {}  
}
// return "" or nexturl like "https://exhentai.org/?next=2603193"
function getVar(v:string, x:string, d:string):string {
  return eval(x+`if (typeof ${v} === 'undefined') var ${v}="${d}";`+v)
}
// helper function to make RouteLocationObj
function urlString2RouteLocationObj(urlString:string) {
// function urlString2RouteLocationObj(urlString:string):[{path:string, query:Record<string,string>}, string] {
  // resolve the raw rul
  const url = new URL(urlString)
  // console.log(url.pathname)
  // console.log(url.search)
  const urlSearchString = url.search

  // get query object from search string
  const query:Record<string,string> = {};
  const params = new URLSearchParams(urlSearchString);
  params.forEach((value,key) => {
    // console.log(key,value)
    query[key] = value;
  });
  // set link
  const location = {
    path: url.pathname,
    query: query,
  }
  return location
}
// helper function to make urlSearchString(compare to route.fullpath to see should show botton or not)
function urlString2urlSearchString(urlString:string) {
  const url = new URL(urlString)
  const urlSearchString = url.search
  return "/"+urlSearchString
}
function clearPreviewList() {
  console.log("clearPreviewList")
  previewList.value = [] 
}

function prevPageHandler() {
  prevPageButtonShow.value = false
  nextPageButtonShow.value = false
  clearPreviewList()
  updateIndex(prevURLString.value)
}
function nextPageHandler() {
  nextPageButtonShow.value = false
  // clearPreviewList()
  updateIndex(nextURLString.value)
}

function updateIndex(url:string) {
  // for no ducaple requests
  if (url === '' || lastPath === url) return 
  lastPath = url
  console.log("!!updateIndex", url)
  fetchIndex(url)
  .then((index) => { // then is accutally a call back?
    router.push(url)
    // set prevPage and nextPage
    // console.log(index)
    const nexturl:string = getVar("nexturl", index.nextPage, HOST+route.fullPath)
    const prevurl:string = getVar("prevurl", index.nextPage, HOST+route.fullPath)
    // console.log(prevurl)
    // console.log(nexturl)
    // const nxv = urlString2RouteLocationObj(nexturl)
    // console.log("nxv,", nxv)
    nextLocation.value = urlString2RouteLocationObj(nexturl)
    nextURLString.value = urlString2urlSearchString(nexturl)
    // console.log(nextURLString.value, route.fullPath)    
    prevLocation.value = urlString2RouteLocationObj(prevurl)
    prevURLString.value = urlString2urlSearchString(prevurl)
    // console.log(prevURLString.value, route.fullPath)    

    console.log(index)
    previewList.value = [...previewList.value, ...index.galleries]

    nextPageButtonShow.value = true
  })
  console.log("updateIndex", url, "end")

}
function handleScroll() {
  // var st = window.pageYOffset || document.documentElement.scrollTop; // Credits: "https://github.com/qeremy/so/blob/master/so.dom.js#L426"
  const st = document.documentElement.scrollTop; 
  // console.log(isIndex, st, lastScrollTop, st-lastScrollTop)
  if (st > lastScrollTop + 10) {
    // downscroll code
    showSearchBox.value = false
  } else if (st < lastScrollTop - 10) {
    // upscroll code
    showSearchBox.value = true
  } // else was horizontal scroll
  lastScrollTop = st <= 0 ? 0 : st; // For Mobile or negative scrolling
}
function search(searchString:string) {
  console.log(searchString)
  prevPageButtonShow.value = false
  nextPageButtonShow.value = false
  prevURLString.value = ''
  nextURLString.value = ''
  previewList.value = []
  updateIndex(searchString)
}

onBeforeMount(async () => {
  document.title = "EhWebViewer"
  console.log("indexView: onBeforeMount")
  await router.isReady() // use this or will get '/' only
  console.log(route.fullPath)
  console.log(route)
  // prevURLString.value = route.fullPath
  // nextURLString.value = route.fullPath
  updateIndex(route.fullPath)

})

// initial hooks
onMounted(() => {
  // auto page turning
  const entry = document.querySelector("#index-bottom-hook")
  console.log(entry)
  if (entry !== null) {// it must not null? dunno
    // observer obj
    const observer = new IntersectionObserver((entries) => {
        entries.forEach((entry) => {
          // console.log(getIdFromUrl(props.image.url), ":" , entry.intersectionRatio)
          const remainingDistance = entry.boundingClientRect.top - entry.rootBounds!.top;
          // seems well 
          // console.log("index:", `Remaining distance to initial entry: ${remainingDistance}px`); 
          if (remainingDistance < 1080) { // will affect initial images loaded
            // seems well
            console.log("observer triggered", nextURLString.value)
            // if (nextPageButtonShow.value)
            //   updateIndex(nextURLString.value)
          }
        });
      }, {
        root: null,
        rootMargin: '0px 0px 0px', // will affect how much images load before, and this is two times when use 'Xpx'
        threshold: [0, 0.2, 0.4, 0.6, 0.8, 1]
      }
    )
    setTimeout( () => observer.observe(entry), 500);
  }
  // scroll
  window.addEventListener('scroll', handleScroll);

})

onBeforeUpdate(async () => {
  console.log("indexView: onBeforeUpdate")
  await router.isReady() // use this or will get '/' only
  console.log(route.fullPath)
  // updateIndex(route.fullPath)
})

onBeforeUnmount(() => {
  // will it run on some day?
  window.removeEventListener('scroll', handleScroll);
})
</script>
  