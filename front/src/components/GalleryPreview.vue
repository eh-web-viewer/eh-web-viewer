<template>

<div style="display: flex;">
  <div @click="clicked" class="gallery-preview" style="align-items: center; width: 100%; margin: 30px 0px">  
    <router-link :to="{ path: summary.url.substring('https://exhentai.org'.length) }">
      <div style="display: flex; width: 100%; height: 280px; align-items:center ">
        <div class="cover-container" style="flex: 1; height: 100%; padding: 10px; display: flex; align-items: center; justify-content: center;">      
          <img :src="src" :alt="src" loading="lazy" style="max-height: 100%; max-width: 100%; width: auto; height: auto" onerror="this.onerror=null; this.src=this.src">
        </div>
        <div style="height: 100%; flex: 2; display: flex; position: relative;">      
          <p style="align-items:baseline; text-align: left;">
            {{ summary.title }}
          </p>
          <p style="position: absolute; bottom:0; right:0;">
            {{ summary.pages }} Pages
          </p>
          <p style="display: none;">
            {{ summary }}
          </p>
        </div>
      </div>
    </router-link>
  </div>
</div>

</template>


<script lang="ts" setup>

import { ref } from 'vue'
import { RouterLink } from 'vue-router'
import { GallerySummary } from '@/functions/store';
import { reactives } from '@/functions/store';

const props = defineProps<{
  summary : GallerySummary,
}>()

function clicked() {
  // reuse the data from summary
  const summary = props.summary
  const galleryData = reactives.galleryMetaData

  galleryData.preview = summary.preview

  galleryData.category = summary.category
  galleryData.title = summary.title
  galleryData.pages = summary.pages
  galleryData.tags = {}

  
}

// function retry() {
//   console.log("GalleryPreview: retry")
//   const temp = src.value
//   src.value = "/favicon.ico"
//   src.value = temp
// }

const src = ref(props.summary.preview.replace(/s.exhentai.org/g, 's-ex.moonchan.xyz'))

console.log(props.summary.preview.replace(/s.exhentai.org/g, 's-ex.moonchan.xyz'))

</script>