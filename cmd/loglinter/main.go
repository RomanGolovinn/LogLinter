package main

// Linter
// Checking logs for compliance with rules

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	loglinter "github.com/RomanGolovinn/loglinter/analyzer"
)

func main() {
	singlechecker.Main(loglinter.Analyzer)
}
