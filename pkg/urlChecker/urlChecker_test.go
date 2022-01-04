package urlChecker

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 체크 테스트
func TestUrlChecker(t *testing.T) {
	urlChecker := NewUrlChekcer("https://www.google.com")
	resp := urlChecker.Run()

	fmt.Println(resp.Request.TraceInfo().ResponseTime.Seconds())

	// fmt.Println(resp.Header())
	fmt.Println(resp.Request.URL)
	fmt.Println(resp.Request.Method)
	h, err := json.MarshalIndent(resp.Header(), "", "    ")
	fmt.Println(err)
	fmt.Println(string(h))

	prettyResp, err := json.MarshalIndent(resp.Request.TraceInfo(), "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(prettyResp))

	assert.Equal(t, resp.StatusCode(), 200)
}

// 헤더설정 테스트
func TestUrlCheckerSetHeader(t *testing.T) {
	urlChecker := NewUrlChekcer("https://www.google.com")

	m := map[string]string{"test": "test"}
	urlChecker.SetHeader(m)

	assert.Equal(t, len(urlChecker.header), 1)
}
