package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// Function to recursively pull all files in a directory, with a string as arguments for src and dst
func dirPuller(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}

	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = dirPuller(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = filePuller(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
