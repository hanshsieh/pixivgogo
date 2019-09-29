package pixivgogo

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/imroc/req"
)

// tokenExpiryDelta determines how earlier a token should be considered
// expired than its actual expiration time. It is used to avoid late
// expiration due to client-server time mismatches.
const (
	tokenExpiryDelta = 10 * time.Second
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
	User  *User  `json:"user"`
}

// Valid reports whether t is non-nil, has an AccessToken, and is not expired.
func (t *Token) Valid() bool {
	return t != nil && t.AccessToken != "" && !t.Expired()
}

func (t *Token) Expired() bool {
	if t.ExpiryTime.IsZero() {
		return false
	}
	return t.ExpiryTime.Round(0).Add(-tokenExpiryDelta).Before(time.Now())
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

// TokenSource is used for providing tokens.
// It references the implementation of https://github.com/golang/oauth2/blob/aaccbc9213b0974828f81aaac109d194880e3014/oauth2.go
type TokenSource interface {
	Token() (*Token, error)
}

type cachedTokenSource struct {
	delegate TokenSource

	lock  sync.Mutex
	token *Token
}

// Token returns the current token if it's still valid, else will
// refresh the current token (using r.Context for HTTP client
// information) and return the new one.
// Notice this implementation won't perform well in multi-threaded environment.
func (s *cachedTokenSource) Token() (*Token, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.token.Valid() {
		return s.token, nil
	}
	newToken, err := s.delegate.Token()
	if err != nil {
		return nil, err
	}
	s.token = newToken
	return newToken, nil
}

type refreshTokenSource struct {
	client       *Client
	refreshToken string
}

func (r *refreshTokenSource) Token() (*Token, error) {
	token, err := r.client.RefreshToken(r.refreshToken)
	if err != nil {
		return nil, err
	}
	r.refreshToken = token.RefreshToken
	return token, err
}

type failingTokenSource struct{}

func (f *failingTokenSource) Token() (*Token, error) {
	return nil, errors.New("login is required")
}

func (c *Client) CreateToken(username, password string) (*TokenResponse, error) {
	clientTime := time.Now().UTC().Format(time.RFC3339)
	clientHash := fmt.Sprintf("%x", md5.Sum([]byte(clientTime+hashSecret)))
	headers := req.Header{
		"User-Agent":    "PixivAndroidApp/5.0.64 (Android 6.0)",
		"X-Client-Time": clientTime,
		"X-Client-Hash": clientHash,
	}
	reqURL := fmt.Sprintf("%s/auth/token", c.authURL)

	reqBody := req.Param{
		"get_secure_url": 1,
		"client_id":      clientID,
		"client_secret":  clientSecret,
		"grant_type":     "password",
		"username":       username,
		"password":       password,
	}
	req.Debug = true
	resp, err := c.client.Post(reqURL, headers, reqBody)
	tokenResp := &TokenResponse{}
	if err := c.unmarshalResponse(resp, err, tokenResp); err != nil {
		return nil, err
	}
	return tokenResp, nil
}

func (c *Client) Login(username, password string) error {
	tokenResp, err := c.CreateToken(username, password)
	if err != nil {
		return err
	}
	refreshToken := tokenResp.Response.RefreshToken
	c.tokenSource = &cachedTokenSource{
		delegate: &refreshTokenSource{
			client:       c,
			refreshToken: refreshToken,
		},
		token: tokenResp.Response,
	}
	return nil
}

func (c *Client) Logout() error {
	// TODO Should call API to logout
	c.tokenSource = &failingTokenSource{}
	return nil
}

func (c *Client) RefreshToken(refreshToken string) (*Token, error) {
	// TODO
	return nil, nil
}
