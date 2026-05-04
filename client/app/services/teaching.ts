import type { Album, AlbumWithTracks } from '~/types'
import { useTeachingsStore } from '../stores/teaching'

export function useTeachingsService() {
  const store = useTeachingsStore()
  const config = useRuntimeConfig()

  async function fetchAlbums() {
    if (store.albums.length > 0) return

    store.setLoading(true)
    store.setError(null)
    try {
      const res = await $fetch<{ data: Album[] }>(
        `${ config.public.apiUrl }/api/v1/albums`
      )
      store.setAlbums(res.data)
    } catch (e) {
      store.setError('Could not load teachings')
    } finally {
      store.setLoading(false)
    }
  }

  async function fetchAlbumByID(id: string) {
    if (store.currentAlbum?.id === id) return

    store.setLoading(true)
    store.setError(null)
    try {
      const res = await $fetch<{ data: AlbumWithTracks }>(
        `${ config.public.apiUrl }/api/v1/albums/${ id }`
      )
      store.setCurrentAlbum(res.data)
    } catch (e) {
      store.setError('Could not load teaching')
    } finally {
      store.setLoading(false)
    }
  }

  return {
    fetchAlbums,
    fetchAlbumByID
  }
}