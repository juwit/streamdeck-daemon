package libxdo

/*
#include <stdlib.h>
#include <xdo.h>
#cgo LDFLAGS: -lxdo
*/
import "C"

const CURRENTWINDOW = 0

func TypeKeys(text string) {
	xdo := C.xdo_new(nil)

	C.xdo_enter_text_window (xdo, C.Window(CURRENTWINDOW), C.CString(text), C.useconds_t(12000))
}

