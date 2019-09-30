package pixivgogo

type Account struct {
	ID            int64             `json:"id"`
	ProfileImages *ProfileImageURLs `json:"profile_image_urls"`
	Name          string            `json:"name"`
	AccountName   string            `json:"account"`
	Followed      bool              `json:"is_followed"`
}

type ProfileImageURLs struct {
	SizeMedium string `json:"medium,omitempty"`
}
