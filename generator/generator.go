package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"

	"github.com/rancher/go-kubernetes-client/model"
)

const (
	CLIENT_OUTPUT_DIR = "../model"
)

var (
	underscoreRegexp *regexp.Regexp = regexp.MustCompile(`([a-z])([A-Z])`)
	versionRegex     *regexp.Regexp = regexp.MustCompile(`.*\.`)
)

var (
	simpleTypeConversions = map[string]string{
		"any":     "map[string]interface{}",
		"boolean": "bool",
	}
	typeConversionOverrides = map[string]string{
		"WatchEvent.Object": "interface{}",
	}
)

func main() {
	err := generateFiles()
	if err != nil {
		log.Fatal(err)
	}
}

func generateFiles() error {
	err := setupDirectory(CLIENT_OUTPUT_DIR)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	schemaBytes, err := ioutil.ReadFile("schemas.json")
	if err != nil {
		return err
	}
	var schema model.Schema

	err = json.Unmarshal(schemaBytes, &schema)
	if err != nil {
		return err
	}

	for _, model := range schema.Models {
		err = generateType(model)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateType(model model.Model) error {
	modelType := stripVersion(model.ID)

	output, err := os.Create(path.Join(CLIENT_OUTPUT_DIR, strings.ToLower("generated_"+addUnderscore(modelType))+".go"))
	if err != nil {
		return err
	}
	defer output.Close()

	data := map[string]interface{}{
		"model":           model,
		"typeCapitalized": capitalize(modelType),
		"typeUpper":       strings.ToUpper(addUnderscore(modelType)),
		"structFields":    getTypeMap(model, modelType),
	}

	funcMap := template.FuncMap{
		"toLowerCamelCase":  toLowerCamelCase,
		"toLowerUnderscore": addUnderscore,
		"capitalize":        capitalize,
		"upper":             strings.ToUpper,
	}

	typeTemplate, err := template.New("type.template").Funcs(funcMap).ParseFiles("type.template")
	if err != nil {
		return err
	}

	return typeTemplate.Execute(output, data)
}

func getTypeMap(modelResource model.Model, modelType string) map[string]string {
	result := map[string]string{}
	for name, prop := range modelResource.Properties {
		fieldName := capitalize(name)

		fullName := fmt.Sprintf("%s.%s", modelType, fieldName)
		if override, ok := typeConversionOverrides[fullName]; ok {
			result[fieldName] = override
		} else if prop.Type != "" {
			if t, ok := simpleTypeConversions[prop.Type]; ok {
				result[fieldName] = t
			} else if prop.Type == "array" {
				if prop.Items.Type != "" {
					result[fieldName] = "[]" + prop.Items.Type
				} else if prop.Items.Ref != "" {
					result[fieldName] = "[]" + stripVersion(prop.Items.Ref)
				}
			} else if prop.Type == "integer" {
				if prop.Format != "" {
					result[fieldName] = prop.Format
				} else {
					result[fieldName] = "int"
				}
			} else {
				result[fieldName] = prop.Type
			}
		} else if prop.Ref != "" {
			result[fieldName] = stripVersion(prop.Ref)
		}

		if _, ok := result[fieldName]; !ok {
			log.Printf("WARNING: Could not parse type for [%s]. Model: [%#v]", fieldName, modelResource)
		}
	}
	return result
}

func setupDirectory(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.Mkdir(dir, 0755)
	}

	return nil
}

func toLowerCamelCase(input string) string {
	return (strings.ToLower(input[:1]) + input[1:])
}

func addUnderscore(input string) string {
	return strings.Replace(strings.ToLower(underscoreRegexp.ReplaceAllString(input, `${1}_${2}`)), ".", "_", -1)
}

func stripVersion(input string) string {
	return versionRegex.ReplaceAllString(input, "")
}

func capitalize(s string) string {
	if len(s) <= 1 {
		return strings.ToUpper(s)
	}

	return strings.ToUpper(s[:1]) + s[1:]
}
