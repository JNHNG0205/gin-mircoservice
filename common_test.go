package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []Article

// this function runs before all tests in the package
func TestMain(m *testing.M) {
	// set the router to test mode
	gin.SetMode(gin.TestMode)
	// run the other tests
	os.Exit(m.Run())
}

// helper function to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
		r.Use(gin.Recovery())
	}
	return r
}

// helper function to process a request and test the response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// create a new recorder to record the response
	w := httptest.NewRecorder()
	// creates a service and processes the request to the recorder
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// save the lists to the temporary list
func saveLists() {
	tmpArticleList = articleList
}

// restore the lists from the temporary list
func restoreLists() {
	articleList = tmpArticleList
}
