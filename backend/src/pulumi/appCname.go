package pulumi

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
const appCname = new shipa.AppCname("app-cname-1", {
    app: "%s",
    cname: "%s",
    encrypt: %s
});
`, cfg.AppName, cfg.Cname, strconv.FormatBool(cfg.Encrypt))
}
