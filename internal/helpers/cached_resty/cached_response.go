package cached_resty

import "github.com/go-resty/resty/v2"

type CachedResponse struct {
	*resty.Response

	CachedBody []byte
}

func (r *CachedResponse) IsCached() bool {
	return len(r.CachedBody) > 0
}

func (r *CachedResponse) Body() []byte {
	if r.IsCached() {
		return r.CachedBody
	}
	return r.Response.Body()
}
