package service

import "io"

type StorageService interface {
	SaveImage(image io.ReadCloser, ext string) (string, error)
}
