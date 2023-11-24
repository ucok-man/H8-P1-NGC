package sample

import "embed"

//go:embed *.txt
var Samplefile embed.FS
