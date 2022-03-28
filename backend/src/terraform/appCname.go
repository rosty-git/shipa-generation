package terraform

import (
	"fmt"
	"shipa-gen/src/shipa"
	"strconv"
)

func hasAppCname(cfg shipa.Config) bool {
	return cfg.AppName != "" && cfg.Cname != ""
}

func genAppCname(cfg shipa.Config) string {
	return fmt.Sprintf(`
# Set app cname
resource "shipa_app_cname" "tf" {
  app = %s
  cname = "%s"
  encrypt = %s
}
`, getAppName(cfg), cfg.Cname, strconv.FormatBool(cfg.Encrypt))
}

func getAppName(cfg shipa.Config) string {
	//if hasApp(cfg) {
	//	return "shipa_app.tf.app[0].name"
	//}

	return fmt.Sprintf(`"%s"`, cfg.AppName)
}
