package main

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/zjy282/stoverride/app/template"
	"os"
	"strings"
)

var MappingRule = map[string]string{
	"PROXY":  "🚀 节点选择",
	"DIRECT": "🎯 全球直连",
	"US":     "🇺🇲 美国节点",
	"QSKX":   "🧑🏼‍💻 科学网络",
}

func main() {
	rules, err := readConf()
	if err != nil {
		color.Red.Println(err)
		return
	}

	resultInfo := genResult(rules)

	err = writeConf(resultInfo)
	if err != nil {
		color.Red.Println(err)
		return
	}
	color.Green.Println("WriteFile Done...")
}

func writeConf(resultInfo map[string]string) error {
	for target, value := range resultInfo {
		err := os.WriteFile(target, []byte(value), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func genResult(rules []string) (resultInfo map[string]string) {
	resultInfo = make(map[string]string)
	for _, conf := range template.GetTemplateConf() {
		var resultRules []string
		for _, ruleOne := range rules {
			resultRules = append(resultRules, fmt.Sprintf(conf.LineTemplate, ruleOne))
		}
		resultRules[len(resultRules)-1] = strings.TrimRight(resultRules[len(resultRules)-1], ",")
		result := fmt.Sprintf(conf.Template, strings.Join(resultRules, "\n"))
		resultInfo[conf.Target] = result
	}
	return
}

func readConf() (rules []string, err error) {
	data, err := os.ReadFile("./rules.txt")
	if err != nil {
		return
	}

	ruleList := strings.Split(string(data), "\n")

	for _, ruleOne := range ruleList {
		ruleOne = strings.TrimSpace(ruleOne)
		if ruleOne == "" {
			continue
		}
		ruleInfo := strings.Split(ruleOne, ",")
		if target, ok := MappingRule[ruleInfo[2]]; ok {
			ruleInfo[2] = target
			rules = append(rules, strings.Join(ruleInfo, ","))
		}
	}
	return
}
