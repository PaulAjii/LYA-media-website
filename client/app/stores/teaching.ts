import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Album, AlbumWithTracks } from '~/types'

export const useTeachingsStore = defineStore('teachings', () => {
  const albums = ref<Album[]>([])
  const currentAlbum = ref<AlbumWithTracks | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  function setAlbums(data: Album[]) { albums.value = data }
  function setCurrentAlbum(data: AlbumWithTracks) { currentAlbum.value = data }
  function setLoading(val: boolean) { loading.value = val }
  function setError(val: string | null) { error.value = val }

  return {
    albums,
    currentAlbum,
    loading,
    error,
    setAlbums,
    setCurrentAlbum,
    setLoading,
    setError
  }
})