package route

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestApi(t *testing.T) {
	a := assert.New(t)

	resp, err := http.Get("http://localhost:3000/ping")

	a.Nil(err)
	a.NotNil(resp)
}
