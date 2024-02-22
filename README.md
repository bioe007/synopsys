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

Show the critical parts from each of those in legible fashion with a
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
