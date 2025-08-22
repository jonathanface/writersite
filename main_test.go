package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// stub handlers used to avoid hitting real DynamoDB-backed endpoints
func stubOK(body string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Data(http.StatusOK, "application/json", []byte(body))
	}
}

// helper to perform GET requests
func doGet(r http.Handler, path string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestBuildRouter_HealthAndAPI(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := buildRouter(
		stubOK(`{"news":"ok"}`),
		stubOK(`{"books":"ok"}`),
	)

	tests := []struct {
		name           string
		path           string
		wantStatus     int
		wantBodySubstr string
	}{
		{"health ok", "/health", http.StatusOK, "null"}, // your handler returns JSON nil => "null"
		{"news ok", "/api/news", http.StatusOK, `"news":"ok"`},
		{"books ok", "/api/books", http.StatusOK, `"books":"ok"`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			w := doGet(r, tc.path)
			if w.Code != tc.wantStatus {
				t.Fatalf("status=%d want=%d body=%s", w.Code, tc.wantStatus, w.Body.String())
			}
			if tc.wantBodySubstr != "" && !strings.Contains(w.Body.String(), tc.wantBodySubstr) {
				t.Fatalf("body=%q does not contain %q", w.Body.String(), tc.wantBodySubstr)
			}
		})
	}
}

func TestBuildRouter_SPAFallbackServesIndex(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a temp working dir containing ./ui/index.html
	tmp := t.TempDir()
	uiPath := filepath.Join(tmp, uiDirectory)
	if err := os.MkdirAll(uiPath, 0o755); err != nil {
		t.Fatal(err)
	}
	indexContent := "<!doctype html><html><body>SPA INDEX</body></html>"
	if err := os.WriteFile(filepath.Join(uiPath, "index.html"), []byte(indexContent), 0o644); err != nil {
		t.Fatal(err)
	}

	// Change working dir so static.Serve("ui", ...) finds files
	oldWD, _ := os.Getwd()
	t.Cleanup(func() { _ = os.Chdir(oldWD) })
	if err := os.Chdir(tmp); err != nil {
		t.Fatal(err)
	}

	r := buildRouter(stubOK(`{}`), stubOK(`{}`))

	// Unknown route should serve index.html (NoRoute)
	w := doGet(r, "/some/unknown/route")
	if w.Code != http.StatusOK {
		t.Fatalf("status=%d want=%d body=%s", w.Code, http.StatusOK, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "SPA INDEX") {
		t.Fatalf("expected SPA index content, got: %s", w.Body.String())
	}
}

func TestBuildRouter_StaticFileServed(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tmp := t.TempDir()
	uiPath := filepath.Join(tmp, uiDirectory)
	if err := os.MkdirAll(filepath.Join(uiPath, "public"), 0o755); err != nil {
		t.Fatal(err)
	}
	// create a static asset inside ui/public
	if err := os.WriteFile(filepath.Join(uiPath, "public", "hello.txt"), []byte("hello world"), fs.FileMode(0o644)); err != nil {
		t.Fatal(err)
	}

	oldWD, _ := os.Getwd()
	t.Cleanup(func() { _ = os.Chdir(oldWD) })
	if err := os.Chdir(tmp); err != nil {
		t.Fatal(err)
	}

	r := buildRouter(stubOK(`{}`), stubOK(`{}`))

	// static middleware is mounted at "/" with LocalFile("ui", false)
	w := doGet(r, "/public/hello.txt")
	if w.Code != http.StatusOK {
		t.Fatalf("status=%d want=%d", w.Code, http.StatusOK)
	}
	if strings.TrimSpace(w.Body.String()) != "hello world" {
		t.Fatalf("unexpected body: %q", w.Body.String())
	}
}
