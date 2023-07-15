package main_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Output struct {
	Total int `json:"total"`
}

func TestMain(t *testing.T) {
	assert := assert.New(t)

	t.Run(`Should to test the API`, func(t *testing.T) {
		response, _ := http.Get(`http://localhost:3000/cards/1234/invoices`)
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)

		var output Output
		json.Unmarshal(body, &output)

		assert.Equal(1500, output.Total)
	})
}
