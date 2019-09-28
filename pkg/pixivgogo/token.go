package pixivgogo

import (
	"errors"
	"sync"

	"github.com/imroc/req"
)

const (
	// Client ID used for talking with Pixiv server.
	// Copied from https://github.com/upbit/pixivpy/blob/bcde6af7bf3590124c89e1d1b8818f915db8f16d/pixivpy3/api.py#L14-L16
	clientID     = "MOBrBDS8blbauoSck0ZfDbtuzpyT"
	clientSecret = "lsACyCD94FhDUtGTXi3QzcFE2uU1hqtDaKeqrdwj"
	hashSecret   = "28c1fdd170a5204386cb1313c7077b34f83e4aaf4aa829ce78c231e05b0bae2c"
)

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

type Client struct {
	tokenSource TokenSource
	client      *req.Req
}

func NewClient() *Client {
	return &Client{
		tokenSource: &failingTokenSource{},
		client:      req.New(),
	}
}

func (c *Client) GetToken(username, password string) (*Token, error) {
	// TODO
	return nil, nil
}

func (c *Client) Login(username, password string) error {
	token, err := c.GetToken(username, password)
	if err != nil {
		return err
	}
	c.tokenSource = &cachedTokenSource{
		delegate: &refreshTokenSource{
			client:       c,
			refreshToken: token.RefreshToken,
		},
		token: token,
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
