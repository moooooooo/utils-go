// Copyright Peter Moore Jan 30, 2020

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	//	"io"
	//	"io/ioutil"
	//	"os"
	//	"os/signal"
	"path/filepath"
	//	"regexp"
	//	"sort"
	//	"strings"
	//	"syscall"
	//	"github.com/digitalocean/godo"
)

var gitPaths [2]string

func main() {
	name := "Git Pull All"
	fmt.Println("Starting", name)
	// Mac - add in OS test for dual platform or $HOME var
	// gitPaths[0] = "/Users/peter/src/github"
	// gitPaths[1] = "/Users/peter/go/src"
	gitPaths[0] = "/home/peter/src/github"
	gitPaths[1] = "/home/peter/go/src"
	//createGitArray2(gitPaths[0])
	createGitArray()
}

func getEnvs() {
	// this function will pull in the envs needed
	fmt.Println("On Unix:")
	fmt.Println(filepath.Dir("/foo/bar/baz.js"))
	fmt.Println(filepath.Dir("/foo/bar/baz"))
	fmt.Println(filepath.Dir("/foo/bar/baz/"))
	fmt.Println(filepath.Dir("/dirty//path///"))
	fmt.Println(filepath.Dir("dev.txt"))
	fmt.Println(filepath.Dir("../todo.txt"))
	fmt.Println(filepath.Dir(".."))
	fmt.Println(filepath.Dir("."))
	fmt.Println(filepath.Dir("/"))
	fmt.Println(filepath.Dir(""))
}

func createGitArray() {
	// hardcode for now until i figure out
	// creating an array of dirnames to iterate

	// Mac - add in OS test for dual platform or $HOME var
	//err := filepath.Walk("/Users/peter/src/github",
	err := filepath.Walk("/home/peter/src/github",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//var gitPaths []string
			// fmt.Println(path, info.Size())
			if info.Name() == ".git" {
				// now we need ot get the path without .git - done
				// now put it into an array
				var gitPath = strings.TrimRight(path, ".git")
				fmt.Println("Found Git ")
				gitPaths := []string{gitPath}
				for i, v := range gitPaths {
					fmt.Println(i, v)
				}

			}

			return nil
		})
	// pass the gitPaths array to another function
	fmt.Println("Sending the array to be processed")
	if err != nil {
		log.Println(err)
	}

}

func createGitArray2(thePath string) {
	// hardcode for now until i figure out
	// creating an array of dirnames to iterate

	filepath.Walk(thePath,
		func(path string, info os.FileInfo, err error) error {

			//var aaa int = 0
			//var gitPaths []string
			// fmt.Println(path, info.Size())
			if info.Name() == ".git" {
				// now we need ot get the path without .git - done
				// now put it into an array
				var gitPath = strings.TrimRight(path, ".git")
				fmt.Println("Found Git ", gitPath)
				// this is only working for one entry - need to push all
				//gitPaths := []string{gitPath}
				//gitPaths[aaa] = gitPath

			}
			fmt.Println("before for loop...")
			for i, v := range gitPaths {
				fmt.Println(i, v)
			}
			fmt.Println("after for loop...")
			return nil
		})
	// pass the gitPaths array to another function
	fmt.Println("Sending the array to be processed")

	//fmt.Println("test ", len(gitPaths))
}
