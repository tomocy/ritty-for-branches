package fs

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func NewLocal() *Local {
	return new(Local)
}

type Local struct{}

func (l *Local) SaveImage(image io.ReadCloser, ext string) (string, error) {
	defer image.Close()
	projPath, _ := filepath.Abs(".")
	name := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	path := filepath.Join(projPath, imageDir, name)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0744)
	if err != nil {
		return "", err
	}
	defer file.Close()

	io.Copy(file, image)
	return strings.TrimPrefix(path, filepath.Join(projPath, "resource/public")), nil
}

const imageDir = "resource/public/image"
