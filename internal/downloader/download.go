package downloader

import (
	"io"
	"net/http"
	"os"
)

func Download(asset Asset) error {

	resp, err := http.Get(asset.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(asset.Path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}
