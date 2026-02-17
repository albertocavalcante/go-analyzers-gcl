// Package goanalyzers registers go-analyzers as a golangci-lint v2 module plugin.
package goanalyzers

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	"github.com/albertocavalcante/go-analyzers/clampcheck"
	"github.com/albertocavalcante/go-analyzers/makecopy"
	"github.com/albertocavalcante/go-analyzers/searchmigrate"
	"github.com/albertocavalcante/go-analyzers/sortmigrate"
)

func init() {
	register.Plugin("go-analyzers", newPlugin)
}

func newPlugin(any) (register.LinterPlugin, error) {
	return &plugin{}, nil
}

type plugin struct{}

func (p *plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		makecopy.Analyzer,
		searchmigrate.Analyzer,
		clampcheck.Analyzer,
		sortmigrate.Analyzer,
	}, nil
}

func (p *plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
