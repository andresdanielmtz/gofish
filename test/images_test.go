package api_test

import (
	"gofish/api"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestGetImagesHandler(t *testing.T) {
	dummyFolder := "images"
	os.MkdirAll(dummyFolder, 0755)
	dummyFile := "images/testfile.png"
	os.WriteFile(dummyFile, []byte("dummy"), 0644)
	defer os.RemoveAll(dummyFolder)

	// Request to /images
	req, err := http.NewRequest("GET", "/images", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.GetFilesHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	body := rr.Body.String()
	if !strings.Contains(body, "testfile.png") {
		t.Errorf("handler response does not contain expected filename: %v", body)
	}
	if !strings.Contains(body, `<a href="/">Home</a>`) {
		t.Errorf("handler response does not contain Home link")
	}
}
