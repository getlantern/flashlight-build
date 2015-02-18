#!/bin/bash

###############################################################################
#
# This script regenerates the source file that embeds systray.dll
#
###############################################################################

function die() {
  echo $*
  exit 1
}
which 2goarray > /dev/null || go get github.com/cratonica/2goarray/...
which 2goarray > /dev/null || die "Please install 2goarray manually, then try again"

2goarray systraydll systray < systray.dll > systraydll_windows.go