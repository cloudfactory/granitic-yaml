// Copyright 2016-2018 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

/*

The grnc-yaml-bind tool - used to convert Granitic's YAML component definition files into Go source.

This tool is a port of grnc-yaml-bind refer to https://godoc.org/github.com/graniticio/granitic/cmd/grnc-bind for usage instructions

*/
package main

import (
	"github.com/graniticio/granitic-yaml"
	"github.com/graniticio/granitic/cmd/grnc-bind/binder"
	"github.com/graniticio/granitic/config"
	"github.com/graniticio/granitic/logging"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func main() {

	b := new(binder.Binder)
	b.ToolName = "grnc-yaml-bind"
	b.Loader = new(YamlDefinitionLoader)
	b.Bind()

}

// Loads YAML files from local files and remote URLs and provides a mechanism for writing the resulting merged
// file to disk
type YamlDefinitionLoader struct {
}

// LoadAndMerge reads one or more YAML from local files or HTTP URLs and merges them into a single data structure
func (ydl *YamlDefinitionLoader) LoadAndMerge(files []string) (map[string]interface{}, error) {

	jm := config.NewJsonMergerWithDirectLogging(new(logging.ConsoleErrorLogger), new(granitic_yaml.YamlContentParser))
	jm.MergeArrays = true

	return jm.LoadAndMergeConfig(files)

}

// WriteMerged converts the supplied data structure to YAML and writes to disk at the specified location
func (ydl *YamlDefinitionLoader) WriteMerged(data map[string]interface{}, path string) error {

	b, err := yaml.Marshal(data)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, b, 0644)

	if err != nil {
		return err
	}

	os.Exit(0)

	return nil
}
