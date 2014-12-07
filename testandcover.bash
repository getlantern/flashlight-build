#!/bin/bash

# This script tests multiple packages and creates a consolidated cover profile
# See https://gist.github.com/hailiang/0f22736320abe6be71ce for inspiration.
# The list of packages to test is specified in testpackages.txt.

function die() {
  echo $*
  exit 1
}

export GOPATH=`pwd`:$GOPATH

# Remove profile.cov from previous run if necessary
if [ -e profile.cov ]
then
    rm profile.cov
fi

# Test each package and append coverage profile info to profile.cov
for pkg in `cat testpackages.txt`
do
    #$HOME/gopath/bin/
    go test -v -covermode=count -coverprofile=profile_tmp.cov $pkg || die "Error testing $pkg"
    tail -n +2 profile_tmp.cov >> profile.cov || die "Error appending coverage for $pkg to profile.cov"
done

#- GOPATH=`pwd`:$GOPATH $HOME/gopath/bin/goveralls -v -service travis-ci github.com/getlantern/buuid