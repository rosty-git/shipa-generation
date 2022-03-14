package cloudformation

type App struct {
	Resources AppResource `yaml:"Resources"`
}

type AppResource struct {
	MyShipaApp MyShipaApp `yaml:"MyShipaApp"`
}

type MyShipaApp struct {
	Type       string        `yaml:"Type"`
	Properties AppProperties `yaml:"Properties"`
}

type AppProperties struct {
	Name       string   `yaml:"Name"`
	Teamowner  string   `yaml:"Teamowner"`
	Framework  string   `yaml:"Framework"`
	Plan       string   `yaml:"Plan"`
	Tags       []string `yaml:"Tags"`
	ShipaHost  string   `yaml:"ShipaHost"`
	ShipaToken string   `yaml:"ShipaToken"`
}
