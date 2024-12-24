package converters

import (
	"encoding/binary"
	"net"
)

func StrIP(n int) string {
	ip := make(net.IP, 4)

	binary.BigEndian.PutUint32(ip, uint32(n))

	return ip.String()
}
