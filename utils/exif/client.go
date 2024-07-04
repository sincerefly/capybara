package exif

import (
	"errors"
	"fmt"
	"github.com/barasher/go-exiftool"
)

type Client struct {
	etClient *exiftool.Exiftool
}

func NewExifClient() (*Client, error) {
	client, err := exiftool.NewExiftool()
	if err != nil {
		return nil, fmt.Errorf("failed to create ExifTool instance: %w", err)
	}
	return &Client{etClient: client}, nil
}

func (c *Client) ExtractMetadata(paths []string) []Meta {

	metas := make([]Meta, 0, len(paths))
	for _, etMetadata := range c.etClient.ExtractMetadata(paths...) {
		metas = append(metas, NewExifMeta(etMetadata))
	}
	return metas
}

func (c *Client) GetFilesMetaBySlice(paths []string) []Meta {
	metas := make([]Meta, 0, len(paths))
	for _, etMetadata := range c.etClient.ExtractMetadata(paths...) {
		metas = append(metas, Meta{etMetadata: etMetadata})
	}
	return metas
}

func (c *Client) GetFileMeta(path string) Meta {
	metas := c.etClient.ExtractMetadata(path)
	if len(metas) == 1 {
		return Meta{etMetadata: metas[0]}
	}
	return Meta{
		etMetadata: exiftool.FileMetadata{Err: errors.New("extract, but got empty result")},
	}
}

type Meta struct {
	etMetadata exiftool.FileMetadata
}

func NewExifMeta(etMetadata exiftool.FileMetadata) Meta {
	return Meta{
		etMetadata: etMetadata,
	}
}

// PrimitiveMeta Get exiftool.FileMetadata
func (m *Meta) PrimitiveMeta() exiftool.FileMetadata {
	return m.etMetadata
}

func (m *Meta) GetString(k string) (string, error) {
	return m.etMetadata.GetString(k)
}

func (m *Meta) GetStringSafe(k string) string {
	v, err := m.etMetadata.GetString(k)
	if err != nil {
		return ""
	}
	return v
}

func (m *Meta) GetStrings(k string) ([]string, error) {
	return m.etMetadata.GetStrings(k)
}

func (m *Meta) GetStringsSafe(k string) []string {
	v, err := m.etMetadata.GetStrings(k)
	if err != nil {
		return []string{}
	}
	return v
}

func (m *Meta) GetInt(k string) (int64, error) {
	return m.etMetadata.GetInt(k)
}

func (m *Meta) GetIntSafe(k string) int64 {
	v, err := m.etMetadata.GetInt(k)
	if err != nil {
		return 0
	}
	return v
}

func (m *Meta) GetFloat(k string) (float64, error) {
	return m.etMetadata.GetFloat(k)
}

func (m *Meta) GetFloatSafe(k string) float64 {
	v, err := m.etMetadata.GetFloat(k)
	if err != nil {
		return 0
	}
	return v
}
