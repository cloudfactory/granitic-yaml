package granitic_yaml

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestObjectListParsing(t *testing.T) {

	p := filepath.Join("testdata", "object-list.yml")
	b, err := ioutil.ReadFile(p)

	if err != nil {
		t.Fatalf(err.Error())
	}

	cp := new(YamlContentParser)

	var loadedConfig interface{}

	err = cp.ParseInto(b, &loadedConfig)

	tl, okay := loadedConfig.(map[string]interface{})

	if !okay {
		t.Fatalf("Top level parsed object not a map[string]interface{}")
	}

	outer, okay := tl["Outer"].(map[string]interface{})

	if !okay {
		t.Fatalf("Outer object not a map[string]interface{}")
	}

	fmt.Printf("%T\n", outer["ObjectList"])

	fmt.Printf("%v\n", outer["ObjectList"].([]interface{}))

}
