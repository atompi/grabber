package tools

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/atompi/grabber/internal/options"
)

func GenerateDownloadList(opts []options.SourceOptions) (downloadList []options.FileOptions) {
	for _, option := range opts {
		for _, file := range option.Files {
			downloadUrl := fmt.Sprintf("%s/%s?access_token=%s", option.Url, file.Src, option.Auth)
			download := &options.FileOptions{
				Src:  downloadUrl,
				Dest: file.Dest,
			}
			downloadList = append(downloadList, *download)
		}
	}
	return
}

func DownloadFile(filePath string, url string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode >= 400 {
		err = fmt.Errorf("bad response: %s", resp.Status)
		return
	}
	defer resp.Body.Close()

	_ = os.MkdirAll(filepath.Dir(filePath), 0755)
	out, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return
}
