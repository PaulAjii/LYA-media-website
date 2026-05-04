<template>
    <div v-if="player.currentTrack || player.currentSong"
        class="fixed bottom-0 left-0 right-0 bg-gray-900 text-white px-4 py-3 flex flex-col gap-2 shadow-lg z-50">
        <!-- Track info -->
        <div class="flex items-center justify-between">
            <div v-if="player.currentTrack" class="flex flex-col">
                <span class=" text-sm font-semibold truncate max-w-xs">{{ player.currentTrack.title }}</span>
                <span class="text-xs text-gray-400 truncate max-w-xs">{{ player.currentTrack.albumTitle }}</span>
            </div>
            <div v-else class="flex flex-col">
                <span class=" text-sm font-semibold truncate max-w-xs">{{ player.currentSong!.songTitle }}</span>
                <span class="text-xs text-gray-400 truncate max-w-xs">{{ player.currentSong!.songWriter }}</span>
            </div>

            <div class="flex items-center gap-4">
                <!-- Play/Pause -->
                <button @click="togglePlay" class="text-gray-400 hover:text-white focus:outline-none cursor-pointer">
                    <span title="Pause" v-if="player.isPlaying" class="text-2xl">
                        <PhPause size="25" />
                    </span>
                    <span title="Play" v-else class="text-2xl">
                        <PhPlay size="25" />
                    </span>
                </button>

                <!-- Stop -->
                <button v-if="player.currentTrack" @click="player.stop()" title="Stop"
                    class="text-gray-400 hover:text-white focus:outline-none cursor-pointer">
                    <span class="text-xl">
                        <PhStop size="25" />
                    </span>
                </button>
                <button v-else @click="player.stopSong()" title="Stop"
                    class="text-gray-400 hover:text-white focus:outline-none cursor-pointer">
                    <span class="text-xl">
                        <PhStop size="25" />
                    </span>
                </button>

                <!-- Download -->
                <button @click="handleDownload" title="Download"
                    class="text-gray-400 hover:text-white focus:outline-none cursor-pointer">
                    <span class="text-xl">
                        <PhDownloadSimple size="25" />
                    </span>
                </button>
            </div>
        </div>

        <!-- Progress bar -->
        <div class="flex items-center gap-2">
            <span title="Current Time" class="text-xs text-gray-400">{{ formattedCurrentTime }}</span>
            <input type="range" :max="player.duration" :value="player.currentTime" @input="seek"
                class="flex-1 h-1 accent-white cursor-pointer" />
            <span title="Track Duration" class="text-xs text-gray-400">{{ formattedDuration }}</span>
        </div>

        <!-- Hidden audio element -->
        <audio ref="audioRef" @timeupdate="onTimeUpdate" @loadedmetadata="onLoadedMetadata" @ended="onEnded"
            class="hidden" />
    </div>
</template>

<script setup lang="ts">
import { ref, watch, computed } from "vue"
import { usePlayerStore } from "../stores/player"
import { PhDownloadSimple, PhStop, PhPlay, PhPause } from "@phosphor-icons/vue"

const player = usePlayerStore()
const audioRef = ref<HTMLAudioElement | null>(null)

watch(() => player.currentTrack, (track) => {
    if (!track) return
    nextTick(() => {
        if (audioRef.value) {
            audioRef.value.src = track.audioUrl
            audioRef.value.play()
        }
    })
})

watch(() => player.currentSong, (song) => {
    if (!song) return
    nextTick(() => {
        if (audioRef.value) {
            audioRef.value.src = song.audioURL
            audioRef.value.play()
        }
    })
})

watch(() => player.isPlaying, (playing) => {
    if (!audioRef.value) return
    if (playing) {
        audioRef.value.play()
    } else {
        audioRef.value.pause()
    }
})

function onTimeUpdate() {
    if (audioRef.value) {
        player.setCurrentTime(audioRef.value.currentTime)
    }
}

function onLoadedMetadata() {
    if (audioRef.value) {
        player.setDuration(audioRef.value.duration)
    }
}

function onEnded() {
    player.pause()
}

function togglePlay() {
    if (player.isPlaying) {
        player.pause()
    } else {
        player.resume()
    }
}

function seek(e: Event) {
    const input = e.target as HTMLInputElement
    if (audioRef.value) {
        audioRef.value.currentTime = Number(input.value)
        player.setCurrentTime(Number(input.value))
    }
}

async function download() {
    if (!player.currentTrack) return
    const res = await fetch(player.currentTrack?.audioUrl)
    const blob = await res.blob()
    const blobURL = URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = blobURL
    a.download = player.currentTrack.title
    a.target = "_blank"
    a.click()
}

async function downloadSong() {
    if (!player.currentSong) return
    const res = await fetch(player.currentSong?.audioURL)
    const blob = await res.blob()
    const blobURL = URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = blobURL
    a.download = player.currentSong.songTitle
    a.target = "_blank"
    a.click()
}

async function handleDownload() {
    if (player.currentTrack) {
        await download()
    } else {
        await downloadSong()
    }
}


const formattedCurrentTime = computed(() => formatTime(player.currentTime))
const formattedDuration = computed(() => formatTime(player.duration))

function formatTime(seconds: number) {
    const m = Math.floor(seconds / 60)
    const s = Math.floor(seconds % 60)
    return `${m}:${s.toString().padStart(2, '0')}`
}
</script>