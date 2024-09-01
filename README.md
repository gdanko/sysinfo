# Free
sysinfo is small utility that displays basic information about your system.

## Installation and upgrade
Free is delivered via [Homebrew](https://brew.sh/). To install free, follow those steps:

`brew tap intuit/gdanko https://github.intuit.com/gdanko/homebrew`

`brew update`

`brew install sysinfo`

To upgrade free, follow these steps:

`brew update`

`brew upgrade sysinfo`

## Free Runtime Options
```
sysinfo --help
Usage:
  sysinfo [OPTIONS]

Application Options:
  -a, --all      Display all system info.
  -c, --cpu      Display system CPU usage.
  -d, --disk     Display system disk usage.
      --host     Display host statistics.
  -l, --load     Display system load averages.
  -m, --memory   Display system memory usage.
  -s, --swap     Display swap memory usage.
  -V, --version  Output version information and exit.

CPU Options:
  -p, --per-cpu  Show information per-CPU.

Disk Options:
  -i, --inode    Show disk inode information.
  -u, --usage    Show disk usage information.

Help Options:
  -h, --help     Show this help message
  ```

## Need Help?
You can find me on Slack [@gdanko](https://intuit-slack-connect.slack.com/app_redirect?channel=W8GN53CA3) or email <Gary_Danko@intuit.com>
