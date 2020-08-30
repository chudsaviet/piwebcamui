// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

func Proto() error {
	fmt.Println("Generating Protobuf sources...")
	os.MkdirAll("build/gen", 0700)
	return sh.Run("protoc", "-I=proto", "--go_out=.", "proto/webcam_config.proto")
}

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(InstallDeps)

	if err := Proto(); err != nil {
		return err
	}

	fmt.Println("Building...")
	cmd := exec.Command("go", "build", "-o", "piwebcamui", ".")
	return cmd.Run()
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing Deps...")

	return sh.Run("go", "mod", "download")
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.Remove("piwebcamui")
	os.RemoveAll("build/gen")
}

// Format all source files
func Fmt() error {
	return sh.Run("gofmt", "-s", "-w", ".")
}

// Run the app
func Run() error {
	if err := Build(); err != nil {
		return err
	}

	fmt.Println("Running application...")
	return sh.RunV("./piwebcamui")
}
