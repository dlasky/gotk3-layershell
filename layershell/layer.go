package layershell

// #cgo pkg-config: gtk+-3.0 gtk-layer-shell-0
// #include <gtk/gtk.h>
// #include "gtk-layer-shell.h"
import "C"
import (
	"github.com/gotk3/gotk3/gdk"
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

type LayerShellKeyboardMode int

const (
	LAYER_SHELL_KEYBOARD_MODE_NONE LayerShellKeyboardMode= iota
	LAYER_SHELL_KEYBOARD_MODE_EXCLUSIVE
	LAYER_SHELL_KEYBOARD_MODE_ON_DEMAND
	LAYER_SHELL_KEYBOARD_MODE_ENTRY_NUMBER
)

func nativeWindow(window *gtk.Window) *C.GtkWindow {
	w := window.Native()
	wp := (*C.GtkWindow)(unsafe.Pointer(w))
	return wp;
}

func nativeMonitor(monitor *gdk.Monitor) *C.GdkMonitor {
	m := monitor.Native()
	mp := (*C.GdkMonitor)(unsafe.Pointer(m))
	return mp;
}

func boolConv(b bool) C.int {
	if (b) { return C.int(1)}
	return C.int(0)
}

func InitForWindow(window *gtk.Window) {
	w := nativeWindow(window);
	C.gtk_layer_init_for_window(w)
}

func SetLayer(window *gtk.Window, layer LayerShellLayerFlags) {
	w := nativeWindow(window)
	C.gtk_layer_set_layer(w, C.GtkLayerShellLayer(layer))
}

func AutoExclusiveZoneEnable(window *gtk.Window) {
	w := nativeWindow(window)
	C.gtk_layer_auto_exclusive_zone_enable(w)
}

func SetExclusiveZone(window *gtk.Window, size int) {
	w := nativeWindow(window)
	C.gtk_layer_set_exclusive_zone(w, C.int(size))
}

func SetAnchor(window *gtk.Window, side LayerShellEdgeFlags, pinned bool) {
	w := nativeWindow(window)
	C.gtk_layer_set_anchor(w, C.GtkLayerShellEdge(side), boolConv(pinned))
}

func SetMargin(window *gtk.Window, side LayerShellEdgeFlags, margin int) {
	w := nativeWindow(window)
	C.gtk_layer_set_margin(w, C.GtkLayerShellEdge(side), C.int(margin))
}

func SetMonitor(window *gtk.Window, monitor *gdk.Monitor) {
	C.gtk_layer_set_monitor(nativeWindow(window), nativeMonitor(monitor))
}

func SetKeyboardMode(window *gtk.Window, mode LayerShellKeyboardMode) {
	w := nativeWindow(window)
	C.gtk_layer_set_keyboard_mode(w, C.GtkLayerShellKeyboardMode(mode))
}
