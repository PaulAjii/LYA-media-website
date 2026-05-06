<template>
    <main class="min-h-screen text-black pb-20 pt-40 bg-gray-50 px-8">
        <header class="-mt-12">
            <p class="uppercase pb-1 font-oswald text-gray-800 md:text-lg">worship</p>
            <hr />
            <h2 v-if="store.currentWorshipSession"
                class="font-monstserrat font-bold uppercase mt-4 text-xl text-gray-700">Worship - {{
                    formatDate(store.currentWorshipSession.date) }}
            </h2>
            <h3 class="font-montserrat font-bold uppercase my-3 text-sm text-gray-600">Worship Leader: {{
                store.currentWorshipSession?.worshipLeader }}</h3>
        </header>

        <section class="mt-6 font-montserrat sm:mt-8">
            <p v-if="store.loading" class="text-gray-400 text-sm w-full text-center">Loading...</p>
            <div v-else-if="store.currentWorshipSession">
                <div class="bg-white/80 rounded-xl p-4 transition shadow-sm cursor-pointer mb-4 hover:bg-white
                            hover:shadow-md md:w-1/3" @click="handlePlay">
                    <div class="w-full h-36 bg-white/80 rounded-lg mb-3 overflow-hidden">
                        <img v-if="store.currentWorshipSession.thumbnailURL"
                            :src="store.currentWorshipSession.thumbnailURL"
                            :alt="store.currentWorshipSession.worshipLeader" class="w-full h-full object-cover" />
                        <div v-else class="w-full h-full flex items-center justify-center text-gray-500 text-sm">
                            <PhMusicNote size="150" weight="thin" />
                        </div>
                    </div>
                    <h3 class="font-semibold text-gray-800 text-sm truncate">Worship - {{
                        formatDate(store.currentWorshipSession.date) }}</h3>
                    <p class="text-gray-400 text-xs mt-1">Uploaded on: {{
                        formatDate(store.currentWorshipSession.createdAt) }}</p>
                </div>
                <div class="mb-2">
                    <span class="text-xs font-bold text-gray-400 mb-2">BACKUP
                        SINGERS</span>
                    <ul class="pl-4 list-disc">
                        <li v-for="bs in store.currentWorshipSession.backupSingers" :key="bs.id"
                            class="font-thin font-montserrat text-gray-600 text-sm">
                            {{ bs.name }}
                        </li>
                    </ul>
                </div>

                <div>
                    <div class="flex items-center justify-between mb-2 p-2 bg-blue-50">
                        <span class="text-xs font-bold text-gray-500">WORSHIP SONGS</span>
                        <span class="text-xs font-bold text-gray-500">
                            SONG COUNT: {{ store.currentWorshipSession.worshipSongs.length }}
                        </span>
                    </div>
                    <div class="grid grid-cols-1 gap-6 font-montserrat px-3 sm:grid-cols-2 lg:grid-cols-4">
                        <div v-for="ws in store.currentWorshipSession.worshipSongs" :key="ws.id">
                            <h4 class="flex gap-2 font-black text-gray-600 mb-2">
                                <span>{{ ws.songOrder }}.</span>
                                <span>{{ ws.songTitle }}</span>
                            </h4>

                            <p v-html="parsedLyrics(ws)" class="text-sm text-gray-500" />
                        </div>
                    </div>
                </div>
            </div>

            <p v-else class="text-gray-500 text-sm w-full text-center">No track yet.</p>
        </section>
    </main>
</template>

<script setup lang="ts">
import { PhMusicNote } from "@phosphor-icons/vue"
import { useWorshipSessionService } from "../../services/worship-sessions"
import type { WorshipWithBackupAndSongs, WorshipSong } from "~/types"

const { setLight } = useHeaderTheme()
setLight()

const store = useWorshipSessionStore()
const playerStore = usePlayerStore()
const { fetchWorshipSessionByID } = useWorshipSessionService()
const route = useRoute()
const id = route.params.id as string

const parsedLyrics = computed(() => (worshipSong: WorshipSong) => {
    if (!worshipSong.lyrics) return ""

    return worshipSong.lyrics.replace(/\n/g, "<br />")
})

function formatDate(date: string) {
    return new Date(date).toLocaleDateString('en-NG', {
        month: 'long',
        year: 'numeric',
        day: 'numeric',
    })
}

function handlePlay() {
    playerStore.playWorship(store.currentWorshipSession as WorshipWithBackupAndSongs)
    console.log(playerStore.currentWorshipSession)
    console.log(store.currentWorshipSession)
}

onMounted(() =>
    fetchWorshipSessionByID(id)
)
</script>