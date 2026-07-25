import tailwindcss from "@tailwindcss/vite";

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  app: {
    head: {
      script: [
        {
          src: "https://www.sabilytics.com/script.js",
          'data-site': "t0bc2ojllsz2",
          'data-domain': "lya-media-website.vercel.app",
          async: true,
        }
      ],
    },
  },
  vite: {
    plugins: [tailwindcss() as any],
    optimizeDeps: {
      include: ["@phosphor-icons/vue"],
    },
  },
  css: [`~/assets/css/main.css`],
  devtools: { enabled: true },
  modules: ["@pinia/nuxt", "@nuxt/fonts", "@vercel/analytics"],
  runtimeConfig: {
    public: {
      apiUrl: "http://localhost:8080",
    },
  },
  $development: {
    runtimeConfig: {
      public: {
        apiUrl: "http://localhost:8080",
      },
    },
  },
  $production: {
    runtimeConfig: {
      public: {
        apiUrl: "https://lya-media-website-production.up.railway.app",
      },
    },
  },
});
