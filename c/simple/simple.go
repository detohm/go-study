package main

// #include<stdio.h>
// #include<stdlib.h>
// void printInC(){
//   printf("This is in C code.\n");
// }
//
// int add(int a, int b){
//   return a+b;
// }
//
// char* readString(char *str){
//   printf("Received string from Go: %s\n", str);
//   char *ret = "string from C";
//   return ret;
// }
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// --- basic function call ---
	fmt.Println("This is in Go code.")
	C.printInC()

	// --- passing value arguments and get return value
	result := C.add(C.int(1), C.int(2))
	fmt.Printf("return: %d\n", result)

	// --- passing string pointer and get return pointer
	cstr := C.CString("string from Go")
	defer C.free(unsafe.Pointer(cstr))
	cstrOut := C.readString(cstr)
	goStr := C.GoString(cstrOut)
	fmt.Printf("Received from C: %s\n", goStr)
}

/* output
This is in Go code.
This is in C code.
return: 3
Received string from Go: string from Go
Received from C: string from C
*/
