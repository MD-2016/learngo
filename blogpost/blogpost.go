package blogpost

import (
	"io/fs"
)

func NewFSPosts(fss fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fss, ".")
	if err != nil {
		return nil, err
	}

	var pts []Post
	for _, f := range dir {
		pt, err := getPost(fss, f)
		if err != nil {
			return nil, err
		}
		pts = append(pts, pt)
	}
	return pts, nil
}

func getPost(fss fs.FS, fde fs.DirEntry) (Post, error) {
	ptf, err := fss.Open(fde.Name())
	if err != nil {
		return Post{}, err
	}
	defer ptf.Close()

	return newPost(ptf)
}
