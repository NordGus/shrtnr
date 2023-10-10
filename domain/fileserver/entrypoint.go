package fileserver

import "io/fs"

var (
	files fs.FS
)

func Start(content fs.FS) {
	files = content
}
