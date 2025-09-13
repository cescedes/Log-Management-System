package main

import (
	"fmt"
	"io/ioutil"
  "os"
)

func createLogDirectory() {
	fmt.Println("Enter the name of the log directory you'd like to create:")
	var dirName string
	fmt.Scanln(&dirName)

	err:= os.Mkdir(dirName, 0755)

  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("Log directory '%s' created successfully.\n", dirName)

}

func changeLogDirectory() {
	cwd := getWorkingDirectory()
	fmt.Println("Current Directory:", cwd)

	fmt.Print("Enter the name of the log directory you want to change to: ")
	var newDirName string
	fmt.Scanln(&newDirName)

	err := os.Chdir(newDirName)
  if err != nil {
    fmt.Println(err)
    return
  }


	newCwd := getWorkingDirectory()
	fmt.Println("Directory changed successfully. Current Directory:", newCwd)

  cwd, err := os.Getwd()
  if err != nil {
    fmt.Println(err)
    return ""
  }
  return title

}

func printWorkingDirectory() {
	cwd := getWorkingDirectory()
	fmt.Println("Current Working Directory:", cwd)
}

func listDirectoryContents() {
	cwd := getWorkingDirectory()
	fmt.Println("Current Working Directory:", cwd)

	files, err := ioutil.ReadDir(cwd)
	if err != nil {
		fmt.Println("Error reading directory contents:", err)
		return
	}

	fmt.Println("Directory Contents:")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func getWorkingDirectory() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return cwd
}
