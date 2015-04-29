package main

// #include "em35x/stack/include/message.h"
// #include "em35x/app/builder/myLight/client-command-macro.h"
import "C"
import "fmt"
import "builtin"

func main() {
  light_address = 0xdedf

  // Make sure the ZCL command is ready
  C.emberAfFillCommandOnOffClusterToggle() // message.h
  
  // Might have to edit the source and destination point
  
  // Send out the signal
  C.emberAfSendCommandUnicast(C.EMBER_OUTGOING_DIRECT, light_address);
}
