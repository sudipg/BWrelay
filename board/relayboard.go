package main

import (
    "fmt"
    "time"
    "strconv"
    bw "gopkg.in/immesys/bw2bind.v5"
    "github.com/kidoman/embd"
    _ "github.com/kidoman/embd/host/rpi" // This loads the RPi driver
)

func main() {
    //Connect
    cl := bw.ConnectOrExit("")
    cl.SetEntityFromEnvironOrExit()

    uri := "scratch.ns/relay/"
    svc := cl.RegisterService(uri, "s.relayState")

    embd.InitGPIO()
    defer embd.CloseGPIO()

    relay1pin := 16
    relay2pin := 18

    embd.SetDirection(relay1pin, embd.Out)
    embd.SetDirection(relay2pin, embd.Out)
    embd.DigitalWrite(relay1pin, embd.Low)
    embd.DigitalWrite(relay2pin, embd.Low)

    //This sets a metadata key on the service
    svc.SetMetadata("relayCtrlApp", "set relay state")

    iface := svc.RegisterInterface("ctrlRelay", "i.echo")

    // assume temperature is always in 'F
    var relay1 int
    var relay2 int
    relay1 = 0
    relay2 = 0

    //Users can change the state of either relay by pinging either slots with the new 
    iface.SubscribeSlot("ctrl1", func(m *bw.SimpleMessage) {
        if newmsg := m.GetOnePODF(bw.PODFString); newmsg != nil {
            fmt.Println("got a new state command for relay 1")
            command, err := strconv.Atoi(newmsg.(bw.TextPayloadObject).Value())
            if err == nil {
                fmt.Println("new state is :", command)
                relay1 = command
                if relay1==0 {
                    embd.DigitalWrite(relay1pin, embd.Low)
                } else {
                    embd.DigitalWrite(relay1pin, embd.High)
                }
            } else {
                fmt.Println("parsing error")
            }
        }
    })

    //People can set what the message should be
    iface.SubscribeSlot("ctrl2", func(m *bw.SimpleMessage) {
        if newmsg := m.GetOnePODF(bw.PODFString); newmsg != nil {
            fmt.Println("got a new state command for relay 2")
            command, err := strconv.Atoi(newmsg.(bw.TextPayloadObject).Value())
            if err == nil {
                fmt.Println("new state is :", command)
                relay2 = command
                if relay2==0 {
                    embd.DigitalWrite(relay2pin, embd.Low)
                } else {
                    embd.DigitalWrite(relay2pin, embd.High)
                }
            } else {
                fmt.Println("parsing error")
            }
        }
    })

    //Also, every five seconds, we publish the message
    for {
        fmt.Println("current relay1 state is ", relay1)
        fmt.Println("current relay2 state is ", relay2)
        po := bw.CreateTextPayloadObject(bw.PONumString, "relay1 = " + strconv.Itoa(relay1) + ", relay2 = " + strconv.Itoa(relay2))
        err := iface.PublishSignal("current", po)
        fmt.Println("Published, error was ", err)
        time.Sleep(2 * time.Second)
    }
}
