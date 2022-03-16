package shipa

type AppsConfig struct {
	Apps []Config `json:"apps"`
}

type Config struct {
	Provider  string `json:"provider"`
	AppName   string `json:"appName"`
	Team      string `json:"team"`
	Framework string `json:"framework"`
	Plan      string `json:"plan"`
	Tags      string `json:"tags"`

	Image          string `json:"image"`
	RegistryUser   string `json:"registryUser"`
	RegistrySecret string `json:"registrySecret"`
	Port           string `json:"port"`

	Cname   string `json:"cname"`
	Encrypt bool   `json:"encrypt"`

	// deprecated
	EnvName string `json:"envName"`
	// deprecated
	EnvValue string `json:"envValue"`

	Envs      []Env `json:"envs"`
	Norestart bool  `json:"norestart"`
	Private   bool  `json:"private"`
}

type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Result struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
