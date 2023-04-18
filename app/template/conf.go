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
    let domains = {};

    for (let i = 0; i < customs.length; i++) {
        let domainItem = customs[i].split(",")[1]
        domains[domainItem] = true
    }
    for (let i = 0; i < obj.rules.length; i++) {
        let domainItem = obj.rules[i].split(",")[1]
        if (domains[domainItem]) {
            delete obj.rules[i]
        }
    }
    obj.rules = obj.rules.slice(0,-2).concat(customs.concat(obj.rules.slice(-2)))

    const proxies = ["DIRECT"];
    for (let i = 0; i < obj.proxies.length; i++) {
        if (obj.proxies[i].name.search("美国") !== -1) {
            proxies.push(obj.proxies[i].name)
        }
    }
    obj["proxy-groups"].push({
        "name": "🇺🇲 美国节点",
        "type": "select",
        url: "http://www.gstatic.com/generate_204",
        interval: 300,
        proxies: proxies
    })
	obj["proxy-groups"].push({
        "name": "🧑🏼‍💻 科学网络",
        "type": "select",
        url: "http://www.gstatic.com/generate_204",
        interval: 300,
        proxies: ["DIRECT", "🚀 节点选择"]
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

type Conf struct {
	Template     string
	Target       string
	LineTemplate string
}

func GetTemplateConf() (targetList []Conf) {
	targetList = append(targetList, Conf{
		Template:     clash,
		Target:       "./clash/my.js",
		LineTemplate: lineTemplate1,
	})
	targetList = append(targetList, Conf{
		Template:     `%s`,
		Target:       "./quantumultx/Kevin",
		LineTemplate: `%s`,
	})
	targetList = append(targetList, Conf{
		Template:     stash,
		Target:       "./stash/my.stoverride",
		LineTemplate: lineTemplate2,
	})
	return
}
