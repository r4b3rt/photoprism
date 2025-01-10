package search

import (
	"encoding/json"

	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/entity/search/viewer"
	"github.com/photoprism/photoprism/internal/form"
	"github.com/photoprism/photoprism/internal/thumb"
)

// PhotosViewerResults finds photos based on the search form provided and returns them as viewer.Results.
func PhotosViewerResults(f form.SearchPhotos, contentUri, apiUri, previewToken, downloadToken string) (viewer.Results, int, error) {
	return UserPhotosViewerResults(f, nil, contentUri, apiUri, previewToken, downloadToken)
}

// UserPhotosViewerResults finds photos based on the search form and user session and returns them as viewer.Results.
func UserPhotosViewerResults(f form.SearchPhotos, sess *entity.Session, contentUri, apiUri, previewToken, downloadToken string) (viewer.Results, int, error) {
	if results, count, err := searchPhotos(f, sess, PhotosColsView); err != nil {
		return viewer.Results{}, count, err
	} else {
		return results.ViewerResults(contentUri, apiUri, previewToken, downloadToken), count, err
	}
}

// ViewerResult returns a new photo viewer result.
func (m Photo) ViewerResult(contentUri, apiUri, previewToken, downloadToken string) viewer.Result {
	return viewer.Result{
		UID:          m.PhotoUID,
		Type:         m.PhotoType,
		Title:        m.PhotoTitle,
		Description:  m.PhotoDescription,
		Lat:          m.PhotoLat,
		Lng:          m.PhotoLng,
		TakenAtLocal: m.TakenAtLocal,
		Favorite:     m.PhotoFavorite,
		Playable:     m.IsPlayable(),
		Duration:     m.PhotoDuration,
		Width:        m.FileWidth,
		Height:       m.FileHeight,
		Hash:         m.FileHash,
		Thumbs: thumb.Public{
			Fit720:  thumb.New(m.FileWidth, m.FileHeight, m.FileHash, thumb.Sizes[thumb.Fit720], contentUri, previewToken),
			Fit1280: thumb.New(m.FileWidth, m.FileHeight, m.FileHash, thumb.Sizes[thumb.Fit1280], contentUri, previewToken),
			Fit1920: thumb.New(m.FileWidth, m.FileHeight, m.FileHash, thumb.Sizes[thumb.Fit1920], contentUri, previewToken),
			Fit2560: thumb.New(m.FileWidth, m.FileHeight, m.FileHash, thumb.Sizes[thumb.Fit2560], contentUri, previewToken),
			Fit4096: thumb.New(m.FileWidth, m.FileHeight, m.FileHash, thumb.Sizes[thumb.Fit4096], contentUri, previewToken),
			Fit7680: thumb.New(m.FileWidth, m.FileHeight, m.FileHash, thumb.Sizes[thumb.Fit7680], contentUri, previewToken),
		},
		DownloadUrl: viewer.DownloadUrl(m.FileHash, apiUri, downloadToken),
	}
}

// ViewerJSON returns the results as photo viewer JSON.
func (m PhotoResults) ViewerJSON(contentUri, apiUri, previewToken, downloadToken string) ([]byte, error) {
	return json.Marshal(m.ViewerResults(contentUri, apiUri, previewToken, downloadToken))
}

// ViewerResults returns the results photo viewer formatted.
func (m PhotoResults) ViewerResults(contentUri, apiUri, previewToken, downloadToken string) (results viewer.Results) {
	results = make(viewer.Results, 0, len(m))

	for _, p := range m {
		results = append(results, p.ViewerResult(contentUri, apiUri, previewToken, downloadToken))
	}

	return results
}

// ViewerResult creates a new photo viewer result.
func (m GeoResult) ViewerResult(contentUri, apiUri, previewToken, downloadToken string) viewer.Result {
	return viewer.Result{
		UID:          m.PhotoUID,
		Type:         m.PhotoType,
		Title:        m.PhotoTitle,
		Description:  m.PhotoDescription,
		Lat:          m.PhotoLat,
		Lng:          m.PhotoLng,
		TakenAtLocal: m.TakenAtLocal,
		Favorite:     m.PhotoFavorite,
		Playable:     m.IsPlayable(),
		Duration:     m.PhotoDuration,
		Width:        m.FileWidth,
		Height:       m.FileHeight,
		Hash:         m.FileHash,
		Thumbs: thumb.Public{
			Fit720:  thumb.New(m.FileWidth, m.FileHeight, m.FileHash, thumb.Sizes[thumb.Fit720], contentUri, previewToken),
			Fit1280: thumb.New(m.FileWidth, m.FileHeight, m.FileHash, thumb.Sizes[thumb.Fit1280], contentUri, previewToken),
			Fit1920: thumb.New(m.FileWidth, m.FileHeight, m.FileHash, thumb.Sizes[thumb.Fit1920], contentUri, previewToken),
			Fit2560: thumb.New(m.FileWidth, m.FileHeight, m.FileHash, thumb.Sizes[thumb.Fit2560], contentUri, previewToken),
			Fit4096: thumb.New(m.FileWidth, m.FileHeight, m.FileHash, thumb.Sizes[thumb.Fit4096], contentUri, previewToken),
			Fit7680: thumb.New(m.FileWidth, m.FileHeight, m.FileHash, thumb.Sizes[thumb.Fit7680], contentUri, previewToken),
		},
		DownloadUrl: viewer.DownloadUrl(m.FileHash, apiUri, downloadToken),
	}
}

// ViewerJSON returns the results as photo viewer JSON.
func (photos GeoResults) ViewerJSON(contentUri, apiUri, previewToken, downloadToken string) ([]byte, error) {
	results := make(viewer.Results, 0, len(photos))

	for _, p := range photos {
		results = append(results, p.ViewerResult(contentUri, apiUri, previewToken, downloadToken))
	}

	return json.Marshal(results)
}
