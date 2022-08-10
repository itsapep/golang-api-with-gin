package repository

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type FileRepository interface {
	Save(file multipart.File, fileName string) (string, error)
}

type fileRepository struct {
	path string
}

// Save implements FileRepository
func (f *fileRepository) Save(file multipart.File, fileName string) (string, error) {
	location := filepath.Join(f.path, fileName)
	fmt.Println(location)
	output, err := os.OpenFile(location, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			panic(err)
		}
	}(output)
	_, err = io.Copy(output, file)
	if err != nil {
		return "", err
	}
	return location, nil
}

func NewFileRepository(path string) FileRepository {
	return &fileRepository{
		path: path,
	}
}
