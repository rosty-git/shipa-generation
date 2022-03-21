package pulumi

import (
	"shipa-gen/src/shipa"
)

func Generate(cfg shipa.Config) *shipa.Result {
	content := genMain()

	if hasApp(cfg) {
		content += genApp(cfg)
	}

	if hasAppCname(cfg) {
		content += genAppCname(cfg)
	}

	if hasAppEnv(cfg) {
		content += genAppEnv(cfg)
	}

	if hasAppDeploy(cfg) {
		content += genAppDeploy(cfg)
	}

	return &shipa.Result{
		Name:    "index.ts",
		Content: content,
	}
}
