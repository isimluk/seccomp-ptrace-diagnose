#!/bin/bash

set -x

strace ls

getpcaps $$
