# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=waveapi
GOFILES=\
    blip.go\
    debug.go\
    event.go\
    ops.go\
    robot.go\
    wavelet.go\

include $(GOROOT)/src/Make.pkg
