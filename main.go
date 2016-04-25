package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type FindReplaceVisitor struct {
	find    string
	replace string
}

func (v FindReplaceVisitor) visit(path string, f os.FileInfo, err error) error {
	if !f.Mode().IsRegular() {
		return nil
	}

	input, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	output := strings.Replace(string(input), v.find, v.replace, -1)
	err = ioutil.WriteFile(path, []byte(output), f.Mode())
	return err
}

func printUsage() {
	fmt.Println("findreplace\n===========")
	fmt.Println("\tFind and replace a string recursively within the current directory's hierarchy.")

	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("\tfindreplace findString replaceString")

}

func main() {
	if len(os.Args) < 3 {
		printUsage()
		os.Exit(1)
	}

	findString := os.Args[1]
	replaceString := os.Args[2]

	visitor := FindReplaceVisitor{findString, replaceString}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	err = filepath.Walk(cwd, visitor.visit)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
