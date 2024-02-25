synopsys - A SYNOPsis of your SYStem

The goal is to create an overview of a system. Instead of running:

- uptime
- dmesg | tail
- free
- vmstat
- mpstat
- iostat
- sar DEV, TCP, ETCP
- top

Show the _critical_ parts from each of those in legible fashion with a
single command.

Some might say this breaks the one-tool philosophy but i disagree, this
isn't intended to 

Ideally this would be doable from a single command that doesn't need to
be pre-installed so either a script or self-contained binary. 

Because any scripted language (e.g. python, lua) doing this would
require either a *ton* of boilerplate or a long list of dependencies
that might not be available to a sysadmin during a critical outage I'm
favoring something that's easy to build/store/distribute as a single
binary so even if it's not resident on a struggling system a simple the
tool is only a simple `scp synopsys` away


## Critical Parts
These are the parameters I'd like to show 

- Uptime 
- CPU: cores, overall % useage, then % sys, usr, guest, ...
-  _wonders_ any way to make mpstat type of info here?
  - by default, show the overall cpu utilization
  - nice to have - show hot cores with 80% utilization over x sec
- Number of processes: run|able, sleep, unint sleep, zombies
- 'Errors' from dmesg and ~ dmesg | tail (or journalctl -b | tail)
- Memory: free/used (proc/meminfo)
- Swap: free/used (proc/meminfo)
- Disk activity: rw/wr in MBs and queue size
  - /proc/diskstats
  - /proc/partitions
  - [/proc/diskstats](https://www.kernel.org/doc/html/latest/admin-guide/iostats.html)
- Network In/Out (per device?)
- connections - active, passive, trans/retrans stats
- top 'few' processes consuming CPU | memory

## Display 

initially this will just output some rolling format. will have to think
about something like a tui to properly place things for readability
though


## References

- [/proc/loadavg](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/4/html/reference_guide/s2-proc-loadavg)


