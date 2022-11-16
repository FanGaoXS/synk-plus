package main

import (
	"embed"
)

var (
	//go:embed frontend/dist/*
	FS embed.FS
)
