// Package goanalyzers registers go-analyzers as golangci-lint v2 module plugins.
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
	register.Plugin("makecopy", newMakecopyPlugin)
	register.Plugin("searchmigrate", newSearchmigratePlugin)
	register.Plugin("clampcheck", newClampcheckPlugin)
	register.Plugin("sortmigrate", newSortmigratePlugin)
}

// makecopy plugin

func newMakecopyPlugin(any) (register.LinterPlugin, error) {
	return &makecopyPlugin{}, nil
}

type makecopyPlugin struct{}

func (p *makecopyPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{makecopy.Analyzer}, nil
}

func (p *makecopyPlugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}

// searchmigrate plugin

func newSearchmigratePlugin(any) (register.LinterPlugin, error) {
	return &searchmigratePlugin{}, nil
}

type searchmigratePlugin struct{}

func (p *searchmigratePlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{searchmigrate.Analyzer}, nil
}

func (p *searchmigratePlugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}

// clampcheck plugin

func newClampcheckPlugin(any) (register.LinterPlugin, error) {
	return &clampcheckPlugin{}, nil
}

type clampcheckPlugin struct{}

func (p *clampcheckPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{clampcheck.Analyzer}, nil
}

func (p *clampcheckPlugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}

// sortmigrate plugin

func newSortmigratePlugin(any) (register.LinterPlugin, error) {
	return &sortmigratePlugin{}, nil
}

type sortmigratePlugin struct{}

func (p *sortmigratePlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{sortmigrate.Analyzer}, nil
}

func (p *sortmigratePlugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
