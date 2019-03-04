package main

import (
	"bufio"
	"fmt"
	"github.com/graniticio/granitic/v2/cmd/grnc-project/generate"
	"path/filepath"
)

func main() {

	pg := new(generate.ProjectGenerator)
	pg.CompWriterFunc = writeComponentsFile
	pg.ConfWriterFunc = writeConfigFile
	pg.MainFileFunc = writeMainFile
	pg.ModFileFunc = writeModFile
	pg.ToolName = "grnc-yaml-project"

	s := generate.SettingsFromArgs(pg.ExitError)

	pg.Generate(s)

}

func writeConfigFile(confDir string, pg *generate.ProjectGenerator) {

	compFile := filepath.Join(confDir, "base.yml")
	f := pg.OpenOutputFile(compFile)

	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString("# Configuration you want to make available to your components\n")
	w.Flush()

}

func writeComponentsFile(compDir string, pg *generate.ProjectGenerator) {

	compFile := filepath.Join(compDir, "common.yml")
	f := pg.OpenOutputFile(compFile)

	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString("packages:\n")
	w.WriteString("  # List of package names (e.g granitic.ws) referenced by components in this file.\n")
	w.WriteString("components:\n")
	w.WriteString("  # Definition of components you want to be managed by Granitic")
	w.Flush()

}

func writeMainFile(w *bufio.Writer, projectPackage string) {

	w.WriteString("package main\n\n")
	w.WriteString("import gy \"github.com/graniticio/granitic-yaml/v2\"\n")
	w.WriteString("import \"")
	w.WriteString(projectPackage)
	w.WriteString("/bindings\"")
	w.WriteString("\n\n")
	w.WriteString("func main() {\n")
	w.WriteString("\tgy.StartGraniticWithYaml(bindings.Components())\n")
	w.WriteString("}\n")

}

func writeModFile(baseDir string, moduleName string, pg *generate.ProjectGenerator) {

	modFile := filepath.Join(baseDir, "go.mod")

	f := pg.OpenOutputFile(modFile)

	defer f.Close()

	w := bufio.NewWriter(f)

	fmt.Fprintf(w, "module %s\n\n", moduleName)
	fmt.Fprintf(w, "require github.com/graniticio/granitic-yaml/v2 v2\n")
	fmt.Fprintf(w, "require github.com/graniticio/granitic/v2 v2\n")

	w.Flush()

}
