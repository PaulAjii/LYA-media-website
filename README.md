![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![Fiber](https://img.shields.io/badge/Fiber-v3-white?style=for-the-badge&logo=go)
![Supabase](https://img.shields.io/badge/Supabase-Storage-3ECF8E?style=for-the-badge&logo=supabase)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql)

# LYA Media Website

A modern, RESTful media management API built with Go, Fiber, PostgreSQL, and Supabase Storage for managing albums and audio tracks with secure file uploads.

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

The API supports full CRUD operations for albums and tracks, with integrated Supabase Storage for secure and efficient audio file management.

## ✨ Features

### Album Management
- ✅ Create albums with title, date, summary, and thumbnail
- ✅ Retrieve all albums with pagination support
- ✅ Get album by ID with associated tracks
- ✅ Search albums by title
- ✅ Update album information
- ✅ Delete albums (with cascade track deletion)

### Track Management
- ✅ Upload audio files (MP3, WAV, M4A, FLAC, OGG, WEBM) up to 50MB
- ✅ Create tracks with automatic upload to Supabase Storage
- ✅ Retrieve tracks by album ID with ordered track listing
- ✅ Update track metadata (title, track number)
- ✅ Delete tracks with automatic file cleanup
- ✅ Public URL generation for audio streaming

### Technical Features
- ✅ Clean Architecture (Handlers → UseCases → Repositories)
- ✅ PostgreSQL database with connection pooling
- ✅ Supabase Storage integration for audio files
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
│         (Handlers - fiber.Router)      │
└─────────────────┬───────────────────────┘
                  ↓
┌─────────────────────────────────────────┐
│         Use Cases Layer                 │
│      (Business Logic - UseCases)       │
└─────────────────┬───────────────────────┘
                  ↓
┌─────────────────────────────────────────┐
│       Repository Layer                  │
│   (Data Access - Repositories)         │
└─────────────────┬───────────────────────┘
                  ↓
┌─────────────────────────────────────────┐
│      Infrastructure Layer               │
│ (DB, Storage, Config - PostgreSQL,     │
│  Supabase, Environment)                │
└─────────────────────────────────────────┘
```

### Dependency Flow

```
HTTP Request → Handler → UseCase → Repository → Database/Storage
     ↓           ↓         ↓           ↓             ↓
   Router  →  DTO  →  Entity  →   Entity       PostgreSQL
                                    ↓
                              Supabase Storage
```

## 🛠️ Tech Stack

### Core Framework
- **[Go](https://golang.org/)** - Programming language (v1.21+)
- **[Fiber v3](https://gofiber.io/)** - Web framework inspired by Express.js
- **[pgx/v5](https://github.com/jackc/pgx)** - PostgreSQL driver and toolkit

### Database & Storage
- **[PostgreSQL](https://www.postgresql.org/)** - Relational database
- **[Supabase](https://supabase.com/)** - Storage for audio files
- **[UUID](https://github.com/google/uuid)** - UUID generation

### Utilities
- **[godotenv](https://github.com/joho/godotenv)** - Environment configuration
- **[validator](https://github.com/go-playground/validator)** - (Implied validation)

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
│   │   └── usecase_test.go
│   ├── storage/                 # Supabase Storage
│   │   ├── supabase.go         # Storage operations
│   │   └── config.go           # Storage configuration
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
- PostgreSQL 14+
- Supabase account (for storage)
- Git

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
DATABASE_URL=postgresql://username:password@localhost:5432/lyamedia

# Supabase Configuration
SUPABASE_URL=https://your-project.supabase.co
SUPABASE_SERVICE_ROLE_KEY=your-service-role-key-here
SUPABASE_BUCKET_NAME=audio
```

### 4. Database Setup

Create the database tables:

```sql
-- Albums table
CREATE TABLE albums (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    date TIMESTAMP NOT NULL,
    summary TEXT,
    thumbnail_url VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tracks table
CREATE TABLE tracks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    album_id UUID NOT NULL REFERENCES albums(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    track_number INTEGER NOT NULL,
    audio_url VARCHAR(500) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for performance
CREATE INDEX idx_tracks_album_id ON tracks(album_id);
CREATE INDEX idx_tracks_track_number ON tracks(track_number);
CREATE INDEX idx_albums_created_at ON albums(created_at DESC);
```

### 5. Supabase Storage Setup

1. Create a Supabase project at [supabase.com](https://supabase.com)
2. Create a new bucket called `audio`
3. Set bucket permissions to "public" for serving audio files:
   - Go to Storage → Buckets → audio → Edit Permissions
   - Add policy: `create` for authenticated users
   - Add policy: `read` for public access
   - Add policy: `delete` for authenticated users

4. Get your Service Role Key from: Project Settings → API → Service Role Key

### 6. Run the application

```bash
go run cmd/api/main.go
```

Or using Air for live reloading:

```bash
# Install Air
go install github.com/cosmtrek/air@latest

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
  }
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
    }
  ]
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
    "album": {
      "id": "uuid",
      "title": "Summer Vibes",
      "date": "2024-06-01T00:00:00Z",
      "summary": "A collection of summer hits",
      "thumbnail_url": "https://example.com/thumbnail.jpg",
      "created_at": "2024-06-01T00:00:00Z"
    },
    "tracks": [
      {
        "id": "uuid",
        "albumId": "uuid",
        "title": "Track 1",
        "trackNumber": 1,
        "audioUrl": "https://supabase.co/storage/v1/object/public/audio/tracks/track1.mp3",
        "createdAt": "2024-06-01T00:00:00Z"
      }
    ]
  }
}
```

#### Search Albums by Title

```http
GET /api/v1/albums/search/:title

Response: 200 OK
{
  "status": "success",
  "message": "Albums retrieved successfully",
  "data": [...]
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
  "message": "Album updated successfully"
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
    "audioUrl": "https://supabase.co/storage/v1/object/public/audio/tracks/AlbumTitle/Track1.mp3",
    "createdAt": "2024-06-01T00:00:00Z"
  }
}
```

#### Get Tracks by Album

```http
GET /api/v1/tracks/album/:albumID

Response: 200 OK
{
  "status": "success",
  "message": "Tracks retrieved successfully",
  "data": [
    {
      "id": "uuid",
      "albumId": "uuid",
      "title": "Track 1",
      "trackNumber": 1,
      "audioUrl": "https://supabase.co/storage/v1/object/public/audio/tracks/AlbumTitle/Track1.mp3",
      "createdAt": "2024-06-01T00:00:00Z"
    }
  ]
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

### File Storage Organization

Files are stored in Supabase Storage with the following structure:

```
{bucket}/
└── {albumTitle}/
    └── {trackTitle}{extension}

Example:
audio/
└── Summer Vibes/
    └── Track 1.mp3
```

### Upload Process

1. Client sends `multipart/form-data` with track info and audio file
2. Server validates file type and size (< 50MB)
3. Server sanitizes filename and creates unique path
4. Server uploads to Supabase Storage at `/tracks/{albumTitle}/{trackTitle}.mp3`
5. Supabase returns public URL
6. Server stores track metadata (including audio URL) in PostgreSQL
7. Response returns track object with public URL for streaming

## 🐳 Docker Support

### Build and run with Docker

```bash
# Build image
docker build -t lya-media-api .

# Run container
docker run -p 8080:8080 --env-file .env lya-media-api
```

### Docker Compose

```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    env_file:
      - .env
    restart: unless-stopped
```

## 🧪 Testing

Run tests:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

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
- Audio files served directly from Supabase CDN
- Public URLs enable caching and efficient delivery
- 50MB file size limit prevents abuse while supporting high-quality audio

### API
- CORS enabled for frontend applications
- Request logging for monitoring
- Graceful shutdown ensures no request drops

## 🛡️ Security

- UUIDs instead of sequential IDs prevent enumeration
- Supabase service role key securely stored in environment
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
