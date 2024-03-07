package go_link_name_src

import (
	"fmt"
	_ "unsafe"
)

//go:linkname golinknamesrc lab/src/go_link_name_interface.GolinknameinterfaceTest
func golinknamesrc() {
	fmt.Println(111111)
}
