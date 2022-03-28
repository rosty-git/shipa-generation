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
	App            string                   `json:"-" yaml:"app"`
	Image          string                   `json:"image" yaml:"image"`
	AppConfig      *AppDeployConfig         `json:"appConfig" yaml:"appConfig"`
	CanarySettings *AppDeployCanarySettings `json:"canarySettings,omitempty" yaml:"canarySettings,omitempty"`
	PodAutoScaler  *AppDeployPodAutoScaler  `json:"podAutoScaler,omitempty" yaml:"podAutoScaler,omitempty"`
	Port           *AppDeployPort           `json:"port,omitempty" yaml:"port,omitempty"`
	Registry       *AppDeployRegistry       `json:"registry,omitempty" yaml:"registry,omitempty"`
	Volumes        []*AppDeployVolume       `json:"volumesToBind,omitempty" yaml:"volumesToBind,omitempty"`
}

type AppDeployConfig struct {
	Team        string   `json:"team" yaml:"team"`
	Framework   string   `json:"framework" yaml:"framework"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
	Env         []string `json:"env,omitempty" yaml:"env,omitempty"`
	Plan        string   `json:"plan,omitempty" yaml:"plan,omitempty"`
	Router      string   `json:"router,omitempty" yaml:"router,omitempty"`
	Tags        []string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

type AppDeployCanarySettings struct {
	StepInterval int64 `json:"stepInterval" yaml:"stepInterval"`
	StepWeight   int64 `json:"stepWeight" yaml:"stepWeight"`
	Steps        int64 `json:"steps" yaml:"steps"`
}

type AppDeployPodAutoScaler struct {
	MaxReplicas                    int64 `json:"maxReplicas" yaml:"maxReplicas"`
	MinReplicas                    int64 `json:"minReplicas" yaml:"minReplicas"`
	TargetCPUUtilizationPercentage int64 `json:"targetCPUUtilizationPercentage" yaml:"targetCPUUtilizationPercentage"`
}

type AppDeployPort struct {
	Number   int64  `json:"number" yaml:"number"`
	Protocol string `json:"protocol" yaml:"protocol"`
}

type AppDeployRegistry struct {
	User   string `json:"user" yaml:"user"`
	Secret string `json:"secret" yaml:"secret"`
}

type AppDeployVolume struct {
	Name    string         `json:"volumeName" yaml:"volumeName"`
	Path    string         `json:"volumeMountPath" yaml:"volumeMountPath"`
	Options *VolumeOptions `json:"volumeMountOptions" yaml:"volumeMountOptions"`
}

type VolumeOptions struct {
	Prop1 string `json:"additionalProp1" yaml:"additionalProp1"`
	Prop2 string `json:"additionalProp2" yaml:"additionalProp2"`
	Prop3 string `json:"additionalProp3" yaml:"additionalProp3"`
}
