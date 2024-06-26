package exif_utils

import (
	"errors"
	"fmt"
	"github.com/barasher/go-exiftool"
	"github.com/sincerefly/capybara/utils/fileitem"
)

type ExifClient struct {
	etClient *exiftool.Exiftool
}

func NewExifClient() (*ExifClient, error) {
	client, err := exiftool.NewExiftool()
	if err != nil {
		return nil, fmt.Errorf("failed to create ExifTool instance: %w", err)
	}
	return &ExifClient{etClient: client}, nil
}

func (c *ExifClient) GetFilesMetaByStore(store *fileitem.Store) []ExifMeta {
	sourceKeys := store.GetSourceKeys()
	metas := make([]ExifMeta, 0, len(sourceKeys))
	for _, etMetadata := range c.etClient.ExtractMetadata(sourceKeys...) {
		metas = append(metas, ExifMeta{etMetadata: etMetadata})
	}
	return metas
}

func (c *ExifClient) GetFilesMetaBySlice(paths []string) []ExifMeta {
	metas := make([]ExifMeta, 0, len(paths))
	for _, etMetadata := range c.etClient.ExtractMetadata(paths...) {
		metas = append(metas, ExifMeta{etMetadata: etMetadata})
	}
	return metas
}

func (c *ExifClient) GetFileMeta(path string) ExifMeta {
	metas := c.etClient.ExtractMetadata(path)
	if len(metas) == 1 {
		return ExifMeta{etMetadata: metas[0]}
	}
	return ExifMeta{
		etMetadata: exiftool.FileMetadata{Err: errors.New("extract, but got empty result")},
	}
}

type ExifMeta struct {
	etMetadata exiftool.FileMetadata
}

// PrimitiveMeta Get exiftool.FileMetadata
func (m *ExifMeta) PrimitiveMeta() exiftool.FileMetadata {
	return m.etMetadata
}

func (m *ExifMeta) GetString(k string) (string, error) {
	return m.etMetadata.GetString(k)
}

func (m *ExifMeta) GetStringSafe(k string) string {
	v, err := m.etMetadata.GetString(k)
	if err != nil {
		return ""
	}
	return v
}

func (m *ExifMeta) GetStrings(k string) ([]string, error) {
	return m.etMetadata.GetStrings(k)
}

func (m *ExifMeta) GetStringsSafe(k string) []string {
	v, err := m.etMetadata.GetStrings(k)
	if err != nil {
		return []string{}
	}
	return v
}

func (m *ExifMeta) GetInt(k string) (int64, error) {
	return m.etMetadata.GetInt(k)
}

func (m *ExifMeta) GetIntSafe(k string) int64 {
	v, err := m.etMetadata.GetInt(k)
	if err != nil {
		return 0
	}
	return v
}

func (m *ExifMeta) GetFloat(k string) (float64, error) {
	return m.etMetadata.GetFloat(k)
}

func (m *ExifMeta) GetFloatSafe(k string) float64 {
	v, err := m.etMetadata.GetFloat(k)
	if err != nil {
		return 0
	}
	return v
}
