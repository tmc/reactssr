package reactssr_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/tmc/reactssr"
)

func Example_newerverSideRenderer() {
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
	// output: <div class="App"><header class="App-header"><img
}

func BenchmarkRender(b *testing.B) {
	r, err := reactssr.NewServerSideRenderer("./testdata/test-app-1/build/out.js")
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		_, err := r.Render()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRenderFast(b *testing.B) {
	r, err := reactssr.NewServerSideRenderer("./testdata/test-app-1/build/out.js")
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		_, err := r.RenderFast()
		if err != nil {
			b.Fatal(err)
		}
	}
}
