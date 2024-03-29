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
	App        string             `yaml:"App"`
	Image      string             `yaml:"Image"`
	AppConfig  AppConfig          `yaml:"AppConfig"`
	Registry   *Registry          `yaml:"Registry,omitempty"`
	Port       *Port              `yaml:"Port,omitempty"`
	Volumes    []*AppDeployVolume `yaml:"Volumes,omitempty"`
	ShipaHost  string             `yaml:"ShipaHost"`
	ShipaToken string             `yaml:"ShipaToken"`
}

type AppDeployVolume struct {
	Name    string         `yaml:"Name"`
	Path    string         `yaml:"MountPath"`
	Options *VolumeOptions `yaml:"MountOptions,omitempty"`
}

type VolumeOptions struct {
	Prop1 string `yaml:"AdditionalProp1,omitempty"`
	Prop2 string `yaml:"AdditionalProp2,omitempty"`
	Prop3 string `yaml:"AdditionalProp3,omitempty"`
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
