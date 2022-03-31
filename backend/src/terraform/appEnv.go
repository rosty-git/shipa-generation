package terraform

import (
	"fmt"
	"shipa-gen/src/shipa"
	"strconv"
	"strings"
)

func hasAppEnv(cfg shipa.Config) bool {
	return cfg.AppName != "" && len(cfg.Envs) > 0
}

func genAppEnv(cfg shipa.Config) string {
	return fmt.Sprintf(`
# Set app envs
resource "shipa_app_env" "tf" {
  app = %s
  app_env {
%s
   norestart = %s
   private = %s
  }
  %s
}
`, getAppName(cfg), getEnvs(cfg), strconv.FormatBool(cfg.Norestart), strconv.FormatBool(cfg.Private), dependsOnApp(cfg))
}

func getEnvs(cfg shipa.Config) string {
	var envs []string
	for _, e := range cfg.Envs {
		envs = append(envs, fmt.Sprintf(`   envs {
     name = "%s"
     value = "%s"
   }`, e.Name, e.Value))
	}
	return strings.Join(envs, "\n")
}

func dependsOnApp(cfg shipa.Config) string {
	if hasApp(cfg) && !hasAppDeploy(cfg) {
		return "depends_on = [shipa_app.tf]"
	}
	return ""
}
