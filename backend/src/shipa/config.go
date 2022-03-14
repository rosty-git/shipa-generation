package shipa

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

	EnvName   string `json:"envName"`
	EnvValue  string `json:"envValue"`
	Norestart bool   `json:"norestart"`
	Private   bool   `json:"private"`
}

type Result struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
