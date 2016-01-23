package icmp

// TCP packet
type Packet struct {}

type Socket struct {
	local  net.TCPAddr
	remote net.TCPAddr
	in chan packet
	out chan packet
	err chan error
}

type ICMP struct {}

func New(ipaddr string) ICMP {
	i := &ICMP{}
	go func() {
		for {
			// Get TCP packet from Data field of
			// ICMP packet and send it to Conns' sockets
		}
	}()
	return i
}

// Get listener socket and write incoming ACK packages to it
func (p ICMP) Listen(s Socket) error {
	
}

// Get connection socket and write incoming DATA packages to it
func (p ICMP) Connect(s Socket) error {
	
}