//Package fileshandler to handle files in cleaner
package fileshandler

import "fmt"

//FileRaw is subitem of raw data for json preprocessor
type fileRaw struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Parent int    `json:"parent"`
	Keep   bool   `json:"keep"`
}

//rawStruct is subitem of raw data for json preprocessor
type rawStruct struct {
	fileRaw
	ChildsDirectories []int `json:"childsd,omitempty"`
	ChildsFiles       []int `json:"childsf,omitempty"`
}

//Raw is an array of raw not processed json data
type Raw []struct {
	rawStruct
}

//File is def of single item in file system, directory or file
//With pointers references
type File struct {
	ID      int
	Name    string
	PParent *File
	Keep    bool
}

//Directory is def of single item in file system, directory or file
//With pointers references
type Directory struct {
	File
	PChildsF []*File
	PChildsD []*Directory
}

func copyCommon(src *rawStruct, dest *File) {
	dest.ID = src.ID
	dest.Name = src.Name
	dest.Keep = src.Keep
}

//func checkIfExistsAndUpdate verify if given related item exists if not creates
func checkIfExistsAndUpdate(srcID int, allsrc *Raw, resFile *[]File) {
	if srcID >= 0 && (*resFile)[srcID].ID == 0 {
		parentPtr := &((*allsrc)[srcID].fileRawItem)
		(*resFile)[srcID].update(parentPtr, allsrc, resFile)
	}
}
func (f *Directory) update(src *rawStruct, allsrc *Raw, parents *[]Directory, childs *[]File) {
	copyCommon(src, f)
	f.PChilds = make([]*File, len(src.Childs), len(src.Childs))
	for k, v := range src.Childs {
		childID := v
		checkIfExistsAndUpdate(childID, allsrc, resFile)
		f.PChilds[k] = &((*resFile)[childID])
		fmt.Println(k, v)
	}
}
func (f *File) update(src *rawStruct, allsrc *Raw, parents *[]Directory, childs *[]File) {
	copyCommon(src, f)
	parentID := src.Parent
	checkIfExistsAndUpdate(parentID, allsrc, parents)
	if parentID >= 0 {
		f.PParent = &((*parents)[parentID])
	}
}

//Discover recurently files in given JSON file
func (f Raw) Discover() (*[]File, *[]Directory) {
	var resFiles = make([]File, 0, len(f))
	var resDirs = make([]Directory, 0, len(f))
	for _, v := range f {
		switch len(v.ChildsDirectories) + len(v.ChildsFiles) {
		case 0:
			// its a file
			file := new(File)
			file.update(&v.rawStruct, &f, &resDirs, &resFiles)
			resFiles = append(resFiles, *file)
		default:
			// its a folder
			dir := new(Directory)
			dir.update(&v.rawStruct, &f, &resDirs, &resFiles)
			resDirs = append(resDirs, *dir)

		}
	}
	return &resFiles, &resDirs
}
