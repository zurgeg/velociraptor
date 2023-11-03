package velociraptor

/*
#define TARGET_EXTENSION
#include <pd_api.h>
#undef TARGET_EXTENSION
*/
import "C"

type PlaydateAPI = *C.PlaydateAPI
type PDSystemEvent = C.PDSystemEvent
