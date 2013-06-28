package netlink

import (
	"bytes"
	"encoding/binary"
	"fmt"
	//"net"
	"syscall"
)

func MakeInetDiagMessage(socktype int) (msg RawNetlinkMessage) {
	msg.Header.Type = TCPDIAG_GETSOCK
	msg.Header.Flags = syscall.NLM_F_ROOT | syscall.NLM_F_REQUEST

	buf := make([]byte, 0)
	w := bytes.NewBuffer(buf)

	request := InetDiagReq{
		Family: syscall.AF_INET,
		States: SS_ALL,
	}

	err := binary.Write(w, binary.LittleEndian, request)
	if err != nil {
		fmt.Println(err)
	}

	msg.Data = w.Bytes()

	return msg
}

const (
	TCPDIAG_GETSOCK = 18
)

const (
	SS_UNKNOWN = iota
	SS_ESTABLISHED
	SS_SYN_SENT
	SS_SYN_RECV
	SS_FIN_WAIT1
	SS_FIN_WAIT2
	SS_TIME_WAIT
	SS_CLOSE
	SS_CLOSE_WAIT
	SS_LAST_ACK
	SS_LISTEN
	SS_CLOSING
	SS_MAX
	SS_ALL = ((1 << SS_MAX) - 1)
)

type IPWrap struct {
	IP []byte `netlink:"1" type:"bytes"`
}

type InetDiagReq struct {
	Family    byte           `netlink:"1" type:"fixed"`
	SourceLen byte           `netlink:"2" type:"fixed"`
	DestLen   byte           `netlink:"3" type:"fixed"`
	Ext       byte           `netlink:"4" type:"fixed"`
	Id        InetDiagSockId `netlink:"5" type:"nested"`
	States    uint32         `netlink:"6" type:"fixed"`
	Dbs       uint32         `netlink:"7" type:"fixed"`
}

type InetDiagSockId struct {
	SourcePort    uint16    `netlink:"1" type:"fixed"`
	DestPort      uint16    `netlink:"2" type:"fixed"`
	SourceAddress [4]byte   `netlink:"3" type:"bytes"`
	DestAddress   [4]byte   `netlink:"4" type:"bytes"`
	If            uint32    `netlink:"5" type:"fixed"`
	Cookie        [2]uint32 `netlink:"6" type:"fixed"`
}

type InetDiagMsg struct {
	Header syscall.NlMsghdr `netlink:"1" type:"nested"`

	Family  byte `netlink:"2" type:"fixed"`
	State   byte `netlink:"3" type:"fixed"`
	Timer   byte `netlink:"4" type:"fixed"`
	Retrans byte `netlink:"5" type:"fixed"`

	Id InetDiagSockId `netlink:"6" type:"nested"`

	Expires uint32 `netlink:"7" type:"fixed"`
	Rqueue  uint32 `netlink:"8" type:"fixed"`
	Wqueue  uint32 `netlink:"9" type:"fixed"`
	Uid     uint32 `netlink:"10" type:"fixed"`
	Inode   uint32 `netlink:"11" type:"fixed"`
}

func ParseInetDiagMessage(msg syscall.NetlinkMessage) (ParsedNetlinkMessage,
	error) {
	m := new(InetDiagMsg)
	m.Header = msg.Header
	buf := bytes.NewBuffer(msg.Data)

	err := ReadManyAttributes(buf, m)
	return m, err
}
