package main

/*
#include <httpd.h>
*/
import "C"

import (
	"github.com/garrettsickles/goapache"
)

//export handler
func handler(rec *C.request_rec) C.int {
	return 0
}
