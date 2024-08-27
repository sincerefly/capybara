package fileitem

// Store FileItem collections
type Store struct {
	items []FileItem
}

// NewFileItemStore create new store
//
// e.g.
// NewFileItemStore()     lens 0, caps 0
// NewFileItemStore(x)    lens x, caps x
// NewFileItemStore(x, y) lens x, caps y
func NewFileItemStore(args ...int) Store {
	var items []FileItem
	switch len(args) {
	case 0:
		items = make([]FileItem, 0)
	case 1:
		items = make([]FileItem, args[0])
	case 2:
		items = make([]FileItem, args[0], args[1])
	default:
		return Store{}
	}
	return Store{
		items: items,
	}
}

func (f *Store) Len() int {
	return len(f.items)
}

func (f *Store) GetItems() []FileItem {
	return f.items
}

func (f *Store) Add(fileItem FileItem) {
	f.items = append(f.items, fileItem)
}

func (f *Store) Set(i int, fileItem FileItem) {
	f.items[i] = fileItem
}

func (f *Store) GetInnerKeys() []string {
	list := make([]string, len(f.items))
	for i, item := range f.items {
		list[i] = item.GetInnerKey()
	}
	return list
}

func (f *Store) GetInnerPaths() []string {
	list := make([]string, len(f.items))
	for i, item := range f.items {
		list[i] = item.GetInnerPath()
	}
	return list
}

func (f *Store) GetSourceKeys() []string {
	list := make([]string, len(f.items))
	for i, item := range f.items {
		list[i] = item.GetSourceKey()
	}
	return list
}

func (f *Store) GetTargetPaths() []string {
	list := make([]string, 0, len(f.items))
	m := make(map[string]struct{}, 0)
	for _, item := range f.items {
		path := item.GetTargetPath()
		if _, ok := m[path]; ok {
			continue
		}
		list = append(list, path)
	}
	return list
}
