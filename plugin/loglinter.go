package main

import (
	"golang.org/x/tools/go/analysis"

	"github.com/RomanGolovinn/loglinter"
)

type analyzerPlugin struct{}

func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		loglinter.Analyzer,
	}
}

var AnalyzerPlugin analyzerPlugin
