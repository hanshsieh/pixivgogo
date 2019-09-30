package pixivgogo

import "net/url"

// IllustsRanking returns the ranking of illustrations.
// The given filter can be used for filtering the illustrations, and control which kind of
// ranking should be used.
// Login is required.
func (c *Client) IllustsRanking(filter *RankingIllustsFilter) (*RankingIllustrations, error) {
	illustrations := &RankingIllustrations{}
	if err := c.doGetRequest("/v1/illust/ranking", filter, illustrations); err != nil {
		return nil, err
	}
	if illustrations.NextURL != "" {
		illustrations.NextFilter = new(RankingIllustsFilter)
		if err := c.convertNextURLToFilter(illustrations.NextURL, illustrations.NextFilter); err != nil {
			return nil, err
		}
	}
	return illustrations, nil
}

// IllustsRecommend returns the recommended illustrations based on the logged in user's preference.
// It needs login.
// There's another API for getting recommended illustrations for not logged-in users.
func (c *Client) IllustsRecommend(filter *RecommendIllustsFilter) (*RecommendIllustrations, error) {
	illustrations := &RecommendIllustrations{}
	if err := c.doGetRequest("/v1/illust/recommended", filter, illustrations); err != nil {
		return nil, err
	}
	if illustrations.NextURL != "" {
		illustrations.NextFilter = new(RecommendIllustsFilter)
		if err := c.convertNextURLToFilter(illustrations.NextURL, illustrations.NextFilter); err != nil {
			return nil, err
		}
	}
	return illustrations, nil
}

func (c *Client) convertNextURLToFilter(nextURLStr string, filter interface{}) error {
	nextURL, err := url.Parse(nextURLStr)
	if err != nil {
		return err
	}
	return c.urlValuesDecoder.Decode(filter, nextURL.Query())
}

// IllustDetail returns the details of an illustration.
// It needs login.
func (c *Client) IllustDetail(filter *IllustDetailFilter) (*IllustrationDetail, error) {
	illustration := &IllustrationDetail{}
	if err := c.doGetRequest("/v1/illust/detail", filter, illustration); err != nil {
		return nil, err
	}
	return illustration, nil
}
