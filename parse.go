// Copyright 2018 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.
package granitic_yaml

import (
	"errors"
	"fmt"
	"github.com/graniticio/granitic/v2/config"
	"gopkg.in/yaml.v2"
)

// Handles the conversion of YAML files into the intermediate JSON-like structures used by Granitic's configuration
// and component definition processes
//
// Implements the jsonmerge.ContentParser interface in Granitic
type YamlContentParser struct {
}

// Implements jsonmerge.ContentParser
func (ycp *YamlContentParser) ParseInto(data []byte, target interface{}) error {

	var firstPass interface{}

	if err := yaml.Unmarshal(data, &firstPass); err != nil {
		return err
	}

	if _, okay := firstPass.(map[interface{}]interface{}); !okay {
		//Symptom of empty config file
		empty := map[string]interface{}{}
		*target.(*interface{}) = empty

		return config.EmptyFileError{Message: "YAML config file logically empty after parsing"}
	}

	if retyped, err := ycp.convertToStringKeyed(firstPass.(map[interface{}]interface{})); err != nil {
		return err
	} else {
		*target.(*interface{}) = retyped
	}

	return nil
}

// Converts the map[interface{}]interface{} maps output by the YAML parser into the map[string]interface{} maps
// expected by Granitic
func (ycp *YamlContentParser) convertToStringKeyed(im map[interface{}]interface{}) (map[string]interface{}, error) {

	converted := make(map[string]interface{}, len(im))

	for k, v := range im {
		var newKey string
		var newValue interface{}

		//Check that the key is actually a string
		switch k := k.(type) {
		case string:
			newKey = k
		default:
			m := fmt.Sprintf("Key %v is not a string", k)
			return nil, errors.New(m)
		}

		switch v := v.(type) {
		case map[interface{}]interface{}:
			if c, err := ycp.convertToStringKeyed(v); err != nil {
				return nil, err
			} else {
				newValue = c
			}
		default:
			newValue = v
		}

		converted[newKey] = newValue

	}

	return converted, nil

}

// Extensions returns the file name extensions accepted as representing YAML files
//
// Implements jsonmerge.ContentParser
func (ycp *YamlContentParser) Extensions() []string {
	return []string{"yaml", "yml"}
}

// ContentTypes returns the HTTP content-types (MIME types) accepted as representing YAML files
//
// Implements jsonmerge.ContentParser
func (ycp *YamlContentParser) ContentTypes() []string {
	return []string{"text/x-yaml", "application/yaml", "text/yaml", "application/x-yaml"}
}
