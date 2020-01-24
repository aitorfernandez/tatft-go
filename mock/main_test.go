package main

import (
	"encoding/json"
	"net/http"
	"testing"
)

var writedData []byte

func TestData(t *testing.T) {
	var writer mockResponseWriter
	writer.header = make(map[string][]string)

	Data(writer, &http.Request{})

	t.Run("Header", func(t *testing.T) {
		if writer.header.Get("Content-Type") != "application/json" {
			t.Error("Incorrect Content-Type")
		}
	})

	t.Run("Data", func(t *testing.T) {
		foo := &Placeholder{Name: "foo"}
		d, _ := json.Marshal(foo)

		if string(d) != string(writedData) {
			t.Errorf("got %v want %v", string(d), string(writedData))
		}
	})
}

type mockResponseWriter struct {
	header http.Header
}

func (m mockResponseWriter) Header() http.Header {
	return m.header
}

func (m mockResponseWriter) Write(data []byte) (int, error) {
	writedData = data
	return len(data), nil
}

func (m mockResponseWriter) WriteHeader(code int) {}
