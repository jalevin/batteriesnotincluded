package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func Test_DefaultHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8080", nil)
	res := httptest.NewRecorder()

	var a App

	a.Default(res, req)
	got := res.Body.String()
	want := "Hello World!"

	fmt.Println("Want length:", len(want))
	if got != want {
		t.Log("Want:", want)
		t.Log("Got:", got)
		t.Fatal("Mismatch")
	}
}
