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
      title: "Loyola Youth Assembly (LYA) Media Website",
      htmlAttrs: {
        lang: "en",
      },
      charset: "utf-8",
      viewport: "width=device-width, initial-scale=1.0",
      link: [
        {
          rel: "icon",
          type: "image/x-icon",
          href: "/favicon.ico",
        },
        {
          rel: "canonical",
          href: "https://lya-media-website.vercel.app",
        },
      ],
      meta: [
        {
          name: "description", 
          content:
            "Welcome to the official website of Loyola Youth Assembly Media Team. Here you will find all our teachings, choir ministrations, and worship sessions. LYA is the Youth Arm of CAC Oke Alafia Loyola, Ibadan.",
        }
      ]
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
