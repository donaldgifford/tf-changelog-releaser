package internal

// TemplateData holds common data for all templates
type TemplateData struct {
	SiteName        string
	SiteDescription string
	ModuleName      string
	ModulePath      string
	Version         string
	Date            string
	IsNewModule     bool
	Changes         []string
	Features        []string
	Fixes           []string
	Others          []string
}

// MkDocsConfigTemplate generates mkdocs.yml configuration
const MkDocsConfigTemplate = `site_name: {{.SiteName}}
site_description: {{.SiteDescription}}
plugins:
  - techdocs-core
nav:
  - Getting Started: index.md
  - Latest Changes: CHANGELOG.md{{if .Modules}}
  - Modules:{{range .Modules}}
    - {{.}}: modules/{{.}}.md{{end}}{{end}}
`

// IndexTemplate generates the main index.md file
const IndexTemplate = `# {{.SiteName}}

Welcome to the {{.SiteName}} documentation.

## Overview

This documentation contains information about our Terraform modules and their usage.

## Modules

Browse the modules section to see detailed documentation for each available module.`

// ChangelogHeaderTemplate generates changelog header
const ChangelogHeaderTemplate = `# Changelog

All notable changes to the Terraform modules will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

`

// ChangelogEntryTemplate generates individual changelog entries
const ChangelogEntryTemplate = `## [{{.ModuleName}}/{{.Version}}] - {{.Date}}{{if .IsNewModule}} - 🆕 New Module{{end}}

{{if .Features}}### Added

{{range .Features}}- {{.}}
{{end}}
{{end}}{{if .Fixes}}### Fixed

{{range .Fixes}}- {{.}}
{{end}}
{{end}}{{if .Others}}### Changed

{{range .Others}}- {{.}}
{{end}}
{{end}}`
