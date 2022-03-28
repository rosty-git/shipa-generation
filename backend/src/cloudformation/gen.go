package cloudformation

import (
	"gopkg.in/yaml.v2"
	"shipa-gen/src/shipa"
	"shipa-gen/src/utils"
	"strconv"
)

const (
	shipaHost  = "{{resolve:secretsmanager:ShipaHost}}"
	shipaToken = "{{resolve:secretsmanager:ShipaToken}}"
)

func Generate(cfg shipa.Config) *shipa.Result {
	var resource []interface{}

	appDeploy := genAppDeploy(cfg)
	if appDeploy != nil {
		resource = append(resource, appDeploy)
	} else {
		app := genApp(cfg)
		if app != nil {
			resource = append(resource, app)
		}
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
	p.AppConfig = AppConfig{
		Team:      cfg.Team,
		Framework: cfg.Framework,
		Plan:      cfg.Plan,
		Tags:      utils.ParseValues(cfg.Tags),
	}
	p.Registry = genAppDeployRegistry(cfg)
	p.Port = genAppDeployPort(cfg)

	return appDeploy
}

func genAppDeployRegistry(cfg shipa.Config) *Registry {
	if cfg.RegistryUser == "" || cfg.RegistrySecret == "" {
		return nil
	}

	return &Registry{
		User:   cfg.RegistryUser,
		Secret: cfg.RegistrySecret,
	}
}

func genAppDeployPort(cfg shipa.Config) *Port {
	if cfg.Port == "" {
		return nil
	}

	port, err := strconv.ParseInt(cfg.Port, 10, 64)
	if err != nil {
		return nil
	}

	return &Port{
		Number:   port,
		Protocol: "TCP",
	}
}

func genAppEnv(cfg shipa.Config) *AppEnv {
	if cfg.AppName == "" || len(cfg.Envs) == 0 {
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
	for _, env := range cfg.Envs {
		p.Envs = append(p.Envs, Env{
			Name:  env.Name,
			Value: env.Value,
		})
	}

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
