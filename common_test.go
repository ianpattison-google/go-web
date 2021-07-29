package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []article

// TestMain performs setup before executing tests
func TestMain(m *testing.M) {
	// set Gin to test mode
	gin.SetMode(gin.TestMode)
	// run tests
	os.Exit(m.Run())
}

// getRouter creates a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

// testHTTPResponse processes a request and tests its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// create a response recorder
	w := httptest.NewRecorder()
	// create the service and process the request
	r.ServeHTTP(w, req)
	if !f(w) {
		t.Fail()
	}
}

// saveLists stores the main list into the temp one for testing
func saveLists() {
	tmpArticleList = articleList
}

// restoreLists restores the main list from the temp one
func restoreLists() {
	articleList = tmpArticleList
}
