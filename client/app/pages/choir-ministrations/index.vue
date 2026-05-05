<template>
    <main class="min-h-screen text-black pb-20 pt-40 bg-gray-50 px-8">
        <header class="flex flex-col items-center justify-center -mt-12">
            <span
                class="inline-block text-[10px] tracking-widest uppercase text-center text-white bg-gray-500 rounded-full px-3 py-1 mb-1 font-montserrat sm:text-xs">
                LISTEN
            </span>
            <h2 class="text-2xl font-oswald font-bold md:text-3xl">CHOIR MINISTRATIONS</h2>

            <p class="my-4 text-md font-montserrat md:text-lg">The Emissaries songs and ministrations</p>

            <div class="w-full border border-gray-700 rounded-3xl flex overflow-hidden md:w-3/5">
                <input type="search" name="search" id="search" placeholder="Search by date and song title..."
                    class="w-full text-sm font-montserrat outline-none py-2 pl-5 pr-1 md:py-3 md:text-md" />
                <button type="button" class="bg-blue-600 px-4 py-3 cursor-pointer text-white/80">
                    <PhMagnifyingGlass weight="bold" size="20" />
                </button>
            </div>
        </header>

        <section class="mt-12 font-montserrat">
            <div v-if="store.loading" class="text-gray-400 text-sm">Loading...</div>

            <div v-else-if="store.choirMinistrations && store.choirMinistrations.length > 0"
                class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
                <NuxtLink v-for="song in store.choirMinistrations" :key="song.id"
                    :to="`/choir-ministrations/${song.id}`"
                    class="bg-white/80 rounded-xl p-4 transition shadow-sm hover:bg-white hover:shadow-md">
                    <div class="w-full h-36 bg-white/80 rounded-lg mb-3 overflow-hidden">
                        <img v-if="song.thumbnailURL" :src="song.thumbnailURL" :alt="song.songTitle"
                            class="w-full h-full object-cover" />
                        <div v-else class="w-full h-full flex items-center justify-center text-gray-500 text-sm">
                            <PhFolder size="150" weight="thin" />
                        </div>
                    </div>
                    <div class="flex justify-between items-center">
                        <h3 class="font-semibold text-gray-800 text-sm truncate">{{ song.songTitle }}</h3>
                        <p class="text-xs text-gray-500 truncate">Song Writer: {{ song.songWriter }}</p>
                    </div>
                    <p class="text-gray-400 text-xs mt-1">Ministered on: {{ formatDate(song.date) }}</p>
                    <p class="text-gray-400 text-xs mt-1">Uploaded on: {{ formatDate(song.createdAt) }}</p>
                </NuxtLink>
            </div>

            <p v-else class="text-gray-500 text-sm w-full text-center">No songs yet.</p>
        </section>
    </main>
</template>

<script setup lang="ts">
import { PhFolder, PhMagnifyingGlass } from '@phosphor-icons/vue'
import { useChoirMinistrationsService } from '~/services/choir-ministrations'
const { setLight } = useHeaderTheme()

const store = useChoirMinistrationsStore()
const { fetchChoirMinistrations } = useChoirMinistrationsService()

function formatDate(date: string) {
    return new Date(date).toLocaleDateString('en-NG', {
        month: 'long',
        year: 'numeric',
        day: 'numeric',
    })
}

onMounted(() => {
    setLight()
    fetchChoirMinistrations()
})
</script>