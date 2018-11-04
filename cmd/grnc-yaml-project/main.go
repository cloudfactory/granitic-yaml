package main

import (
	"bufio"
	"github.com/graniticio/granitic/cmd/grnc-project/generate"
	"path/filepath"
)

func main() {

	pg := new(generate.ProjectGenerator)
	pg.CompWriterFunc = writeComponentsFile
	pg.ConfWriterFunc = writeConfigFile
	pg.ToolName = "grnc-yaml-project"
	pg.Generate()

}

func writeConfigFile(confDir string, pg *generate.ProjectGenerator) {

	compFile := filepath.Join(confDir, "config.yml")
	f := pg.OpenOutputFile(compFile)

	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString("# Configuration you want to make available to your components")
	w.Flush()

}

func writeComponentsFile(compDir string, pg *generate.ProjectGenerator) {

	compFile := filepath.Join(compDir, "components.yml")
	f := pg.OpenOutputFile(compFile)

	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString("packages:\n")
	w.WriteString("  # List of package names (e.g granitic.ws) referenced by components in this file.\n")
	w.WriteString("components:\n")
	w.WriteString("  # Definition of components you want to be managed by Granitic")
	w.Flush()

}
