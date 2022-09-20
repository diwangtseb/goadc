/*
Copyright © 2022 diwangtseb <diwang839639311@gmail.com>

*/
package cmd

import (
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"

	"github.com/diwangtseb/goadc/helper"
	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var path string

func init() {
	rootCmd.Flags().StringVarP(&path, "path", "p", "./example/demo.go", "add context to the file")
}

var rootCmd = cobra.Command{
	Use:   "goadc",
	Short: "go add context",
	Long:  "go add context by cli",
	Run:   pathFunc,
}

func pathFunc(_ *cobra.Command, _ []string) {
	err := astParse()
	if err != nil {
		return
	}
}

func astParse() error {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		return err
	}
	fmt.Println(file)
	for _, node := range file.Decls {
		v := &helper.Visitor{}
		vs := v.Visit(node)
		fmt.Println(vs)
	}

	var output []byte
	buffer := bytes.NewBuffer(output)
	err = format.Node(buffer, fset, file)
	if err != nil {
		log.Fatal(err)
	}
	// 输出Go代码
	fmt.Println(buffer.String())
	return nil
}
