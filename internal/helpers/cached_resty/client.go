package cached_resty

import "github.com/go-resty/resty/v2"

type Client struct {
	*resty.Client
}

func (c *Client) CachedR(dirPath string) *CachedRequest {
	return &CachedRequest{
		Request:      c.R(),
		cacheDirPath: dirPath,
	}
}

func (c *Client) NewCachedRequest(dirPath string) *CachedRequest {
	return c.CachedR(dirPath)
}

func New() *Client {
	return &Client{Client: resty.New()}
}
