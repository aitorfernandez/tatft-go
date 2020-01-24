package main

import (
	"bytes"
	"html/template"
	"testing"
)

func BenchmarkFoo(b *testing.B) {
	t, _ := template.New("").Parse("Hello, world")

	// Ignore everithing before.
	b.ResetTimer()

	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		t.Execute(&buf, nil)
		buf.Reset()
	}
}
