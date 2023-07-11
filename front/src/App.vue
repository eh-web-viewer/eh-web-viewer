<template>
  <!-- <header> -->
  <header style="display:none">
    <span><router-link :to="{path: '/'}">home</router-link></span>
    <span><router-link :to="{path: '/g/2597892/20f9db69dd/'}">gallery</router-link></span>
    <span><router-link :to="{path: '/s/b028d14f3d/2599914-1'}">image</router-link></span>
  </header>
  <!-- app.vue do not update. -->
  <!-- <search-bar :show-search-box="showSearchBox"></search-bar> -->
  <!-- <router-view> can no longer be used directly inside <transition> or <keep-alive>. -->
  <!-- <keep-alive> -->
    <router-view />    
  <!-- </keep-alive> -->
</template>

<script lang="ts" setup>

import { onBeforeUnmount, onMounted, onUpdated } from 'vue'
import { RouterView, RouterLink, useRoute, useRouter } from 'vue-router'
const route = useRoute()
const router = useRouter()

// let isIndex = false
// let lastScrollTop = 0

// const showSearchBox = ref(false)
// watch(route.fullPath, (oldValue, newValue) => {})

// not used here
// function handleScroll() {
//   // var st = window.pageYOffset || document.documentElement.scrollTop; // Credits: "https://github.com/qeremy/so/blob/master/so.dom.js#L426"
//   const st = document.documentElement.scrollTop; 
//   // console.log(isIndex, st, lastScrollTop, st-lastScrollTop)
//   if (st > lastScrollTop + 10) {
//     // downscroll code
//     showSearchBox.value = false
//   } else if (st < lastScrollTop - 10) {
//     // upscroll code
//     showSearchBox.value = isIndex
//   } // else was horizontal scroll
//   lastScrollTop = st <= 0 ? 0 : st; // For Mobile or negative scrolling
// }

onMounted(async () => {
  // scroll listener
  // window.addEventListener('scroll', handleScroll);
  await router.isReady()
  // fetch home page
  // do not remove this for it do the first time 
  // if (route.fullPath.startsWith('/g') || route.fullPath.startsWith('/s')){
  //   isIndex = false
  // }else{
  //   isIndex = true
  // }
  // console.log("isIndex" , isIndex)
  // showSearchBox.value = isIndex
})

onUpdated(async () => {
  console.log("App on Upaated")
  await router.isReady()
  console.log(route.fullPath)
  // fetch home page
  // do not remove this for it do the first time 
  // if (route.fullPath.startsWith('/g') || route.fullPath.startsWith('/s')){
  //   isIndex = false
  //   showSearchBox.value = false
  // }else{
  //   isIndex = true
  // }
  // console.log("isIndex" , isIndex)
})

onBeforeUnmount(() => {
  // will it run on some day?
  // window.removeEventListener('scroll', handleScroll);
})
</script>