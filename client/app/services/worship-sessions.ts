import type { WorshipSession, WorshipWithBackupAndSongs } from "~/types"
import { useWorshipSessionStore } from "~/stores/worship-sessions"

export function useWorshipSessionService() {
    const store = useWorshipSessionStore()
    const config = useRuntimeConfig()

    async function fetchWorshipSessions() {
        if (store.worshipSessions.length > 0) return

        store.setLoading(true)
        store.setError(null)
        try {
            const res = await $fetch<{ data: WorshipSession[] }>(
                `${ config.public.apiUrl }/api/v1/worship-sessions`
            )
            store.setWorshipSessions(res.data)
        } catch(err) {
            store.setError("Could not fetch Worship sessions")
        } finally {
            store.setLoading(false)
        }
    }

    async function fetchWorshipSessionByID(id: string) {
        if (store.currentWorshipSession?.id === id) return

        store.setLoading(true)
        store.setError(null)
        try {
            const res = await $fetch<{ data: WorshipWithBackupAndSongs }>(
                `${ config.public.apiUrl }/api/v1/worship-sessions/${ id }`
            )
            store.setCurrentWorshipSession(res.data)
        } catch(err) {
            store.setError("Could not fetch Worship session")
        } finally {
            store.setLoading(false)
        }
    }

    return {
        fetchWorshipSessions,
        fetchWorshipSessionByID,
    }
}