package http

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 一个简单的回显处理器：把下游看到的 r.URL.Path 写回响应体
var echoPathHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, r.URL.Path)
})

func TestServeMux_TrailingSlash_Matching(t *testing.T) {
	mux := http.NewServeMux()

	// 精确匹配：只匹配 /images
	mux.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "exact")
	})
	// 前缀匹配：匹配 /images/ 及其子路径
	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "prefix")
	})

	tests := []struct {
		path       string
		wantBody   string
		wantStatus int
	}{
		{"/images", "exact", http.StatusOK},          // 命中精确规则
		{"/images/", "prefix", http.StatusOK},        // 命中前缀规则
		{"/images/foo.png", "prefix", http.StatusOK}, // 命中前缀规则
		{"/image", "", http.StatusNotFound},          // 没有匹配
	}

	for _, tc := range tests {
		req := httptest.NewRequest(http.MethodGet, tc.path, nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != tc.wantStatus {
			t.Fatalf("path=%q status=%d want=%d body=%q",
				tc.path, rr.Code, tc.wantStatus, rr.Body.String())
		}
		if tc.wantBody != "" && rr.Body.String() != tc.wantBody {
			t.Fatalf("path=%q body=%q want=%q",
				tc.path, rr.Body.String(), tc.wantBody)
		}
	}
}

func TestStripPrefix_CorrectRegistration_WithTrailingSlash(t *testing.T) {
	mux := http.NewServeMux()

	// ✅ 正确：注册为 /static/（前缀匹配），并 StripPrefix("/static/")
	mux.Handle("/static/",
		http.StripPrefix("/static/", echoPathHandler))

	req := httptest.NewRequest(http.MethodGet, "/static/css/app.css", nil)
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status=%d want=%d", rr.Code, http.StatusOK)
	}
	got := rr.Body.String()
	want := "css/app.css" // 去掉了 /static/ 前缀
	if got != want {
		t.Fatalf("echoed path=%q want=%q", got, want)
	}
}

func TestStripPrefix_WrongRegistration_WithoutTrailingSlash(t *testing.T) {
	mux := http.NewServeMux()

	// ❌ 错误：注册为 /static（只精确匹配 /static），虽然 StripPrefix("/static") 看似对，
	// 但 /static/app.js 根本不会路由到这个 handler，直接 404。
	mux.Handle("/static", http.StripPrefix("/static", echoPathHandler))

	req := httptest.NewRequest(http.MethodGet, "/static/app.js", nil)
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("status=%d want=%d (因为没有使用 /static/ 做前缀匹配)",
			rr.Code, http.StatusNotFound)
	}
}
