package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	supabase "github.com/supabase-community/supabase-go"
)

type SupabaseStorage struct {
	client     *supabase.Client
	bucketName string
}

func NewSupabaseStorage(supabaseURL, supabaseKey, bucketName string) (*SupabaseStorage, error) {
	client, err := supabase.NewClient(supabaseURL, supabaseKey, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create supabase client: %w", err)
	}

	return &SupabaseStorage{
		client:     client,
		bucketName: bucketName,
	}, nil
}

// UploadFile uploads a file to Supabase Storage and returns the public URL
func (s *SupabaseStorage) UploadFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, folder, title string) (string, error) {
	// Validate file type
	if !isAllowedAudioType(fileHeader.Header.Get("Content-Type")) {
		return "", fmt.Errorf("invalid file type. Only MP3, WAV, M4A, FLAC are allowed")
	}

	// Validate file size (max 50MB)
	const maxFileSize = 50 * 1024 * 1024 // 50MB
	if fileHeader.Size > maxFileSize {
		return "", fmt.Errorf("file too large. Maximum size is 50MB")
	}

	// Create unique filename
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	filename := fmt.Sprintf("%s/%s%s", folder, title, ext)

	// Upload to Supabase
	_, err := s.client.Storage.UploadFile(s.bucketName, filename, file)
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	// Get public URL
	publicURL := s.client.Storage.GetPublicUrl(s.bucketName, filename)

	return publicURL.SignedURL, nil
}

// DeleteFile deletes a file from Supabase Storage
func (s *SupabaseStorage) DeleteFile(ctx context.Context, fileURL string) error {
	// Extract filename from URL
	filename := extractFilenameFromURL(fileURL)
	if filename == "" {
		return fmt.Errorf("invalid file URL")
	}

	_, err := s.client.Storage.RemoveFile(s.bucketName, []string{filename})
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

func isAllowedAudioType(contentType string) bool {
	allowedTypes := map[string]bool{
		"audio/mpeg":   true, // MP3
		"audio/wav":    true, // WAV
		"audio/x-wav":  true, // WAV
		"audio/mp4":    true, // M4A
		"audio/x-m4a":  true, // M4A
		"audio/flac":   true, // FLAC
		"audio/x-flac": true, // FLAC
		"audio/ogg":    true, // OGG
		"audio/webm":   true, // WEBM
	}
	return allowedTypes[contentType]
}

func extractFilenameFromURL(url string) string {
	// Extract filename from full URL
	// Example: https://ywtantwnktrzsrukoorm.supabase.co/storage/v1/object/public/audio/tracks/123.mp3
	parts := strings.Split(url, "/")
	if len(parts) < 2 {
		return ""
	}
	return parts[len(parts)-2] + "/" + parts[len(parts)-1]
}
