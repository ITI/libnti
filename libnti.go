package libnti

import (
    "fmt"
    "errors"
    "net"
)

var ControlCodes = map[string]string {
    "ReadSize" :            "RU",
    "ReadOutput" :          "RO ",
    "ConnectSource" :       "CS ",
    "ConnectAll" :          "CA ",
    "ExamineConnections" :  "SX",
    "CloseConnection" :     "XX",
}

const EndCommand = "\r"



type Veemux struct {
    IP string
    Port int
    Debug bool
}

func (v *Veemux) SendCommand (cmd string, opt string) (err error) {
    if v.IP == ""  || v.Port == 0 {
        return errors.New("No IP or Port")
    }

    command, ok := ControlCodes[cmd]
    if !ok {
        return errors.New(fmt.Sprintf("%v is not an available command", cmd))
    }

    command = fmt.Sprintf("%v%v%v", command, opt, EndCommand)

    if v.Debug {
        fmt.Printf("%v\n", command)
    } else {
        addr,err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%v:%v", v.IP, v.Port))
        if err != nil {
            fmt.Printf("ERR::::%v", err);
            return err
        }
        con,err := net.DialTCP("tcp", nil, addr)
        if err != nil {
            fmt.Printf("ERR::::%v", err);
            return err
        }
        l, err := con.Write([]byte(command))
        if err != nil {
            fmt.Printf("ERR::::%v", err);
            return err
        }
        con.Close()
        fmt.Printf("WROTE:::::%v", l)
    }
    return nil
}

