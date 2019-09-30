package pixivgogo

import (
	"github.com/sleepingpig/pixivgogo/pkg/pixivgogo/datetime"
)

type RankingIllustsFilter struct {
	// Filter is the filter for the illustrations.
	// Possible values: "for_ios", "for_android"
	// This field is optional.
	Filter IllustFilter `schema:"filter,omitempty"`

	// Mode of the ranking.
	// This field is optional.
	// TODO What's the default?
	Mode RankMode `schema:"mode,omitempty"`

	// Date can be used to query the ranking in the past.
	// This field is optional.
	Date *datetime.Date `schema:"date,omitempty"`

	// Offset is the offset of the illustrations to query.
	// Starting from 0.
	// This field is optional.
	Offset int `schema:"offset,omitempty"`
}

type IllustFilter string

const (
	RANK_FILTER_FOR_IOS     IllustFilter = "for_ios"
	RANK_FILTER_FOR_ANDROID IllustFilter = "for_android"
)

type RankMode string

const (
	RANK_MODE_DAY            RankMode = "day"
	RANK_MODE_DAY_MALE       RankMode = "day_male"
	RANK_MODE_DAY_FEMALE     RankMode = "day_female"
	RANK_MODE_WEEK_ORIGINAL  RankMode = "week_original"
	RANK_MODE_WEEK_ROOKIE    RankMode = "week_rookie"
	RANK_MODE_MONTH_ROOKIE   RankMode = "month_rookie"
	RANK_MODE_WEEK           RankMode = "week"
	RANK_MODE_MONTH          RankMode = "month"
	RANK_MODE_DAY_R18        RankMode = "day_r18"
	RANK_MODE_WEEK_R18       RankMode = "week_r18"
	RANK_MODE_WEEK_R18G      RankMode = "week_r18g"
	RANK_MODE_DAY_MALE_R18   RankMode = "day_male_r18"
	RANK_MODE_DAY_FEMALE_R18 RankMode = "day_female_r18"
	RANK_MODE_DAY_MANGA      RankMode = "day_manga"
	RANK_MODE_WEEK_MANGA     RankMode = "week_manga"
	RANK_MODE_MONTH_MANGA    RankMode = "month_manga"
)

type RankingIllustrations struct {
	Illustrations []*Illustration       `json:"illusts"`
	NextURL       string                `json:"next_url,omitempty"`
	NextFilter    *RankingIllustsFilter `json:"-"`
}

type RecommendIllustrations struct {
	Illustrations        []*Illustration         `json:"illusts"`
	RankingIllustrations []*Illustration         `json:"ranking_illusts"`
	ContestExists        bool                    `json:"contest_exists"`
	PrivacyPolicy        *PrivacyPolicy          `json:"privacy_policy,omitempty"`
	NextURL              string                  `json:"next_url,omitempty"`
	NextFilter           *RecommendIllustsFilter `json:"-"`
}

type PrivacyPolicy struct {
	Version string `json:"version,omitempty"`
	Message string `json:"message,omitempty"`
	URL     string `json:"url,omitempty"`
}

type IllustrationDetail struct {
	Illustration *Illustration `json:"illust"`
}

type Illustration struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
	// ImageURLs contains the thumbnail image URLs of the illustration.
	// To get the original image URLs, see "MetaSinglePage" and "MetaPages".
	ImageURLs   *IllustImageURLs `json:"image_urls,omitempty"`
	Caption     string           `json:"caption"`
	Restrict    int              `json:"restrict"`
	User        *Account         `json:"user"`
	Tags        []Tag            `json:"tags"`
	Tools       []string         `json:"tools"`
	CreateDate  datetime.Time    `json:"create_date"`
	PageCount   int              `json:"page_count"`
	Width       int              `json:"width"`
	Height      int              `json:"height"`
	SanityLevel int              `json:"sanity_level"`
	XRestrict   int              `json:"x_restrict"`
	Series      *Series          `json:"series,omitempty"`

	// MetaSinglePage will contain non-empty URLs when the illustration contain
	// only one image.
	MetaSinglePage *MetaSinglePage `json:"meta_single_page,omitempty"`

	// MetaPages will be non-empty when the illustration contains multiple
	// images.
	MetaPages      []*MetaPage `json:"meta_pages,omitempty"`
	TotalView      int         `json:"total_view"`
	TotalBookmarks int         `json:"total_bookmarks"`
	Bookmarked     bool        `json:"is_bookmarked"`
	Visible        bool        `json:"visible"`
	Muted          bool        `json:"is_muted"`
}

// MetaPage contains the information of one of the multiple pages of an illustration.
type MetaPage struct {
	ImageURLs *MetaPageImageURLs `json:"image_urls"`
}

// MetaPageImageURLs contains the image URLs of a meta page.
type MetaPageImageURLs struct {
	SquareMedium string `json:"square_medium,omitempty"`
	Medium       string `json:"medium,omitempty"`
	Large        string `json:"large,omitempty"`
	Original     string `json:"original,omitempty"`
}

type Series struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

// MetaSinglePage contains the image URLs when the illustration contain only a single image.
type MetaSinglePage struct {
	// OriginalImageURL is the URL to the original size of the illustration.
	OriginalImageURL string `json:"original_image_url,omitempty"`
}

type IllustImageURLs struct {
	SquareMedium string `json:"square_medium,omitempty"`
	Medium       string `json:"medium,omitempty"`
	Large        string `json:"large,omitempty"`
}

type Tag struct {
	Name           string `json:"name,omitempty"`
	TranslatedName string `json:"translated_name,omitempty"`
}

type IllustDetailFilter struct {
	ID int64 `schema:"illust_id,omitempty"`
}

type RecommendIllustsFilter struct {
	Filter                       IllustFilter `schema:"filter,omitempty"`
	MinBookmarkIDForRecentIllust string       `schema:"min_bookmark_id_for_recent_illust,omitempty"`
	MaxBookmarkIDForRecommend    string       `schema:"max_bookmark_id_for_recommend,omitempty"`
	Offset                       int          `schema:"offset"`

	// IncludeRankingIllusts indicates that whether the illustration ranking should also be included
	// in the result.
	// Notice that even if it's set to true, the "next_url" in the response will still
	// set it to false.
	IncludeRankingIllusts bool `schema:"include_ranking_illusts"`

	// IncludePrivacyPolicy indicates that whether the privacy policy info should also be included
	// in the result.
	// Notice that even if it's set to true, the "next_url" in the response will still
	// set it to false.
	IncludePrivacyPolicy bool `schema:"include_privacy_policy"`
}

type ContentType string

const (
	TypeIllust ContentType = "illust"
	TypeManga  ContentType = "manga"
)
