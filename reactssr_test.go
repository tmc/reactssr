package reactssr_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/tmc/reactssr"
)

func Example_newServerSideRenderer() {
	r, err := reactssr.NewServerSideRenderer("./testdata/test-app-1/build/out.js")
	if err != nil {
		panic(err)
	}
	output, err := r.Render()
	if err != nil {
		panic(err)
	}
	untilFirstImage := regexp.MustCompile(`^(.*)<img`).FindAllString(output, -1)
	fmt.Println(untilFirstImage[0])
	// This is the expected output from a vanilla create-react-app app:
	// output: <div class="App"><header class="App-header"><img
}

func BenchmarkAWarmup(b *testing.B) {
	r, err := reactssr.NewServerSideRenderer("./testdata/test-app-1/build/out.js")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := r.Render()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRender(b *testing.B) {
	r, err := reactssr.NewServerSideRenderer("./testdata/test-app-1/build/out.js")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := r.Render()
		if err != nil {
			b.Fatal(err)
		}
	}
}
