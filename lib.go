package cefresources

import (
	"embed"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//go:embed resources/*
var content embed.FS

func Extract() string {
	dir, err := ioutil.TempDir("", "wv")
	if err != nil {
		log.Fatal(err)
	}
	os.RemoveAll(dir)
	spdir := strings.Split(dir, `\`)
	spdir[len(spdir)-1] = "cef-res-2785"
	dir = strings.Join(spdir, `\`)
	if _, err := os.Stat(dir); err != nil {
		os.Mkdir(dir, 0755)
	}

	for _, fname := range []string{
		"cef.pak",
		"cef_100_percent.pak",
		"cef_200_percent.pak",
		"d3dcompiler_47.dll",
		"icudtl.dat",
		"libEGL.dll",
		"libcef.dll",
		"natives_blob.bin",
		"snapshot_blob.bin",
		"en-US.pak",
	} {
		file, err := content.Open("resources/" + fname)
		if err != nil {
			log.Fatal("Err 01: ", err)
		}

		destPath := filepath.Join(dir, fname)
		if _, err := os.Stat(destPath); err != nil {
			destination, err := os.Create(destPath)
			if err != nil {
				log.Fatal("Err 02: ", err)
			}
			_, err = io.Copy(destination, file)
			if err != nil {
				log.Fatal("Err 03: ", err)
			}
			destination.Close()
		}
	}

	return dir
}
