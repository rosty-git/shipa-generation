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
	App            string `yaml:"App"`
	Image          string `yaml:"Image"`
	PrivateImage   bool   `yaml:"PrivateImage,omitempty"`
	RegistryUser   string `yaml:"RegistryUser,omitempty"`
	RegistrySecret string `yaml:"RegistrySecret,omitempty"`
	Port           string `yaml:"Port,omitempty"`
	ShipaHost      string `yaml:"ShipaHost"`
	ShipaToken     string `yaml:"ShipaToken"`
}
