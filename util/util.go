package util

import (
	"fmt"
	"math"
	"os"

	"github.com/gdanko/sysinfo/globals"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// ConfigureLogger : Configure the logger
func ConfigureLogger(logLevel logrus.Level, nocolorFlag bool) (logger *logrus.Logger) {
	disableColors := false
	if nocolorFlag {
		disableColors = true
	}
	logger = &logrus.Logger{
		Out:   os.Stderr,
		Level: logLevel,
		Formatter: &prefixed.TextFormatter{
			DisableColors:    disableColors,
			DisableTimestamp: true,
			TimestampFormat:  "2006-01-02 15:04:05",
			FullTimestamp:    true,
			ForceFormatting:  false,
		},
	}
	logger.SetLevel(logLevel)

	return logger
}

func ValidateOpts(opts *globals.Options) (err error) {
	if !opts.All && !opts.CPU && !opts.Disk && !opts.Host && !opts.Load && !opts.Memory && !opts.Swap && !opts.Process && !opts.Net {
		opts.All = true
	}

	if len(opts.ProcessOptions.ByName) > 0 {
		if len(opts.ProcessOptions.ByUser) > 0 || len(opts.ProcessOptions.ByPid) > 0 {
			return fmt.Errorf("--by-name cannot be used with: --by-pid, --by-user")
		}
	} else if len(opts.ProcessOptions.ByPid) > 0 {
		if len(opts.ProcessOptions.ByName) > 0 || len(opts.ProcessOptions.ByUser) > 0 {
			return fmt.Errorf("--by-pid cannot be used with: --by-name, --by-user")
		}
	} else if len(opts.ProcessOptions.ByUser) > 0 {
		if len(opts.ProcessOptions.ByName) > 0 || len(opts.ProcessOptions.ByPid) > 0 {
			return fmt.Errorf("--by-name cannot be used with: --by-pid, --by-user")
		}
	}

	if opts.All {
		opts.CPU = true
		opts.CpuOptions.PerCPU = false
		opts.Disk = true
		opts.DiskOptions.Inode = true
		opts.DiskOptions.Usage = true
		opts.Host = true
		opts.HostOptions.HostTemps = true
		opts.HostOptions.HostUsers = true
		opts.Load = true
		opts.Memory = true
		opts.Swap = true
		opts.Process = true
		opts.Net = true
		opts.NetworkOptions.FilterCounters = true
		opts.NetworkOptions.IOStats = true
		opts.NetworkOptions.Interfaces = true
		opts.NetworkOptions.ProtocolCounters = true
	}
	return nil
}

func RoundTo(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}
