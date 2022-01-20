package urlChecker

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
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
	count  int
}

func NewUrlChekcer(url string, count int) *UrlChecker {
	return &UrlChecker{
		url:   url,
		count: count,
	}
}

// http 체크 실행
func (uc *UrlChecker) Run() *resty.Response {
	uc.setClient()
	resp, err := uc.client.R().
		EnableTrace().
		Get(uc.url)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// 체크 후 출력
func (uc *UrlChecker) Print() {
	w := tabwriter.NewWriter(os.Stdout, 1, 2, 2, ' ', 0)
	fmt.Fprintln(w, "URL\tStatus\tRemoteAddr\tTotal_Time\tDNSLookup\tTCPConnTime\tTLSHandshake\tServerTime\tResponseTime\t")
	for i := 0; i < uc.count; i++ {
		fmt.Printf("Running ... %dst\n", i+1)
		resp := uc.Run()
		ti := resp.Request.TraceInfo()
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t\n",
			uc.url, resp.Status(), ti.RemoteAddr.String(), resp.Time(), ti.DNSLookup, ti.TCPConnTime, ti.TLSHandshake, ti.ServerTime, ti.ResponseTime)
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
	w.Flush()
}

// 헤더 설정
func (uc *UrlChecker) SetHeader(m map[string]string) {
	uc.header = m
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
