package terraform

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
# Deploy app
resource "shipa_app_deploy" "tf" {
  app = %s
  deploy {
%s
  }
  %s
}
`, getAppName(cfg), genAppDeployParams(cfg), dependsOnAppEnvOrApp(cfg))
}

func genAppDeployParams(cfg shipa.Config) string {
	const indent = "   "
	out := []string{
		fmt.Sprintf(`%s image = "%s"`, indent, cfg.Image),
		genAppDeployConfig(cfg),
	}

	if cfg.Port != "" {
		out = append(out, genAppDeployPort(cfg))
	}

	if cfg.RegistryUser != "" && cfg.RegistrySecret != "" {
		out = append(out, genAppDeployRegistry(cfg))
	}

	return strings.Join(out, "\n")
}

func dependsOnAppEnvOrApp(cfg shipa.Config) string {
	switch {
	case hasAppEnv(cfg):
		return "depends_on = [shipa_app_env.tf]"
	default:
		return ""
	}
}

func genAppDeployConfig(cfg shipa.Config) string {
	const indent = "     "
	out := []string{
		fmt.Sprintf(`%s team = "%s"`, indent, cfg.Team),
		fmt.Sprintf(`%s framework = "%s"`, indent, cfg.Framework),
	}

	if cfg.Plan != "" {
		out = append(out, fmt.Sprintf(`%s plan = "%s"`, indent, cfg.Plan))
	}

	tags := genTags(cfg)
	if tags != "" {
		out = append(out, fmt.Sprintf(`%s %s`, indent, tags))
	}

	params := strings.Join(out, "\n")
	return fmt.Sprintf(`    app_config {
%s
    }`, params)
}

func genAppDeployPort(cfg shipa.Config) string {
	if cfg.Port == "" {
		return ""
	}

	return fmt.Sprintf(`    port {
      number = %s
      protocol = "TCP"
    }`, cfg.Port)
}

func genAppDeployRegistry(cfg shipa.Config) string {
	if cfg.RegistryUser == "" || cfg.RegistrySecret == "" {
		return ""
	}

	return fmt.Sprintf(`    registry {
      user = "%s"
      secret = "%s"
    }`, cfg.RegistryUser, cfg.RegistrySecret)
}
