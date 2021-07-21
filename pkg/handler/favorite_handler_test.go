package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFavorites(t *testing.T) {
	reqBody := bytes.NewBufferString("")
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000", reqBody)
	res := httptest.NewRecorder()

	GetFavorites(res, req)

	assert.Equal(t, res.Code, 200)
}

func TestCreateFavorite(t *testing.T) {
	values := url.Values{}
	values.Set("video_id", "smash_sp")

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000", bytes.NewBufferString(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res := httptest.NewRecorder()

	CreateFavorite(res, req)

	assert.Equal(t, res.Code, 200)
	assert.True(t, strings.Contains(res.Body.String(), "smash_sp"))
}
