package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strings"
)

// generateAlbumCacheKey hashes artist+album for on-disk artwork filenames.
func generateAlbumCacheKey(artist, album string) string {
	key := fmt.Sprintf("%s|%s", strings.ToLower(artist), strings.ToLower(album))
	hash := md5.Sum([]byte(key))
	return fmt.Sprintf("%x", hash)
}

// saveAlbumArtwork writes embedded art to disk and frees ArtData from memory.
func (app *MiyooPod) saveAlbumArtwork(album *Album) error {
	if album == nil || album.ArtData == nil || len(album.ArtData) == 0 {
		return fmt.Errorf("no artwork data to save")
	}

	if err := os.MkdirAll(ARTWORK_DIR, 0755); err != nil {
		return fmt.Errorf("failed to create artwork directory: %v", err)
	}

	hash := generateAlbumCacheKey(album.Artist, album.Name)
	filename := fmt.Sprintf("%s.%s", hash, album.ArtExt)
	filepath := ARTWORK_DIR + filename

	if err := os.WriteFile(filepath, album.ArtData, 0644); err != nil {
		return fmt.Errorf("failed to write artwork file: %v", err)
	}

	album.ArtPath = filepath
	album.ArtData = nil

	logMsg(fmt.Sprintf("[ARTWORK] Saved to disk: %s", filepath))
	return nil
}

// loadAlbumArtwork reads previously saved album artwork from disk.
func (app *MiyooPod) loadAlbumArtwork(album *Album) error {
	if album == nil || album.ArtPath == "" {
		return fmt.Errorf("no artwork path set")
	}

	if _, err := os.Stat(album.ArtPath); os.IsNotExist(err) {
		return fmt.Errorf("artwork file not found: %s", album.ArtPath)
	}

	artData, err := os.ReadFile(album.ArtPath)
	if err != nil {
		return fmt.Errorf("failed to read artwork file: %v", err)
	}

	if strings.HasSuffix(album.ArtPath, ".jpg") || strings.HasSuffix(album.ArtPath, ".jpeg") {
		album.ArtExt = "jpg"
	} else if strings.HasSuffix(album.ArtPath, ".png") {
		album.ArtExt = "png"
	}

	album.ArtData = artData
	return nil
}
