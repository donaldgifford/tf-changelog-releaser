package templates

import (
	"embed"
	"fmt"
	"os"
	"text/template"

	"github.com/donaldgifford/tf-changelog-releaser/internal"
	"github.com/spf13/viper"
)

//go:embed default/mkdocs.yml.tmpl
//go:embed default/index.md.tmpl
//go:embed default/module.md.tmpl
var DefaultTemplates embed.FS

const (
	EMBEDDED_MKDOCS_TEMPLATE string = "default/mkdocs.yml.tmpl"
	EMBEDDED_INDEX_TEMPLATE  string = "default/index.md.tmpl"
	EMBEDDED_MODULE_TEMPLATE string = "default/module.md.tmpl"
)

// TemplateExecutor provides methods to execute templates
type TemplateExecutor struct{}

// NewTemplateExecutor creates a new template executor
func NewTemplateExecutor() *TemplateExecutor {
	return &TemplateExecutor{}
}

// ExecuteTemplate executes a template with the provided data
func (te *TemplateExecutor) ExecuteTemplate(
	templateType string,
	outputFile string,
	data interface{},
) error {
	tmpl, err := template.ParseFS(DefaultTemplates, templateType)
	if err != nil {
		return err
	}

	// check directories
	docsPath := viper.GetString("docs.directory")
	modsPath := viper.GetString("modules.docs_path")

	if _, err := os.Stat(docsPath); os.IsNotExist(err) {
		fmt.Printf("%s does not exist\n", docsPath)
		return err
	} else {
		if _, err := os.Stat(modsPath); os.IsNotExist(err) {
			fmt.Printf("%s does not exist\n", modsPath)
			fmt.Printf("%s exists\n", modsPath)
			err = os.Mkdir(modsPath, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			fmt.Printf("%s exists\n", modsPath)
		}
		fmt.Printf("%s exists\n", docsPath)
	}

	var f *os.File
	// open file
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		fmt.Printf("%s does not exist, creating...\n", outputFile)
		f, err = os.Create(outputFile)
		if err != nil {
			fmt.Printf("%s couldnt be created...\n", outputFile)
			return err
		}
	} else if err != nil {
		return err
	} else {
		f, err = os.Open(outputFile)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("File %s does not exist\n", outputFile)
				return err
			} else {
				return err
			}
		}
	}

	if err := tmpl.Execute(f, data); err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}

//
// // ExecuteMkDocsConfig executes the MkDocs configuration template
// func (te *TemplateExecutor) ExecuteMkDocsConfig(
// 	data interface{},
// ) (string, error) {
// 	return te.ExecuteTemplate(EMBEDDED_MKDOCS_TEMPLATE, data)
// }
//
// // ExecuteIndex executes the index template
// func (te *TemplateExecutor) ExecuteIndex(data interface{}) (string, error) {
// 	return te.ExecuteTemplate(EMBEDDED_INDEX_TEMPLATE, data)
// }

// ExecuteModule executes the module template
func (te *TemplateExecutor) ExecuteModule(
	mod *internal.TfModule,
) error {
	modulesDocsPath := viper.GetString("modules.docs_path")
	modulesOutputFileName := modulesDocsPath + "/" + mod.D.Name() + ".md"
	return te.ExecuteTemplate(
		EMBEDDED_MODULE_TEMPLATE,
		modulesOutputFileName,
		mod,
	)
}
