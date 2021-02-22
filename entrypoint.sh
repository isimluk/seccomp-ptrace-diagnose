#!/bin/bash

set -x

uname -a

strace --seccomp-bpf ls /dev/null

getpcaps $$

diagnose
