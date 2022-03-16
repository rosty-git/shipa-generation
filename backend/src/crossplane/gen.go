package crossplane

import (
	"gopkg.in/yaml.v2"
	"shipa-gen/src/shipa"
	"shipa-gen/src/utils"
)

const apiVersion = "shipa.crossplane.io/v1alpha1"

func Generate(cfg shipa.Config) *shipa.Result {
	var resource []interface{}

	app := genApp(cfg)
	if app != nil {
		resource = append(resource, app)
	}

	appDeploy := genAppDeploy(cfg)
	if appDeploy != nil {
		resource = append(resource, appDeploy)
	}

	appCname := genAppCname(cfg)
	if appCname != nil {
		resource = append(resource, appCname)
	}

	appEnv := genAppEnv(cfg)
	if appEnv != nil {
		resource = append(resource, appEnv)
	}

	if len(resource) == 0 {
		return nil
	}

	data, _ := yaml.Marshal(resource)
	return &shipa.Result{
		Name:    "crossplane.yaml",
		Content: string(data),
	}
}

func genApp(cfg shipa.Config) *App {
	if cfg.AppName == "" || cfg.Team == "" || cfg.Framework == "" || cfg.Plan == "" {
		return nil
	}

	app := &App{
		ApiVersion: apiVersion,
		Kind:       "App",
	}
	app.Metadata.Name = cfg.AppName
	app.Spec.ForProvider = AppForProvider{
		Name:      cfg.AppName,
		TeamOwner: cfg.Team,
		Framework: cfg.Framework,
		Plan:      cfg.Plan,
		Tags:      utils.ParseValues(cfg.Tags),
	}

	return app
}

func genAppDeploy(cfg shipa.Config) *AppDeploy {
	if cfg.AppName == "" || cfg.Image == "" {
		return nil
	}

	appDeploy := &AppDeploy{
		ApiVersion: apiVersion,
		Kind:       "AppDeploy",
	}
	appDeploy.Metadata.Name = cfg.AppName
	appDeploy.Spec.ForProvider = AppDeployForProvider{
		App:            cfg.AppName,
		Image:          cfg.Image,
		RegistryUser:   cfg.RegistryUser,
		RegistrySecret: cfg.RegistrySecret,
		Port:           cfg.Port,
		PrivateImage:   cfg.RegistryUser != "" || cfg.RegistrySecret != "",
	}

	return appDeploy
}

func genAppCname(cfg shipa.Config) *AppCname {
	if cfg.AppName == "" || cfg.Cname == "" {
		return nil
	}

	appCname := &AppCname{
		ApiVersion: apiVersion,
		Kind:       "AppCname",
	}
	appCname.Metadata.Name = cfg.AppName
	appCname.Spec.ForProvider = AppCnameForProvider{
		App:     cfg.AppName,
		Cname:   cfg.Cname,
		Encrypt: cfg.Encrypt,
	}

	return appCname
}

func genAppEnv(cfg shipa.Config) *AppEnv {
	if cfg.AppName == "" || len(cfg.Envs) == 0 {
		return nil
	}

	appEnv := &AppEnv{
		ApiVersion: apiVersion,
		Kind:       "AppEnv",
	}
	appEnv.Metadata.Name = cfg.AppName
	p := &appEnv.Spec.ForProvider
	p.App = cfg.AppName
	p.AppEnv.Norestart = cfg.Norestart
	p.AppEnv.Private = cfg.Private
	for _, env := range cfg.Envs {
		p.AppEnv.Envs = append(p.AppEnv.Envs, Env{
			Name:  env.Name,
			Value: env.Value,
		})
	}

	return appEnv
}
