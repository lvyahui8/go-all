package main

import (
	"bytes"
	"fmt"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/pointer"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func main() {
	//	var ssaFunc *ssa.Function
	cfg := &packages.Config{
		Dir:   "./sample_app",
		Tests: false,
	}
	initPkgs, err := packages.Load(cfg, "...")
	if err != nil {
		fmt.Println(fmt.Errorf("load packages failed. %+v", err))
		return
	}
	_, pkgs := ssautil.AllPackages(initPkgs, 0)
	mainPkgs := ssautil.MainPackages(pkgs)
	result, err := pointer.Analyze(&pointer.Config{
		Mains:          mainPkgs,
		BuildCallGraph: true,
	})
	if err != nil {
		fmt.Println(fmt.Errorf("create callgraph failed. %+v", err))
		return
	}
	for ssaFunc, _ := range result.CallGraph.Nodes {
		bytes.NewBuffer()
		ssa.WriteFunction()
	}
}
