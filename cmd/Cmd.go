package cmd

import (
	"github.com/z-bool/Venom/internal/entity"
)

func Execute() entity.CmdOptions {
	// 	flagSet := goflags.NewFlagSet()
	// 	flagSet.SetDescription(`

	// --       description: 基于Golang的一款被动式集成漏扫工具
	// --       version: v1.0
	// `)
	// 	opt := &entity.CmdOptions{}

	// 	flagSet.CreateGroup("Target", "用户启动相关配置",
	// 		flagSet.StringVarP(&opt.Target, "target", "t", "127.0.0.1:1234", "配置监听的IP:Port，默认127.0.0.1:1234"),
	// 		flagSet.StringVarP(&opt.Proxy, "proxy", "px", "", "配置http/socks代理信息"),
	// 		flagSet.StringVarP(&opt.ProxyUrl, "proxy-url", "pu", "", "配置代理的接口地址,如https://ip:4321/random"),
	// 	)
	// 	flagSet.CreateGroup("PathScan", "路径扫描相关配置",
	// 		flagSet.BoolVarP(&opt.Headless, "headless", "hl", false, "是否通过chrome强化爬行,默认false(开启影响速度)"),
	// 		flagSet.IntVarP(&opt.Deep, "deep", "d", 1, "设置爬行深度,以被动为主默认为1"),
	// 		flagSet.StringSliceVarP(&opt.CrawlCode, "scawl-code", "sc", nil, "设置响应包判别，只保留响应为模糊设定数组的结果(以,分割),默认为空,example: \"20X,30X,40X\"", goflags.CommaSeparatedStringSliceOptions),
	// 		flagSet.BoolVarP(&opt.DirSearch, "dirsearch", "ds", false, "是否开启dirseach遍历url,默认false(容易封IP)"),
	// 		flagSet.StringSliceVarP(&opt.DirSearchCode, "dirseach-code", "dc", nil, "设置响应包判别，只保留dirseach响应为模糊设定数组的结果(以,分割),默认为不限制,example: \"20X,30X,40X\"", goflags.CommaSeparatedStringSliceOptions),
	// 		flagSet.StringSliceVarP(&opt.IncludeScope, "include-scope", "is", nil, "白名单响应后缀判别，只保留响应为设定数组的结果(以,分割),默认为不限制,example: \"js,php,jsp,asp...\"", goflags.CommaSeparatedStringSliceOptions),
	// 		flagSet.StringSliceVarP(&opt.ExclusionScope, "exclusion-scope", "es", nil, "黑名单响应后缀判别，只保留响应为设定数组的结果(以,分割),默认为不限制,example: \"jpeg,img,png...\"", goflags.CommaSeparatedStringSliceOptions),
	// 		flagSet.BoolVarP(&opt.SameDomain, "same-domain", "sd", false, "是否只扫描同域名下的结果,默认为false(被动层次不深不影响)"),
	// 	)
	// 	flagSet.CreateGroup("PocScan", "漏洞验证相关配置",
	// 		flagSet.StringSliceVarP(&opt.AllowTag, "allow-tag", "at", nil, "漏扫等级只扫有等级标记的,默认为不限制,example: \"info,low,medium,high\"", goflags.CommaSeparatedStringSliceOptions),
	// 		flagSet.StringSliceVarP(&opt.IncludeCve, "include-cve", "ic", nil, "指定漏扫POC，example: \"CVE-2020-0109,CVE-2021-0101\"", goflags.CommaSeparatedStringSliceOptions),
	// 		flagSet.StringVarP(&opt.DnsLog, "dnslog", "dl", "", "配置DNSLog验证FastJson、Log4j等,example: sah1jk.test.com"),
	// 		flagSet.StringVarP(&opt.TemplatePath, "template-path", "tp", "", "配置自定义模板路径"),
	// 	)
	// 	flagSet.CreateGroup("Configuration", "其他相关配置",
	// 		flagSet.BoolVarP(&opt.GenerateCA, "generate-ca", "gc", false, "是否下载证书,默认false"),
	// 		flagSet.StringVarP(&opt.KeyCAPath, "key-ca-path", "kcp", "./ca.key", "证书安装路径，默认./ca.key"),
	// 		flagSet.StringVarP(&opt.CrtCAPath, "crt-ca-path", "ccp", "./ca.crt", "证书安装路径，默认./ca.crt"),
	// 		flagSet.StringVarP(&opt.PemCAPath, "pem-ca-path", "pcp", "./ca.pem", "证书安装路径，默认./ca.pem"),
	// 		flagSet.BoolVarP(&opt.ManagerSevice, "manager-service", "ms", true, "由于文件写入消耗性能改为rpc传输结果到服务端另处理，服务端来处理损耗(默认开启系统)"),
	// 		flagSet.StringVarP(&opt.ManagerTarget, "manager-target", "mt", "http://127.0.0.1:4321/", "漏洞结果管理平台的路径,默认http://127.0.0.1:4321/"),
	// 		flagSet.BoolVarP(&opt.LoadBalancing, "load-balancing", "lb", false, "如果有多台主机负载均衡进行漏洞验证及爬行，可以有效减少ip封禁及性能损耗的问题,默认false(实现滞后)"),
	// 		flagSet.StringSliceVarP(&opt.LoadBalancingTarges, "load-balancing-targets", "lbt", nil, "指定负载均衡的主机数组，example: \"192.168.8.12:1234,10.10.8.12:1234\"(实现滞后)", goflags.CommaSeparatedStringSliceOptions),
	// 	)
	// 	if err := flagSet.Parse(); err != nil {
	// 		log.Fatalf("Could not parse flags: %s\n", err)
	// 	}
	// 	if opt.Config != "" {
	// 		if err := flagSet.MergeConfigFile(opt.Config); err != nil {
	// 			log.Fatalf("Could not merge config file: %s\n", err)
	// 		}
	// 	}
	// 	return *opt
	return entity.CmdOptions{}
}
