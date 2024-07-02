package fileitem

import (
	"github.com/sincerefly/capybara/utils/exif"
	"path/filepath"
	"strings"
)

type FileItem struct {
	filename    string // e.g. 01.jpg
	size        int64
	hashValue   string
	content     []byte
	contentType string
	innerPath   string // e.g. lite/01.jpg
	sourceBase  string // e.g. /home/input/ or input/
	targetBase  string // e.g. /home/target/

	exifMeta exif.ExifMeta
}

func NewFileItem(filename string) FileItem {
	return FileItem{
		filename: filename,
	}
}

func (f *FileItem) GetFileBasename() string {
	return strings.TrimSuffix(f.filename, f.GetFileExt())
}

func (f *FileItem) GetFilename() string {
	return f.filename
}

func (f *FileItem) GetFileExt() string {
	return filepath.Ext(f.filename)
}

func (f *FileItem) SetHash(hash string) *FileItem {
	f.hashValue = hash
	return f
}

func (f *FileItem) GetHash() string {
	return f.hashValue
}

func (f *FileItem) SetSize(size int64) *FileItem {
	f.size = size
	return f
}

func (f *FileItem) GetSize() int64 {
	return f.size
}

func (f *FileItem) SetContent(content []byte) *FileItem {
	f.content = content
	return f
}

func (f *FileItem) GetContent() []byte {
	return f.content
}

func (f *FileItem) SetContentType(contentType string) *FileItem {
	f.contentType = contentType
	return f
}

func (f *FileItem) GetContentType() string {
	return f.contentType
}

func (f *FileItem) SetInnerPath(path string) *FileItem {
	f.innerPath = path
	return f
}

func (f *FileItem) GetInnerPath() string {
	return f.innerPath
}

func (f *FileItem) GetInnerKey() string {
	return filepath.Join(f.innerPath, f.filename)
}

func (f *FileItem) SetSourceBase(base string) *FileItem {
	f.sourceBase = base
	return f
}

func (f *FileItem) GetSourcePath() string {
	return filepath.Join(f.sourceBase, f.innerPath)
}

func (f *FileItem) GetSourceKey() string {
	return filepath.Join(f.sourceBase, f.innerPath, f.filename)
}

func (f *FileItem) SetTargetBase(base string) *FileItem {
	f.targetBase = base
	return f
}

func (f *FileItem) GetTargetPath() string {
	return filepath.Join(f.targetBase, f.innerPath)
}

func (f *FileItem) GetTargetKey() string {
	return filepath.Join(f.targetBase, f.innerPath, f.filename)
}

func (f *FileItem) SetExifMeta(meta exif.ExifMeta) {
	f.exifMeta = meta
}

func (f *FileItem) GetExifMeta() exif.ExifMeta {
	return f.exifMeta
}
