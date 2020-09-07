package utils

import (
	"errors"
	"io"
	"net/http"
	"os"
)

// Download file into given filepath from given url
func Download(url, dest string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("Node version not found")
	}
	defer resp.Body.Close()
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}
