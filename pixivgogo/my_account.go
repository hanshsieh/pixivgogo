package pixivgogo

type MyAccount struct {
	ID              string              `json:"id"`
	ProfileImages   *MyProfileImageURLs `json:"profile_image_urls"`
	Name            string              `json:"name"`
	AccountName     string              `json:"account"`
	Email           string              `json:"mail_address"`
	Premium         *bool               `json:"is_premium,omitempty"`
	XRestrict       *int                `json:"x_restrict,omitempty"`
	EmailAuthorized *bool               `json:"is_mail_authorized,omitempty"`
	Followed        *bool               `json:"is_followed"`
}

type MyProfileImageURLs struct {
	Size16x16   string `json:"px_16x16,omitempty"`
	Size50x50   string `json:"px_50x50,omitempty"`
	Size170x170 string `json:"px_170x170,omitempty"`
	SizeMedium  string `json:"medium,omitempty"`
}
