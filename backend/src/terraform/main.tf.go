package terraform

func genMain() string {
	return `terraform {
  required_providers {
    shipa = {
      version = "0.0.8"
      source = "shipa-corp/shipa"
    }
  }
}

provider "shipa" {}
`
}
