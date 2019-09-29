package pixivgogo

type User struct {
	ID              string           `json:"id"`
	ProfileImages   ProfileImageURLs `json:"profile_image_urls"`
	Name            string           `json:"name"`
	AccountName     string           `json:"account"`
	Email           string           `json:"mail_address"`
	Premium         bool             `json:"is_premium"`
	XRestrict       int              `json:"x_restrict"`
	EmailAuthorized bool             `json:"is_mail_authorized"`
}

type ProfileImageURLs struct {
	URL16x16   string `json:"px_16x16"`
	URL50x50   string `json:"px_50x50"`
	URL170x170 string `json:"px_170x170"`
}
