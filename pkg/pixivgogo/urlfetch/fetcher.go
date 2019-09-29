package urlfetch

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/sleepingpig/pixivgogo/pkg/pixivgogo"

	"github.com/imroc/req"
)

type Fetcher struct {
	client   *req.Req
	referrer string
}

func NewFetcher() *Fetcher {
	return &Fetcher{
		client:   req.New(),
		referrer: pixivgogo.DefaultAPIURL,
	}
}

func (f *Fetcher) FetchToWriter(srcURL string, output io.Writer) error {
	resp, err := f.fetchFromURL(srcURL)
	if err != nil {
		return err
	}
	input := resp.Response().Body
	defer closeQuietlyFunc(input)()
	_, err = io.Copy(output, input)
	return err
}

func (f *Fetcher) FetchToFile(srcURL string, filePath string) error {
	resp, err := f.fetchFromURL(srcURL)
	if err != nil {
		return err
	}
	return resp.ToFile(filePath)
}

// FetchToDir fetch the file from the given URL, and save to a file in the given directory.
// The file name is decided from the last path segment of the URL.
// TODO The method used here is incorrect. It will having timing problem.
func (f *Fetcher) FetchToDir(srcURL string, dirPath string, overwrite bool) error {
	filePath, err := f.findFilePathToStore(srcURL, dirPath, overwrite)
	if err != nil {
		return err
	}
	return f.FetchToFile(srcURL, filePath)
}

func (f *Fetcher) findFilePathToStore(srcURLStr string, dirPath string, overwrite bool) (string, error) {
	if err := f.shouldBeDir(dirPath); err != nil {
		return "", err
	}
	srcURL, err := url.Parse(srcURLStr)
	if err != nil {
		return "", err
	}
	lastSlashIdx := strings.LastIndex(srcURL.Path, "/")
	if lastSlashIdx < 0 {
		return "", fmt.Errorf("unable to decide the file name from url %s", srcURLStr)
	}
	fileName := srcURL.Path[lastSlashIdx+1:]
	if fileName == "" {
		return "", fmt.Errorf("unable to decide the file name from url %s", srcURLStr)
	}
	filePath := path.Join(dirPath, fileName)
	if err := f.shouldAbleToSaveFile(filePath, overwrite); err != nil {
		return "", err
	}
	return filePath, nil
}

func (f *Fetcher) shouldAbleToSaveFile(path string, overwrite bool) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	} else if fileInfo.IsDir() {
		return fmt.Errorf("target path %q is a directory", path)
	} else if !overwrite {
		return fmt.Errorf("file %q already exists", path)
	}
	return nil
}

func (f *Fetcher) shouldBeDir(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%q doesn't exist", path)
		}
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("%q isn't a directory", path)
	}
	return nil
}

func (f *Fetcher) fetchFromURL(srcURL string) (*req.Resp, error) {
	headers := req.Header{
		"Referer": f.referrer,
	}
	resp, err := req.Get(srcURL, headers)
	if err != nil {
		return nil, err
	}
	if resp.Response().StatusCode != 200 {
		defer closeQuietlyFunc(resp.Response().Body)()
		return nil, err
	}
	return resp, err
}

func closeQuietlyFunc(closer io.Closer) func() {
	return func() {
		err := closer.Close()
		if err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}
}
