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
    obj.rules = customs.concat(obj.rules)
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
