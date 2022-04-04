package ccall

/*
#cgo CFLAGS: -Iunflatten
#include "unflatten.h"
*/
import "C"

func Transform(g *Agraph, maxMinlen int, chainLimit int, doFans bool) {
	var doFansValue int
	if doFans {
		doFansValue = 1
	} else {
		doFansValue = 0
	}
	C.transform(g.c, C.int(maxMinlen), C.int(chainLimit), C.int(doFansValue))
}
