package entity

import "github.com/projectdiscovery/goflags"

type CmdOptions struct {
	Target              string
	Proxy               string
	ProxyUrl            string
	Config              string
	Headless            bool
	Deep                int
	CrawlCode           *goflags.StringSlice
	DirSearch           bool
	DirSearchCode       goflags.StringSlice
	IncludeScope        goflags.StringSlice
	ExclusionScope      goflags.StringSlice
	SameDomain          bool
	AllowTag            goflags.StringSlice
	IncludeCve          goflags.StringSlice
	DnsLog              string
	TemplatePath        string
	GenerateCA          bool
	KeyCAPath           string
	CrtCAPath           string
	PemCAPath           string
	ManagerSevice       bool
	ManagerTarget       string
	LoadBalancing       bool
	LoadBalancingTarges goflags.StringSlice
}
