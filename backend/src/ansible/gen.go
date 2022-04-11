package ansible

import (
	"gopkg.in/yaml.v2"
	"shipa-gen/src/shipa"
	"shipa-gen/src/utils"
	"strconv"
)

func Generate(cfg shipa.Config) *shipa.Result {
	play := newPlay()

	appEnv := genAppEnv(cfg)
	if appEnv != nil {
		play.Tasks = append(play.Tasks, appEnv)
	}

	appCname := genAppCname(cfg)
	if appCname != nil {
		play.Tasks = append(play.Tasks, appCname)
	}

	policy := genNetworkPolicy(cfg)
	if policy != nil {
		play.Tasks = append(play.Tasks, policy)
	}

	appDeploy := genAppDeploy(cfg)
	if appDeploy != nil {
		play.Tasks = append(play.Tasks, appDeploy)
	} else {
		app := genApp(cfg)
		if app != nil {
			play.Tasks = append(play.Tasks, app)
		}
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

	t := newAppDeployTask()
	t.AppDeploy = AppDeploy{
		Shipa: credentials,
		App:   cfg.AppName,
		Image: cfg.Image,
		AppConfig: &AppConfig{
			Team:      cfg.Team,
			Framework: cfg.Framework,
			Plan:      cfg.Plan,
			Tags:      utils.ParseValues(cfg.Tags),
		},
		Registry: genAppDeployRegistry(cfg),
		Port:     genAppDeployPort(cfg),
		Volumes:  genAppDeployVolumes(cfg),
	}
	return t
}

func genAppDeployVolumes(cfg shipa.Config) (volumes []*AppDeployVolume) {
	for _, volume := range cfg.Volumes {
		volumes = append(volumes, genAppDeployVolume(volume))
	}
	return
}

func genAppDeployVolume(volume *shipa.Volume) *AppDeployVolume {
	if volume == nil {
		return nil
	}

	return &AppDeployVolume{
		Name:    volume.Name,
		Path:    volume.Path,
		Options: genVolumeOptions(volume.Opts),
	}
}

func genVolumeOptions(opts *shipa.VolumeOptions) *VolumeOptions {
	if opts == nil {
		return nil
	}

	return &VolumeOptions{
		Prop1: opts.Prop1,
		Prop2: opts.Prop2,
		Prop3: opts.Prop3,
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

func genAppDeployRegistry(cfg shipa.Config) *Registry {
	if cfg.RegistryUser == "" || cfg.RegistrySecret == "" {
		return nil
	}

	return &Registry{
		User:   cfg.RegistryUser,
		Secret: cfg.RegistrySecret,
	}
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

func genNetworkPolicy(cfg shipa.Config) *NetworkPolicyTask {
	if cfg.AppName == "" || cfg.NetworkPolicy == nil {
		return nil
	}

	t := newNetworkPolicyTask()
	t.NetworkPolicy = NetworkPolicy{
		Shipa: credentials,
		App:   cfg.AppName,
	}

	utils.CopyJsonData(cfg.NetworkPolicy, &t.NetworkPolicy)

	return t
}
