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
	app.Spec.ForProvider.Name = cfg.AppName
	app.Spec.ForProvider.TeamOwner = cfg.Team
	app.Spec.ForProvider.Framework = cfg.Framework
	app.Spec.ForProvider.Plan = cfg.Plan
	app.Spec.ForProvider.Tags = utils.ParseValues(cfg.Tags)

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
	appDeploy.Spec.ForProvider.App = cfg.AppName
	appDeploy.Spec.ForProvider.Image = cfg.Image
	appDeploy.Spec.ForProvider.RegistryUser = cfg.RegistryUser
	appDeploy.Spec.ForProvider.RegistrySecret = cfg.RegistrySecret
	appDeploy.Spec.ForProvider.Port = cfg.Port
	appDeploy.Spec.ForProvider.PrivateImage = cfg.RegistryUser != "" || cfg.RegistrySecret != ""

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
	appCname.Spec.ForProvider.App = cfg.AppName
	appCname.Spec.ForProvider.Cname = cfg.Cname
	appCname.Spec.ForProvider.Encrypt = cfg.Encrypt

	return appCname
}

func genAppEnv(cfg shipa.Config) *AppEnv {
	if cfg.AppName == "" || cfg.EnvName == "" || cfg.EnvValue == "" {
		return nil
	}

	appEnv := &AppEnv{
		ApiVersion: apiVersion,
		Kind:       "AppEnv",
	}
	appEnv.Metadata.Name = cfg.AppName
	appEnv.Spec.ForProvider.App = cfg.AppName
	appEnv.Spec.ForProvider.AppEnv.Norestart = cfg.Norestart
	appEnv.Spec.ForProvider.AppEnv.Private = cfg.Private
	appEnv.Spec.ForProvider.AppEnv.Envs = append(appEnv.Spec.ForProvider.AppEnv.Envs, Env{
		Name:  cfg.EnvName,
		Value: cfg.EnvValue,
	})

	return appEnv
}
