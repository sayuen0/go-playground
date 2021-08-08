package main

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	before()
	files := find()

	b:= compress(files)
	if err := save(b); err != nil {
		panic(err)
	}
}

func before() {
	f,  err := os.Create("sample.txt")
	if err != nil {
		panic(err)
	}
	f.Write([]byte("hello world"))
	f.Close()
}

func find() []string{
	return []string{"sample.txt"}
}

func compress(files []string) *bytes.Buffer{
	b := new(bytes.Buffer)
	w := zip.NewWriter(b)
	for _, file := range files{
		info,_ := os.Stat(file)
		hdr , _ := zip.FileInfoHeader(info)
		hdr.Name = filepath.Join("files/", file)
		f, err := w.CreateHeader(hdr)
		if err != nil {
			panic(err)
		}
		body, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		f.Write(body)
	}
	w.Close()
	return b
}

func save(b *bytes.Buffer) error{
	zf, err:= os.Create("sample.zip")
	if err != nil {
		return err
	}
	zf.Write(b.Bytes())
	zf.Close()
	return nil
}