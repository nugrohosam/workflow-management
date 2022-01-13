package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"runtime"
	"time"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type Print struct {
	Status            bool
	Data              string
	StatusIntermitten bool
}

func main() {
	_, base, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(base)

	ruleByte, err := ioutil.ReadFile(basepath + "/rules/PrintRule.grl")
	if err != nil {
		fmt.Println(err)
	}

	print := &Print{
		Status:            true,
		Data:              "",
		StatusIntermitten: false,
	}
	for {
		print.businessCall(ruleByte)
		time.Sleep(time.Second / 5)
	}

}

func (print *Print) businessCall(rule []byte) {
	dataContext := ast.NewDataContext()

	print.StatusIntermitten = rand.Intn(100)%2 == 0
	dataContext.Add("Print", print)

	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)
	ruleBuilder.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource(rule))

	kb := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	eng1 := &engine.GruleEngine{MaxCycle: 100, ReturnErrOnFailedRuleEvaluation: true}
	eng1.Execute(dataContext, kb)

	fmt.Println(print.Data)
}
