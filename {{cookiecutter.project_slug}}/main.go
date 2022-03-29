package main

import (
	"github.com/mt-inside/go-usvc"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.project_slug}}/pkg/health"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.project_slug}}/version"
)

func main() {
	log := usvc.GetLogger(false, 10)
	log.Info("Starting", "app", "{{cookiecutter.project_slug}}", "version", version.Version)

	go health.ListenHealthChecks(":8090")
}
