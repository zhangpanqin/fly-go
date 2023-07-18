package test

import (
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

func init() {
	os.Setenv("a", "1")
}

type FooTestSuite struct {
	suite.Suite
}

func (suite *FooTestSuite) BeforeTest(suiteName, testName string) {

}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(FooTestSuite))
}

func (suite *FooTestSuite) Test_Sum1() {
	sum := Sum1(2)
	suite.Equal(3, sum)
}
