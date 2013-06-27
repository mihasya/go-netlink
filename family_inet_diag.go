package netlink

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"syscall"
)

func MakeInetDiagMessage(socktype, family int) (msg RawNetlinkMessage) {
  msg.Header.Type = socktype
  msg.Header.Flags = syscall.NLM_F_ROOT | syscall.NLM_F_MATCH |
  syscall.NLM_F_REQUEST
  msg.Data = make([]byte, 0)
  msg.Data[0] = uint8(syscall.AF_INET)
  msg.Data[1] = uint8(SS_ALL)

  return msg
}

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
  SS_ALL = (1<<SS_MAX)-1
)

type InetDiagSockId struct {
  SourcePort  uint16
  DestPort    uint16
  SourceAddress net.IP
  DestAddress net.IP
  If  uint32
  Cookie [2]uint32
}

type InetDiagMsg struct {
  Family  byte  `netlink:"1" type:"fixed"`
  State   byte  `netlink:"2" type:"fixed"`
  Timer   byte  `netlink:"3" type:"fixed"`
  Retrans byte  `netlink:"4" type:"fixed"`

  Id      InetDiagSockId `netlink:"5" type:"nested"`

  Expires uint32 `netlink:"6" type:"fixed"`
  Rqueue  uint32 `netlink:"7" type:"fixed"`
  Wqueue  uint32 `netlink:"8" type:"fixed"`
  Uid     uint32 `netlink:"9" type:"fixed"`
  Inode   uint32 `netlink:"10" type:"fixed"`
}

func ParseRouteLinkMessage(msg syscall.NetlinkMessage) (ParsedNetlinkMessage, error) {
	m := new(RouteLinkMessage)
	m.Header = msg.Header
	buf := bytes.NewBuffer(msg.Data)
	binary.Read(buf, SystemEndianness, &m.IfInfo)
	// read link attributes
	er := ReadManyAttributes(buf, m)
	return m, er
}

// Parse address messages for family NETLINK_ROUTE
func ParseRouteAddrMessage(msg syscall.NetlinkMessage) (ParsedNetlinkMessage, error) {
	m := new(RouteAddrMessage)
	m.Header = msg.Header
	buf := bytes.NewBuffer(msg.Data)

	binary.Read(buf, SystemEndianness, &m.IfAddr)
	// read Address attributes
	er := ReadManyAttributes(buf, m)
	return m, er
}

func ParseInetDiagMessage(msg syscall.NetlinkMessage) (ParsedNetlinkMessage,
error) {
  m := new()
  m.HEader = msg.Header
  buf := bytes.NewBuffer(msg.Data)

  binary.Read(buf, SystemEndiannes, &)
  er := ReadManyAttributes(buf, m)
  return m, err
}

func ParseInetDiagMessage(msg syscall.NetlinkMessage) (ParsedNetlinkMessage, error) {
  return ParseInetDiagMessage(msg)
}
