package ansible

func newPlay() *Play {
	return &Play{
		Hosts: "localhost",
		Vars: []map[string]string{
			{"shipa_host": "<host>"},
			{"shipa_token": "<token>"},
		},
	}
}

type Play struct {
	Hosts string `yaml:"hosts"`
	Vars  []map[string]string
	Tasks []interface{}
}

type AppTask struct {
	Name string `yaml:"name"`
	App  App    `yaml:"shipa_application"`
}

func newAppTask() *AppTask {
	return &AppTask{
		Name: "Create shipa application",
	}
}

type Shipa struct {
	ShipaHost  string `yaml:"shipa_host"`
	ShipaToken string `yaml:"shipa_token"`
}

type App struct {
	Shipa     `yaml:",inline"`
	Name      string   `yaml:"name"`
	Teamowner string   `yaml:"teamowner,omitempty"`
	Framework string   `yaml:"framework,omitempty"`
	Plan      string   `yaml:"plan,omitempty"`
	Tags      []string `yaml:"tags,omitempty"`
}

type AppEnvTask struct {
	Name   string `yaml:"name"`
	AppEnv AppEnv `yaml:"shipa_app_env"`
}

func newAppEnvTask() *AppEnvTask {
	return &AppEnvTask{
		Name: "Create shipa app env",
	}
}

type AppEnv struct {
	Shipa     `yaml:",inline"`
	App       string `yaml:"app"`
	Envs      []Env  `yaml:"envs"`
	Norestart bool   `yaml:"norestart"`
	Private   bool   `yaml:"private"`
}

type Env struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type AppCnameTask struct {
	Name     string   `yaml:"name"`
	AppCname AppCname `yaml:"shipa_app_cname"`
}

func newAppCnameTask() *AppCnameTask {
	return &AppCnameTask{
		Name: "Create shipa app cname",
	}
}

type AppCname struct {
	Shipa   `yaml:",inline"`
	App     string `yaml:"app"`
	Cname   string `yaml:"cname"`
	Encrypt bool   `yaml:"encrypt"`
}

type AppDeployTask struct {
	Name      string    `yaml:"name"`
	AppDeploy AppDeploy `yaml:"shipa_app_deploy"`
}

func newAppDeployTask() *AppDeployTask {
	return &AppDeployTask{
		Name: "Deploy shipa application",
	}
}

type AppDeploy struct {
	Shipa          `yaml:",inline"`
	App            string `yaml:"app"`
	Image          string `yaml:"image"`
	PrivateImage   bool   `yaml:"private-image,omitempty"`
	RegistryUser   string `yaml:"registry-user,omitempty"`
	RegistrySecret string `yaml:"registry-secret,omitempty"`
	Steps          int64  `yaml:"steps,omitempty"`
	StepWeight     int64  `yaml:"step-weight,omitempty"`
	StepInterval   string `yaml:"step-interval,omitempty"`
	Port           int64  `yaml:"port,omitempty"`
	Detach         bool   `yaml:"detach"`
	Message        string `yaml:"message,omitempty"`
	ShipaYaml      string `yaml:"shipayaml,omitempty"`
}
