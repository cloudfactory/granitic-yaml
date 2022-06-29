package main

import (
	granitic_yaml "github.com/cloudfactory/granitic-yaml/v2"
	"github.com/cloudfactory/granitic-yaml/v2/test/mapping/bindings"
)

func main() {
	granitic_yaml.StartGraniticWithYaml(bindings.Components())
}
