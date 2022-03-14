package crossplane

type AppDeploy struct {
	ApiVersion string        `yaml:"apiVersion"`
	Kind       string        `yaml:"kind"`
	Metadata   Metadata      `yaml:"metadata"`
	Spec       AppDeploySpec `yaml:"spec"`
}

type AppDeploySpec struct {
	ForProvider AppDeployForProvider `yaml:"forProvider"`
}

type AppDeployForProvider struct {
	App            string `yaml:"app"`
	Image          string `yaml:"image"`
	PrivateImage   bool   `yaml:"private-image,omitempty"`
	RegistryUser   string `yaml:"registry-user,omitempty"`
	RegistrySecret string `yaml:"registry-secret,omitempty"`
	Port           string `yaml:"port,omitempty"`
}
