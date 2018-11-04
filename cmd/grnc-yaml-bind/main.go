package main

import (
	"fmt"
	"github.com/graniticio/granitic-yaml"
	"github.com/graniticio/granitic/cmd/grnc-bind/binder"
	"github.com/graniticio/granitic/config"
	"github.com/graniticio/granitic/logging"
	"os"
)

func main() {

	b := new(binder.Binder)
	b.ToolName = "grnc-yaml-bind"
	b.Loader = new(YamlDefinitionLoader)
	b.Bind()

}

type YamlDefinitionLoader struct {
}

func (ydl *YamlDefinitionLoader) LoadAndMerge(files []string) (map[string]interface{}, error) {

	jm := new(config.JsonMerger)
	jm.MergeArrays = true
	jm.Logger = new(logging.ConsoleErrorLogger)
	jm.DefaultParser = new(granitic_yaml.YamlContentParser)

	return jm.LoadAndMergeConfig(files)

}

func (ydl *YamlDefinitionLoader) WriteMerged(data map[string]interface{}, path string) error {
	fmt.Println(data)
	os.Exit(-1)

	return nil
}
