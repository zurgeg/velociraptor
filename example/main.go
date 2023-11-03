package main

/*
#include <stdint.h>
*/
import "C"
import (
	"fmt"

	"github.com/zurgeg/velociraptor"
)

//export VProgEvent
func VProgEvent(playdate velociraptor.PlaydateAPI, event velociraptor.PDSystemEvent, arg C.uint32_t) int {

	fmt.Println("Program initialized!")
	return 0
}
