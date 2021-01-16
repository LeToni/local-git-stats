package main

import (
	"flag"
	"os"
	"strings"
)

func scan(folder string) {
	repositories := recursiveScanFolder(folder)
	filepath := getDotFilePath()
	addNewFoundRepositories(filePath, repositories)
}

func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

func scanGitFolders(folders []string, folder string) []string {
	folder = strings.TrimSuffix(folder, "/")
	f, err := os.Open(folder)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	files, err := f.Readdir(-1)
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
			}
			folders = scanGitFolders(folders, path)
		}
	}

	return folders
}

func stats(email string) {
	print("Displaying stats...")
}

func main() {
	var (
		folder string
		email  string
	)

	flag.StringVar(&folder, "add", "", "Add new folder to scan for git repositories")
	flag.StringVar(&email, "email", "user@email.com", "Email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)
}
