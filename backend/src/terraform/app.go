package terraform

import (
	"fmt"
	"shipa-gen/src/shipa"
	"shipa-gen/src/utils"
	"strings"
)

func hasApp(cfg shipa.Config) bool {
	return cfg.AppName != "" && cfg.Team != "" && cfg.Framework != ""
}

func genApp(cfg shipa.Config) string {
	app := fmt.Sprintf(`
# Create app
resource "shipa_app" "tf" {
  app {
    name = "%s"
    teamowner = "%s"
    framework = "%s"`, cfg.AppName, cfg.Team, cfg.Framework)

	tags := genTags(cfg)
	if tags != "" {
		app = fmt.Sprintf(`%s
    %s`, app, tags)
	}

	app = fmt.Sprintf(`%s
  }
}
`, app)

	return app
}

func genTags(cfg shipa.Config) string {
	tags := utils.ParseValues(cfg.Tags)
	if len(tags) == 0 {
		return ""
	}

	return fmt.Sprintf(`tags = ["%s"]`, strings.Join(tags, `", "`))
}
