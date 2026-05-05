<template>
    <main class="min-h-screen text-black pb-20 pt-40 bg-gray-50 px-8">
        <header class="-mt-12">
            <p class="uppercase pb-1 font-oswald text-gray-800 md:text-lg">Teaching</p>
            <hr />
            <h2 class="font-monstserrat font-bold uppercase mt-4 text-xl text-gray-700 md:text2xl">{{
                store.currentAlbum?.title }}
            </h2>
        </header>

        <section class="mt-6 font-montserrat sm:mt-8">
            <p v-if="store.loading" class="text-gray-400 text-sm w-full text-center">Loading...</p>
            <div v-else-if="store.currentAlbum && store.currentAlbum.tracks.length > 0">
                <div class="flex items-center justify-between">
                    <span class="text-sm font-bold text-gray-500">TRACKS</span>
                    <span class="text-sm font-bold text-gray-500">TOTAL TRACKS: {{ store.currentAlbum.tracks.length
                    }}</span>
                </div>

                <div class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3 mt-3">
                    <div v-for="track in store.currentAlbum.tracks" :key="track.id""
                    class=" bg-white/80 rounded-xl p-4 transition shadow-sm cursor-pointer hover:bg-white
                        hover:shadow-md" @click="() => playerStore.play(track, store.currentAlbum!.title)">
                        <div class="w-full h-36 bg-white/80 rounded-lg mb-3 overflow-hidden">
                            <img v-if="store.currentAlbum.thumbnail_url" :src="store.currentAlbum.thumbnail_url"
                                :alt="store.currentAlbum.title" class="w-full h-full object-cover" />
                            <div v-else class="w-full h-full flex items-center justify-center text-gray-500 text-sm">
                                <PhMusicNote size="150" weight="thin" />
                            </div>
                        </div>
                        <h3 class="font-semibold text-gray-800 text-sm truncate">{{ track.title }} {{ track.trackNumber
                        }}</h3>
                        <p class="text-gray-400 text-xs mt-1">Uploaded on: {{ formatDate(track.createdAt) }}</p>
                    </div>
                </div>
            </div>

            <p v-else class="text-gray-500 text-sm w-full text-center">No track yet.</p>
        </section>
    </main>
</template>

<script setup lang="ts">
import { PhMusicNote } from "@phosphor-icons/vue"
import { useTeachingsService } from "../../services/teaching"

const { setLight } = useHeaderTheme()
setLight()

const store = useTeachingsStore()
const playerStore = usePlayerStore()
const { fetchAlbumByID } = useTeachingsService()
const route = useRoute()
const id = route.params.id as string

function formatDate(date: string) {
    return new Date(date).toLocaleDateString('en-NG', {
        month: 'long',
        year: 'numeric',
        day: 'numeric',
    })
}
onMounted(() =>
    fetchAlbumByID(id)
)
</script>