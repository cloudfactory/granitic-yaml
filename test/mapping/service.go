package main

import "github.com/graniticio/granitic-yaml/v2"
import "github.com/graniticio/granitic-yaml/v2/test/mapping/bindings"

func main() {
	granitic_yaml.StartGraniticWithYaml(bindings.Components())
}
