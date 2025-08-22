package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gin-gonic/gin"
)

// ---- mock client ----
type mockDDB struct {
	out *dynamodb.QueryOutput
	err error
}

func (m *mockDDB) Query(_ context.Context, _ *dynamodb.QueryInput, _ ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	return m.out, m.err
}

// helper to run a GET request against a handler and return status + response bytes
func performRequest(h gin.HandlerFunc, path string) (int, []byte) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET(path, h)

	req := httptest.NewRequest(http.MethodGet, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w.Code, w.Body.Bytes()
}

func TestGetNews(t *testing.T) {
	// isolate env used by handler
	t.Setenv("AWS_REGION", "us-east-1")
	t.Setenv("NEWS_TABLE", "jonathanface_news")

	tests := []struct {
		name           string
		mockClient     ddbAPI
		overrideClient bool
		wantStatus     int
		wantArrayLen   int // for 200 OK
	}{
		{
			name:           "client init error",
			overrideClient: false, // force factory to return error
			wantStatus:     http.StatusInternalServerError,
		},
		{
			name:           "query error",
			overrideClient: true,
			mockClient:     &mockDDB{out: nil, err: errors.New("boom")},
			wantStatus:     http.StatusInternalServerError,
		},
		{
			name:           "success returns list",
			overrideClient: true,
			mockClient: &mockDDB{
				out: &dynamodb.QueryOutput{
					// two empty items are enough; handler will unmarshal into []models.News (zero values)
					Items: []map[string]types.AttributeValue{
						{}, {},
					},
				},
				err: nil,
			},
			wantStatus:   http.StatusOK,
			wantArrayLen: 2,
		},
	}

	// save and restore original factory
	origFactory := newDynamoClientFn
	t.Cleanup(func() { newDynamoClientFn = origFactory })

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.overrideClient {
				newDynamoClientFn = func() (ddbAPI, error) {
					return tt.mockClient, nil
				}
			} else {
				// make the factory return an error
				newDynamoClientFn = func() (ddbAPI, error) {
					return nil, errors.New("init error")
				}
			}

			status, body := performRequest(GetNews, "/api/news")

			if status != tt.wantStatus {
				t.Fatalf("status = %d, want %d. body=%s", status, tt.wantStatus, string(body))
			}

			if status == http.StatusOK {
				var arr []map[string]any
				if err := json.Unmarshal(body, &arr); err != nil {
					t.Fatalf("invalid JSON array: %v; body=%s", err, string(body))
				}
				if len(arr) != tt.wantArrayLen {
					t.Fatalf("array len = %d, want %d; body=%s", len(arr), tt.wantArrayLen, string(body))
				}
			}
		})
	}
}

func TestGetBooks(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-1")
	t.Setenv("BOOKS_TABLE", "jonathanface_books2")

	tests := []struct {
		name           string
		mockClient     ddbAPI
		overrideClient bool
		wantStatus     int
		wantArrayLen   int
	}{
		{
			name:           "client init error",
			overrideClient: false,
			wantStatus:     http.StatusInternalServerError,
		},
		{
			name:           "query error",
			overrideClient: true,
			mockClient:     &mockDDB{out: nil, err: errors.New("boom")},
			wantStatus:     http.StatusInternalServerError,
		},
		{
			name:           "success returns list",
			overrideClient: true,
			mockClient: &mockDDB{
				out: &dynamodb.QueryOutput{
					Items: []map[string]types.AttributeValue{
						{}, {}, {}, {},
					},
				},
			},
			wantStatus:   http.StatusOK,
			wantArrayLen: 4,
		},
	}

	origFactory := newDynamoClientFn
	t.Cleanup(func() { newDynamoClientFn = origFactory })

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.overrideClient {
				newDynamoClientFn = func() (ddbAPI, error) {
					return tt.mockClient, nil
				}
			} else {
				newDynamoClientFn = func() (ddbAPI, error) {
					return nil, errors.New("init error")
				}
			}

			status, body := performRequest(GetBooks, "/api/books")

			if status != tt.wantStatus {
				t.Fatalf("status = %d, want %d. body=%s", status, tt.wantStatus, string(body))
			}

			if status == http.StatusOK {
				var arr []map[string]any
				if err := json.Unmarshal(body, &arr); err != nil {
					t.Fatalf("invalid JSON array: %v; body=%s", err, string(body))
				}
				if len(arr) != tt.wantArrayLen {
					t.Fatalf("array len = %d, want %d; body=%s", len(arr), tt.wantArrayLen, string(body))
				}
			}
		})
	}
}
