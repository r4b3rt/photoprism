package viewer

import (
	"time"

	"github.com/photoprism/photoprism/internal/thumb"
)

// Result represents a photo viewer result.
type Result struct {
	UID          string        `json:"UID"`
	Type         string        `json:"Type,omitempty"`
	Title        string        `json:"Title,omitempty"`
	Description  string        `json:"Description,omitempty"`
	Lat          float64       `json:"Lat"`
	Lng          float64       `json:"Lng"`
	TakenAtLocal time.Time     `json:"TakenAtLocal"`
	Favorite     bool          `json:"Favorite"`
	Playable     bool          `json:"Playable"`
	Duration     time.Duration `json:"Duration,omitempty"`
	Width        int           `json:"Width"`
	Height       int           `json:"Height"`
	Hash         string        `json:"Hash"`
	Thumbs       thumb.Public  `json:"Thumbs"`
	DownloadUrl  string        `json:"DownloadUrl,omitempty"`
}

// Results represents a list of viewer search results.
type Results []Result
