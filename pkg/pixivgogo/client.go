package pixivgogo

import (
	"github.com/imroc/req"
)

const (
	DefaultAuthURL = "https://oauth.secure.pixiv.net"
	DefaultAPIURL  = "https://app-api.pixiv.net"
	// Client ID used for talking with Pixiv server.
	// Copied from https://github.com/upbit/pixivpy/blob/bcde6af7bf3590124c89e1d1b8818f915db8f16d/pixivpy3/api.py#L14-L16
	clientID             = "MOBrBDS8blbauoSck0ZfDbtuzpyT"
	clientSecret         = "lsACyCD94FhDUtGTXi3QzcFE2uU1hqtDaKeqrdwj"
	hashSecret           = "28c1fdd170a5204386cb1313c7077b34f83e4aaf4aa829ce78c231e05b0bae2c"
	minSuccessStatusCode = 200
	maxSuccessStatusCode = 299
)

type Client struct {
	tokenSource TokenSource
	client      *req.Req
	authURL     string
	apiURL      string
}

func NewClient() *Client {
	return &Client{
		tokenSource: &emptyTokenSource{},
		client:      req.New(),
		authURL:     DefaultAuthURL,
		apiURL:      DefaultAPIURL,
	}
}

func (c *Client) isErrorStatusCode(statusCode int) bool {
	return statusCode < 200 || statusCode >= 300
}

func (c *Client) unmarshalAuthResponse(
	resp *req.Resp,
	reqError error,
	successResp interface{}) error {
	errResp := &AuthError{}
	return c.unmarshalResponse(resp, reqError, successResp, errResp)
}

func (c *Client) unmarshalAPIResponse(resp *req.Resp,
	reqError error,
	successResp interface{}) error {
	errResp := &APIError{}
	return c.unmarshalResponse(resp, reqError, successResp, errResp)
}

func (c *Client) unmarshalResponse(
	resp *req.Resp,
	reqError error,
	successResp interface{},
	errorResp error) error {
	if reqError != nil {
		return reqError
	}
	statusCode := resp.Response().StatusCode
	if statusCode < minSuccessStatusCode || statusCode > maxSuccessStatusCode {
		if err := resp.ToJSON(errorResp); err != nil {
			return err
		}
		return errorResp
	} else {
		if err := resp.ToJSON(successResp); err != nil {
			return err
		}
		return nil
	}
}

func (c *Client) createHeaders() (req.Header, error) {
	token, err := c.tokenSource.Token()
	if err != nil {
		return nil, err
	}
	if token == nil {
		return req.Header{}, nil
	}
	return req.Header{
		"Authorization": "Bearer " + token.AccessToken,
	}, nil
}
