package netlink

import (
	"bytes"
	"encoding/binary"
	"fmt"
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

type InetDiagReq struct {
	Family    uint8
	SourceLen uint8
	DestLen   uint8
	Ext       uint8
	Id        InetDiagSockId
	States    uint32
	Dbs       uint32
}

type InetDiagSockId struct {
	SourcePort    uint16
	DestPort      uint16
	SourceAddress [4]uint32
	DestAddress   [4]uint32
	If            uint32
	Cookie        [2]uint32
}

type InetDiagMsg struct {
  Header       syscall.NlMsghdr
  InetDiagData InetDiagMsgData
}

type InetDiagMsgData struct {
	Family  uint8
	State   uint8
	Timer   uint8
	Retrans uint8
	Id InetDiagSockId
	Expires uint32
	Rqueue  uint32
	Wqueue  uint32
	Uid     uint32
	Inode   uint32
}

func readInetDiagMsg(r *bytes.Buffer, m *InetDiagMsgData) (er error) {
    /* Reminder:
     *
     *  This data comes off the wire as it would come off the wire had we been
     *  using the C netlink API.
     *
     *  This means, specifically, that if you want to actually print the ports
     *  or IP address, you will first need to do the equivalent of ntohs(port)
     *  or ntohl(ipaddress) prior to outputting or bit shifting.
     */
    binary.Read(r, SystemEndianness, m)
    return
}

func ParseInetDiagMessage(msg syscall.NetlinkMessage) (ParsedNetlinkMessage,
	error) {

	m := new(InetDiagMsg)
	m.Header = msg.Header
	buf := bytes.NewBuffer(msg.Data)

	err := readInetDiagMsg(buf, &m.InetDiagData)
	return m, err
}
