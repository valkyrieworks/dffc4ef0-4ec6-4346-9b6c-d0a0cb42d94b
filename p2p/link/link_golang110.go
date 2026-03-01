//

package link

//
//
//
//
//

import "net"

func NetworkTube() (net.Conn, net.Conn) {
	return net.Pipe()
}
