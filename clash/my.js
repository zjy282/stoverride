/*
parsers: # array
  - reg: https://suo.st/*
    remote:
      url: https://raw.githubusercontent.com/zjy282/stoverride/main/clash/my.js
      cache: true
*/
module.exports.parse = async (raw, { axios, yaml, notify, console }, { name, url, interval, selected }) => {
    var customs = [
        'DOMAIN-KEYWORD,rockstargames,🚀 节点选择',
'DOMAIN-SUFFIX,jetbrains.com,🚀 节点选择',
'DOMAIN-SUFFIX,playstation.com,🚀 节点选择',
'DOMAIN-SUFFIX,playstation.net,🚀 节点选择',
'DOMAIN-SUFFIX,playstationnetwork.com,🚀 节点选择',
'DOMAIN-SUFFIX,sublimetext.com,🚀 节点选择',
'DOMAIN-KEYWORD,typora,🚀 节点选择',
'DOMAIN-KEYWORD,postman,🚀 节点选择',
'DOMAIN-SUFFIX,tunnelblick.net,🚀 节点选择',
'DOMAIN-SUFFIX,skicat.net,🚀 节点选择',
'DOMAIN-SUFFIX,skimeow.com,🚀 节点选择',
'DOMAIN-SUFFIX,redisdesktop.com,🚀 节点选择',
'DOMAIN-SUFFIX,resp.app,🚀 节点选择',
'DOMAIN-SUFFIX,doesitarm.com,🚀 节点选择',
'DOMAIN-SUFFIX,isapplesiliconready.com,🚀 节点选择',
'DOMAIN-SUFFIX,gravatar.com,🚀 节点选择',
'DOMAIN-SUFFIX,brew.sh,🚀 节点选择',
'DOMAIN,thesecretlivesofdata.com,🚀 节点选择',
'DOMAIN,www.amazon.com,🚀 节点选择',
'DOMAIN,www.parallels.com,🚀 节点选择',
'DOMAIN-SUFFIX,nexitally.net,🚀 节点选择',
'DOMAIN-SUFFIX,vox.rocks,🚀 节点选择',
'DOMAIN-SUFFIX,azureedge.net,🚀 节点选择',
'DOMAIN-SUFFIX,paoche.info,🚀 节点选择',
'DOMAIN-SUFFIX,yodobashi.com,🚀 节点选择',
'DOMAIN-KEYWORD,fuli,🚀 节点选择',
'DOMAIN-SUFFIX,52.mk,🚀 节点选择',
'DOMAIN-SUFFIX,id9.cc,🚀 节点选择',
'DOMAIN-SUFFIX,suo.st,🚀 节点选择',
'DOMAIN-SUFFIX,suo.yt,🚀 节点选择',
'DOMAIN,api.subcloud.xyz,🚀 节点选择',
'DOMAIN-SUFFIX,nvidia.com,🚀 节点选择',
'DOMAIN-SUFFIX,jsdelivr.net,🚀 节点选择',
'DOMAIN-SUFFIX,convertio.me,🚀 节点选择',
'DOMAIN-SUFFIX,pythonhosted.org,🚀 节点选择',
'DOMAIN-KEYWORD,depay,🚀 节点选择',
'DOMAIN-SUFFIX,githubusercontent.com,🧑🏼‍💻 科学网络',
'DOMAIN-KEYWORD,github,🧑🏼‍💻 科学网络',
'DOMAIN-SUFFIX,github.com,🧑🏼‍💻 科学网络',
'DOMAIN-SUFFIX,midjourney.com,🧑🏼‍💻 科学网络',
'DOMAIN-SUFFIX,discord.com,🧑🏼‍💻 科学网络',
'DOMAIN-SUFFIX,stripe.com,🧑🏼‍💻 科学网络',
'DOMAIN-SUFFIX,openai.azure.com,🧑🏼‍💻 科学网络',
'DOMAIN-SUFFIX,openai.com,🧑🏼‍💻 科学网络',
'DOMAIN-KEYWORD,openai,🧑🏼‍💻 科学网络',
'DOMAIN-SUFFIX,cloudfront.net,🇺🇲 美国节点',
'DOMAIN-SUFFIX,claude.ai,🇺🇲 美国节点',
'DOMAIN,bard.google.com,🇺🇲 美国节点',
'IP-CIDR,52.58.0.0/15,🧑🏼‍💻 科学网络',
'DOMAIN-SUFFIX,intellij.net,🧑🏼‍💻 科学网络',
'DOMAIN,bard.google.com,🇺🇲 美国节点',
'GEOIP,US,🚀 节点选择'
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

    const proxies = ["DIRECT"];
	const proxies2 = ["DIRECT"];
    for (let i = 0; i < obj.proxies.length; i++) {
        if (obj.proxies[i].name.search("美国") !== -1) {
            proxies.push(obj.proxies[i].name)
        }
		if (obj.proxies[i].name.search("新加坡") !== -1) {
            proxies2.push(obj.proxies[i].name)
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
        "name": "🇸🇬 新加坡节点",
        "type": "select",
        url: "http://www.gstatic.com/generate_204",
        interval: 300,
        proxies: proxies2
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
