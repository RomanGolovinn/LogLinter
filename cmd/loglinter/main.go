package main

// Linter
// Checking logs for compliance with rules

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/RomanGolovinn/loglinter"
)

func main() {
	singlechecker.Main(loglinter.Analyzer)
}
