package pulumi

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
const appDeploy = new shipa.AppEnv("app-env-1", {
    app: "%s",
    appEnv: {
        envs: [
%s
        ],
        norestart: %s,
        private: %s
    }
});
`, cfg.AppName, getEnvs(cfg), strconv.FormatBool(cfg.Norestart), strconv.FormatBool(cfg.Private))
}

func getEnvs(cfg shipa.Config) string {
	const indent = "           "
	var envs []string
	for _, e := range cfg.Envs {
		envs = append(envs, fmt.Sprintf(`%s {name: "%s", value: "%s"},`, indent, e.Name, e.Value))
	}
	return strings.Join(envs, "\n")
}
