package cached_resty

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/go-resty/resty/v2"
	"github.com/legion-zver/premier-one-bleve-search/internal/utils"
)

type CachedRequest struct {
	*resty.Request

	cacheDirPath string
}

func (r *CachedRequest) SetDoNotParseResponse(parse bool) *CachedRequest {
	r.Request.SetDoNotParseResponse(parse)
	return r
}

func (r *CachedRequest) SetHeader(header, value string) *CachedRequest {
	r.Header.Set(header, value)
	return r
}

func (r *CachedRequest) SetCookies(rs []*http.Cookie) *CachedRequest {
	r.Request.SetCookies(rs)
	return r
}

// Get method does GET HTTP request. It's defined in section 4.3.1 of RFC7231.
func (r *CachedRequest) Get(url string) (*CachedResponse, error) {
	return r.Execute(resty.MethodGet, url)
}

// Head method does HEAD HTTP request. It's defined in section 4.3.2 of RFC7231.
func (r *CachedRequest) Head(url string) (*CachedResponse, error) {
	return r.Execute(resty.MethodHead, url)
}

// Post method does POST HTTP request. It's defined in section 4.3.3 of RFC7231.
func (r *CachedRequest) Post(url string) (*CachedResponse, error) {
	return r.Execute(resty.MethodPost, url)
}

// Put method does PUT HTTP request. It's defined in section 4.3.4 of RFC7231.
func (r *CachedRequest) Put(url string) (*CachedResponse, error) {
	return r.Execute(resty.MethodPut, url)
}

// Delete method does DELETE HTTP request. It's defined in section 4.3.5 of RFC7231.
func (r *CachedRequest) Delete(url string) (*CachedResponse, error) {
	return r.Execute(resty.MethodDelete, url)
}

// Options method does OPTION HTTP request. It's defined in section 4.3.7 of RFC7231.
func (r *CachedRequest) Options(url string) (*CachedResponse, error) {
	return r.Execute(resty.MethodOptions, url)
}

// Patch method does PATCH HTTP request. It's defined in section 2 of RFC5789.
func (r *CachedRequest) Patch(url string) (*CachedResponse, error) {
	return r.Execute(resty.MethodPatch, url)
}

func (r *CachedRequest) Execute(method, url string) (*CachedResponse, error) {
	cacheFilePath := path.Join(r.cacheDirPath, fmt.Sprintf("%s_%s.json", method, utils.SHA1(url)))
	if _, err := os.Stat(cacheFilePath); err == nil || os.IsExist(err) {
		body, err := os.ReadFile(cacheFilePath)
		if err == nil {
			return &CachedResponse{
				Response: &resty.Response{
					Request:     r.Request,
					RawResponse: nil,
				},
				CachedBody: body,
			}, nil
		}
	}
	resp, err := r.Request.Execute(method, url)
	if err != nil {
		return nil, err
	}
	if body := resp.Body(); len(body) > 0 {
		_ = os.WriteFile(cacheFilePath, body, os.ModePerm)
	}
	return &CachedResponse{Response: resp}, nil
}
