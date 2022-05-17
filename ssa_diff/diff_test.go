package ssa_diff

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go/build"
	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/pointer"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
	"strings"
	"testing"
)

func inStd(node *callgraph.Node) bool {
	if node.Func.Pkg == nil {
		return false
	}
	pkg, _ := build.Import(node.Func.Pkg.Pkg.Path(), "", 0)
	return pkg.Goroot
}

func LoadGraph(path string) (*callgraph.Graph, error) {
	cfg := &packages.Config{
		Dir:   path,
		Mode:  packages.LoadAllSyntax,
		Tests: false,
		//BuildFlags: build.Default.BuildTags,
	}
	initPkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return nil, fmt.Errorf("load packages failed. %+v", err)
	}
	prog, pkgs := ssautil.AllPackages(initPkgs, 0)
	prog.Build()
	mainPkgs := ssautil.MainPackages(pkgs)
	result, err := pointer.Analyze(&pointer.Config{
		Mains:          mainPkgs,
		BuildCallGraph: true,
	})
	if err != nil {
		return nil, fmt.Errorf("create callgraph failed. %+v", err)
	}
	return result.CallGraph, nil
}

func ssaFunc2String(ssaFunc *ssa.Function) string {
	var b bytes.Buffer
	ssa.WriteFunction(&b, ssaFunc)
	lines := strings.Split(b.String(), "\n")
	output := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			// 忽略ssa注释内容
			continue
		}
		output += line + "\n"
	}
	return output
}

func FindFunc(graph *callgraph.Graph, funcName string) *ssa.Function {
	for ssaFunc, node := range graph.Nodes {

		if inStd(node) {
			continue
		}
		if strings.Contains(ssaFunc.Name(), funcName) {
			return ssaFunc
		}
	}
	return nil
}

func TestStructInsertedDiff(t *testing.T) {
	sourceGraph, err := LoadGraph("./source")
	assert.Nil(t, err)
	targetGraph, err := LoadGraph("./target")
	assert.Nil(t, err)
	sourceFunc := FindFunc(sourceGraph, "StInsertFunc")
	targetFunc := FindFunc(targetGraph, "StInsertFunc")
	srcFuncStr := ssaFunc2String(sourceFunc)
	tgtFuncStr := ssaFunc2String(targetFunc)
	if !strings.EqualFold(srcFuncStr, tgtFuncStr) {
		t.Errorf("func modified")
		t.Logf("==== src func begin ======")
		t.Logf("%s", srcFuncStr)
		t.Logf("==== src func end ======")
		t.Logf("==== tgt func begin ======")
		t.Logf("%s", tgtFuncStr)
		t.Logf("==== tgt func end ======")
	}
}
