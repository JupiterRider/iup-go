//go:build ((aix || dragonfly || freebsd || linux || netbsd || openbsd || solaris || illumos) && !motif) || gtk

package iup

/*
#include "external/src/gtk/iupgtk_text.c"
*/
import "C"
