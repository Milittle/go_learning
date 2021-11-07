package composite

import "fmt"

type File struct {
	name string
}

func (f *File) Search(keyword string) {
	fmt.Printf("Searching the keyword %s in file %s\n", keyword, f.name)
}

func (f *File) GetName() string {
	return f.name
}