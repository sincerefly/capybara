package main

import (
	"fmt"
	"github.com/sincerefly/capybara/utils/exif_utils"
	"github.com/sincerefly/capybara/utils/fileitem"
	"os"
	"path/filepath"
)

/*
	Usage
		bin/exifviewer input/nikon-z30.jpg
*/

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No arguments provided.")
		fmt.Println("Usage: ./exifviewer input/photo.jpg")
	}

	imgPath := args[0]

	// parser exif meta data
	etClient, err := exif_utils.NewExifClient()
	if err != nil {
		fmt.Println(err)
	}

	filename := filepath.Base(imgPath)
	fi := fileitem.NewFileItem(filename)
	fi.SetInnerPath(filepath.Dir(imgPath))

	store := fileitem.NewFileItemStore()
	store.Add(fi)
	metas := etClient.GetFilesMetaByStore(&store)

	for field, value := range metas[0].PrimitiveMeta().Fields {
		fmt.Printf("[%s] %v\n", field, value)
	}
}
