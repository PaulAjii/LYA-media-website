import { defineStore } from "pinia"
import { ref } from "vue"
import type { ChoirMinistration } from "~/types"

export const useChoirMinistrationsStore = defineStore("choir-ministrations", () => {
    const choirMinistrations = ref<ChoirMinistration[]>([])
    const currentSong = ref<ChoirMinistration | null>(null)
    const loading = ref(false)
    const error = ref<string | null>(null)

    function setChoirMinistrations(data: ChoirMinistration[]) { choirMinistrations.value = data }
    function setCurrentSong(data: ChoirMinistration) { currentSong.value = data }
    function setLoading(val: boolean) { loading.value = val }
    function setError(val: string | null) { error.value = val }

    return {
        choirMinistrations,
        setChoirMinistrations,
        currentSong,
        setCurrentSong,
        loading,
        setLoading,
        error,
        setError
    }
})