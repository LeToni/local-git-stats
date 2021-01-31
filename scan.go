package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

func scan(folder string) {
	repositories := recursiveScanFolder(folder)
	filePath := getDotFilePath()
	addNewFoundRepositories(filePath, repositories)
}

func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

func scanGitFolders(folders []string, folder string) []string {
	folder = strings.TrimSuffix(folder, "/")
	f, err := os.Open(folder)
	if err != nil {
		panic(err)
	}
	files, err := f.Readdir(-1)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	var path string

	for _, file := range files {
		if file.IsDir() {
			path = folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				folders = append(folders, path)
				continue
			}
			folders = scanGitFolders(folders, path)
		}
	}

	return folders
}

func getDotFilePath() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	dotFile := usr.HomeDir + "/.gogitlocalstats"

	return dotFile
}

func addNewFoundRepositories(filePath string, newRepos []string) {
	existingRepos := parseDotFile(filePath)
	repos := joinRepos(newRepos, existingRepos)
	WriteToDotFile(repos, filePath)
}

func parseDotFile(filePath string) []string {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create(filePath)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	var repos []string
	for scanner := bufio.NewScanner(f); scanner.Scan(); {
		repo := scanner.Text()
		repos = append(repos, repo)
	}

	return repos
}

func joinRepos(newRepos []string, existingRepos []string) []string {
	for _, new := range newRepos {
		if !reposContains(existingRepos, new) {
			existingRepos = append(existingRepos, new)
		}
	}
	return existingRepos
}

func reposContains(repos []string, repo string) bool {
	for _, r := range repos {
		if r == repo {
			return true
		}
	}
	return false
}

func WriteToDotFile(repos []string, filePath string) {
	contentFile := strings.Join(repos, "\n")
	err := ioutil.WriteFile(filePath, []byte(contentFile), 0644)
	if err != nil {
		panic(err)
	}
}
