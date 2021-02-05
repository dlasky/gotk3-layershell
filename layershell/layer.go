package layershell

// #cgo pkg-config: gtk+-3.0 gtk-layer-shell-0
// #include <gtk/gtk.h>
// #include "gtk-layer-shell.h"
import "C"
import (
	"github.com/gotk3/gotk3/gtk"
	"unsafe"
)

type LayerShellEdgeFlags int

const (
	LAYER_SHELL_EDGE_LEFT   LayerShellEdgeFlags = iota
	LAYER_SHELL_EDGE_RIGHT
	LAYER_SHELL_EDGE_TOP
	LAYER_SHELL_EDGE_BOTTOM
)

type LayerShellLayerFlags int

const (
	LAYER_SHELL_LAYER_BACKGROUND LayerShellLayerFlags = iota
	LAYER_SHELL_LAYER_BOTTOM
	LAYER_SHELL_LAYER_TOP
	LAYER_SHELL_LAYER_OVERLAY    
)

func native(window *gtk.Window) *C.GtkWindow {
	w := window.Native()
	wp := (*C.GtkWindow)(unsafe.Pointer(w))
	return wp;
}

func boolConv(b bool) C.int {
	if (b) { return C.int(1)}
	return C.int(0)
}

func InitForWindow(window *gtk.Window) {
	w := native(window);
	C.gtk_layer_init_for_window(w)
}

func SetLayer(window *gtk.Window, layer LayerShellLayerFlags) {
	w := native(window)
	C.gtk_layer_set_layer(w, C.GtkLayerShellLayer(layer))
}

func AutoExclusiveZoneEnable(window *gtk.Window) {
	w := native(window)
	C.gtk_layer_auto_exclusive_zone_enable(w)
}

func SetAnchor(window *gtk.Window, side LayerShellEdgeFlags, pinned bool) {
	w := native(window)
	C.gtk_layer_set_anchor(w, C.GtkLayerShellEdge(side), boolConv(pinned))
}

func SetMargin(window *gtk.Window, side LayerShellEdgeFlags, margin int) {
	w := native(window)
	C.gtk_layer_set_margin(w, C.GtkLayerShellEdge(side), C.int(margin))
}
