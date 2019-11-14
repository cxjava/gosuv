package main

import (
	"fmt"
)

var (
	Version            = "dev"
	Commit             = "none"
	RepoUrl            = "unknown"
	BuildDate          = "unknown"
	BuiltBy            = "unknown"
	BuiltWithGoVersion = "unknown"
)

func showVersion() {
	fmt.Println("Version: \t" + Version)
	fmt.Println("Build Time: \t" + BuildDate)
	fmt.Println("Go Version: \t" + BuiltWithGoVersion)
	fmt.Println("Repo URL: \t" + RepoUrl)
	fmt.Println("Commit Info: \t" + Commit)
	fmt.Println("Build By: \t" + BuiltBy)
}
