package granitic_yaml

import (
	"github.com/graniticio/granitic"
	"github.com/graniticio/granitic/config"
	"github.com/graniticio/granitic/ioc"
)

// StartGranitic starts the IoC container and populates it with the supplied list of prototype components. Any settings
// required during the initial startup of the container are expected to be provided via command line arguments (see
// this page's header for more details). This function will run until the application is halted by an interrupt (ctrl+c) or
// a runtime control shutdown command.
func StartGraniticWithYaml(cs *ioc.ProtoComponents) {

	is := config.InitialSettingsFromEnvironment()

	is.ConfigParsers = []config.ContentParser{new(YamlContentParser)}

	is.BuiltInConfig = cs.FrameworkConfig

	granitic.StartGraniticWithSettings(cs, is)
}
