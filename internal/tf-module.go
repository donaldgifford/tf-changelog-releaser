package internal

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type TerraformModules struct {
	Path    string
	Modules []TfModule
}

type TfModule struct {
	D    fs.DirEntry
	Path string
}

func NewTerraformModule() *TerraformModules {
	return &TerraformModules{}
}

func (t *TerraformModules) GenerateDocs() (string, error) {
	// 	tempString := `
	// # My super cool template
	//
	// <!-- BEGIN_TF_DOCS -->
	// {{ .Content }}
	// <!-D- END_TF_DOCS -->
	// `

	for _, tf := range t.Modules {
		d, err := BuildTerraformDocs(tf)
		// d, err := BuildTerraformDocs(t.Path, tempString)
		if err != nil {
			return "", err
		}

		fmt.Println(d)

	}

	return "", nil
}

func (t *TerraformModules) tfModuleDirWalker(
	fsys fs.FS,
) fs.WalkDirFunc {
	return func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if readDirFS, ok := fsys.(fs.ReadDirFS); ok {
			dirs, err := readDirFS.ReadDir(d.Name())
			if err != nil {
				return err
			}

			for _, k := range dirs {
				if strings.Contains(k.Name(), ".tf") {
					fmt.Printf(
						"Terraform file found at %s/%s\n",
						d.Name(),
						k.Name(),
					)
					t.Modules = append(
						t.Modules,
						TfModule{
							D:    d,
							Path: fmt.Sprintf("%s/%s", t.Path, d.Name()),
						},
					)
					return fs.SkipDir
				}
			}
		}
		fmt.Println(path)
		return nil
	}
}

func (t *TerraformModules) GenerateModules() error {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cwd)
	t.Path = viper.GetString("modules.directory")
	moduleFS := os.DirFS(cwd + "/" + t.Path)

	err = fs.WalkDir(
		moduleFS,
		".",
		t.tfModuleDirWalker(moduleFS),
	)
	if err != nil {
		return err
	}

	fmt.Println(t.Modules)

	return nil
}

// find docs dir
