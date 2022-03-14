package cloudformation

type AppCname struct {
	Resources AppCnameResource `yaml:"Resources"`
}

type AppCnameResource struct {
	MyShipaAppCname MyShipaAppCname `yaml:"MyShipaAppCname"`
}

type MyShipaAppCname struct {
	Type       string             `yaml:"Type"`
	Properties AppCnameProperties `yaml:"Properties"`
}

type AppCnameProperties struct {
	App        string `yaml:"App"`
	Cname      string `yaml:"Cname"`
	Encrypt    bool   `yaml:"Encrypt,omitempty"`
	ShipaHost  string `yaml:"ShipaHost"`
	ShipaToken string `yaml:"ShipaToken"`
}
