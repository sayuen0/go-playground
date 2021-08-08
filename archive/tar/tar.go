package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

/**
tar archive つかてみた
*/

func main() {
	openTarArchive()
}

func openTarArchive() {
	file, _ := os.Open("./dist/archive.tar")
	defer file.Close()
	tr := tar.NewReader(file)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		fmt.Println(header.Name)
		var buf bytes.Buffer
		buf.ReadFrom(tr)
		fmt.Println(buf.String())
	}
}

func createTarArchive() {
	dist, err := os.Create("dist/archive.tar")
	if err != nil {
		panic(err)
	}
	defer dist.Close()

	tw := tar.NewWriter(dist)
	defer tw.Close()

	if err := filepath.Walk("./dir", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// ignore directories
		if info.IsDir() {
			return nil
		}
		// write header
		if err := tw.WriteHeader(&tar.Header{
			Name:    path,
			Mode:    int64((info.Mode())),
			ModTime: info.ModTime(),
			Size:    info.Size(),
		}); err != nil {
			return err
		}

		// write body
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := io.Copy(tw, f); err != nil {
			return err
		}
		return nil
	}); err != nil {
		panic(err)
	}
}
