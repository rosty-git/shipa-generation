package github

import (
	"gopkg.in/yaml.v2"
	"shipa-gen/src/shipa"
	"shipa-gen/src/utils"
	"strconv"
)

func Generate(cfg shipa.Config) *shipa.Result {
	var action Action

	action.AppEnv = genAppEnv(cfg)
	action.AppCname = genAppCname(cfg)
	action.AppDeploy = genAppDeploy(cfg)
	if action.AppDeploy == nil {
		action.App = genApp(cfg)
	}

	if action.App == nil && action.AppEnv == nil &&
		action.AppCname == nil && action.AppDeploy == nil {
		return nil
	}

	data, _ := yaml.Marshal(action)
	return &shipa.Result{
		Name:    "shipa-action.yml",
		Content: string(data),
	}
}

func genAppDeploy(cfg shipa.Config) *AppDeploy {
	if cfg.AppName == "" || cfg.Image == "" {
		return nil
	}

	return &AppDeploy{
		App:   cfg.AppName,
		Image: cfg.Image,
		AppConfig: &AppDeployConfig{
			Team:      cfg.Team,
			Framework: cfg.Framework,
			Plan:      cfg.Plan,
			Tags:      utils.ParseValues(cfg.Tags),
		},
		Registry: genAppDeployRegistry(cfg),
		Port:     genAppDeployPort(cfg),
	}
}

func genAppDeployRegistry(cfg shipa.Config) *AppDeployRegistry {
	if cfg.RegistryUser == "" || cfg.RegistrySecret == "" {
		return nil
	}

	return &AppDeployRegistry{
		User:   cfg.RegistryUser,
		Secret: cfg.RegistrySecret,
	}
}

func genAppDeployPort(cfg shipa.Config) *AppDeployPort {
	if cfg.Port == "" {
		return nil
	}

	port, err := strconv.ParseInt(cfg.Port, 10, 64)
	if err != nil {
		return nil
	}

	return &AppDeployPort{
		Number:   port,
		Protocol: "TCP",
	}
}

func genAppCname(cfg shipa.Config) *AppCname {
	if cfg.AppName == "" || cfg.Cname == "" {
		return nil
	}

	return &AppCname{
		App:       cfg.AppName,
		Cname:     cfg.Cname,
		Encrypted: cfg.Encrypt,
	}
}

func genApp(cfg shipa.Config) *App {
	if cfg.AppName == "" || cfg.Team == "" || cfg.Framework == "" || cfg.Plan == "" {
		return nil
	}

	return &App{
		Name:      cfg.AppName,
		TeamOwner: cfg.Team,
		Pool:      cfg.Framework,
		Plan:      cfg.Plan,
		Tags:      utils.ParseValues(cfg.Tags),
	}
}

func genAppEnv(cfg shipa.Config) *AppEnv {
	if cfg.AppName == "" || len(cfg.Envs) == 0 {
		return nil
	}

	var envs []*Env
	for _, env := range cfg.Envs {
		envs = append(envs, &Env{
			Name:  env.Name,
			Value: env.Value,
		})
	}

	return &AppEnv{
		App:       cfg.AppName,
		Envs:      envs,
		NoRestart: cfg.Norestart,
		Private:   cfg.Private,
	}
}
