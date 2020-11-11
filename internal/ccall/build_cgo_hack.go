// +build required

package ccall

// This file exists purely to prevent the golang toolchain from stripping
// away the c source directories and files when `go mod vendor` is used
// to populate a `vendor/` directory of a project depending on `goccy/go-graphviz`.
//
// How it works:
//  - every directory which only includes c source files receives a dummy.go file.
//  - every directory we want to preserve is included here as a _ import.
//  - this file is given a build to exclude it from the regular build.
import (
	// Prevent go tooling from stripping out the c source files.
	_ "github.com/goccy/go-graphviz/internal"
	_ "github.com/goccy/go-graphviz/internal/ccall/ast"
	_ "github.com/goccy/go-graphviz/internal/ccall/cdt"
	_ "github.com/goccy/go-graphviz/internal/ccall/cgraph"
	_ "github.com/goccy/go-graphviz/internal/ccall/circogen"
	_ "github.com/goccy/go-graphviz/internal/ccall/common"
	_ "github.com/goccy/go-graphviz/internal/ccall/dotgen"
	_ "github.com/goccy/go-graphviz/internal/ccall/edgepaint"
	_ "github.com/goccy/go-graphviz/internal/ccall/expr"
	_ "github.com/goccy/go-graphviz/internal/ccall/fdpgen"
	_ "github.com/goccy/go-graphviz/internal/ccall/glcomp"
	_ "github.com/goccy/go-graphviz/internal/ccall/gvc"
	_ "github.com/goccy/go-graphviz/internal/ccall/gvpr"
	_ "github.com/goccy/go-graphviz/internal/ccall/ingraphs"
	_ "github.com/goccy/go-graphviz/internal/ccall/label"
	_ "github.com/goccy/go-graphviz/internal/ccall/mingle"
	_ "github.com/goccy/go-graphviz/internal/ccall/neatogen"
	_ "github.com/goccy/go-graphviz/internal/ccall/ortho"
	_ "github.com/goccy/go-graphviz/internal/ccall/osage"
	_ "github.com/goccy/go-graphviz/internal/ccall/pack"
	_ "github.com/goccy/go-graphviz/internal/ccall/patchwork"
	_ "github.com/goccy/go-graphviz/internal/ccall/pathplan"
	_ "github.com/goccy/go-graphviz/internal/ccall/rbtree"
	_ "github.com/goccy/go-graphviz/internal/ccall/sfdpgen"
	_ "github.com/goccy/go-graphviz/internal/ccall/sfio"
	_ "github.com/goccy/go-graphviz/internal/ccall/sfio/Sfio_f"
	_ "github.com/goccy/go-graphviz/internal/ccall/sparse"
	_ "github.com/goccy/go-graphviz/internal/ccall/spine"
	_ "github.com/goccy/go-graphviz/internal/ccall/topfish"
	_ "github.com/goccy/go-graphviz/internal/ccall/twopigen"
	_ "github.com/goccy/go-graphviz/internal/ccall/vmalloc"
	_ "github.com/goccy/go-graphviz/internal/ccall/vpsc"
	_ "github.com/goccy/go-graphviz/internal/ccall/vpsc/pairingheap"
	_ "github.com/goccy/go-graphviz/internal/ccall/xdot"
	_ "github.com/goccy/go-graphviz/internal/expat"
	_ "github.com/goccy/go-graphviz/internal/plugin/core"
	_ "github.com/goccy/go-graphviz/internal/plugin/devil"
	_ "github.com/goccy/go-graphviz/internal/plugin/dot_layout"
	_ "github.com/goccy/go-graphviz/internal/plugin/gd"
	_ "github.com/goccy/go-graphviz/internal/plugin/gdiplus"
	_ "github.com/goccy/go-graphviz/internal/plugin/gdk"
	_ "github.com/goccy/go-graphviz/internal/plugin/glitz"
	_ "github.com/goccy/go-graphviz/internal/plugin/gs"
	_ "github.com/goccy/go-graphviz/internal/plugin/gtk"
	_ "github.com/goccy/go-graphviz/internal/plugin/lasi"
	_ "github.com/goccy/go-graphviz/internal/plugin/ming"
	_ "github.com/goccy/go-graphviz/internal/plugin/neato_layout"
	_ "github.com/goccy/go-graphviz/internal/plugin/pango"
	_ "github.com/goccy/go-graphviz/internal/plugin/poppler"
	_ "github.com/goccy/go-graphviz/internal/plugin/quartz"
	_ "github.com/goccy/go-graphviz/internal/plugin/rsvg"
	_ "github.com/goccy/go-graphviz/internal/plugin/visio"
	_ "github.com/goccy/go-graphviz/internal/plugin/webp"
	_ "github.com/goccy/go-graphviz/internal/plugin/xlib"
)
