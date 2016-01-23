package tcp

import (
	"time"

	"github.com/Maksadbek/tcpovericmp/icmp"
)

type Conn struct {
	socket icmp.Socket
}

// Addr is host:port
func Dial(addr string) (Conn, error) {

}

func (c Conn) Read(b []byte) (int, error) {

}

func (c Conn) Write(b []byte) (int, error) {

}

func (c Conn) Close() {

}

func (c Conn) LocalAddr() string {

}

func (c Conn) RemoteAddr() string {

}

func (c Conn) SetDeadline(t time.Time) error {

}

func (c Conn) SetReadDeadline(t time.Time) error {

}

func (c Conn) SetWriteDeadline(t time.Time) error {

}

type Listener struct {
	socket icmp.Socket
}

// Set local TCP address
func Listen(addr string) (Listener, error) {
	s, err := icmp.NewListenSocket(addr)

}

// Handshake implementation will be here.
// Wait for ACK from some icmp subpackege type.
func (l Listener) Accept() (Conn, error) {

}

func (l Listener) Close() {

}

func (l Listener) Addr() string {

}
