package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sincerefly/capybara/service/style/styles_common"
	"github.com/sincerefly/capybara/structure/fileitem"
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

	filename := filepath.Base(imgPath)
	fi := fileitem.NewFileItem(filename)
	fi.SetInnerPath(filepath.Dir(imgPath))

	store := fileitem.NewFileItemStore()
	store.Add(fi)
	metas, err := styles_common.GetFilesMetaByStore(&store)
	if err != nil {
		fmt.Println(err)
	}

	for field, value := range metas[0].PrimitiveMeta().Fields {
		fmt.Printf("[%s] %v\n", field, value)
	}
}
