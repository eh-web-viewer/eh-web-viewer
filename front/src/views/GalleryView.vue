<template>
  <!-- metaData : -->
  <br />
  <!-- {{ metaData }} -->
  <template v-for="imageList in reactives.galleryImageLists.value">
    <!-- 【{{ imageList }}】 -->
    <template v-for="image in imageList">
      <!-- 【{{ image }}】 -->
      <gallery-image :image="image"></gallery-image>
    </template>
  </template>
  <div id="gallery-view-hook"></div>
</template>

<script setup lang="ts">
// components
import GalleryImage from "@/components/GalleryImage.vue";
// vue-core
import { onBeforeUpdate, onMounted, Ref } from "vue";
// vue-route
import { useRoute, useRouter } from "vue-router";
const route = useRoute();
const router = useRouter();
// get datas
import { reactives, Image } from "@/functions/store";
import { fetchGallery } from "@/functions/api";


const metaData = reactives.galleryMetaData; // Record
// const imageLists = reactives.galleryImageLists // reactive<IImage[][]>
// let lastPath : string

// fetch next page
async function updateGallery(path: string, nextPage: string) {
  // const nextPage = ((reactives.galleryImageLists as Ref<Image[][]>).value.length).toString()
  // it seems that the await will with the inner function finish.
  await fetchGallery(path + "?p=" + nextPage).then((gallery) => {
    const metaData = reactives.galleryMetaData;
    // metaData.query = gallery.query;
    metaData.title = gallery.title;
    document.title = gallery.title + ' - EhWebView';
    metaData.originalTitle = gallery.originalTitle;
    metaData.tags = gallery.tags;
    metaData.preview = gallery.preview;
    metaData.category = gallery.category;
    metaData.pages = gallery.pages;
    const imageList = Array<number>(gallery.images.length)
      .fill(0)
      .map((_, k) => {
        return {
          preview: gallery.images[k],
          url: gallery.url[k],
          image: "",
        } as Image;
      });
    reactives.galleryImageLists.value.push(imageList);
  });
}

function galleryShouldStop(pages: string, len: number) {
  if (pages === "") return false;
  const maxPages = Math.floor((parseInt(pages) - 0.5) / 20);
  return len > maxPages;
}

onMounted(async () => {
  console.log("Gallery: onMount");
  await router.isReady(); // use this or will get '/' only
  // do not renew if click the same gallery.
  metaData.query = route.path;
  console.log(route)
  if (metaData.query === metaData.lastQuery) {
    return;
  }
  (reactives.galleryImageLists as Ref<Image[][]>).value.length = 0;
  metaData.lastQuery = metaData.query;
  // fetch all gallery pages.
  while (!galleryShouldStop( metaData.pages, (reactives.galleryImageLists as Ref<Image[][]>).value.length) ) {
    
    console.log(metaData.pages, (reactives.galleryImageLists as Ref<Image[][]>).value.length)
    console.log(!galleryShouldStop( metaData.pages, (reactives.galleryImageLists as Ref<Image[][]>).value.length))
      
    await updateGallery( metaData.lastQuery, (reactives.galleryImageLists as Ref<Image[][]>).value.length.toString() );
  }  
});

onBeforeUpdate(async () => {
  // console.log("Gallery: onBeforeUpdate")
  // await router.isReady()
  // updateGallery(route.fullPath)
  // lastPath = route.fullPath
});
</script>
