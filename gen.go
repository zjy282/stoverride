package main

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"github.com/zjy282/stoverride/app/template"
	"os"
	"strings"
)

var MappingRule = map[string]string{
	"PROXY":  "🚀 节点选择",
	"DIRECT": "🎯 全球直连",
	"TW":     "🇨🇳 台湾节点",
	"US":     "🇺🇲 美国节点",
	"SG":     "🇸🇬 新加坡节点",
	"QSKX":   "🧑🏼‍💻 科学网络",
	"GMAIL":  "📧 Gmail邮件",
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

func genResult(rules []*RuleInfo) (resultInfo map[string]string) {
	resultInfo = make(map[string]string)
	reverseRuleMapping := make(map[string]string)
	for key, value := range MappingRule {
		reverseRuleMapping[value] = key
	}

	for _, conf := range template.GetTemplateConf() {
		if conf.FileType == template.AllRuleWithProxy {
			var resultRules []string
			for _, ruleOne := range rules {
				rule := ruleOne.DefaultRule
				if ruleOne.AdvancedRule != nil {
					if ruleOne.AdvancedRule[conf.Key] != "" {
						rule = ruleOne.AdvancedRule[conf.Key]
					}
				}
				resultRules = append(resultRules, fmt.Sprintf(conf.LineTemplate, rule))
			}
			resultRules[len(resultRules)-1] = strings.TrimRight(resultRules[len(resultRules)-1], ",")
			result := fmt.Sprintf(conf.Template, strings.Join(resultRules, "\n"))
			resultInfo[conf.Target] = result
		} else {
			resultRules := make(map[string][]string)
			for _, ruleOne := range rules {
				rule := ruleOne.DefaultRule
				if ruleOne.AdvancedRule != nil {
					if ruleOne.AdvancedRule[conf.Key] != "" {
						rule = ruleOne.AdvancedRule[conf.Key]
					}
				}
				ruleSplit := strings.Split(rule, ",")
				ruleKey := reverseRuleMapping[ruleSplit[2]]
				resultRules[ruleKey] = append(resultRules[ruleKey], fmt.Sprintf(conf.LineTemplate, strings.Join(ruleSplit[:2], ",")))
			}
			for key, value := range resultRules {
				resultInfo[fmt.Sprintf(conf.Target, key)] = fmt.Sprintf(conf.Template, strings.Join(value, "\n"))
			}
		}

	}
	return
}

type RuleInfo struct {
	DefaultRule  string
	AdvancedRule map[string]string
}

func readConf() (rules []*RuleInfo, err error) {
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
		ruleSplit := strings.Split(ruleOne, ",")
		ruleInfo := RuleInfo{}
		ruleInfo.DefaultRule = strings.Join([]string{ruleSplit[0], ruleSplit[1], MappingRule[ruleSplit[2]]}, ",")
		if len(ruleSplit) > 3 {
			var advancedRule map[string]string
			UnmarshalErr := json.Unmarshal([]byte(ruleSplit[3]), &advancedRule)
			if UnmarshalErr != nil {
				continue
			}
			if len(advancedRule) > 0 {
				ruleInfo.AdvancedRule = make(map[string]string)
			}
			for key, value := range advancedRule {
				ruleInfo.AdvancedRule[key] = strings.Join([]string{ruleSplit[0], ruleSplit[1], MappingRule[value]}, ",")
			}
		}
		rules = append(rules, &ruleInfo)
	}
	return
}
