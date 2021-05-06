package main

import (
	"flag"
	"fmt"
)

func main() {

	var (
		list           bool
		email          string
		folderToAdd    string
		folderToDelete string
		//		scanFolderOnly string
	)

	flag.BoolVar(&list, "list", false, "List folders that will be scanned for Git repositories")
	flag.StringVar(&email, "email", "", "the email to scan")
	flag.StringVar(&folderToAdd, "add", "", "add new folder to list to scan for Git repositories")
	flag.StringVar(&folderToDelete, "delete", "", "remove folder from list to scan for Git repositories")
	//	flag.StringVar(&scanFolderOnly, "filter", "", "filtered path - display only Git repositories in path and its subfolders")

	flag.Parse()

	if list {
		fmt.Print("List all repositories that we are scanning and going to display the Git stats")
		listAllRepos()
		return
	}

	if folderToAdd != "" {
		fmt.Println("Following folder will be added to the scanning list:", folderToAdd)
		addFolder(folderToAdd)
	}

	if folderToDelete != "" {
		fmt.Println("Following folder will be removed from scanning list:", folderToDelete)
		deleteFolder(folderToDelete)
	}

	if email == "" {
		fmt.Println("No specific user defined")
		fmt.Println("Stats contribution of all users will be displayed")
		fmt.Println("NOT YET IMPLEMENTED")
	} else {
		fmt.Println("Specific user defined")
		scanRepos(email)
	}

}

func deleteFolder(folderToDelete string) {
	fmt.Println("Remove folder", folderToDelete, "from list for scanning Git repositories")
}

func addFolder(folderToAdd string) {
	fmt.Println("Add folder", folderToAdd, "to list for scanning Git repositories")
}

func listAllRepos() {
	fmt.Println("List all Repos on console")
}

func scanRepos(email string) {
	fmt.Println("Scan repo with user:", email)
}
