package cloudformation

import (
	"encoding/json"
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

	policy := genNetworkPolicy(cfg)
	if policy != nil {
		resource = append(resource, policy)
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

	return &App{
		Resources: AppResource{
			MyShipaApp: MyShipaApp{
				Type: "Shipa::Application::Item",
				Properties: AppProperties{
					Name:       cfg.AppName,
					ShipaHost:  shipaHost,
					ShipaToken: shipaToken,
					Teamowner:  cfg.Team,
					Framework:  cfg.Framework,
					Plan:       cfg.Plan,
					Tags:       utils.ParseValues(cfg.Tags),
				},
			},
		},
	}
}

func genAppDeploy(cfg shipa.Config) *AppDeploy {
	if cfg.AppName == "" || cfg.Image == "" {
		return nil
	}

	return &AppDeploy{
		Resources: AppDeployResource{
			MyShipaAppDeploy: MyShipaAppDeploy{
				Type: "Shipa::AppDeploy::Item",
				Properties: AppDeployProperties{
					App:        cfg.AppName,
					ShipaHost:  shipaHost,
					ShipaToken: shipaToken,
					Image:      cfg.Image,
					Registry:   genAppDeployRegistry(cfg),
					Port:       genAppDeployPort(cfg),
					AppConfig: AppConfig{
						Team:      cfg.Team,
						Framework: cfg.Framework,
						Plan:      cfg.Plan,
						Tags:      utils.ParseValues(cfg.Tags),
					},
				},
			},
		},
	}
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

	var envs []Env
	for _, env := range cfg.Envs {
		envs = append(envs, Env{
			Name:  env.Name,
			Value: env.Value,
		})
	}

	return &AppEnv{
		Resources: AppEnvResource{
			MyShipaAppEnv: MyShipaAppEnv{
				Type: "Shipa::AppEnv::Item",
				Properties: AppEnvProperties{
					App:        cfg.AppName,
					ShipaHost:  shipaHost,
					ShipaToken: shipaToken,
					Norestart:  cfg.Norestart,
					Private:    cfg.Private,
					Envs:       envs,
				},
			},
		},
	}
}

func genAppCname(cfg shipa.Config) *AppCname {
	if cfg.AppName == "" || cfg.Cname == "" {
		return nil
	}

	return &AppCname{
		Resources: AppCnameResource{
			MyShipaAppCname: MyShipaAppCname{
				Type: "Shipa::AppCname::Item",
				Properties: AppCnameProperties{
					App:        cfg.AppName,
					ShipaHost:  shipaHost,
					ShipaToken: shipaToken,
					Cname:      cfg.Cname,
					Encrypt:    cfg.Encrypt,
				},
			},
		},
	}
}

func genNetworkPolicy(cfg shipa.Config) *NetworkPolicy {
	if cfg.AppName == "" || cfg.Cname == "" {
		return nil
	}

	policy := &NetworkPolicy{
		Resources: NetworkPolicyResource{
			MyShipaNetworkPolicy: MyShipaNetworkPolicy{
				Type: "Shipa::NetworkPolicy::Item",
				Properties: NetworkPolicyProperties{
					App:        cfg.AppName,
					ShipaHost:  shipaHost,
					ShipaToken: shipaToken,
				},
			},
		},
	}

	data, _ := json.Marshal(cfg.NetworkPolicy)
	json.Unmarshal(data, &policy.Resources.MyShipaNetworkPolicy.Properties)

	return policy
}
