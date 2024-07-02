package border

import (
	"fmt"
	"github.com/sincerefly/capybara/base/log"
	"github.com/sincerefly/capybara/service/border/styles"
	"github.com/sincerefly/capybara/structure"
	"github.com/sincerefly/capybara/utils"
	"github.com/sincerefly/capybara/utils/fileitem"
	"github.com/spf13/afero"
	"path/filepath"
	"slices"
	"strings"
)

type StyleProcessor struct {
	style  Style
	params Parameter
}

func NewStyleProcessor(style Style, params Parameter) *StyleProcessor {
	return &StyleProcessor{
		style:  style,
		params: params,
	}
}

func (s *StyleProcessor) Run() {

	style := s.style
	if !slices.Contains(s.Supports(), style) { // check
		log.Fatal("unsupported style")
	}

	fiStore, err := s.collectInputs()
	if err != nil {
		log.Fatal(err)
	}

	if err := s.prepareOutputDirs(fiStore); err != nil {
		log.Fatal(err)
	}

	switch style {
	case StyleSimple:
		params := s.params.(*styles.SimpleParameter)
		err = styles.NewSimpleProcessor(params, fiStore).Run()
	case StyleTextBottom:
		params := s.params.(*styles.TextBottomParameter)
		err = styles.NewTextBottomProcessor(params, fiStore).Run()
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Info("finished")
}

func (s *StyleProcessor) Supports() []Style {
	return []Style{StyleSimple, StyleTextBottom}
}

func (s *StyleProcessor) SupportExtensions() []string {
	return []string{structure.ExtJPG, structure.ExtPNG, structure.ExtJPEG}
}

// collect input dir images path
func (s *StyleProcessor) collectInputs() (*fileitem.Store, error) {
	store := fileitem.NewFileItemStore()

	input := s.params.Input()
	output := s.params.Output()

	fs := afero.NewOsFs()
	srcImagePaths, err := utils.GetAllFiles(fs, s.params.Input())
	if err != nil {
		return nil, fmt.Errorf("load input folder failed, %v", err)
	}

	for _, srcImagePath := range srcImagePaths {

		filename := filepath.Base(srcImagePath)

		if !slices.Contains(s.SupportExtensions(), strings.ToLower(filepath.Ext(filename))) {
			continue
		}

		relativePath, err := filepath.Rel(input, srcImagePath)
		if err != nil {
			return nil, err
		}
		innerPath := filepath.Dir(relativePath)
		if err != nil {
			return nil, err
		}

		fi := fileitem.NewFileItem(filename)
		fi.SetInnerPath(innerPath)
		fi.SetSourceBase(input)
		fi.SetTargetBase(output)

		store.Add(fi)

	}
	return &store, nil
}

// prepare output dirs
func (s *StyleProcessor) prepareOutputDirs(fiStore *fileitem.Store) error {
	for _, outputDir := range fiStore.GetTargetPaths() {
		if err := utils.MkdirAll(outputDir); err != nil {
			return err
		}
	}
	return nil
}
