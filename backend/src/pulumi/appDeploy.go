package pulumi

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
const appDeploy = new shipa.AppDeploy("app-deploy-1", {
    app: "%s",
    deploy: {
%s
    }
});
`, cfg.AppName, genAppDeployParams(cfg))
}

func genAppDeployParams(cfg shipa.Config) string {
	const indent = "       "
	out := []string{
		fmt.Sprintf(`%s image: "%s"`, indent, cfg.Image),
	}

	if cfg.Port != "" {
		out = append(out, fmt.Sprintf(`%s port = %s`, indent, cfg.Port))
	}

	if cfg.RegistryUser != "" || cfg.RegistrySecret != "" {
		out = append(out, fmt.Sprintf(`%s private_image = %s`, indent, strconv.FormatBool(true)))
	}

	if cfg.RegistryUser != "" {
		out = append(out, fmt.Sprintf(`%s registry_user = "%s"`, indent, cfg.RegistryUser))
	}

	if cfg.RegistrySecret != "" {
		out = append(out, fmt.Sprintf(`%s registry_secret = "%s"`, indent, cfg.RegistrySecret))
	}

	return strings.Join(out, ",\n")
}
