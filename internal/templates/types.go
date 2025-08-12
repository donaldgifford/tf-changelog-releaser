package templates

// MkDocsData holds data for MkDocs configuration template
type MkDocsData struct {
	SiteName        string
	SiteDescription string
	Modules         []string
}

// IndexData holds data for index.md template
type IndexData struct {
	SiteName string
}

// ChangelogEntryData holds data for changelog entry template
type ChangelogEntryData struct {
	ModuleName  string
	Version     string
	Date        string
	IsNewModule bool
	Features    []string
	Fixes       []string
	Others      []string
}

// // ModuleData holds data for module documentation templates
// type ModuleData struct {
// 	ModuleName  string
// 	Providers   []ProviderInfo
// 	Resources   []ResourceInfo
// 	DataSources []DataSourceInfo
// 	Variables   []VariableInfo
// 	Outputs     []OutputInfo
// }
//
// // ProviderInfo holds provider information
// type ProviderInfo struct {
// 	Name    string
// 	Version string
// }
//
// // ResourceInfo holds resource information
// type ResourceInfo struct {
// 	Name         string
// 	Type         string
// 	ProviderName string
// 	ResourceType string
// }
//
// // DataSourceInfo holds data source information
// type DataSourceInfo struct {
// 	Name           string
// 	Type           string
// 	ProviderName   string
// 	DataSourceType string
// }
//
// // VariableInfo holds variable information
// type VariableInfo struct {
// 	Name        string
// 	Description string
// 	Type        string
// 	Default     string
// }
//
// // OutputInfo holds output information
// type OutputInfo struct {
// 	Name        string
// 	Description string
// }
//
// // ReleaseNotesData holds data for release notes template
// type ReleaseNotesData struct {
// 	ModuleName  string
// 	Version     string
// 	IsNewModule bool
// 	Features    []string
// 	Fixes       []string
// 	Others      []string
// }
