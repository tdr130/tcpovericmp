package icmp

import (
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"errors"
	"net"
)

const (
	_PROTOCOL_ICMP = 1
)

var (
	ErrWrongType = errors.New("Got ICMP datagram of wrong type.")
)


type ICMP struct {
	listener *icmp.PacketConn
	raddr *net.IPAddr
}

func New(laddr, raddr string) (*ICMP, error) {
	addr, err := net.ResolveIPAddr("ip4", laddr)
	if err != nil {
		return nil, err
	}
	listener, err := icmp.ListenPacket("ip4:icmp", addr.String())
	if err != nil {
		return nil, err
	}
	r, err := net.ResolveIPAddr("ip4", raddr)
	i := &ICMP{
		listener: listener,
		raddr: r,
	}
	return i, nil
}

func (i *ICMP) Read() ([]byte, error) {
	buf := make([]byte, 10000)
	n, _, err := i.listener.ReadFrom(buf)
	if err != nil {
		return nil, err
	}
	input, err := icmp.ParseMessage(_PROTOCOL_ICMP, buf[:n])
	if err != nil {
		return nil, err
	}
	if input.Type != ipv4.ICMPTypeEchoReply {
		return nil, nil
	}
	echo, ok := input.Body.(*icmp.Echo)
	if !ok {
		return nil, nil
	}
	return echo.Data, nil
}

func (i *ICMP) Write(b []byte) error {
	message := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			Data: b,
		},
	}
	outgoing, err := message.Marshal(nil)
	if err != nil {
		return err
	}
	_, err = i.listener.WriteTo(outgoing, i.raddr)
	if err != nil {
		return err
	}
	return nil
}

func (i *ICMP) Close() error {
	return i.listener.Close()
}

