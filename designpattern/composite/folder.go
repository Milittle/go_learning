package composite

import (
	"fmt"
)

type Folder struct {
	components []Component
	name string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Searching recurively for keyword %s in folder %s\n", keyword, f.name)
	for _, com := range f.components {
		com.Search(keyword)
	}
}

func (f *Folder) AddComponent(com Component) {
	f.components = append(f.components, com)
}

func (f *Folder) GetName() string {
	return f.name
}
