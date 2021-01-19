// Copyright 2016-2019 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

/*

The grnc-yaml-bind tool - used to convert Granitic's YAML component definition files into Go source.

This tool is a port of grnc-yaml-bind refer to https://godoc.org/github.com/graniticio/granitic/v2/cmd/grnc-bind for usage instructions

*/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/graniticio/granitic-yaml/v2"
	"github.com/graniticio/granitic/v2/cmd/grnc-bind/binder"
	"github.com/graniticio/granitic/v2/config"
	"github.com/graniticio/granitic/v2/logging"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	b := new(binder.Binder)
	b.ToolName = "grnc-yaml-bind"
	b.Loader = new(YamlDefinitionLoader)
	b.SupportedExtensions = new(granitic_yaml.YamlContentParser).Extensions()

	b.SupportedExtensions = append(b.SupportedExtensions, new(config.JSONContentParser).Extensions()...)

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
	jm.RegisterContentParser(new(config.JSONContentParser))
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

func (ydl *YamlDefinitionLoader) FacilityManifest(path string) (*binder.Manifest, error) {

	mf, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("unable to open manifest file at %s: %s", path, err.Error())
	}

	defer mf.Close()

	b, err := ioutil.ReadAll(mf)

	if err != nil {
		return nil, fmt.Errorf("unable to read manifest file at %s: %s", path, err.Error())
	}

	lp := strings.ToLower(path)
	m := new(binder.Manifest)

	if strings.HasSuffix(lp, "json") {
		//Support manifests in JSON as well as YAML

		err = json.Unmarshal(b, m)

		if err != nil {
			return nil, fmt.Errorf("unable to parse manifest file at %s: %s", path, err.Error())
		}

	} else if strings.HasSuffix(lp, "yml") || strings.HasSuffix(lp, "yaml") {

		var looseParsed interface{}

		ycp := new(granitic_yaml.YamlContentParser)

		err = ycp.ParseInto(b, &looseParsed)

		if err != nil {
			return nil, fmt.Errorf("unable to parse manifest file at %s: %s", path, err.Error())
		}

		return parseManifest(looseParsed, path)

	} else {
		return nil, fmt.Errorf("%s does not appear to be a YAML or JSON file", err.Error())
	}

	return m, nil

}

func parseManifest(i interface{}, path string) (*binder.Manifest, error) {

	b, err := json.Marshal(i)

	if err != nil {
		return nil, fmt.Errorf("unable to convert the YAML manifest to JSON: %s", err.Error())
	}

	m := new(binder.Manifest)

	err = json.Unmarshal(b, m)

	if err != nil {
		return nil, fmt.Errorf("unable to parse manifest file at %s: %s", path, err.Error())
	}

	fmt.Printf("%#v", m)

	return m, nil
}
