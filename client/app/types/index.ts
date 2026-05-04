export interface Album {
    id: string,
    title: string,
    date: string,
    summary: string,
    thumbnail_url: string,
    createdAt: string
}

export interface Track {
    id: string,
    albumId: string,
    title: string,
    trackNumber: number,
    audioUrl: string,
    createdAt: string
}

export interface AlbumWithTracks extends Album {
    tracks: Track[]
}

export interface ChoirMinistration {
  id: string
  songTitle: string
  date: string
  songWriter: string | null
  lyrics: string | null
  audioURL: string
  thumbnailURL: string | null
  createdAt: string
}
