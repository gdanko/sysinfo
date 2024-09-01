package gather

import (
	"github.com/gdanko/sysinfo/stats"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetCpuPercent(c chan func() ([]stats.PercentStat, error)) {
	c <- (func() ([]stats.PercentStat, error) {
		cpuPercent, err := stats.GetCpuPercent(false)
		return cpuPercent, err
	})
}

func GetDiskUsage(c chan func() ([]stats.DiskUsageData, error)) {
	c <- (func() ([]stats.DiskUsageData, error) {
		diskUsage, err := stats.GetDiskUsage()
		return diskUsage, err
	})
}

func GetHostInformation(c chan func() (stats.HostInformation, error)) {
	c <- (func() (stats.HostInformation, error) {
		hostInformation, err := stats.GetHostInformation()
		return hostInformation, err
	})
}

func GetLoadAverages(c chan func() (*load.AvgStat, error)) {
	c <- (func() (*load.AvgStat, error) {
		loadAverages, err := stats.GetLoadAverages()
		return loadAverages, err
	})
}

func GetMemoryUsage(c chan func() (*mem.VirtualMemoryStat, error)) {
	c <- (func() (*mem.VirtualMemoryStat, error) {
		memoryUsage, err := stats.GetMemoryUsage()
		return memoryUsage, err
	})
}

func GetNetworkData(c chan func(filterData, interfaceData, iostatData, protocolData bool) (stats.NetworkData, error)) {
	c <- (func(filterData, interfaceData, iostatData, protocolData bool) (stats.NetworkData, error) {
		networkData, err := stats.GetNetworkData(filterData, interfaceData, iostatData, protocolData)
		return networkData, err
	})
}

func GetSwapUsage(c chan func() (*mem.SwapMemoryStat, error)) {
	c <- (func() (*mem.SwapMemoryStat, error) {
		swapUsage, err := stats.GetSwapUsage()
		return swapUsage, err
	})
}

func GetProcessList(c chan func() ([]stats.Process, error)) {
	c <- (func() ([]stats.Process, error) {
		processList, err := stats.GetProcessList()
		return processList, err
	})
}
