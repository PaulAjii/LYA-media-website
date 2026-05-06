import { defineStore } from "pinia"
import type { WorshipSession, WorshipWithBackupAndSongs } from "~/types"
import { ref } from "vue"

export const useWorshipSessionStore = defineStore("worship-session", () => {
    const worshipSessions = ref<WorshipSession[]>([])
    const currentWorshipSession = ref<WorshipWithBackupAndSongs | null>(null)
    const loading = ref(false)
    const error = ref<string | null>(null)

    function setWorshipSessions(data: WorshipSession[]) { worshipSessions.value = data }
    function setCurrentWorshipSession(data: WorshipWithBackupAndSongs) { currentWorshipSession.value = data }
    function setLoading(val: boolean) { loading.value = val }
    function setError(val: string | null) { error.value = val }

    return {
        worshipSessions,
        currentWorshipSession,
        loading,
        error,
        setWorshipSessions,
        setCurrentWorshipSession,
        setLoading,
        setError
    }
})