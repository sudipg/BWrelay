

package main

import (
	"fmt"
	"time"
	"strconv"
	bw "gopkg.in/immesys/bw2bind.v5"
)

func main() {
	//Connect
	cl := bw.ConnectOrExit("")
	cl.SetEntityFromEnvironOrExit()

  uri := "scratch.ns/relay/"
	svc := cl.RegisterService(uri, "s.relayState")

	//This sets a metadata key on the service
	svc.SetMetadata("relayCtrlApp", "set relay state")

	//You can have multiple interfaces per service. The second parameter
	//is the interface type, the first is the name of that instance of the
	//interface. We only have one interface, so underscore is a placeholder
	iface := svc.RegisterInterface("ctrlRelay", "i.echo")

	// assume temperature is always in 'F
	var relay1 int
  var relay2 int
  relay1 = 0
  relay2 = 0

	//People can set what the message should be
	iface.SubscribeSlot("ctrl1", func(m *bw.SimpleMessage) {
		if newmsg := m.GetOnePODF(bw.PODFString); newmsg != nil {
			fmt.Println("got a new state command for relay 1")
      command, err := strconv.Atoi(newmsg.(bw.TextPayloadObject).Value())
			if err == nil {
				relay1 = command
				if command != 0 {
	        fmt.Println("got a command: turn relay 1 on")
	      } else {
	        fmt.Println("got a command: turn relay 1 off")
	      }
			} else {
				fmt.Println("ERROR while parsing command")
			}
		}
	})

	//Also, every five seconds, we publish the message
	for {
    fmt.Println("current relay1 state is ", relay1)
		fmt.Println("current relay2 state is ", relay2)
		po := bw.CreateTextPayloadObject(bw.PONumString, strconv.Itoa(relay1))
		err := iface.PublishSignal("current ", po)
		fmt.Println("Published, error was ", err)
		time.Sleep(1 * time.Second)
	}
}
