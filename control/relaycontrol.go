// Simple test application to change the echo message on the bw2 example hellosvc

package main

import (
  "fmt"
  //"github.com/immesys/spawnpoint/spawnable"
  bw "gopkg.in/immesys/bw2bind.v5"
)

func main() {
  // connect
  cl := bw.ConnectOrExit("")
  cl.SetEntityFromEnvironOrExit()
  uri := "scratch.ns/relay/s.relayState/ctrlRelay/i.echo/slot/ctrl1"
	fmt.Println("Enter your desired state for relay 1")
	var newState1 string
	fmt.Scanf("%s",&newState1)
  po := []bw.PayloadObject{bw.CreateStringPayloadObject(newState1),}
  fmt.Println("made a po, ready to publish")
  err :=  cl.Publish(&bw.PublishParams{
                                  URI : uri,
                            			AutoChain : true,
                            			PayloadObjects : po,
                            	    })
  fmt.Println("published, err was", err)
}
