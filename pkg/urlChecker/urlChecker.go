package urlChecker

import (
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	TIMEOUT_SECONDS = 2 * time.Second
)

type UrlChecker struct {
	url    string
	header map[string]string
	client *resty.Client
}

func NewUrlChekcer(url string) *UrlChecker {
	return &UrlChecker{
		url: url,
	}
}

// http 체크 실행
func (uc *UrlChecker) Run() *resty.Response {
	uc.setClient()
	resp, err := uc.client.
		R().
		EnableTrace().
		Get(uc.url)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// 헤더 설정
func (uc *UrlChecker) SetHeader(m map[string]string) *UrlChecker {
	uc.header = m
	return uc
}

// resty http client 설정 (internal)
func (uc *UrlChecker) setClient() *UrlChecker {
	uc.client = resty.New()
	uc.client.SetTimeout(TIMEOUT_SECONDS)

	reqp := uc.client.R()
	if len(uc.header) > 0 {
		reqp.SetHeaders(uc.header)
	}
	return uc
}
