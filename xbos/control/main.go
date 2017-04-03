// Simple test application to change the echo message on the bw2 example hellosvc

package main

import (
  "fmt"
  //"github.com/immesys/spawnpoint/spawnable"
  bw "gopkg.in/immesys/bw2bind.v5"
  "strconv"
)

const (
    PONUM = "2.1.1.2"
)

func main() {
  // connect
  cl := bw.ConnectOrExit("")
  cl.SetEntityFromEnvironOrExit()
  uri := "scratch.ns/relay/s.vplug/vplug/i.xbos.plug/slot/state"
  b, err := strconv.ParseBool("true")
  po, err := bw.CreateMsgPackPayloadObject(bw.FromDotForm(PONUM), map[string]interface{} {"state":b})
  fmt.Println("made a po, ready to publish")
  po.SetPONum(bw.FromDotForm(PONUM))
  pos := []bw.PayloadObject{po,}
  cl.Publish(&bw.PublishParams{
    URI : uri,
    AutoChain : true,
    PayloadObjects : pos,
    })
  fmt.Println("published, err was", err)
}
