package main

import (
	"flag"
)

func scan(folder string) {
	repositories := recursiveScanFolder(folder)
	filepath := getDotFilePath()
	addNewFoundRepositories(filePath, repositories)
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
