package main

import "github.com/graniticio/granitic-yaml"
import "github.com/graniticio/granitic-yaml/test/mapping/bindings"

func main() {
	granitic_yaml.StartGraniticWithYaml(bindings.Components())
}
