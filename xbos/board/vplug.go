package main

type Vplug struct {
    status bool
}

func NewVplug() *Vplug {
    // TODO init embed 
    return &Vplug{
        status: false,
    }
}

func (v *Vplug) ActuatePlug(status bool) {
    v.status = status
    // TODO flip pins
}

func (v *Vplug) GetStatus() bool {
    return v.status
}