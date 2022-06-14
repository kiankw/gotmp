package main

import (
	"fmt"
	"io"
	"io/fs"
)

type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	io.WriterAt
}

type UnionFile struct {
	Base  File
	Layer File
}

func (f *UnionFile) Seek(o int64, w int) (pos int64, err error) {
	if f.Layer != nil {
		pos, err = f.Layer.Seek(o, w)
		if (err == nil || err == io.EOF) && f.Base != nil {
			_, err = f.Base.Seek(o, w)
		}
		return pos, err
	}
	if f.Base != nil {
		return f.Base.Seek(o, w)
	}
	return 0, nil
}

type RegexpFile struct {
	f File
}

func (f *RegexpFile) Seek(o int64, w int) (int64, error) {
	return f.f.Seek(o, w)
}

type fromIOFSFile struct {
	fs.File
}

func (f fromIOFSFile) Seek(offset int64, whence int) (int64, error) {
	seeker, ok := f.File.(io.Seeker)
	if !ok {
		return -1, nil
	}
	return seeker.Seek(offset, whence)
}

func main() {
	var f fromIOFSFile

	seeker, ok := f.File.(io.Seeker)
	if ok {
		fmt.Println(ok)
	}

	seeker.Seek(1, 1)

	fmt.Println("Hello World!")
}
