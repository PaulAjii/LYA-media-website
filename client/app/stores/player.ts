import { defineStore } from "pinia"
import type { Track, ChoirMinistration } from "../types"
import { ref } from "vue"

export interface PlayerTrack extends Track {
    albumTitle: string
}

export const usePlayerStore = defineStore("player", () => {
    const currentTrack = ref<PlayerTrack | null>(null)
    const currentSong = ref<ChoirMinistration | null>(null)
    const isPlaying = ref(false)
    const currentTime = ref(0)
    const duration = ref(0)

    function play(track: Track, albumTitle: string) {
        const setTrack = {
            albumTitle,
            id: track?.id,
            albumId: track?.albumId,
            title: track?.title,
            trackNumber: track?.trackNumber,
            audioUrl: track?.audioUrl,
            createdAt: track?.createdAt
        } as PlayerTrack
        currentTrack.value = setTrack
        isPlaying.value = true
    }

    function playSong(song: ChoirMinistration) {
        currentSong.value = song
        isPlaying.value = true
    }

    function pause() {
        isPlaying.value = false
    }

    function resume() {
        isPlaying.value = true
    }

    function stop() {
        currentTrack.value = null
        isPlaying.value = false
        currentTime.value = 0
        duration.value = 0
    }

    function stopSong() {
        currentSong.value = null
        isPlaying.value = false
        currentTime.value = 0
        duration.value = 0
    }

    function setCurrentTime(time: number) {
        currentTime.value = time
    }

    function setDuration(d: number) {
        duration.value = d
    }

    return {
        currentTrack,
        currentSong,
        isPlaying,
        currentTime,
        duration,
        play,
        playSong,
        pause,
        resume,
        stop,
        stopSong,
        setCurrentTime,
        setDuration
    }
})