package maker

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gobeam/stringy"
	"github.com/spf13/cobra"
)

type fileNode struct {
	Name     string
	Type     string // file or dir
	Children []*fileNode
}

var extensionMakerRootCommand = &cobra.Command{
	Use:   "make",
	Short: "Maker for boilerplate code in a shopware extension",
}

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(extensionMakerRootCommand)
}

func renderTemplates(templatePath, outputPath string, data interface{}, force bool, namespace string) error {
	parsedOutputPath, err := replaceVariablesInString(outputPath, data)
	if err != nil {
		return err
	}

	if err := createOutputDirectory(parsedOutputPath); err != nil {
		return err
	}

	templateFiles, err := os.ReadDir(templatePath)
	if err != nil {
		return err
	}

	for _, file := range templateFiles {
		println("Rendering: ", templatePath, file.Name())
		if file.IsDir() {
			subdir := filepath.Join(templatePath, file.Name())
			suboutput := filepath.Join(parsedOutputPath, file.Name())
			subNamespace := namespace + "\\" + file.Name()

			if err := renderTemplates(subdir, suboutput, data, force, subNamespace); err != nil {
				return err
			}
			continue
		}

		if err := renderTemplateFile(templatePath, parsedOutputPath, file, data, force, namespace); err != nil {
			return err
		}
	}

	return nil
}

func renderTemplateFile(templatePath, outputPath string, file os.DirEntry, data interface{}, force bool, namespace string) error {
	newData := data.(map[string]interface{})
	newData["namespace"] = namespace

	println("Rendering File: ", templatePath, file.Name())

	templateContent, err := os.ReadFile(filepath.Join(templatePath, file.Name()))
	if err != nil {
		return err
	}

	tmpl, err := parseTemplate(string(templateContent))
	if err != nil {
		return err
	}

	templateName := string(file.Name())
	newTemplateName, err := replaceVariablesInString(templateName, newData)
	if err != nil {
		return err
	}

	if !force {
		if _, err := os.Stat(filepath.Join(outputPath, newTemplateName)); !os.IsNotExist(err) {
			fmt.Println("File exists, skipping: ", newTemplateName)
			return nil
		}
	}

	outputFile, err := os.Create(filepath.Join(outputPath, newTemplateName))
	if err != nil {
		return err
	}
	defer outputFile.Close()

	if err := tmpl.Execute(outputFile, data); err != nil {
		return err
	}

	return nil
}

func createOutputDirectory(outputPath string) error {
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		return os.MkdirAll(outputPath, os.ModePerm)
	}
	return nil
}

func printTree(node *fileNode, indent string) {
	if node.Type == "file" {
		fmt.Printf("%s└ \033[32m%s\033[0m\n", indent, node.Name)
	} else {
		fmt.Printf("%s└── \033[34m%s\033[0m\n", indent, node.Name)
	}

	if len(node.Children) > 0 {
		subIndent := strings.Replace(indent+"  ", "└", " ", -1)
		for i, child := range node.Children {
			if i < len(node.Children)-1 {
				printTree(child, indent+"│ ")
			} else {
				printTree(child, subIndent)
			}
		}
	}
}

func parseTemplate(templateContent string) (*template.Template, error) {
	return template.New("template").Funcs(getFuncMap()).Parse(templateContent)
}

// create new template and return string from template
func replaceVariablesInString(input string, data interface{}) (string, error) {
	tmpl, err := parseTemplate(input)
	if err != nil {
		return "", err
	}

	var output bytes.Buffer

	if err := tmpl.Execute(&output, data); err != nil {
		return "", err
	}

	return output.String(), nil
}

func SnakeCase(s string) string {
	return stringy.New(s).SnakeCase().ToLower()
}

func CamelCase(s string) string {
	return stringy.New(s).CamelCase().Get()
}

func PascalCase(s string) string {
	return stringy.New(s).PascalCase().Get()
}

func KebabCase(s string) string {
	return stringy.New(s).KebabCase().ToLower()
}

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"SnakeCase":  SnakeCase,
		"CamelCase":  CamelCase,
		"KebabCase":  KebabCase,
		"PascalCase": PascalCase,
	}
}
