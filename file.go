package gome

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CopyDir recursively copies a directory from source to destination.
func CopyDir(src, dst string) error {
	if err := os.MkdirAll(dst, 07555); err != nil {
		return err
	}

	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, file := range files {
		sp := filepath.Join(src, file.Name())
		dp := filepath.Join(dst, file.Name())

		if file.IsDir() {
			if err := CopyDir(sp, dp); err != nil {
				fmt.Println("Error copying directory:", err.Error())
			}
		} else {
			if err := CopyFile(sp, dp); err != nil {
				fmt.Println("Error copying file:", err.Error())
			}
		}
	}

	return nil
}

// CopyFile copies a file from the source to the destination paths.
func CopyFile(src, dst string) error {
	sf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sf.Close()

	df, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer df.Close()

	if _, err := io.Copy(df, sf); err != nil {
		return err
	}

	return df.Sync()
}

// Copy is kept for retro-compatibility.
func Copy(src, dst string) error {
	return CopyFile(src, dst)
}
