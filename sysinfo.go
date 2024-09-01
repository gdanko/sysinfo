package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
	"github.intuit.com/gdanko/sysinfo/gather"
	"github.intuit.com/gdanko/sysinfo/globals"
	"github.intuit.com/gdanko/sysinfo/internal"
	"github.intuit.com/gdanko/sysinfo/stats"
	"github.intuit.com/gdanko/sysinfo/util"
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
