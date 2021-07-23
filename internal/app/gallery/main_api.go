package gallery

import (
	"io/fs"
	"os"
	"strings"

	"github.com/maldan/go-cmhp"
)

type FileApi struct {
}

var FormatList = []string{".jpg", ".jpeg", ".png", ".gif", ".bmp"}

func (f FileApi) GetIndex(args Image) *os.File {
	file, err := os.Open(args.Path)
	if err != nil {
		return nil
	}
	return file
}

func (f FileApi) GetInfo() interface{} {
	list, _ := cmhp.FileList(Folder)
	filterList := cmhp.SliceFilterR(list, func(i interface{}) bool {
		for _, format := range FormatList {
			if strings.HasSuffix(i.(fs.FileInfo).Name(), format) {
				return true
			}
		}
		return false
	})

	out := make([]Image, 0)
	for _, file := range filterList {
		out = append(out, Image{Path: Folder + "/" + file.(fs.FileInfo).Name()})
	}

	return map[string]interface{}{
		"file":  File,
		"files": out,
	}
}
