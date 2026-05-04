<template>
    <main class="min-h-screen text-black pb-20 pt-40 bg-gray-50 px-8">
        <header class="flex flex-col items-center justify-center -mt-12">
            <span
                class="inline-block text-[10px] tracking-widest uppercase text-center text-white bg-gray-500 rounded-full px-3 py-1 mb-1 font-montserrat sm:text-xs">
                LISTEN
            </span>
            <h2 class="text-2xl font-oswald font-bold md:text-3xl">TEACHINGS</h2>

            <p class="my-4 text-md font-montserrat md:text-lg">Our latest teachings for your growth in the faith</p>

            <div class="w-full border border-gray-700 rounded-3xl flex overflow-hidden md:w-3/5">
                <input type="search" name="search" id="search" placeholder="Search by album and track title..."
                    class="w-full text-sm font-montserrat outline-none py-2 pl-5 pr-1 md:py-3 md:text-md" />
                <button type="button" class="bg-blue-600 px-4 py-3 cursor-pointer text-white/80">
                    <PhMagnifyingGlass weight="bold" size="20" />
                </button>
            </div>
        </header>

        <section class="mt-12 font-montserrat">
            <div v-if="store.loading" class="text-gray-400 text-sm">Loading...</div>

            <div v-else-if="store.albums && store.albums.length > 0"
                class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
                <NuxtLink v-for="album in store.albums" :key="album.id" :to="`/teachings/${album.id}`"
                    class="bg-white/80 rounded-xl p-4 transition shadow-sm hover:bg-white hover:shadow-md">
                    <div class="w-full h-36 bg-white/80 rounded-lg mb-3 overflow-hidden">
                        <img v-if="album.thumbnail_url" :src="album.thumbnail_url" :alt="album.title"
                            class="w-full h-full object-cover" />
                        <div v-else class="w-full h-full flex items-center justify-center text-gray-500 text-sm">
                            <PhMusicNote size="150" weight="thin" />
                        </div>
                    </div>
                    <h3 class="font-semibold text-gray-800 text-sm truncate">{{ album.title }}</h3>
                    <p class="text-gray-400 text-xs mt-1">{{ formatDate(album.date) }}</p>
                    <p class="text-gray-400 text-xs mt-1">Uploaded on: {{ formatDate(album.createdAt) }}</p>
                </NuxtLink>
            </div>

            <p v-else class="text-gray-500 text-sm w-full text-center">No teachings yet.</p>
        </section>
    </main>
</template>

<script setup lang="ts">
import { PhMagnifyingGlass, PhMusicNote } from '@phosphor-icons/vue';
import { useTeachingsService } from '~/services/teaching';

const { setLight } = useHeaderTheme()

const store = useTeachingsStore()
const { fetchAlbums } = useTeachingsService()

function formatDate(date: string) {
    return new Date(date).toLocaleDateString('en-NG', {
        month: 'long',
        year: 'numeric',
        day: 'numeric',
    })
}

onMounted(() => {
    setLight()
    fetchAlbums()
})
</script>