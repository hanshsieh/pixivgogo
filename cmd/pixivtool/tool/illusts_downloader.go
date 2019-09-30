package tool

import (
	"errors"
	"log"

	"github.com/sleepingpig/pixivgogo/pkg/pixivgogo"
	"github.com/sleepingpig/pixivgogo/pkg/pixivgogo/urlfetch"
)

type DownloadIllustConfig struct {
	Username     string
	Password     string
	DstDirectory string
	Count        int
}

type IllustsDownloader struct {
	client *pixivgogo.Client
}

func NewIllustsDownloader() *IllustsDownloader {
	return &IllustsDownloader{
		client: pixivgogo.NewClient(),
	}
}

func (p *IllustsDownloader) login(username, password string) error {
	if username == "" || password == "" {
		return errors.New("username and password are required")
	}
	return p.client.Login(username, password)
}

func (p *IllustsDownloader) DownloadIllustrations(config *DownloadIllustConfig) error {
	log.Print("Logging in...")
	if err := p.login(config.Username, config.Password); err != nil {
		return err
	}
	fetcher := urlfetch.NewFetcher()
	numIllusts := 0
	filter := &pixivgogo.RecommendIllustsFilter{}

	for {
		log.Print("Getting recommended illustrations...")
		illusts, err := p.client.IllustsRecommend(filter)
		if err != nil {
			return err
		}
		for _, illust := range illusts.Illustrations {
			if numIllusts >= config.Count {
				return nil
			}
			if illust.Type != pixivgogo.ILLUST_TYPE_ILLUST {
				log.Printf("%d is skipped because it's type %q != %q",
					illust.ID, illust.Type, pixivgogo.ILLUST_TYPE_ILLUST)
				continue
			}
			if err := p.downloadIllustration(illust, fetcher, config.DstDirectory); err != nil {
				return err
			}
			numIllusts++
		}
		filter = illusts.NextFilter
	}
}

func (p *IllustsDownloader) downloadIllustration(
	illust *pixivgogo.Illustration,
	fetcher *urlfetch.Fetcher,
	dstDir string) error {
	var urls []string
	if illust.MetaSinglePage != nil {
		urls = append(urls, illust.MetaSinglePage.OriginalImageURL)
	}
	for _, metaPage := range illust.MetaPages {
		urls = append(urls, metaPage.ImageURLs.Original)
	}
	urls = p.filterEmptyString(urls)
	if err := p.downloadURLs(fetcher, urls, dstDir); err != nil {
		return err
	}
	return nil
}

func (p *IllustsDownloader) filterEmptyString(strs []string) []string {
	result := make([]string, 0, len(strs))
	for _, str := range strs {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

func (p *IllustsDownloader) downloadURLs(
	fetcher *urlfetch.Fetcher,
	urls []string,
	dirPath string) error {
	for _, srcURLStr := range urls {
		log.Printf("Downloading from %s", srcURLStr)
		if err := fetcher.FetchToDir(srcURLStr, dirPath, true); err != nil {
			return err
		}
	}
	return nil
}
