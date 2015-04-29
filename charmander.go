package main

// #include "fucking.h"
// #include "ducking.h"
// #include "em35x/stack/include/message.h"
// #include "em35x/app/builder/myLight/client-command-macro.h"
import "C"
import "fmt"
import "builtin"

func main() {
//  fmt.Printf("Saw %d\n", C.jesus)
//  fmt.Printf("Saw %d\n", C.moses)
//  fmt.Printf("Saw %s\n", C.hahaha)
  fmt.Printf("Saw %d\n", C.myfoo())
  fmt.Printf("Saw %d\n", C.asf())
  light_address = 0xdedf

  // Make sure the ZCL command is ready
  C.emberAfFillCommandOnOffClusterToggle() // message.h
  
  // Might have to edit the source and destination point
  
  // Send out the signal
  C.emberAfSendCommandUnicast(C.EMBER_OUTGOING_DIRECT, light_address);

//    fmt.Printf("Saw %d\n", C.random())
}
