<template>
    <main class="min-h-screen text-black pb-20 pt-40 bg-gray-50 px-8">
        <header class="-mt-12">
            <p class="uppercase pb-1 font-oswald text-gray-800 md:text-lg">Choir Ministration</p>
            <hr />
            <h2 class="font-monstserrat font-bold uppercase mt-4 text-xl text-gray-700 md:text2xl">{{
                store.currentSong?.songTitle }}
            </h2>
        </header>

        <section class="mt-6 font-montserrat sm:mt-8">
            <div v-if="store.loading" class="text-gray-400 text-sm w-full text-center">Loading...</div>
            <div v-else-if="store.currentSong">
                <div class="flex items-center justify-between">
                    <span class="text-xs font-bold text-gray-500">SONG</span>
                    <span class="text-xs font-bold text-gray-500">SONG WRITER: {{ store.currentSong.songWriter }}</span>
                </div>

                <div class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3 mt-3">
                    <div class="bg-white/80 rounded-xl p-4 transition shadow-sm cursor-pointer hover:bg-white
                        hover:shadow-md" @click="() => playerStore.playSong(store.currentSong as ChoirMinistration)">
                        <div class="w-full h-36 bg-white/80 rounded-lg mb-3 overflow-hidden">
                            <img v-if="store.currentSong.thumbnailURL" :src="store.currentSong.thumbnailURL"
                                :alt="store.currentSong.songTitle" class="w-full h-full object-cover" />
                            <div v-else class="w-full h-full flex items-center justify-center text-gray-500 text-sm">
                                <PhMusicNote size="150" weight="thin" />
                            </div>
                        </div>
                        <h3 class="font-semibold text-gray-800 text-sm truncate">{{ store.currentSong.songTitle }}</h3>
                        <p class="text-gray-400 text-xs mt-1">Uploaded on: {{ formatDate(store.currentSong.createdAt) }}
                        </p>
                    </div>
                </div>
            </div>

            <p v-else class="text-gray-500 text-sm w-full text-center">No song yet.</p>
        </section>
    </main>
</template>

<script setup lang="ts">
import { PhMusicNote } from "@phosphor-icons/vue"
import { useChoirMinistrationsService } from "../../services/choir-ministrations"
import type { ChoirMinistration } from "~/types"

const { setLight } = useHeaderTheme()

const store = useChoirMinistrationsStore()
const playerStore = usePlayerStore()
const { fetchChoirMinistrationById } = useChoirMinistrationsService()
const route = useRoute()
const id = route.params.id as string

function formatDate(date: string) {
    return new Date(date).toLocaleDateString('en-NG', {
        month: 'long',
        year: 'numeric',
        day: 'numeric',
    })
}
onMounted(() => {
    setLight()
    fetchChoirMinistrationById(id)
})
</script>