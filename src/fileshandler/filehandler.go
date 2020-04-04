//Package fileshandler to handle files in cleaner
package fileshandler

import (
	"fmt"
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
type FileRaw []struct {
	fileRawItem
}

//File is def of single item in file system, directory or file
//With pointers references
type File struct {
	ID      int
	Name    string
	PParent *File
	PChilds []*File
	Keep    bool
}

//func checkIfExistsAndUpdate verify if given related item exists if not creates
func checkIfExistsAndUpdate(srcID int, allsrc *FileRaw, resFile *[]File) {
	if srcID >= 0 && (*resFile)[srcID].ID == 0 {
		parentPtr := &((*allsrc)[srcID].fileRawItem)
		(*resFile)[srcID].update(parentPtr, allsrc, resFile)
	}
}

func (f *File) update(src *fileRawItem, allsrc *FileRaw, resFile *[]File) {
	f.ID = src.ID
	f.Name = src.Name
	f.Keep = src.Keep
	parentID := src.Parent
	checkIfExistsAndUpdate(parentID, allsrc, resFile)
	if parentID >= 0 {
		f.PParent = &((*resFile)[parentID])
	}
	f.PChilds = make([]*File, len(src.Childs), len(src.Childs))
	for k, v := range src.Childs {
		childID := v
		checkIfExistsAndUpdate(childID, allsrc, resFile)
		f.PChilds[k] = &((*resFile)[childID])
		fmt.Println(k, v)
	}
}

//Discover recurently files in given JSON file
func (f FileRaw) Discover() *[]File {
	var resFile = make([]File, len(f), len(f))
	for k, v := range f {
		resFile[k].update(&v.fileRawItem, &f, &resFile)
		fmt.Println(k, v)
	}
	return &resFile
}
