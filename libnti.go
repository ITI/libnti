package libnti

import (
    "fmt"
    "errors"
)


type veemux struct {
    IP string
    Port int
    Commands map[string]string
    Debug bool
}

func NewVeemux(debug bool) *veemux {
    return &veemux{Commands: map[string]string {
        "unitSize" : "RU\r",
        "outputConnection": "RO%v\r",
        "connectPort": "CS%v\r",
        },
        Debug: debug}
}

func (v *veemux) SendCommand (cmd string, opts ...string) (err error) {
    if v.IP == ""  || v.Port == 0 {
        return errors.New("No IP or Port")
    }

    command, ok := v.Commands[cmd]
    if !ok {
        return errors.New(fmt.Sprintf("%v is not an available command", cmd))
    }

    opstring := ""
    for _, v := range opts {
        opstring = fmt.Sprintf("%v %v", opstring, v)
    }

    if len(opts) > 0 {
        command = fmt.Sprintf(command, opstring)
    }

    if v.Debug {
        fmt.Printf("%v\n", command)
    } else {
        // Shoot it over the wire
    }

    return nil
}

