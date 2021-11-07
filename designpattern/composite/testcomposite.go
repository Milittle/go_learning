package composite

func TestComposite() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{name: "Folder1"}
	folder1.AddComponent(file1)

	folder2 := &Folder{name: "Folder2"}
	folder2.AddComponent(file2)
	folder2.AddComponent(file3)
	folder2.AddComponent(folder1)

	folder2.Search("rose")
}
