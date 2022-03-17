package terraform

import (
	"fmt"
	"shipa-gen/src/shipa"
	"strconv"
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
    image = "%s"
%s
  }
  %s
}
`, getAppName(cfg), cfg.Image, genAppDeployOptional(cfg), dependsOnAppEnvOrApp(cfg))
}

func dependsOnAppEnvOrApp(cfg shipa.Config) string {
	switch {
	case hasAppEnv(cfg):
		return "depends_on = [shipa_app_env.tf]"
	case hasApp(cfg):
		return "depends_on = [shipa_app.tf]"
	default:
		return ""
	}
}

func genAppDeployOptional(cfg shipa.Config) string {
	var out []string

	if cfg.Port != "" {
		out = append(out, fmt.Sprintf(`    port = %s`, cfg.Port))
	}

	if cfg.RegistryUser != "" || cfg.RegistrySecret != "" {
		out = append(out, fmt.Sprintf(`    private_image = %s`, strconv.FormatBool(true)))
	}

	if cfg.RegistryUser != "" {
		out = append(out, fmt.Sprintf(`    registry_user = "%s"`, cfg.RegistryUser))
	}

	if cfg.RegistrySecret != "" {
		out = append(out, fmt.Sprintf(`    registry_secret = "%s"`, cfg.RegistrySecret))
	}

	return strings.Join(out, "\n")
}
