package manager

import (
	"net/url"
	"testing"
)

func BenchmarkStartFast(b *testing.B) {
	testURL, _ := url.Parse("https://cdn.videvo.net/videvo_files/video/premium/video0042/large_watermarked/900-2_900-6334-PD2_preview.mp4")
	// Run the benchmark b.N times
	for i := 0; i < b.N; i++ {
		StartFast(testURL)
	}
}

func BenchmarkStartSlow(b *testing.B) {
	testURL, _ := url.Parse("https://cdn.videvo.net/videvo_files/video/premium/video0042/large_watermarked/900-2_900-6334-PD2_preview.mp4")
	// Run the benchmark b.N times
	for i := 0; i < b.N; i++ {
		StartSlow(testURL)
	}
}
