package urlChecker

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 체크 테스트
func TestUrlChecker(t *testing.T) {
	urlChecker := NewUrlChekcer("https://www.google.com", 5)
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

	fmt.Printf("")

	assert.Equal(t, resp.StatusCode(), 200)
}

// 헤더설정 테스트
func TestUrlCheckerSetHeader(t *testing.T) {
	urlChecker := NewUrlChekcer("https://www.google.com", 5)

	m := map[string]string{"test": "test"}
	urlChecker.SetHeader(m)

	assert.NotEqual(t, len(urlChecker.header), 0)
}

// 프린트 테스트
func TestUrlCheckerPrint(t *testing.T) {
	urlChecker := NewUrlChekcer("https://www.google.com", 5)

	urlChecker.Run()

	assert.NotEqual(t, urlChecker, nil)
}
