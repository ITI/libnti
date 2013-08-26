package libnti

import (
    "fmt"
    "errors"
    "net"
)

var ControlCodes = map[string][]byte {
    "ReadSize" :            []byte{0x52, 0x55},
    "ReadOutput" :          []byte{0x52, 0x4f},
    "ConnectSource" :       []byte{0x43, 0x53},
    "ConnectAll" :          []byte{0x43, 0x41},
    "ExamineConnections" :  []byte{0x53, 0x58},
    "CloseConnection" :     []byte{0x58, 0x58},
}

const EndCommand = byte(0x0d)



type Veemux struct {
    IP string
    Port int
    Debug bool
}

func (v *Veemux) SendCommand (cmd string, opts []byte) (err error) {
    if v.IP == ""  || v.Port == 0 {
        return errors.New("No IP or Port")
    }

    command, ok := ControlCodes[cmd]
    if !ok {
        return errors.New(fmt.Sprintf("%v is not an available command", cmd))
    }

    for _,o := range opts {
        command = append(command, byte(o))
    }
    command = append(command, EndCommand)

    if v.Debug {
        fmt.Printf("%v\n", command)
    } else {
        addr,err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%v:%v", v.IP, v.Port))
        if err != nil {
            return err
        }
        con,err := net.DialTCP("tcp", nil, addr)
        if err != nil {
            return err
        }
        _, err = con.Write(command)
        if err != nil {
            return err
        }
    }

    return nil
}

