require('dotenv').config()
const {
  FIREBASE_API_KEY,
  FIREBASE_AUTH_DOMAIN,
  FIREBASE_DATABASE_URL,
  FIREBASE_STORAGE_BUCKET,
  FIREBASE_PROJECT_ID,
  FIREBASE_MESSAGING_SENDER_ID,
  FIREBASE_APP_ID,
} = process.env

export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'youtube-manager-nuxt',
    htmlAttrs: {
      lang: 'en',
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
    '@fortawesome/fontawesome-free/css/all.css',
    '@fortawesome/fontawesome-free/css/brands.css',
    '@fortawesome/fontawesome-free/css/fontawesome.css',
    '@fortawesome/fontawesome-free/css/regular.css',
    '@fortawesome/fontawesome-free/css/solid.css',
    '@fortawesome/fontawesome-free/css/svg-with-js.css',
    '@fortawesome/fontawesome-free/css/v4-shims.css',
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
    '~/plugins/vue-youtube',
    '~/plugins/cookies-to-state',
    '~/plugins/axios',
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/eslint
    // "@nuxtjs/eslint-module"
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://bulma.io
    '@nuxtjs/bulma',
    // https://go.nuxtjs.dev/axios
    '@nuxtjs/axios',
    'cookie-universal-nuxt',
  ],

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    baseURL: 'http://localhost:8080/',
  },
  proxy: {
    '/api': '/',
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
    postcss: {
      preset: {
        features: {
          customProperties: false,
        },
      },
    },
  },

  env: {
    FIREBASE_API_KEY,
    FIREBASE_AUTH_DOMAIN,
    FIREBASE_DATABASE_URL,
    FIREBASE_STORAGE_BUCKET,
    FIREBASE_PROJECT_ID,
    FIREBASE_MESSAGING_SENDER_ID,
    FIREBASE_APP_ID,
  },
}
