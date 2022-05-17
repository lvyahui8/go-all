package main

import (
	"bytes"
	"fmt"
	"go/build"
	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/pointer"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func inStd(node *callgraph.Node) bool {
	pkg, _ := build.Import(node.Func.Pkg.Pkg.Path(), "", 0)
	return pkg.Goroot
}

func main() {
	//	var ssaFunc *ssa.Function
	cfg := &packages.Config{
		Dir:   "./sample_app",
		Mode:  packages.LoadAllSyntax,
		Tests: false,
		//BuildFlags: build.Default.BuildTags,
	}
	initPkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		fmt.Println(fmt.Errorf("load packages failed. %+v", err))
		return
	}
	prog, pkgs := ssautil.AllPackages(initPkgs, 0)
	prog.Build()
	mainPkgs := ssautil.MainPackages(pkgs)
	result, err := pointer.Analyze(&pointer.Config{
		Mains:          mainPkgs,
		BuildCallGraph: true,
	})
	if err != nil {
		fmt.Println(fmt.Errorf("create callgraph failed. %+v", err))
		return
	}
	for ssaFunc, node := range result.CallGraph.Nodes {
		if ssaFunc.Pkg == nil {
			continue
		}
		if inStd(node) {
			continue
		}
		//pkgPath := ssaFunc.Pkg.Pkg.Path()
		//if !(strings.EqualFold(pkgPath, "sample_app") || strings.HasPrefix(pkgPath, "sample_app/")) {
		//	continue
		//}
		var b bytes.Buffer
		ssa.WriteFunction(&b, ssaFunc)
		fmt.Println("============func begin =====")
		fmt.Println(b.String())
		fmt.Println("============func end   =====")
	}
}
