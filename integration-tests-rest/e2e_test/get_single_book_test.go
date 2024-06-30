package e2e_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"

	_ "github.com/lib/pq"
)

type GetSingleBookSuite struct {
	suite.Suite
}

func TestGetSingleBookSuite(t *testing.T) {
	suite.Run(t, new(GetSingleBookSuite))
}

func (s *GetSingleBookSuite) TestGetBookThatDoesNotExist() {
	c := http.Client{}

	r, _ := c.Get("http://localhost:8080/book/123456789")
	body, _ := ioutil.ReadAll(r.Body)

	s.Equal(http.StatusNotFound, r.StatusCode)
 	s.JSONEq(`{"code": "001", "msg": "No book with ISBN 123456789"}`, string(body))
}

func (s *GetSingleBookSuite) TestGetBookThatDoesExist() {
	c := http.Client{}

	r, _ := c.Get("http://localhost:8080/book/351964302699")
	body, _ := ioutil.ReadAll(r.Body)

	s.Equal(http.StatusOK, r.StatusCode)

	expBody := `{
	"isbn": "351964302699",
	"title": "Call me",
	"image": "ball.jpg",
	"genre": "Computing",
	"year_published": 2000
}`

	s.JSONEq(expBody, string(body))
}
