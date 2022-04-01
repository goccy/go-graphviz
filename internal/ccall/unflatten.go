package ccall

/*
#cgo CFLAGS: -DGVLIBDIR=graphviz
#cgo CFLAGS: -Icdt
#cgo CFLAGS: -Icommon
#cgo CFLAGS: -Igvc
#cgo CFLAGS: -Ipathplan
#cgo CFLAGS: -Icgraph
#cgo CFLAGS: -Iunflatten
#cgo CFLAGS: -Ifdpgen
#cgo CFLAGS: -Isfdpgen
#cgo CFLAGS: -Ixdot
#cgo CFLAGS: -Ilabel
#cgo CFLAGS: -Ipack
#cgo CFLAGS: -Iortho
#cgo CFLAGS: -Iosage
#cgo CFLAGS: -Ineatogen
#cgo CFLAGS: -Isparse
#cgo CFLAGS: -Icircogen
#cgo CFLAGS: -Irbtree
#cgo CFLAGS: -Ipatchwork
#cgo CFLAGS: -Itwopigen
#cgo CFLAGS: -I../
#cgo CFLAGS: -I../libltdl
#cgo CFLAGS: -Wno-unused-result -Wno-format -Wno-pointer-to-int-cast -Wno-attributes
#include "unflatten.h"
*/
import "C"

func Transform(g *Agraph) {
	C.transform(g.c)
}
