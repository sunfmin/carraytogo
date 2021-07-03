package main

// #cgo LDFLAGS: -lstdc++ ${SRCDIR}/../build/libhello.dylib
// #include <stdint.h>
// void fillArray1(const uint8_t *ret);
// void fillArray2(const uint8_t **ret);
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	var ret1 = make([]C.uint8_t, 10)
	// var ret = ([]*C.uint8_t)(C.malloc(C.uint8_t * 32 * 100))
	// for i, _ := range ret {
		// uint8_t* array = malloc( sizeof( double ) * 100 );
		// e := [2]C.uint8_t{}
		// ret[i] = e
	// }
	C.fillArray1(&ret1[0])
	fmt.Println("fillArray1", ret1)

	var ret2 = make([]*C.uint8_t, 10)
	C.fillArray2(&ret2[0])
	fmt.Println("fillArray2", ret2)
	fmt.Println("fillArray2", C.GoBytes(unsafe.Pointer(ret2[0]), 3))

}
