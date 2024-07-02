package styles_common

import (
	"github.com/sincerefly/capybara/utils/exif"
	"github.com/sincerefly/capybara/utils/fileitem"
)

func GetFilesMetaByStore(store *fileitem.Store) ([]exif.ExifMeta, error) {
	client, err := exif.NewExifClient()
	if err != nil {
		return nil, err
	}
	sourceKeys := store.GetSourceKeys()
	return client.ExtractMetadata(sourceKeys), nil
}

func SupplementaryMetaToStore(store *fileitem.Store) (*fileitem.Store, error) {
	client, err := exif.NewExifClient()
	if err != nil {
		return nil, err
	}
	sourceKeys := store.GetSourceKeys()

	metas := client.ExtractMetadata(sourceKeys)

	newStore := fileitem.NewFileItemStore(0, store.Len())
	for i, fi := range store.GetItems() {
		fi.SetExifMeta(metas[i])
		newStore.Add(fi)
	}

	return &newStore, nil
}
