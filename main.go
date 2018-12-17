package main

import (
	"protoc-gen-twirp_cpp/gen"
)

func main() {
	g := newGenerator()
	gen.Main(g)
}
