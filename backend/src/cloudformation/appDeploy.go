package cloudformation

type AppDeploy struct {
	Resources AppDeployResource `yaml:"Resources"`
}

type AppDeployResource struct {
	MyShipaAppDeploy MyShipaAppDeploy `yaml:"MyShipaAppDeploy"`
}

type MyShipaAppDeploy struct {
	Type       string              `yaml:"Type"`
	Properties AppDeployProperties `yaml:"Properties"`
}

type AppDeployProperties struct {
	App        string    `yaml:"App"`
	Image      string    `yaml:"Image"`
	AppConfig  AppConfig `yaml:"AppConfig"`
	Registry   *Registry `yaml:"Registry,omitempty"`
	Port       *Port     `yaml:"Port,omitempty"`
	ShipaHost  string    `yaml:"ShipaHost"`
	ShipaToken string    `yaml:"ShipaToken"`
}

type Port struct {
	Number   int64  `yaml:"Number"`
	Protocol string `yaml:"Protocol"`
}

type Registry struct {
	User   string `yaml:"User"`
	Secret string `yaml:"Secret"`
}

type AppConfig struct {
	Team      string   `yaml:"Team"`
	Framework string   `yaml:"Framework"`
	Plan      string   `yaml:"Plan,omitempty"`
	Tags      []string `yaml:"Tags,omitempty"`
}
