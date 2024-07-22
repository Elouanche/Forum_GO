package main

import (
	"io"
	"mime/multipart"
	"os"
)

func saveFile(file multipart.File, filePath string) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return err
	}

	return nil
}
