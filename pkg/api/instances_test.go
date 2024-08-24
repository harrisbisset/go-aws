package api_test

import (
	"testing"

	"github.com/harrisbisset/go-aws/pkg/api"
)

func TestGetInstanceList(t *testing.T) {
	_, err := api.GetInstanceList()
	if err != nil {
		panic(err)
	}
}
