// Copyright 2016-2019 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

/*

The grnc-yaml-bind tool - used to convert Granitic's YAML component definition files into Go source.

This tool is a port of grnc-yaml-bind refer to https://godoc.org/github.com/graniticio/granitic/v2/cmd/grnc-bind for usage instructions

*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	granitic_yaml "github.com/cloudfactory/granitic-yaml/v2"
	"github.com/graniticio/granitic/v2/cmd/grnc-bind/binder"
	"github.com/graniticio/granitic/v2/config"
	"github.com/graniticio/granitic/v2/logging"
	"gopkg.in/yaml.v2"
)

func main() {

	b := new(binder.Binder)
	b.ToolName = "grnc-yaml-bind"
	b.Loader = new(YamlDefinitionLoader)

	s, err := binder.SettingsFromArgs()

	if err != nil {
		fmt.Printf("%s: %s\n", b.ToolName, err.Error())
		os.Exit(1)
	}

	pref := fmt.Sprintf("%s: ", b.ToolName)
	b.Log = logging.NewStdoutLogger(s.LogLevel, pref)

	b.Bind(s)

}

// Loads YAML files from local files and remote URLs and provides a mechanism for writing the resulting merged
// file to disk
type YamlDefinitionLoader struct {
}

// LoadAndMerge reads one or more YAML from local files or HTTP URLs and merges them into a single data structure
func (ydl *YamlDefinitionLoader) LoadAndMerge(files []string, log logging.Logger) (map[string]interface{}, error) {

	jm := config.NewJSONMergerWithDirectLogging(log, new(granitic_yaml.YamlContentParser))
	jm.MergeArrays = true

	return jm.LoadAndMergeConfig(files)

}

// WriteMerged converts the supplied data structure to YAML and writes to disk at the specified location
func (ydl *YamlDefinitionLoader) WriteMerged(data map[string]interface{}, path string, log logging.Logger) error {

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
