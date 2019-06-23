package files

import (
	"log"
	"os"
)

//FileScaner scans media folder and find all media files
type FileScaner struct {
	root string
}

//New creates new Filescancer
func New(dir string) *FileScaner {
	if dir == "" {
		panic("media file directory can not be empty")
	}

	finfo, err := os.Lstat(dir)
	if err != nil || !finfo.IsDir() {
		panic("root folder does not exist or not a dir: " + dir)
	}

	log.Printf("initialized filescaner with root dir: %s", dir)

	return &FileScaner{root: dir}
}
