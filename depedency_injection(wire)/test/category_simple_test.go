package test

import (
	"fmt"
	"testing"

	"github.com/depri11/go-learning/restful/simple"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializedService()
	fmt.Println(err)
	fmt.Println(simpleService.SimpleRepository)
}

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializedService()
	fmt.Println(err)
	fmt.Println(simpleService)
}
