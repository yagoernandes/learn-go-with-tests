package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (s *SpyStore) Fetch() string {
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertWasCanceled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCanceled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("it should have not cancelled the store")
	}
}

func TestServer(t *testing.T) {
	t.Run("happy case", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %v, want %v", response.Body.String(), data)
		}

		store.assertWasNotCanceled()
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, the world can be cancelled"
		store := &SpyStore{response: data, t: t}
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		cancel()
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		store.assertWasCanceled()
	})
}
