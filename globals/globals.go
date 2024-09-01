package globals

import (
	"sync"
)

type Options struct {
	All        bool `short:"a" long:"all" description:"Display all system info."`
	CPU        bool `short:"c" long:"cpu" description:"Display system CPU usage."`
	CpuOptions struct {
		PerCPU bool `long:"per-cpu" description:"Display information per-CPU."`
	} `group:"CPU Options"`
	DiskOptions struct {
		ByMount []string `long:"by-mount" description:"Search for disk(s) by mount point. Can be used multiple times."`
		Inode   bool     `short:"i" long:"inode" description:"Display disk inode information."`
		Usage   bool     `short:"u" long:"usage" description:"Display disk usage information."`
	} `group:"Disk Options"`
	HostOptions struct {
		HostTemps bool `long:"temps" description:"Display host sensor temperatures."`
		HostUsers bool `long:"users" description:"Display users on the host."`
	} `group:"Host Options"`
	NetworkOptions struct {
		FilterCounters   bool `long:"filter-counters" description:"Display iptables conntrack statistics."`
		IOStats          bool `long:"io-stats" description:"Display I/O statistics for every installed interface."`
		Interfaces       bool `long:"interfaces" description:"Display network interfaces."`
		ProtocolCounters bool `long:"proto-counters" description:"Display network protocol counters."`
	} `group:"Network Options"`
	ProcessOptions struct {
		ByName []string `long:"by-name" description:"Search for processes by name. Partial names are acceptable. Can be used multiple times."`
		ByPid  []int32  `long:"by-pid" description:"Show process with a given <pid>. Can be used multiple times."`
		ByUser []string `long:"by-user" description:"Show processes belonging to <user>. Can be used multiple times."`
	} `group:"Process Options"`
	Disk         bool `short:"d" long:"disk" description:"Display system disk usage."`
	Host         bool `long:"host" description:"Display host statistics."`
	Load         bool `short:"l" long:"load" description:"Display system load averages."`
	Memory       bool `short:"m" long:"memory" description:"Display system memory usage."`
	Net          bool `short:"n" long:"network" description:"Display network connection information."`
	Swap         bool `short:"s" long:"swap" description:"Display swap memory usage."`
	Process      bool `short:"p" long:"process" description:"Display process information."`
	PrintVersion bool `short:"V" long:"version" description:"Output version information and exit."`
}

var (
	mu      sync.RWMutex
	options Options
)

func GetOptions() (x Options) {
	mu.Lock()
	x = options
	mu.Unlock()
	return x
}
