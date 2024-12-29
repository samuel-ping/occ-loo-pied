package web

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed build/*
var clientFs embed.FS

func GetSvelteFs() fs.FS {
	svelteEmbed, err := fs.Sub(clientFs, "build")
	if err != nil {
		log.Fatal("Error getting svelte embed")
	}

	return svelteEmbed
}
