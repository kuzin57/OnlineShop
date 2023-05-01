package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	pathToConf = "/home/kuzin57/OnlineShop/cmd/config/page_handlers.yaml"
)

type PageHandlersTestSuite struct {
	suite.Suite
	contentOfHTML []byte
	handler       PageHandler
}

func (suite *PageHandlersTestSuite) SetupTest() {
	var err error
	suite.contentOfHTML, err = os.ReadFile("./correct_test_cases/home_page_correct_response.html")
	if err != nil {
		suite.T().Fatal(err)
	}

	mux := http.NewServeMux()
	pagesConfig := GetHandlersParameters(pathToConf)
	suite.handler = AddHomePageHandler(mux, pagesConfig)
}

func (suite *PageHandlersTestSuite) TestAddHomePageHandler() {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		suite.T().Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(suite.handler.Handle)

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
