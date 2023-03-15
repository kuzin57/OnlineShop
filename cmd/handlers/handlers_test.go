package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PageHandlersTestSuite struct {
	suite.Suite
	htmlTemps     []string
	contentOfHTML []byte
}

func (suite *PageHandlersTestSuite) SetupTest() {
	suite.htmlTemps = []string{
		"../../ui/html/home.html",
		"../../ui/html/base.html",
		"../../ui/html/footer.html",
	}

	for _, file := range suite.htmlTemps {
		if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
			suite.T().Fatalf("file %s does not exist", file)
		}
	}

	var err error
	suite.contentOfHTML, err = os.ReadFile("./correct_test_cases/home_page_correct_response.html")
	if err != nil {
		suite.T().Fatal(err)
	}
}

func (suite *PageHandlersTestSuite) TestAddHomePageHandler() {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		suite.T().Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(htmlSources(suite.htmlTemps).homePageHandler)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		suite.T().Errorf("status code is not OK!")
	}

	if rr.Body.String() != string(suite.contentOfHTML) {
		suite.T().Errorf("Unexpected body")
	}
}

func TestPageHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(PageHandlersTestSuite))
}
