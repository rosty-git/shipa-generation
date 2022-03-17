package ansible

import (
	"gopkg.in/yaml.v2"
	"shipa-gen/src/shipa"
	"shipa-gen/src/utils"
	"strconv"
)

func Generate(cfg shipa.Config) *shipa.Result {
	play := newPlay()

	app := genApp(cfg)
	if app != nil {
		play.Tasks = append(play.Tasks, app)
	}

	appEnv := genAppEnv(cfg)
	if appEnv != nil {
		play.Tasks = append(play.Tasks, appEnv)
	}

	appCname := genAppCname(cfg)
	if appCname != nil {
		play.Tasks = append(play.Tasks, appCname)
	}

	appDeploy := genAppDeploy(cfg)
	if appDeploy != nil {
		play.Tasks = append(play.Tasks, appDeploy)
	}

	data, _ := yaml.Marshal([]interface{}{play})
	return &shipa.Result{
		Name:    "play.yml",
		Content: string(data),
	}
}

func genAppDeploy(cfg shipa.Config) *AppDeployTask {
	if cfg.AppName == "" || cfg.Image == "" {
		return nil
	}

	var port int64
	val, err := strconv.ParseInt(cfg.Port, 10, 64)
	if err == nil {
		port = val
	}
	t := newAppDeployTask()
	t.AppDeploy = AppDeploy{
		Shipa:          credentials,
		App:            cfg.AppName,
		Image:          cfg.Image,
		RegistryUser:   cfg.RegistryUser,
		RegistrySecret: cfg.RegistrySecret,
		Port:           port,
		PrivateImage:   cfg.RegistryUser != "" || cfg.RegistrySecret != "",
	}
	return t
}

func genAppCname(cfg shipa.Config) *AppCnameTask {
	if cfg.AppName == "" || cfg.Cname == "" {
		return nil
	}

	t := newAppCnameTask()
	t.AppCname = AppCname{
		Shipa:   credentials,
		App:     cfg.AppName,
		Cname:   cfg.Cname,
		Encrypt: cfg.Encrypt,
	}
	return t
}

func genAppEnv(cfg shipa.Config) *AppEnvTask {
	if cfg.AppName == "" || len(cfg.Envs) == 0 {
		return nil
	}

	var envs []Env
	for _, env := range cfg.Envs {
		envs = append(envs, Env{
			Name:  env.Name,
			Value: env.Value,
		})
	}

	t := newAppEnvTask()
	t.AppEnv = AppEnv{
		Shipa:     credentials,
		App:       cfg.AppName,
		Envs:      envs,
		Norestart: cfg.Norestart,
		Private:   cfg.Private,
	}
	return t
}

var credentials = Shipa{
	ShipaHost:  "{{ shipa_host }}",
	ShipaToken: "{{ shipa_token }}",
}

func genApp(cfg shipa.Config) *AppTask {
	if cfg.AppName == "" || cfg.Team == "" || cfg.Framework == "" || cfg.Plan == "" {
		return nil
	}

	t := newAppTask()
	t.App = App{
		Shipa:     credentials,
		Name:      cfg.AppName,
		Teamowner: cfg.Team,
		Framework: cfg.Framework,
		Plan:      cfg.Plan,
		Tags:      utils.ParseValues(cfg.Tags),
	}
	return t
}
