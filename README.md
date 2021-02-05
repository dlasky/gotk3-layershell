# gotk3-layershell

A simple golang library to provide bindings for the excellent [Gtk Layer Shell](https://github.com/wmww/gtk-layer-shell) library which can be consumed in the also excellent [gotk3 gtk library](https://github.com/gotk3/gotk3).

This allows for GTK windows in linux window managers like swaywm that utilize the Layer Shell proticol in wayland to be positioned relative to the viewport including pinning and layer depth control.

## Getting started

For the moment please see `example.go` for a simple working example

## WIP

While the critical API components are ported (layer, anchors), currently the entirety `gtk_layer_shell.h` is not exported. I will be adding support over the next few months (2/5/2021).