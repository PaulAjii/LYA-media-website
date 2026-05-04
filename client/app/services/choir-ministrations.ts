import type { ChoirMinistration } from "~/types"
import { useChoirMinistrationsStore } from "~/stores/choir-ministrations"

export function useChoirMinistrationsService() {
    const store = useChoirMinistrationsStore()
    const config = useRuntimeConfig()

    async function fetchChoirMinistrations() {
        if (store.choirMinistrations.length > 0) return

        store.setLoading(true)
        store.setError(null)

        try {
            const res = await $fetch<{data: { pagedData: ChoirMinistration[]}}>(
                `${ config.public.apiUrl }/api/v1/choir-ministrations`
            )
            store.setChoirMinistrations(res.data.pagedData)
        } catch (err) {
            store.setError("Could not load choir minstrations")
        } finally {
            store.setLoading(false)
        }
    }

    async function fetchChoirMinistrationById(id: string) {
        if (store.currentSong?.id === id) return

        store.setLoading(true)
        store.setError(null)
        try {
            const res = await $fetch<{ data: ChoirMinistration }>(
                `${ config.public.apiUrl }/api/v1/choir-ministrations/${ id }`
            )
            store.setCurrentSong(res.data)
        } catch(err) {
            store.setError("Could not load song")
        } finally {
            store.setLoading(false)
        }
    }

    return {
        fetchChoirMinistrations,
        fetchChoirMinistrationById
    }
}