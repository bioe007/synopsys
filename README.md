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

Ideally this would be doable from a single command that doesn't need to
be pre-installed so either a script or self-contained binary.

## Current state
This isn't really useable, but it does output:

  - running / total processes, number of vcores, current cpu frequency
  - load average
  - cpus - can show top N cpus sorted by user time
  - memory - free/total, buff/cache
  - disks - total requests, written/read KB
  - uptime

there are also some cli options to limit the number of CPU and disks shown, to
show disks only, etc..

# Notes
This is just thinking out-loud stuff..

Because any scripted language (e.g. python, lua) doing this would
require either a *ton* of boilerplate or a long list of dependencies
that might not be available to a sysadmin during a critical outage I'm
favoring something that's easy to build/store/distribute as a single
binary so even if it's not resident on a struggling system a simple the
tool is only a simple `scp synopsys` away


## Critical Parts
These are the parameters I'd like to show

- *Uptime*: express as hours:min:sec  (done)
- *load average* (done)
  - *todo* Enable mode where R/num_vc and B/num_disks ?

- *CPU:* cores, overall % useage, then % sys, usr, guest, ... (done)
  - *TODO* Have a mode that shows top 'any%' so if/when CPU are above a
  threhshold in any usage category they're shown. So if a system has some high
  guest% cores and some high system% cores all can be shown.
  - by default, show the overall cpu utilization (done)
  -  _wonders_ any way to make mpstat type of info here?

- *Processes:*
  - run|able, sleep, unint sleep, zombies
  - *TODO*

- *'Errors'* from dmesg and ~ dmesg | tail (or journalctl -b | tail)
  - *todo*

- *Memory:* free/used (proc/meminfo)  (done)
- *Swap:* free/used (proc/meminfo)
  - *todo*

- *Disk activity*: rw/wr in MBs and queue size
    - /proc/diskstats  - (done'ish)
    - /proc/partitions - (done)

- *Network* In/Out (per device?) - *TODO*
    - connections - active, passive, trans/retrans stats
    - top 'few' processes consuming CPU | memory

## Display

Initially this will just output some rolling format. will have to think
about something like a tui to properly place things for readability
though

## Random thoughts
Is there a faster way to fetch all this data than reading a text file each time?

lots of repetition with a ThingStats that has 'previous', 'current' and does
math on all the fields. I could have something like a generic `Delta` struct
that accepts the same type twice, does the math between all fields etc.. But
this would require making everything public in the structs.

## References

- [/proc](https://www.man7.org/linux/man-pages/man5/proc.5.html)
- [/proc/loadavg](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/4/html/reference_guide/s2-proc-loadavg)
- [filesystems](https://www.kernel.org/doc/html/latest/filesystems/index.html)
- [/sys/block](https://www.kernel.org/doc/html/latest/block/stat.html)
  - [/proc/diskstats](https://www.kernel.org/doc/html/latest/admin-guide/iostats.html)
