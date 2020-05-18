package main

import (
	"os"
	"testing"
)

func Test_run(t *testing.T) {
	tests := []struct {
		name  string
		count int
	}{
		{
			"run",
			50000,
		},
	}
	os.Stdout, _ = os.Open(os.DevNull)
	setup()
	defer taskExecutor.Close()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			run(tt.count)
		})
	}
}

func Test_run_serial(t *testing.T) {
	tests := []struct {
		name  string
		count int
	}{
		{
			"run_serial",
			50000,
		},
	}
	os.Stdout, _ = os.Open(os.DevNull)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			run_serial(tt.count)
		})
	}
}

func BenchmarkRun(b *testing.B) {
	os.Stdout, _ = os.Open(os.DevNull)
	setup()
	defer taskExecutor.Close()
	b.ResetTimer()
	b.StartTimer()
	run(50000)
}

func BenchmarkRunSerial(b *testing.B) {
	os.Stdout, _ = os.Open(os.DevNull)

	run_serial(50000)
}
