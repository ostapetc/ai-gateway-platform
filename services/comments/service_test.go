package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/config"
	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/handler"
	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/store"
	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/svc"
	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/types"

	"github.com/zeromicro/go-zero/rest"
)

func startServer(t *testing.T) (baseURL string, stop func()) {
	t.Helper()

	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("find free port: %v", err)
	}
	port := l.Addr().(*net.TCPAddr).Port
	l.Close()

	var cfg config.Config
	cfg.RestConf.Name = "comments-test"
	cfg.RestConf.Host = "127.0.0.1"
	cfg.RestConf.Port = port

	svcCtx := &svc.ServiceContext{
		Config:       cfg,
		CommentStore: store.NewCommentStore(),
	}

	srv := rest.MustNewServer(cfg.RestConf)
	handler.RegisterHandlers(srv, svcCtx)
	go srv.Start()

	base := fmt.Sprintf("http://127.0.0.1:%d", port)
	waitReady(t, base+"/comments")
	return base, srv.Stop
}

func waitReady(t *testing.T, url string) {
	t.Helper()
	deadline := time.Now().Add(3 * time.Second)
	for time.Now().Before(deadline) {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
	t.Fatal("server did not become ready in time")
}

func doPost(t *testing.T, url string, body any) *http.Response {
	t.Helper()
	b, _ := json.Marshal(body)
	resp, err := http.Post(url, "application/json", bytes.NewReader(b))
	if err != nil {
		t.Fatalf("POST %s: %v", url, err)
	}
	return resp
}

func decode[T any](t *testing.T, resp *http.Response) T {
	t.Helper()
	var v T
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	return v
}

func TestCreateComment(t *testing.T) {
	cases := []struct {
		name string
		req  types.CreateCommentRequest
	}{
		{"simple comment", types.CreateCommentRequest{UserID: 1, PostID: 1, Body: "great post"}},
		{"another post", types.CreateCommentRequest{UserID: 2, PostID: 5, Body: "nice work"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			base, stop := startServer(t)
			defer stop()

			resp := doPost(t, base+"/comments", tc.req)
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("want 200, got %d", resp.StatusCode)
			}

			got := decode[types.CreateCommentResponse](t, resp)
			if got.ID == 0 {
				t.Fatal("want non-zero id in response")
			}
		})
	}
}

func TestListComments(t *testing.T) {
	t.Run("empty for unknown post", func(t *testing.T) {
		base, stop := startServer(t)
		defer stop()

		resp, err := http.Get(base + "/comments?post_id=999")
		if err != nil {
			t.Fatalf("GET /comments/999: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("want 200, got %d", resp.StatusCode)
		}

		got := decode[types.ListCommentsResponse](t, resp)
		if len(got.Comments) != 0 {
			t.Fatalf("want 0 comments, got %d", len(got.Comments))
		}
		if got.Total != 0 {
			t.Fatalf("want total 0, got %d", got.Total)
		}
	})

	t.Run("created comment appears in list", func(t *testing.T) {
		base, stop := startServer(t)
		defer stop()

		req := types.CreateCommentRequest{UserID: 1, PostID: 10, Body: "integration test comment"}

		createResp := doPost(t, base+"/comments", req)
		defer createResp.Body.Close()
		created := decode[types.CreateCommentResponse](t, createResp)

		listResp, err := http.Get(fmt.Sprintf("%s/comments?post_id=%d", base, req.PostID))
		if err != nil {
			t.Fatalf("GET /comments/%d: %v", req.PostID, err)
		}
		defer listResp.Body.Close()

		if listResp.StatusCode != http.StatusOK {
			t.Fatalf("want 200, got %d", listResp.StatusCode)
		}

		list := decode[types.ListCommentsResponse](t, listResp)
		if len(list.Comments) != 1 {
			t.Fatalf("want 1 comment, got %d", len(list.Comments))
		}
		if list.Total != 1 {
			t.Fatalf("want total 1, got %d", list.Total)
		}

		c := list.Comments[0]
		if c.ID != created.ID {
			t.Errorf("id: want %d, got %d", created.ID, c.ID)
		}
		if c.UserID != req.UserID {
			t.Errorf("user_id: want %d, got %d", req.UserID, c.UserID)
		}
		if c.PostID != req.PostID {
			t.Errorf("post_id: want %d, got %d", req.PostID, c.PostID)
		}
		if c.Body != req.Body {
			t.Errorf("body: want %q, got %q", req.Body, c.Body)
		}
		if c.CreatedAt == "" {
			t.Error("created_at must not be empty")
		}
	})

	t.Run("filters by post id", func(t *testing.T) {
		base, stop := startServer(t)
		defer stop()

		for i := 0; i < 2; i++ {
			r := doPost(t, base+"/comments", types.CreateCommentRequest{UserID: 1, PostID: 1, Body: "post 1"})
			r.Body.Close()
		}
		for i := 0; i < 3; i++ {
			r := doPost(t, base+"/comments", types.CreateCommentRequest{UserID: 1, PostID: 2, Body: "post 2"})
			r.Body.Close()
		}

		check := func(postID uint64, wantCount int) {
			t.Helper()
			resp, err := http.Get(fmt.Sprintf("%s/comments?post_id=%d", base, postID))
			if err != nil {
				t.Fatalf("GET /comments/%d: %v", postID, err)
			}
			defer resp.Body.Close()

			list := decode[types.ListCommentsResponse](t, resp)
			if len(list.Comments) != wantCount {
				t.Errorf("post %d: want %d comments, got %d", postID, wantCount, len(list.Comments))
			}
			if list.Total != wantCount {
				t.Errorf("post %d: want total %d, got %d", postID, wantCount, list.Total)
			}
			for _, c := range list.Comments {
				if c.PostID != postID {
					t.Errorf("comment %d: wrong post_id, want %d got %d", c.ID, postID, c.PostID)
				}
			}
		}

		check(1, 2)
		check(2, 3)
		check(999, 0)
	})
}
