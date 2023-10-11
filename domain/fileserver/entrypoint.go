package fileserver

import "io/fs"

var (
	// files most only contain 2 directories one called dist (for public static files)
	// and another called private (for private static files)
	files fs.FS
)

func Start(content fs.FS) {
	files = content
}
