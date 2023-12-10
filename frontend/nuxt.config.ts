// https://nuxt.com/docs/api/configuration/nuxt-config
// import { resolve } from 'path'
export default defineNuxtConfig({
  // might conflict with namespace declaration - can lead to bugs in nuxt3
  // alias: {
  //   '@': resolve(__dirname, '/')
  // },
  spaLoadingTemplate: "spa-loading-template.html",
  modules: [
    "@pinia/nuxt",
    // '@nuxtjs/tailwindcss',
    "@pinia-plugin-persistedstate/nuxt",
  ],
  routeRules: {
    "/**": { ssr: false },
  },
  devtools: { enabled: true },
});
