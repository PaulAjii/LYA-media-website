package storage

import (
	"os"
)

type Config struct {
	SupabaseURL      string
	SupabaseKey      string
	SupabaseBucket   string
	MaxFileSize      int64
	AllowedMimeTypes []string
}

func LoadConfig() *Config {
	return &Config{
		SupabaseURL:    os.Getenv("SUPABASE_URL"),
		SupabaseKey:    os.Getenv("SUPABASE_SERVICE_ROLE_KEY"),
		SupabaseBucket: os.Getenv("SUPABASE_BUCKET_NAME"),
		MaxFileSize:    50 * 1024 * 1024, // 50MB default
		AllowedMimeTypes: []string{
			"audio/mpeg",
			"audio/wav",
			"audio/mp4",
			"audio/flac",
			"audio/ogg",
			"audio/webm",
		},
	}
}
