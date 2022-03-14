package cloudformation

type AppEnv struct {
	Resources AppEnvResource `yaml:"Resources"`
}

type AppEnvResource struct {
	MyShipaAppEnv MyShipaAppEnv `yaml:"MyShipaAppEnv"`
}

type MyShipaAppEnv struct {
	Type       string           `yaml:"Type"`
	Properties AppEnvProperties `yaml:"Properties"`
}

type AppEnvProperties struct {
	App        string `yaml:"App"`
	Envs       []Env  `yaml:"Envs"`
	Norestart  bool   `yaml:"Norestart"`
	Private    bool   `yaml:"Private"`
	ShipaHost  string `yaml:"ShipaHost"`
	ShipaToken string `yaml:"ShipaToken"`
}

type Env struct {
	Name  string `yaml:"Name"`
	Value string `yaml:"Value"`
}
