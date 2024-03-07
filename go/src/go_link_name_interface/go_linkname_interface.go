package go_link_name_interface

import (
	_ "unsafe"

	_ "lab/src/go_link_name_src"
)

//go:linkname golinknameinterface lab/src/go_link_name_src.golinknamesrc
func golinknameinterface()

func GolinknameinterfaceTest() {
	golinknameinterface()
}
