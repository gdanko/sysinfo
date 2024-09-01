package stats

import (
	"github.com/shirou/gopsutil/v3/process"
)

type Process struct {
	Command         string                      `json:"command"`
	CommandArgs     []string                    `json:"command_args"`
	ContextSwitches *process.NumCtxSwitchesStat `json:"context_switches"`
	CPUAffinity     []int32                     `json:"cpu_affinity"`
	CPUPercent      float64                     `json:"cpu"`
	CreationTime    int64                       `json:"creation_time"`
	CWD             string                      `json:"cwd"`
	Environment     []string                    `json:"environment"`
	GroupIDs        []int32                     `json:"groups_ids"`
	MemoryPercent   float32                     `json:"memory_percent"`
	Name            string                      `json:"name"`
	Nice            int32                       `json:"nice_level"`
	OpenFiles       int32                       `json:"open_file_count"`
	ParentPID       int32                       `json:"ppid"`
	PID             int32                       `json:"pid"`
	Terminal        string                      `json:"terminal"`
	ThreadCount     int32                       `json:"thread_count"`
	UserIDs         []int32                     `json:"user_ids"`
	Username        string                      `json:"username"`
}

func GetProcessList() (processes []Process, err error) {
	procs, err := process.Processes()
	if err != nil {
		return []Process{}, err
	}
	for _, proc := range procs {
		name, _ := proc.Name()
		// children, _ := proc.Children()
		cmdArgs, _ := proc.CmdlineSlice()
		contextSwitches, _ := proc.NumCtxSwitches()
		cpuAffinity, _ := proc.CPUAffinity()
		cpuPercent, _ := proc.CPUPercent()
		createTime, _ := proc.CreateTime()
		cwd, _ := proc.Cwd()
		env, _ := proc.Environ()
		exe, _ := proc.Exe()
		groupIDs, _ := proc.Gids()
		memoryPercent, _ := proc.MemoryPercent()
		nice, _ := proc.Nice()
		openFiles, _ := proc.OpenFiles()
		ppid, _ := proc.Ppid()
		terminal, _ := proc.Terminal()
		threadCount, _ := proc.NumThreads()
		uids, _ := proc.Uids()
		username, _ := proc.Username()

		if len(cmdArgs) < 2 {
			cmdArgs = []string{}
		} else if len(cmdArgs) >= 1 {
			cmdArgs = cmdArgs[1:]
		}

		procObject := Process{
			Command:         exe,
			CommandArgs:     cmdArgs,
			ContextSwitches: contextSwitches,
			CPUAffinity:     cpuAffinity,
			CPUPercent:      cpuPercent,
			CreationTime:    createTime,
			CWD:             cwd,
			Environment:     env,
			GroupIDs:        groupIDs,
			MemoryPercent:   memoryPercent,
			Name:            name,
			Nice:            nice,
			OpenFiles:       int32(len(openFiles)),
			ParentPID:       ppid,
			PID:             proc.Pid,
			Terminal:        terminal,
			ThreadCount:     threadCount,
			UserIDs:         uids,
			Username:        username,
		}
		processes = append(processes, procObject)
		// if len(opts.ProcessOptions.ByName) > 0 {
		// 	byName = true
		// }
		// if len(opts.ProcessOptions.ByPid) > 0 {
		// 	byPid = true
		// }
		// if len(opts.ProcessOptions.ByUser) > 0 {
		// 	byUser = true
		// }

		// if byName {
		// 	for _, element := range opts.ProcessOptions.ByName {
		// 		if strings.Contains(procObject.Name, element) {
		// 			allProcs = append(allProcs, procObject)
		// 		}
		// 	}
		// } else if byPid {
		// 	for _, element := range opts.ProcessOptions.ByPid {
		// 		if procObject.PID == element {
		// 			allProcs = append(allProcs, procObject)
		// 		}
		// 	}
		// } else if byUser {
		// 	for _, element := range opts.ProcessOptions.ByUser {
		// 		if procObject.Username == element {
		// 			allProcs = append(allProcs, procObject)
		// 		}
		// 	}
		// } else {
		// }
	}
	return processes, nil
}
