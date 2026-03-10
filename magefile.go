//go:build mage
// +build mage

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/magefile/mage/sh"

	_ "github.com/elastic/elastic-transport-go/v8/elastictransport"
)

type mutation struct {
	// Copy files
	From, To string

	// Go Replace by regex
	Match, Replace, Glob string

	DeleteGlob string
}

type DependencyGithub struct {
	Repo   string
	Branch string
}

var (
	deps = []DependencyGithub{
		{Repo: "https://github.com/elastic/go-elasticsearch",
			Branch: "9.3"},
	}

	// Transform the codebase so it can build
	mutations = []mutation{}
)

func replace_string_in_file(filename string, old string, new string) error {
	read, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	newContents := strings.Replace(string(read), old, new, -1)
	return ioutil.WriteFile(filename, []byte(newContents), 0644)
}

func maybeClone(dep DependencyGithub) error {
	base := filepath.Base(dep.Repo)
	_, err := os.Lstat(base)
	if err == nil {
		return nil
	}

	branch := "master"
	if dep.Branch != "" {
		branch = dep.Branch
	}

	return sh.RunV("git", "clone", "--depth", "1",
		"--single-branch", "-b", branch, dep.Repo)
}

func Build() error {
	err := os.MkdirAll("build", 0700)
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(cwd)

	err = os.Chdir("build")
	if err != nil {
		return err
	}

	for _, dep := range deps {
		err = maybeClone(dep)
		if err != nil {
			return err
		}
	}

	for _, m := range mutations {
		if m.From != "" {
			fmt.Printf("Copying %v to %v\n", m.From, m.To)
			basedir := filepath.Dir(m.To)
			os.MkdirAll(basedir, 0755)

			err := sh.Copy(m.To, m.From)
			if err != nil {
				return err
			}
		}

		if m.DeleteGlob != "" {
			basepath, pattern := doublestar.SplitPattern(m.DeleteGlob)
			fsys := os.DirFS(basepath)
			matches, err := doublestar.Glob(fsys, pattern)
			if err != nil {
				return err
			}

			for _, match := range matches {
				filename := filepath.Join(basepath, match)
				fmt.Printf("Deleting %v in %v\n", m.Match, filename)
				err = os.Remove(filename)
				if err != nil {
					return err
				}
			}
		}

		if m.Glob != "" {
			basepath, pattern := doublestar.SplitPattern(m.Glob)
			fsys := os.DirFS(basepath)
			matches, err := doublestar.Glob(fsys, pattern)
			if err != nil {
				return err
			}

			for _, match := range matches {
				filename := filepath.Join(basepath, match)
				fmt.Printf("Replacing %v in %v\n", m.Match, filename)
				err = replace_string_in_file(filename, m.Match, m.Replace)
				if err != nil {
					return err
				}
			}
		}

	}

	return nil
}
