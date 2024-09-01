package gather

import "github.com/gdanko/sysinfo/stats"

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
