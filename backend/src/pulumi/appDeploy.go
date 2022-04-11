package pulumi

import (
	"fmt"
	"shipa-gen/src/shipa"
	"strings"
)

func hasAppDeploy(cfg shipa.Config) bool {
	return cfg.AppName != "" && cfg.Image != ""
}

func genAppDeploy(cfg shipa.Config) string {
	return fmt.Sprintf(`
const appDeploy = new shipa.AppDeploy("app-deploy-1", {
    app: "%s",
%s
});
`, cfg.AppName, genAppDeployParams(cfg))
}

func genAppDeployParams(cfg shipa.Config) string {
	const indent = "   "
	out := []string{
		fmt.Sprintf(`%s image: "%s"`, indent, cfg.Image),
		genAppDeployConfig(cfg),
	}

	if cfg.Port != "" {
		out = append(out, genAppDeployPort(cfg))
	}

	if cfg.RegistryUser != "" && cfg.RegistrySecret != "" {
		out = append(out, genAppDeployRegistry(cfg))
	}

	if len(cfg.Volumes) > 0 {
		out = append(out, genAppDeployVolumes(cfg))
	}

	return strings.Join(out, ",\n")
}

func genAppDeployConfig(cfg shipa.Config) string {
	const indent = "       "
	out := []string{
		fmt.Sprintf(`%s team: "%s"`, indent, cfg.Team),
		fmt.Sprintf(`%s framework: "%s"`, indent, cfg.Framework),
	}

	if cfg.Plan != "" {
		out = append(out, fmt.Sprintf(`%s plan: "%s"`, indent, cfg.Plan))
	}

	tags := genTags(cfg)
	if tags != "" {
		out = append(out, fmt.Sprintf(`%s %s`, indent, tags))
	}

	params := strings.Join(out, ",\n")
	return fmt.Sprintf(`    appConfig: {
%s
    }`, params)
}

func genAppDeployPort(cfg shipa.Config) string {
	if cfg.Port == "" {
		return ""
	}

	return fmt.Sprintf(`    port: {
        number: %s,
        protocol: "TCP"
    }`, cfg.Port)
}

func genAppDeployRegistry(cfg shipa.Config) string {
	if cfg.RegistryUser == "" || cfg.RegistrySecret == "" {
		return ""
	}

	return fmt.Sprintf(`    registry: {
        user: "%s",
        secret: "%s"
    }`, cfg.RegistryUser, cfg.RegistrySecret)
}

func genAppDeployVolumes(cfg shipa.Config) string {
	if len(cfg.Volumes) == 0 {
		return ""
	}

	return fmt.Sprintf(`    volumes: [
%s
    ]`, genVolumes(cfg.Volumes))
}

func genVolumes(volumes []*shipa.Volume) string {
	var items []string
	for _, vol := range volumes {
		items = append(items, genVolume(vol))
	}

	return strings.Join(items, ",\n")
}

func genVolume(vol *shipa.Volume) string {
	return fmt.Sprintf(`        {
%s
        }`, genVolumeFields(vol))
}

func genVolumeFields(vol *shipa.Volume) string {
	const indent = "           "
	fields := []string{
		fmt.Sprintf(`%s name: "%s"`, indent, vol.Name),
		fmt.Sprintf(`%s mountPath: "%s"`, indent, vol.Path),
	}

	if opts := genVolumeOpts(vol.Opts); opts != "" {
		fields = append(fields, opts)
	}

	return strings.Join(fields, ",\n")
}

func genVolumeOpts(opts *shipa.VolumeOptions) string {
	if opts == nil {
		return ""
	}

	const indent = "               "
	var fields []string
	if opts.Prop1 != "" {
		fields = append(fields, fmt.Sprintf(`%s additionalProp1: "%s"`, indent, opts.Prop1))
	}
	if opts.Prop2 != "" {
		fields = append(fields, fmt.Sprintf(`%s additionalProp2: "%s"`, indent, opts.Prop2))
	}
	if opts.Prop3 != "" {
		fields = append(fields, fmt.Sprintf(`%s additionalProp3: "%s"`, indent, opts.Prop3))
	}

	if len(fields) == 0 {
		return ""
	}

	return fmt.Sprintf(`            mountOptions: {
%s
            }`, strings.Join(fields, ",\n"))
}
