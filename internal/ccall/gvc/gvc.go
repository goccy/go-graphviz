package gvc

/*
#cgo CFLAGS: -I../common -I../pathplan -I../cgraph -I../cdt
#include "gvc.h"
*/
import "C"

func gvToggle(i int) {
	C.gvToggle(i)
}
