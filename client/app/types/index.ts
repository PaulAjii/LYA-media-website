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

export interface WorshipSession {
    id: string
    date: string
    worshipLeader: string
    audioURL: string
    thumbnailURL: string | null
    createdAt: string
}

export interface WorshipBackupSinger {
    id: string
    name: string
}

export interface WorshipSong {
    id: string,
    songTitle: string,
    lyrics: string | null
    songOrder: number
    createdAt: string
}

export interface WorshipWithBackupAndSongs extends WorshipSession {
    backupSingers: WorshipBackupSinger[]
    worshipSongs: WorshipSong[]
}