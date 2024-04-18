package main_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"lotchiexample.com/customserver/handlers"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	handlers.Home(w, req)

	res := w.Result()
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)

	indexHtml, _ := ioutil.ReadFile("templates/index.html")
	text := string(indexHtml)
	if string(data) != text {
		t.Errorf("expected %v got %v", text, string(data))
	}
}
