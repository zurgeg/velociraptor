// Loads Velociraptor and contains all of the init funcs and whatnot

#include <pd_api.h>
#include <build/velociraptor_program.h>

/// @brief Handles events and passes them to relevant go funcs
/// @param playdate The playdate API
/// @param event The event
/// @param arg An argument passed by said event
/// @return 0

int eventHandler(PlaydateAPI* playdate, PDSystemEvent event, uint32_t arg){
    VProgEvent(playdate, event, arg);
}