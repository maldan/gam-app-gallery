package main

import (
	"embed"

	helloworld "github.com/maldan/gam-app-gallery/internal/app/gallery"
)

//go:embed frontend/build/*
var frontFs embed.FS

func main() {
	helloworld.Start(frontFs) //
}
