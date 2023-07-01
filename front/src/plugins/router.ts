import { createRouter, createWebHistory } from 'vue-router';

const routes = [
    { path: '/g/:id/:key/' , name: 'gallery', component: () => import('@/views/GalleryView.vue') },
    { path: '/g/:id/:key' , name: 'gallery', component: () => import('@/views/GalleryView.vue') },
    { path: '/:pathMatch(.*)*' , name: 'index', component: () => import('@/views/IndexView.vue') }, // catch all (https://router.vuejs.org/zh/guide/migration/)
]

const router = createRouter({
    routes,
    history: createWebHistory(),
})

export default router 