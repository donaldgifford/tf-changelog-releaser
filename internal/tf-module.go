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
	tempString := `
# My super cool template

<!-- BEGIN_TF_DOCS -->
{{ .Content }}
<!-D- END_TF_DOCS -->
`

	d, err := BuildTerraformDocs(t.Path, tempString)
	if err != nil {
		return "", err
	}

	return d, nil
}

// func examineFileType(fsys fs.FS, path string) error {
// 	info, err := fs.Stat(fsys, path)
// 	if err != nil {
// 		return fmt.Errorf("cannot stat %s: %w", path, err)
// 	}
//
// 	if info.IsDir() {
// 		fmt.Printf("%s is a directory\n", path)
// 		fmt.Printf(
// 			"Directory size: %d bytes\n",
// 			info.Size(),
// 		) // Often system-dependent
// 	} else {
// 		fmt.Printf("%s is a regular file\n", path)
// 		fmt.Printf("File size: %d bytes\n", info.Size())
// 	}
//
// 	return nil
// }
//
// func listDirectoryContents(fsys fs.FS, dirPath string) ([]fs.DirEntry, error) {
// 	if readDirFS, ok := fsys.(fs.ReadDirFS); ok {
// 		return readDirFS.ReadDir(dirPath)
// 	}
//
// 	return nil, fmt.Errorf("filesystem does not support directory reading")
// }

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

// func tfModuleChecker() fs.WalkDirFunc {
// 	// func tfModuleChecker(path string, d fs.DirEntry, err error) error {
// 	return func(path string, d fs.DirEntry, err error) error {
// 		if err != nil {
// 			return err
// 		}
// 		fmt.Println(path)
// 		// fmt.Println(d.Info())
// 		return nil
// 	}
// }

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
	// dirs, err := listDirectoryContents(moduleFS, ".")
	// if err != nil {
	// 	return err
	// }
	//
	// for _, dir := range dirs {
	// 	t, err := listDirectoryContents(moduleFS, dir.Name())
	// 	if err != nil {
	// 		return err
	// 	}
	// 	for _, k := range t {
	// 		if strings.Contains(k.Name(), ".tf") {
	// 			fmt.Println("Terraform file found")
	// 		}
	// 	}
	// }
	// if slices.Contains(t, )
	// fmt.Println(t)

	fmt.Println(t.Modules)

	return nil
}
