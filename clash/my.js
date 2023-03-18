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
        'DOMAIN-KEYWORD,aliyun,🎯 全球直连',
        'DOMAIN,pingfore.qq.com,🎯 全球直连',
        'DOMAIN,thesecretlivesofdata.com,🚀 节点选择',
        'DOMAIN,www.amazon.com,🚀 节点选择',
        'DOMAIN,www.parallels.com,🚀 节点选择',
        'DOMAIN-SUFFIX,nexitally.net,🚀 节点选择',
        'DOMAIN-SUFFIX,gstatic.com,🚀 节点选择',
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
        'DOMAIN-SUFFIX,openai.com,🇺🇸 美国01',
        'DOMAIN-SUFFIX,nvidia.com,🚀 节点选择',
        'DOMAIN-SUFFIX,jsdelivr.net,🚀 节点选择',
        'DOMAIN-SUFFIX,convertio.me,🚀 节点选择',
        'GEOIP,US,🚀 节点选择'
    ]
    const obj = yaml.parse(raw)
    obj.rules = customs.concat(obj.rules)
    return yaml.stringify(obj)
}