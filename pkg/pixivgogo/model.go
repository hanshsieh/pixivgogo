package pixivgogo

import (
	"encoding/json"
	"time"
)

// expiryDelta determines how earlier a token should be considered
// expired than its actual expiration time. It is used to avoid late
// expiration due to client-server time mismatches.
const (
	expiryDelta = 10 * time.Second
)

type TokenResponse struct {
	Response *Token `json:"response"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	DeviceToken  string `json:"device_token"`

	// ExpiryTime is the number of seconds the access token will expire.
	ExpiryTime TokenExpireTime `json:"expires_in"`

	// TokenType the token type, such as "bearer".
	TokenType TokenType `json:"token_type"`

	// Scope is the scope of the token, may be empty string.
	Scope string `json:"scope"`
}

// Valid reports whether t is non-nil, has an AccessToken, and is not expired.
func (t *Token) Valid() bool {
	return t != nil && t.AccessToken != "" && !t.Expired()
}

func (t *Token) Expired() bool {
	if t.ExpiryTime.IsZero() {
		return false
	}
	return t.ExpiryTime.Round(0).Add(-expiryDelta).Before(time.Now())
}

type TokenExpireTime struct {
	time.Time
}

func (t *TokenExpireTime) UnmarshalJSON(b []byte) error {
	var expireSec int64
	if err := json.Unmarshal(b, &expireSec); err != nil {
		return err
	}
	newTime := time.Now().Add(time.Duration(expireSec) * time.Second)
	*t = TokenExpireTime{newTime}
	return nil
}

type TokenType string

type Account struct {
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
