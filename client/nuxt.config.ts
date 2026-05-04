import tailwindcss from "@tailwindcss/vite"

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  vite: {
    plugins: [tailwindcss() as any],
    optimizeDeps: {
      include: [
        "@phosphor-icons/vue"
      ]
    }
  },
  css: [`~/assets/css/main.css`],
  devtools: { enabled: true },
  modules: [
    "@pinia/nuxt",
    "@nuxt/fonts"
  ],
  runtimeConfig: {
    public: {
      apiUrl: "http://localhost:8080"
    }
  },
  $development: {
    runtimeConfig: {
      public: {
        apiUrl: "http://localhost:8080"
      }
    }
  },
  $production: {
    runtimeConfig: {
      public: {
        apiUrl: "my-api-url"
      }
    }
  },
})