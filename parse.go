package granitic_yaml

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"strconv"
)

type YamlContentParser struct {
}

func (ycp *YamlContentParser) ParseInto(data []byte, target interface{}) error {

	var firstPass interface{}

	//Parse as YAML - will have interface{} keyed maps and not detect bools and numbers
	if err := yaml.Unmarshal(data, &firstPass); err != nil {
		return err
	}

	if retyped, err := ycp.convertToStringKeyed(firstPass.(map[interface{}]interface{})); err != nil {
		return err
	} else {
		*target.(*interface{}) = retyped
	}

	/*if err := yaml_mapstr.Unmarshal(data, &interim); err != nil {
		return err
	}

	untyped := interim.(map[string]interface{})

	ycp.reinstateTypes(untyped)*/

	//*target.(*interface{}) = untyped

	return nil
}

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
			m := fmt.Sprintf("Key %V is not a string", k)
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
			/*case string:
				newValue = ycp.stringToJsonType(v)
			case []interface{}:
				if err := ycp.correctArrayTypes(v); err != nil {
					return nil, err
				} else {
					newValue = v
				}
			case float64:
				newValue = v
			case bool:
				newValue = v
			case int:
				newValue = v
			default:
				m := fmt.Sprintf("Unsupported type %T for map value %v", v, v)
				return nil, errors.New(m)*/
		}

		converted[newKey] = newValue

	}

	return converted, nil

}

func (ycp *YamlContentParser) correctArrayTypes(a []interface{}) error {

	fmt.Println(a)

	return nil
}

func (ycp *YamlContentParser) stringToJsonType(s string) interface{} {
	if s == "true" {

		fmt.Println("con")
		return true
	}

	if s == "false" {
		return false
	}

	if n, err := strconv.ParseFloat(s, 64); err == nil {
		return n
	}

	return s
}

func (ycp *YamlContentParser) Extensions() []string {
	return []string{"yaml", "yml"}
}

func (ycp *YamlContentParser) ContentTypes() []string {
	return []string{"text/x-yaml", "application/yaml", "text/yaml", "application/x-yaml"}
}
