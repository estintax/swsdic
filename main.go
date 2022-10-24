package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) < 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <filename> <output without .go> <package name> <variable name>\n", os.Args[0])
		os.Exit(1)
		return
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while opening file: %s\n", err.Error())
		os.Exit(2)
		return
	}

	defer file.Close()

	stat, _ := file.Stat()
	bytes := make([]byte, stat.Size())
	n, err := file.Read(bytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while reading file: %s\n", err.Error())
		os.Exit(3)
		return
	}

	fmt.Fprintf(os.Stdout, "Readed %d bytes.\n", n)

	outputFile, err := os.OpenFile(args[1]+".go", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while creating file: %s\n", err.Error())
		os.Exit(4)
		return
	}

	defer outputFile.Close()

	n, err = outputFile.WriteString(fmt.Sprintf("package %s\n\nvar %s []byte = []byte", args[2], args[3]))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while writing to file: %s\n", err.Error())
		os.Exit(5)
		return
	}
	fmt.Fprintf(os.Stdout, "Writing... %d, ", n)
	n, err = outputFile.WriteString(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprint(bytes), "[", "{"), "]", "}"), " ", ","))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while writing to file: %s\n", err.Error())
		os.Exit(5)
		return
	}
	fmt.Fprintf(os.Stdout, "%d. Done\n", n)
}
