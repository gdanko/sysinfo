package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gdanko/sysinfo/gather"
	"github.com/gdanko/sysinfo/globals"
	"github.com/gdanko/sysinfo/internal"
	"github.com/gdanko/sysinfo/stats"
	"github.com/gdanko/sysinfo/util"
	"github.com/jessevdk/go-flags"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/sirupsen/logrus"
)

const (
	PERCPU = false
)

type SysInfo struct {
	PrintVersion bool
	All          bool
	CPU          bool
	Disk         bool
	Host         bool
	Load         bool
	Memory       bool
	Net          bool
	NetFilters   bool
	NetInterface bool
	NetIostat    bool
	NetProto     bool
	Process      bool
	Swap         bool
	Logger       *logrus.Logger
	Output       map[string]interface{}
}

func (s *SysInfo) init(args []string) error {
	var (
		err    error
		opts   globals.Options
		parser *flags.Parser
	)

	opts = globals.GetOptions()
	parser = flags.NewParser(&opts, flags.Default)
	parser.Usage = `[OPTIONS] 
  sysinfo displays information about the system`
	if _, err = parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	s.All = opts.All
	s.CPU = opts.CPU
	s.Disk = opts.Disk
	s.Host = opts.Host
	s.Load = opts.Load
	s.Memory = opts.Memory
	s.Net = opts.Net
	s.NetFilters = opts.NetworkOptions.FilterCounters
	s.NetInterface = opts.NetworkOptions.Interfaces
	s.NetIostat = opts.NetworkOptions.IOStats
	s.NetProto = opts.NetworkOptions.ProtocolCounters
	s.Process = opts.Process
	s.Swap = opts.Swap
	s.Logger = util.ConfigureLogger(logrus.InfoLevel, false)
	s.PrintVersion = opts.PrintVersion

	if len(args) == 1 || s.All {
		s.CPU, s.Disk, s.Host, s.Load, s.Memory, s.Net, s.Swap = true, true, true, true, true, true, true
	}

	return nil
}

func (s *SysInfo) ExitError(errorMessage error) {
	s.Logger.Error(errorMessage.Error())
	os.Exit(1)
}

func (s *SysInfo) ExitCleanly() {
	os.Exit(0)
}

func (s *SysInfo) ParallelTester() {
	s.Output = make(map[string]interface{})
	if s.CPU {
		cpuPercentChannel := make(chan func() ([]stats.PercentStat, error))
		go gather.GetCpuPercent(cpuPercentChannel)
		cpuPercent, err := (<-cpuPercentChannel)()
		if err == nil {
			s.Output["cpu"] = cpuPercent
		}
	}

	if s.Disk {
		diskUsageChannel := make(chan func() ([]stats.DiskUsageData, error))
		go gather.GetDiskUsage(diskUsageChannel)
		diskUsage, err := (<-diskUsageChannel)()
		if err == nil {
			s.Output["disk"] = diskUsage
		}
	}

	if s.Host {
		hostInformationChannel := make(chan func() (stats.HostInformation, error))
		go gather.GetHostInformation(hostInformationChannel)
		hostInformation, err := (<-hostInformationChannel)()
		if err == nil {
			s.Output["host"] = hostInformation
		}
	}

	if s.Load {
		loadAveragesChannel := make(chan func() (*load.AvgStat, error))
		go gather.GetLoadAverages(loadAveragesChannel)
		loadAverages, err := (<-loadAveragesChannel)()
		if err == nil {
			s.Output["load"] = loadAverages
		}
	}

	if s.Memory {
		memoryUsageChannel := make(chan func() (*mem.VirtualMemoryStat, error))
		go gather.GetMemoryUsage(memoryUsageChannel)
		memoryUsage, err := (<-memoryUsageChannel)()
		if err == nil {
			s.Output["memory"] = memoryUsage
		}
	}

	// if s.Net {
	// 	networkThroughputChannel := make(chan func(logger *logrus.Logger, iostatDataOld gather.IOStatData) ([]gather.NetworkInterfaceData, gather.IOStatData, error))
	// 	go gather.GetNetworkThroughput(networkThroughputChannel)
	// 	networkThroughput, iostatDataNew, err := (<-networkThroughputChannel)(s.Logger, s.IostatDataOld)
	// 	if err == nil {
	// 		s.IostatDataOld = iostatDataNew
	// 		output["network"] = networkThroughput
	// 	}
	// }

	if s.Net {
		networkInfoChannel := make(chan func(filterData, interfaceData, iostatData, protocolData bool) (stats.NetworkData, error))
		go gather.GetNetworkData(networkInfoChannel)
		networkData, err := (<-networkInfoChannel)(s.NetFilters, s.NetInterface, s.NetIostat, s.NetProto)
		if err == nil {
			s.Output["network"] = networkData
		}
	}

	if s.Swap {
		swapUsageChannel := make(chan func() (*mem.SwapMemoryStat, error))
		go gather.GetSwapUsage(swapUsageChannel)
		swapUsage, err := (<-swapUsageChannel)()
		if err == nil {
			s.Output["swap"] = swapUsage
		}
	}

	if s.Process {
		processListChannel := make(chan func() ([]stats.Process, error))
		go gather.GetProcessList(processListChannel)
		processList, err := (<-processListChannel)()
		if err == nil {
			s.Output["processes"] = processList
		}
	}
}

func (s *SysInfo) ProcessOutput() {
	jsonBytes, err := json.MarshalIndent(s.Output, "", "    ")
	if err != nil {
		s.ExitError(err)
	}
	fmt.Println(string(jsonBytes))
}

func (s *SysInfo) Run() {
	if s.PrintVersion {
		fmt.Fprintf(os.Stdout, "sysinfo version %s\n", internal.Version(false, true))
		s.ExitCleanly()
	}
	s.ParallelTester()
	s.ProcessOutput()
}

func main() {
	s := &SysInfo{}
	err := s.init(os.Args)
	if err != nil {
		panic(err)
	}
	s.Run()
}
