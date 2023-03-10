package handlers

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PageHandlersTestSuite struct {
	suite.Suite
	htmlTemps []string
}

func (suite *PageHandlersTestSuite) SetupTest() {
	suite.htmlTemps = []string{
		"./ui/html/home.html",
		"./ui/html/base.html",
		"./ui/html/footer.html",
	}
}

func (suite *PageHandlersTestSuite) TestAddHomePageHandler() {

}

func TestPageHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(PageHandlersTestSuite))
}
