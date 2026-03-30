package analyzer_test

import (
	"testing"

	loglinter "github.com/RomanGolovinn/loglinter/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	analysistest.Run(t, testdata, loglinter.Analyzer, "a")
}
