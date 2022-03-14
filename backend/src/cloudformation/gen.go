package cloudformation

import (
	"gopkg.in/yaml.v2"
	"shipa-gen/src/shipa"
	"shipa-gen/src/utils"
)

const (
	shipaHost  = "{{resolve:secretsmanager:ShipaHost}}"
	shipaToken = "{{resolve:secretsmanager:ShipaToken}}"
)

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
		Name:    "cloudformation.yaml",
		Content: string(data),
	}
}

func genApp(cfg shipa.Config) *App {
	if cfg.AppName == "" || cfg.Team == "" || cfg.Framework == "" || cfg.Plan == "" {
		return nil
	}

	app := &App{
		Resources: AppResource{
			MyShipaApp: MyShipaApp{
				Type: "Shipa::Application::Item",
			},
		},
	}
	p := &app.Resources.MyShipaApp.Properties
	p.ShipaHost = shipaHost
	p.ShipaToken = shipaToken
	p.Name = cfg.AppName
	p.Teamowner = cfg.Team
	p.Framework = cfg.Framework
	p.Plan = cfg.Plan
	p.Tags = utils.ParseValues(cfg.Tags)

	return app
}

func genAppDeploy(cfg shipa.Config) *AppDeploy {
	if cfg.AppName == "" || cfg.Image == "" {
		return nil
	}

	appDeploy := &AppDeploy{
		Resources: AppDeployResource{
			MyShipaAppDeploy: MyShipaAppDeploy{
				Type: "Shipa::AppDeploy::Item",
			},
		},
	}
	p := &appDeploy.Resources.MyShipaAppDeploy.Properties
	p.ShipaHost = shipaHost
	p.ShipaToken = shipaToken
	p.App = cfg.AppName
	p.Image = cfg.Image
	p.RegistryUser = cfg.RegistryUser
	p.RegistrySecret = cfg.RegistrySecret
	p.Port = cfg.Port
	p.PrivateImage = cfg.RegistryUser != "" || cfg.RegistrySecret != ""

	return appDeploy
}

func genAppEnv(cfg shipa.Config) *AppEnv {
	if cfg.AppName == "" || cfg.EnvName == "" || cfg.EnvValue == "" {
		return nil
	}

	appEnv := &AppEnv{
		Resources: AppEnvResource{
			MyShipaAppEnv: MyShipaAppEnv{
				Type: "Shipa::AppEnv::Item",
			},
		},
	}

	p := &appEnv.Resources.MyShipaAppEnv.Properties
	p.ShipaHost = shipaHost
	p.ShipaToken = shipaToken
	p.App = cfg.AppName
	p.Norestart = cfg.Norestart
	p.Private = cfg.Private
	p.Envs = append(p.Envs, Env{
		Name:  cfg.EnvName,
		Value: cfg.EnvValue,
	})

	return appEnv
}

func genAppCname(cfg shipa.Config) *AppCname {
	if cfg.AppName == "" || cfg.Cname == "" {
		return nil
	}

	appCname := &AppCname{
		Resources: AppCnameResource{
			MyShipaAppCname: MyShipaAppCname{
				Type: "Shipa::AppCname::Item",
			},
		},
	}

	p := &appCname.Resources.MyShipaAppCname.Properties
	p.ShipaHost = shipaHost
	p.ShipaToken = shipaToken
	p.App = cfg.AppName
	p.Cname = cfg.Cname
	p.Encrypt = cfg.Encrypt

	return appCname
}
