package github

type Action struct {
	App       *App       `yaml:"app,omitempty"`
	AppEnv    *AppEnv    `yaml:"app-env,omitempty"`
	AppCname  *AppCname  `yaml:"app-cname,omitempty"`
	AppDeploy *AppDeploy `yaml:"app-deploy,omitempty"`
}

type App struct {
	Name      string   `json:"name" yaml:"name,omitempty"`
	Pool      string   `json:"pool,omitempty" yaml:"framework,omitempty"`
	TeamOwner string   `json:"teamOwner,omitempty" yaml:"teamOwner,omitempty"`
	Plan      string   `json:"plan,omitempty" yaml:"plan,omitempty"`
	Tags      []string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

type AppEnv struct {
	App       string `json:"-" yaml:"app"`
	Envs      []*Env `json:"envs" yaml:"envs"`
	NoRestart bool   `json:"norestart" yaml:"norestart"`
	Private   bool   `json:"private" yaml:"private"`
}

type Env struct {
	Name  string `json:"name" yaml:"name"`
	Value string `json:"value" yaml:"value"`
}

type AppCname struct {
	App       string `json:"-" yaml:"app"`
	Cname     string `json:"cname" yaml:"cname"`
	Encrypted bool   `json:"-" yaml:"encrypted"`
}

type AppDeploy struct {
	App            string `yaml:"app"`
	Image          string `json:"image" yaml:"image"`
	PrivateImage   bool   `json:"private-image,omitempty" yaml:"private-image,omitempty"`
	RegistryUser   string `json:"registry-user,omitempty" yaml:"registry-user,omitempty"`
	RegistrySecret string `json:"registry-secret,omitempty" yaml:"registry-secret,omitempty"`
	Steps          int64  `json:"steps,omitempty" yaml:"steps,omitempty"`
	StepWeight     int64  `json:"step-weight,omitempty" yaml:"step-weight,omitempty"`
	StepInterval   string `json:"step-interval,omitempty" yaml:"step-interval,omitempty"`
	Port           int64  `json:"port,omitempty" yaml:"port,omitempty"`
	Detach         bool   `json:"detach" yaml:"detach"`
	Message        string `json:"message,omitempty" yaml:"message,omitempty"`
	ShipaYaml      string `json:"shipayaml,omitempty" yaml:"shipayaml,omitempty"`
}
