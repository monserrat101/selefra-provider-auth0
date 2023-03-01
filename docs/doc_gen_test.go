package main

import (
	"os"
	"testing"

	"github.com/selefra/selefra-provider-auth0/resources"
	"github.com/selefra/selefra-provider-sdk/doc_gen"
)

func Test(t *testing.T) {
	docOutputDirectory := "./docs"
	
	_ = os.Mkdir(docOutputDirectory, os.ModePerm)
	err := doc_gen.New(resources.GetSelefraProvider(), docOutputDirectory).Run()

	if err != nil {
		panic(err)
	}
}
