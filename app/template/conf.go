package template

const clash = `/*
parsers: # array
  - reg: https://suo.st/*
    remote:
      url: https://raw.githubusercontent.com/zjy282/stoverride/main/clash/my.js
      cache: true
*/
module.exports.parse = async (raw, { axios, yaml, notify, console }, { name, url, interval, selected }) => {
    var customs = [
        %s
    ]
    const obj = yaml.parse(raw)
    let rules = obj.rules
    let domains = {};
    let customsDomain = []
    let customsOther = []
    let customsLast = []
    let originDomain = []
    let originOther = []
    let originLast = []

    for (let i = 0; i < customs.length; i++) {
        if (customs[i].startsWith("DOMAIN")) {
            domains[customs[i].split(",")[1]] = true
            customsDomain.push(customs[i])
        } else if (customs[i].startsWith("GEOIP") || customs[i].startsWith("MATCH")) {
            customsLast.push(customs[i])
        } else {
            customsOther.push(customs[i])
        }
    }

    for (let i = 0; i < rules.length; i++) {
        let domainItem = rules[i].split(",")[1]
        if (domains[domainItem]) {
            continue
        }
        if (rules[i].startsWith("DOMAIN")) {
            originDomain.push(rules[i])
        } else if (rules[i].startsWith("GEOIP") || rules[i].startsWith("MATCH")) {
            originLast.push(rules[i])
        } else {
            originOther.push(rules[i])
        }
    }

    obj.rules = customsDomain.concat(originDomain)
    obj.rules = obj.rules.concat(customsOther).concat(originOther)
    obj.rules = obj.rules.concat(customsLast).concat(originLast)

    const proxiesUS = ["DIRECT"];
	const proxiesSG = ["DIRECT"];
	const proxiesTW = ["DIRECT"];
    for (let i = 0; i < obj.proxies.length; i++) {
        if (obj.proxies[i].name.search("美国") !== -1) {
            proxiesUS.push(obj.proxies[i].name)
        }
		if (obj.proxies[i].name.search("新加坡") !== -1) {
            proxiesSG.push(obj.proxies[i].name)
        }
        if (obj.proxies[i].name.search("台湾") !== -1) {
            proxiesTW.push(obj.proxies[i].name)
        }
    }
    obj["proxy-groups"].push({
        "name": "🇺🇲 美国节点",
        "type": "select",
        url: "http://www.gstatic.com/generate_204",
        interval: 300,
        proxies: proxiesUS
    })
	obj["proxy-groups"].push({
        "name": "🇸🇬 新加坡节点",
        "type": "select",
        url: "http://www.gstatic.com/generate_204",
        interval: 300,
        proxies: proxiesSG
    })
	obj["proxy-groups"].push({
        "name": "🇨🇳 台湾节点",
        "type": "select",
        url: "http://www.gstatic.com/generate_204",
        interval: 300,
        proxies: proxiesTW
    })
    obj["proxy-groups"].push({
        "name": "🧑🏼‍💻 科学网络",
        "type": "select",
        url: "https://google.com/",
        interval: 1000,
        proxies: ["DIRECT", "🇺🇲 美国节点", "🚀 节点选择"]
    })
    return yaml.stringify(obj)
}
`
const stash = `name: Kevin的重写规则
desc: 仅自用

rules:
%s
`

const lineTemplate1 = `'%s',`
const lineTemplate2 = `  - %s`

type RuleFileType int

const (
	AllRuleWithProxy RuleFileType = iota
	PartRuleWithoutProxy
)

type Conf struct {
	FileType     RuleFileType
	Key          string
	Template     string
	Target       string
	LineTemplate string
}

func GetTemplateConf() (targetList []Conf) {
	targetList = append(targetList, Conf{
		FileType:     AllRuleWithProxy,
		Key:          "clash",
		Template:     clash,
		Target:       "./clash/my.js",
		LineTemplate: lineTemplate1,
	})
	targetList = append(targetList, Conf{
		FileType:     AllRuleWithProxy,
		Key:          "qx",
		Template:     `%s`,
		Target:       "./quantumultx/Kevin",
		LineTemplate: `%s`,
	})
	targetList = append(targetList, Conf{
		FileType:     AllRuleWithProxy,
		Key:          "stash",
		Template:     stash,
		Target:       "./stash/my.stoverride",
		LineTemplate: lineTemplate2,
	})
	targetList = append(targetList, Conf{
		FileType:     PartRuleWithoutProxy,
		Key:          "surge",
		Template:     `%s`,
		Target:       "./surge/Kevin-%s",
		LineTemplate: `%s`,
	})
	return
}
