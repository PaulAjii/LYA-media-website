![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![Fiber](https://img.shields.io/badge/Fiber-v3-white?style=for-the-badge&logo=go)
![R2 Storage](https://img.shields.io/badge/R2-Storage-orange)
[![Supabase](https://img.shields.io/badge/Supabase-3ECF8E?style=for-the-badge&logo=supabase&logoColor=white)](https://supabase.com/)

# LYA Media Website

A modern, RESTful media management API built with Go, Fiber, PostgreSQL, and Cloudflare R2 Storage for managing church teachings, choir ministrations and worship sessions.

## 📋 Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Architecture](#architecture)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [API Documentation](#api-documentation)
- [Database Schema](#database-schema)
- [Contributing](#contributing)
- [License](#license)

## 🎯 Overview

LYA Media Website is a backend API service that provides comprehensive media management capabilities for audio content. It follows clean architecture principles with clear separation of concerns, making it maintainable, testable, and scalable.

The API supports full CRUD operations for sermons in albums and tracks segregations, choir ministrations, and worship sessions with integrated Cloudflare R2 Storage for secure and efficient audio file management.

## ✨ Features

### Sermon Recordings

#### Album Management
- ✅ Create albums with title, date, summary, and thumbnail
- ✅ Retrieve all albums with pagination support
- ✅ Get album by ID with associated tracks
- ✅ Search albums by title
- ✅ Update album information
- ✅ Delete albums (with cascade track deletion)

### #Track Management
- ✅ Upload audio files (MP3, WAV, M4A, FLAC, OGG, WEBM) up to 50MB
- ✅ Create tracks with automatic upload to Supabase Storage
- ✅ Retrieve tracks by album ID with ordered track listing
- ✅ Update track metadata (title, track number)
- ✅ Delete tracks with automatic file cleanup
- ✅ Public URL generation for audio streaming

### Technical Features
- ✅ Clean Architecture (Handlers → UseCases → Repositories)
- ✅ PostgreSQL database with connection pooling
- ✅ Cloudflare R2 Storage integration for audio files
- ✅ UUID primary keys for security
- ✅ CORS support for web clients
- ✅ Request logging and error handling
- ✅ Graceful server shutdown
- ✅ Health check endpoint
- ✅ Environment-based configuration

## 🏗️ Architecture

This project implements **Clean Architecture** principles:

```
┌─────────────────────────────────────────┐
│              API Layer                  │
│         (Handlers - fiber.Router)       │
└─────────────────┬───────────────────────┘
                  ↓
┌─────────────────────────────────────────┐
│         Use Cases Layer                 │
│      (Business Logic - UseCases)        │
└─────────────────┬───────────────────────┘
                  ↓
┌─────────────────────────────────────────┐
│       Repository Layer                  │
│   (Data Access - Repositories)          │
└─────────────────┬───────────────────────┘
                  ↓
┌─────────────────────────────────────────┐
│      Infrastructure Layer               │
│ (DB, Storage, │ Supabase, Environment)  │
└─────────────────────────────────────────┘
```

### Dependency Flow

```
HTTP Request → Handler → UseCase → Repository → Database/Storage
     ↓           ↓         ↓           ↓             ↓
   Router  →  DTO  →  Entity  →   Entity       PostgreSQL
                                    ↓
                            Cloudflare R2 Storage
```

## 🛠️ Tech Stack

### Core Framework
- **[Go](https://golang.org/)** - Programming language (v1.21+)
- **[Fiber v3](https://gofiber.io/)** - Web framework inspired by Express.js
- **[pgx/v5](https://github.com/jackc/pgx)** - PostgreSQL driver and toolkit

### Database & Storage
- **[Supabase](https://supabase.com/)** - Database
- **[Cloudflare R2](https://www.cloudflare.com/en-gb/developer-platform/products/r2/)** - Storage for audio media
- **[UUID](https://github.com/google/uuid)** - UUID generation

### Utilities
- **[godotenv](https://github.com/joho/godotenv)** - Environment configuration

### Development
- Built for production with connection pooling
- CORS enabled for cross-origin requests
- Structured logging
- Graceful shutdown

## 📁 Project Structure

```
LYA-media-website/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── albums/                  # Album domain
│   │   ├── dtos/                # Data Transfer Objects
│   │   │   └── albums_dto.go
│   │   ├── handler.go          # HTTP handlers
│   │   ├── repository.go       # Data access layer
│   │   ├── router.go           # Route definitions
│   │   └── usecase.go          # Business logic
│   ├── tracks/                  # Track domain
│   │   ├── dtos/
│   │   │   └── tracks_dtos.go
│   │   ├── handler.go
│   │   ├── repository.go
│   │   ├── router.go
│   │   ├── usecase.go
│   ├── choir_ministration/                  # Choir ministration domain
│   │   ├── dtos/
│   │   │   └── ministration_dtos.go
│   │   ├── handler.go
│   │   ├── repository.go
│   │   ├── router.go
│   │   ├── usecase.go
│   ├── storage/                 # R2 Storage
│   │   ├── r2.go        # Storage config and handler methods
│   └── api/
│       └── setup_routes.go      # Route initialization
├── pkg/
│   ├── db/
│   │   └── postgres.go         # Database connection
│   └── response/
│       └── response.go         # Response helpers
├── .env                        # Environment variables
├── go.mod                      # Dependencies
├── LICENSE
└── README.md
```

## 🔧 Prerequisites

- Go 1.21 or higher
- Supabase account (for cloud database)
- Git
- Cloudflare R2 account (for storage)

## ⚙️ Installation

### 1. Clone the repository

```bash
git clone https://github.com/PaulAjii/LYA-media-website.git
cd LYA-media-website
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Set up environment variables

Create a `.env` file in the project root:

```bash
cp .env.example .env  # If you have an example file
```

Or create manually:

```env
# Port Configuration
PORT=8080

# Database Configuration
DATABASE_PASSWORD=your-db-password
DATABASE_URL=postgresql://postgres.projectID:${DATABASE_PASSWORD}@db-host:db-port/postgres

# R2 CONFIGURATION
R2_ACCOUNT_ID=R2_ACCOUNT_ID
R2_ACCESS_KEY_ID=R2_ACCESS_KEY_ID
R2_SECRET_ACCESS_KEY=R2_SECRET_ACCESS_KEY
R2_BUCKET_NAME=R2_BUCKET_NAME
R2_PUBLIC_URL=R2_PUBLIC_URL
```

### 4. Database Setup

Create the database tables:

```sql
-- Albums Table
CREATE TABLE albums(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title TEXT NOT NULL,
  date DATE NOT NULL,
  summary TEXT,
  thumbnail_url TEXT,
  created_at TIMESTAMPTZ DEFAULT NOW()
);
CREATE INDEX idx_album_title ON albums(title);
CREATE INDEX idx_album_date ON albums(date);

-- Tracks Table
CREATE TABLE tracks(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  album_id UUID NOT NULL REFERENCES albums(id) ON DELETE CASCADE,
  title TEXT NOT NULL,
  track_number INT NOT NULL,
  audio_url TEXT NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW()
);
CREATE INDEX idx_album_id ON tracks(album_id);

-- Choir Ministrations table
CREATE TABLE choir_ministrations(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  song_title TEXT NOT NULL,
  date DATE NOT NULL,
  songwriter TEXT,
  lyrics TEXT,
  audio_url TEXT NOT NULL,
  thumbnail_url TEXT,
  created_at TIMESTAMPTZ DEFAULT NOW()
);
CREATE INDEX idx_choir_date ON choir_ministrations(date);
CREATE INDEX idx_choir_song ON choir_ministrations(song_title);

-- Worship Sessions table
CREATE TABLE worship_sessions(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  date DATE NOT NULL,
  worship_leader TEXT NOT NULL,
  thumbnail_url TEXT,
  created_at TIMESTAMPTZ DEFAULT NOW()
);
CREATE INDEX idx_worship_leader ON worship_sessions(worship_leader);
CREATE INDEX idx_worship_date ON worship_sessions(date);

-- Worship Songs Table
CREATE TABLE worship_songs(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  session_id UUID NOT NULL REFERENCES worship_sessions(id) ON DELETE CASCADE,
  song_title TEXT NOT NULL,
  lyrics TEXT,
  song_order INT NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW()
);

-- Worship Backup Singers Table
CREATE TABLE worship_backup_singers(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  session_id UUID NOT NULL REFERENCES worship_sessions(id) ON DELETE CASCADE,
  name TEXT NOT NULL
);
```

### 5. Cloudflare R2 Storage Setup

### 6. Run the application

```bash
go run cmd/api/main.go
```

Or using Air for live reloading:

```bash
# Install Air
go install github.com/air-verse/air@latest

# Run with Air
air
```

The server will start on `http://localhost:8080`

## 📖 API Documentation

### Base URL
```
http://localhost:8080/api/v1
```

### Health Check

```http
GET /api/health

Response:
{
  "status": "ok",
  "message": "API is healthy and running"
}
```

### Albums

#### Create Album

```http
POST /api/v1/albums
Content-Type: application/json

{
  "title": "Summer Vibes",
  "date": "2024-06-01T00:00:00Z",
  "summary": "A collection of summer hits",
  "thumbnail_url": "https://example.com/thumbnail.jpg"
}

Response: 201 Created
{
  "status": "success",
  "message": "Album created successfully",
  "data": {
    "id": "uuid",
    "title": "Summer Vibes",
    "date": "2024-06-01T00:00:00Z",
    "summary": "A collection of summer hits",
    "thumbnail_url": "https://example.com/thumbnail.jpg",
    "created_at": "2024-06-01T00:00:00Z"
  },
  "statusCode": 201,
}
```

#### Get All Albums

```http
GET /api/v1/albums

Response: 200 OK
{
  "status": "success",
  "message": "Retrieved albums successfully",
  "data": [
    {
      "id": "uuid",
      "title": "Summer Vibes",
      "date": "2024-06-01T00:00:00Z",
      "summary": "A collection of summer hits",
      "thumbnail_url": "https://example.com/thumbnail.jpg",
      "created_at": "2024-06-01T00:00:00Z"
    },
    {...},
  ],
  "statusCode": 200,
}
```

#### Get Album by ID (with Tracks)

```http
GET /api/v1/albums/:albumID

Response: 200 OK
{
  "status": "success",
  "message": "Album retrieved successfully",
  "data": {
    "id": "uuid",
    "title": "Summer Vibes",
    "date": "2024-06-01T00:00:00Z",
    "summary": "A collection of summer hits",
    "thumbnail_url": "https://example.com/thumbnail.jpg",
    "created_at": "2024-06-01T00:00:00Z",
    "tracks": [
      {
        "id": "uuid",
        "albumId": "uuid",
        "title": "Track 1",
        "trackNumber": 1,
        "audioUrl": "https://supabase.co/storage/v1/object/public/audio/tracks/track1.mp3",
        "createdAt": "2024-06-01T00:00:00Z"
      },
      {...},
    ],
  },
  "statusCode": 200,
}
```

#### Search Albums by Title

```http
GET /api/v1/albums/search?title

Response: 200 OK
{
  "status": "success",
  "message": "Albums retrieved successfully",
  "data": [...],
  "statusCode": 200,
}
```

#### Update Album

```http
PATCH /api/v1/albums/:albumID
Content-Type: application/json

{
  "title": "Updated Title",
  "summary": "Updated summary"
}

Response: 200 OK
{
  "status": "success",
  "message": "Album updated successfully",
  "data": {
    "id": "uuid",
    "title": "Updated Title",
    "date": "0001-01-01T00:00:00Z",
    "summary": "Updated summary",
    "thumbnail_url": null,
    "createdAt": "0001-01-01T00:00:00Z"
  },
  "statusCode": 200
}
```

#### Delete Album

```http
DELETE /api/v1/albums/:albumID

Response: 204 No Content
```

### Tracks

#### Create Track (with Audio Upload)

```http
POST /api/v1/tracks
Content-Type: multipart/form-data

Form Fields:
- albumId: {UUID} (required)
- albumTitle: {String} (required)
- title: {String} (required)
- trackNumber: {Integer} (optional, default: 1)
- audio: {File} (required, MP3/WAV/M4A/FLAC/OGG/WEBM, max 50MB)

Response: 201 Created
{
  "status": "success",
  "message": "Track created successfully",
  "data": {
    "id": "uuid",
    "albumId": "uuid",
    "title": "Track 1",
    "trackNumber": 1,
    "audioUrl": "https://public-url.r2.dev/tracks/Album Title/Album Title - Track 1 - Track Title.mp3",
    "createdAt": "0001-00-00T00:00:00Z"
  },
  "statusCode": 201,
}
```

#### Get Tracks by Album

```http
GET /api/v1/tracks/album/:albumID

Response: 200 OK
{
  "status": "success",
  "message": "Tracks fetched successfully",
  "data": [
    {
      "id": "uuid",
      "albumId": "uuid",
      "title": "Track 1",
      "trackNumber": 1,
      "audioUrl": "https://public-url.r2.dev/tracks/Album Title/Album Title - Track 1 - Track Title.mp3",
      "createdAt": "0001-00-00T00:00:00Z"
    },
    {...},
  ],
  "statusCode": 200,
}
```

#### Update Track

```http
PATCH /api/v1/tracks/:trackID
Content-Type: application/json

{
  "title": "Updated Track Title",
  "trackNumber": 2
}

Response: 200 OK
{
  "status": "success",
  "message": "Tracked updated successfully"
}
```

#### Delete Track

```http
DELETE /api/v1/tracks/:trackID

Response: 204 No Content
```

## 💾 Database Schema

### Albums Table

```sql
CREATE TABLE albums (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title         VARCHAR(255) NOT NULL,
    date          TIMESTAMP NOT NULL,
    summary       TEXT,
    thumbnail_url VARCHAR(500),
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**Indexes:**
- `idx_albums_created_at` - For sorting by creation date

### Tracks Table

```sql
CREATE TABLE tracks (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    album_id      UUID NOT NULL REFERENCES albums(id) ON DELETE CASCADE,
    title         VARCHAR(255) NOT NULL,
    track_number  INTEGER NOT NULL,
    audio_url     VARCHAR(500) NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**Indexes:**
- `idx_tracks_album_id` - For efficient track lookups by album
- `idx_tracks_track_number` - For sorting tracks

**Foreign Key:**
- `album_id` references `albums(id)` with cascade delete

## 🎵 Media Upload Details

### Supported Audio Formats

- **MP3** (`audio/mpeg`) - Primary format, best compression
- **WAV** (`audio/wav`, `audio/x-wav`) - Uncompressed, high quality
- **M4A** (`audio/mp4`, `audio/x-m4a`) - AAC compression, good quality
- **FLAC** (`audio/flac`, `audio/x-flac`) - Lossless compression
- **OGG** (`audio/ogg`) - Open format, good compression
- **WEBM** (`audio/webm`) - Web-optimized format

### Upload Process

1. Client sends `multipart/form-data` with track info and audio file
2. Server validates file type and size (< 50MB)
3. Server sanitizes filename and creates unique path
4. Server uploads to Cloudflare R2 Storage at `/tracks/{albumTitle}/{trackTitle}.mp3`
5. R2 returns public URL
6. Server stores track metadata (including audio URL) in Supabase PostgreSQL database
7. Response returns track object with public URL for streaming

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch: `git checkout -b feature/amazing-feature`
3. Commit your changes: `git commit -m 'Add amazing feature'`
4. Push to the branch: `git push origin feature/amazing-feature`
5. Open a Pull Request

### Development Guidelines

- Follow Go naming conventions
- Use meaningful variable and function names
- Write tests for new features
- Update documentation
- Follow clean architecture principles:
  - Handlers only handle HTTP (no business logic)
  - UseCases contain all business rules
  - Repositories handle data access only

## 📊 Performance Considerations

### Database
- Connection pooling configured for optimal performance
- Indexed foreign keys for efficient joins
- Indexed timestamps for chronological queries

### Storage
- Audio files served directly from Cloudflare R2 CDN
- Public URLs enable caching and efficient delivery
- 50MB file size limit prevents abuse while supporting high-quality audio

### API
- CORS enabled for frontend applications
- Request logging for monitoring
- Graceful shutdown ensures no request drops

## 🛡️ Security

- UUIDs instead of sequential IDs prevent enumeration
- File type validation prevents malicious uploads
- CORS configured for specific origins
- File size limits prevent DoS attacks

## 📚 API Examples

### Create Album with cURL

```bash
curl -X POST http://localhost:8080/api/v1/albums \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Summer Playlist",
    "date": "2024-06-01T00:00:00Z",
    "summary": "Best songs of summer 2024"
  }'
```

### Upload Track with cURL

```bash
curl -X POST http://localhost:8080/api/v1/tracks \
  -F "albumId=uuid-here" \
  -F "albumTitle=Summer Playlist" \
  -F "title=Track 1" \
  -F "trackNumber=1" \
  -F "audio=@/path/to/audio/file.mp3"
```

### Get Album with Tracks

```bash
curl http://localhost:8080/api/v1/albums/uuid-here
```

## 🐛 Troubleshooting

### Database Connection Issues
- Ensure DATABASE_URL is correctly formatted
- Verify PostgreSQL is running and accessible
- Check connection parameters in .env file

### Supabase Storage Upload Fails
- Verify SUPABASE_URL and SUPABASE_SERVICE_ROLE_KEY
- Ensure bucket exists and has correct permissions
- Check file type and size constraints
- Ensure albumTitle doesn't contain special characters that break paths

### CORS Issues
- Update allowed origins in .env or main.go
- Ensure frontend origin is included in CORS configuration

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Support

For support, please open an issue in the GitHub repository or contact the development team.

---

**Built with ❤️ by Paul Ajii**

Currently maintained and actively developed.
