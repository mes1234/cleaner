//Package fileshandler to handle files in cleaner
package fileshandler

import (
	"fmt"
	"os"
	"path/filepath"
)

//FileRawItem is subitem of raw data for json preprocessor
type fileRawItem struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Parent int    `json:"parent"`
	Childs []int  `json:"childs,omitempty"`
	Keep   bool   `json:"keep"`
}

//FileRaw is def of single item in file system, directory or file
//Used for decoding Json files
type FileRaw []fileRawItem

//File is def of single item in file system, directory or file
//With pointers references
type File struct {
	ID      int
	Name    string
	Keep    bool
	PParent *Directory
}

//Directory is def of single item in file system, directory or file
//With pointers references
type Directory struct {
	File
	PChilds []*Directory
}

//Directories is abstraction over slice of Directory
type Directories []Directory

//Files is abstraction over slice of File
type Files []File

//func checkIfExistsAndUpdate verify if given related item exists if not creates
func checkIfExistsAndUpdate(srcID int, allsrc *FileRaw, resFile *[]Directory) {
	if srcID >= 0 && (*resFile)[srcID].ID == 0 {
		parentPtr := &((*allsrc)[srcID])
		go (*resFile)[srcID].update(parentPtr, allsrc, resFile)
	}
}

// Split returns two object one files only, second directory only
func (f Directories) Split() (*[]Directory, *[]File) {

	resDirs := make([]Directory, 0, 0)
	resFiles := make([]File, 0, 0)
	for _, v := range f {
		switch len(v.PChilds) {
		case 0:
			resFiles = append(resFiles, v.File)
		default:
			resDirs = append(resDirs, v)
		}
	}
	return &resDirs, &resFiles
}

//update fils fields in struct and if needed create subitems
func (f *Directory) update(src *fileRawItem, allsrc *FileRaw, resFile *[]Directory) {
	f.ID = src.ID
	f.Name = src.Name
	f.Keep = src.Keep
	parentID := src.Parent
	checkIfExistsAndUpdate(parentID, allsrc, resFile)
	if parentID >= 0 {
		f.PParent = &((*resFile)[parentID])
	}
	f.PChilds = make([]*Directory, len(src.Childs), len(src.Childs))
	for k, v := range src.Childs {
		childID := v
		checkIfExistsAndUpdate(childID, allsrc, resFile)
		f.PChilds[k] = &((*resFile)[childID])
		fmt.Println(k, v)
	}
}

//Discover recurently files in given JSON file
func (f FileRaw) Discover() *[]Directory {
	var resFile = make([]Directory, len(f), len(f))
	for k, v := range f {
		resFile[k].update(&v, &f, &resFile)
		fmt.Println(k, v)
	}
	return &resFile
}

//getPathRec creates
func (f Directory) getPathRec() string {

}

//Create creates folder structure in given root path
func (f Directories) Create(rootPath string) bool {
	for _, v := range f {
		v.create(rootPath)
	}
	return true
}

//create creates folder structure in given root path
// need to work on it
func (f Directory) create(rootPath string) bool {
	dir, _ := os.Getwd()
	switch id := f.PParent; id {
	case nil:
		os.MkdirAll(filepath.Join(dir, rootPath, f.Name), 1777)
	default:
		path := filepath.Join(dir, rootPath, f.Name)
		f.create(path)
	}
	return true
}
