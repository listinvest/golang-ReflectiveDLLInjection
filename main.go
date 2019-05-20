package main

import (
	"fmt"

	"unsafe"

	"github.com/zetamatta/go-outputdebug"
	
)


import "C"


// OnProcessAttach is an async callback (hook).
//export OnProcessAttach
func OnProcessAttach(
	hinstDLL unsafe.Pointer, // handle to DLL module
	fdwReason uint32, // reason for calling function
	lpReserved unsafe.Pointer, // reserved
) {
	
	outputdebug.String("OnProcessAttach(): Begin")
}

func main() {
	// nothing really. xD
	fmt.Printf("test");
}
