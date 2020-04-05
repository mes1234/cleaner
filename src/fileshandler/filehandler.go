//Package fileshandler to handle files in cleaner
package fileshandler

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
	PParent *Directory
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
func checkIfExistsAndUpdate(srcID int, allsrc *Raw, resFile *[]File, resDirs *[]Directory) {
	if srcID >= 0 && (*resFile)[srcID].ID == 0 {
		parentPtr := &((*allsrc)[srcID].fileRawItem)
		(*resFile)[srcID].update(parentPtr, allsrc, resFile)
	}
}
func (d *Directory) update(src *rawStruct, allsrc *Raw, dirs *[]Directory, files *[]File) {
	copyCommon(src, &d.File)
	d.PChildsD = make([]*Directory, 0, len(src.ChildsDirectories))
	d.PChildsF = make([]*File, 0, len(src.ChildsFiles))
	for _, v := range src.ChildsDirectories {
		childID := v
		checkIfExistsAndUpdate(childID, allsrc, files, dirs)
		if childID >= 0 {
			for _, v := range *dirs {
				if v.ID == childID {
					d.PChildsD = append(d.PChildsD, &v)
				}
			}
		}
	}

	for _, v := range src.ChildsFiles {
		childID := v
		checkIfExistsAndUpdate(childID, allsrc, files, dirs)
		if childID >= 0 {
			for _, v := range *files {
				if v.ID == childID {
					d.PChildsF = append(d.PChildsF, &v)
				}
			}
		}
	}
}

func (f *File) update(src *rawStruct, allsrc *Raw, dirs *[]Directory, files *[]File) {
	copyCommon(src, f)
	parentID := src.Parent
	checkIfExistsAndUpdate(parentID, allsrc, files, dirs)
	if parentID >= 0 {
		for _, v := range *dirs {
			if v.ID == parentID {
				f.PParent = &v
			}
		}
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
