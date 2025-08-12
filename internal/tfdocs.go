package internal

import (
	"github.com/terraform-docs/terraform-docs/format"
	"github.com/terraform-docs/terraform-docs/print"
	"github.com/terraform-docs/terraform-docs/terraform"
)

// BuildTerraformDocs for module root `path` and provided content `tmpl`.
func BuildTerraformDocs(tf *TfModule) (string, error) {
	// func BuildTerraformDocs(tf TfModule, tmpl string) (string, error) {
	config := print.DefaultConfig()
	config.ModuleRoot = tf.Path // module root path (can be relative or absolute)

	module, err := terraform.LoadWithOptions(config)
	if err != nil {
		return "", err
	}

	// Generate in Markdown Table format
	formatter := format.NewMarkdownTable(config)

	if err := formatter.Generate(module); err != nil {
		return "", err
	}

	// Note: if you don't intend to provide additional template for the generated
	// content, or the target format doesn't provide templating (e.g. json, yaml,
	// xml, or toml) you can use `Content()` function instead of `Render()`.
	// `Content()` returns all the sections combined with predefined order.
	// return formatter.Content(), nil

	// return formatter.Render(tmpl)
	return formatter.Content(), nil
}
