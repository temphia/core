package sockcore

import (
	"errors"
	"log"
	"time"

	"github.com/k0kubun/pp"
	"github.com/temphia/core/backend/server/btypes/service"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var (
	errConnClosed = errors.New("conn closed")
)

type Conn struct {
	parent  *SockCore
	conn    service.Conn
	ns      string
	closed  bool
	expiry  int
	writeCh chan []byte
}

func (c *Conn) start() {
	go c.writeLoop()
	go c.readLoop()
}

func (c *Conn) write(payload []byte) {
	c.writeCh <- payload
}

func (c *Conn) writeLoop() {

	ticker := time.NewTicker(10 * time.Second)
	defer func() {
		log.Println("closing write loop")
		ticker.Stop()
		c.close()
	}()

	for {
		if c.closed {
			return
		}
		data := <-c.writeCh
		err := c.conn.Write(data)

		if err == errConnClosed {
			return
		}

		if err != nil {
			log.Println(err)
			continue
		}
	}

}

func (c *Conn) readLoop() {

	defer func() {
		log.Println("closing read loop")
		log.Println(c.close())
	}()

	for {
		msg, err := c.conn.Read()
		if err != nil {
			log.Println("PACKET READ ERR", err)
			return
		}

		c.processPacket(msg)
	}
}

func (c *Conn) processPacket(msg []byte) {

	msgRoom := gjson.GetBytes(msg, "room").String()
	msgFromid := gjson.GetBytes(msg, "from_id").String()
	msgType := gjson.GetBytes(msg, "type").String()

	log.Println(msgFromid)

	if msgRoom == "" {
		log.Println("Not allowed in this room", msgRoom)
		return
	}

	pp.Println(string(msg))

	switch msgType {

	case MESSAGE_PEER_DIRECT:
		log.Println("fixme => implement client p2p message")
	case MESSAGE_PEER_BROADCAST:
		c.parent._broadcastFromClient(c.ns, msgRoom, c.conn.Id(), msg)
	case MESSAGE_PEER_PUBLISH:
		rawTargets := gjson.GetBytes(msg, "targets").Array()
		targets := make([]string, 0, len(rawTargets))
		for _, target := range rawTargets {
			// fixme => verify with bloom_ticket
			targets = append(targets, target.String())
		}

		var ticket string
		var err error

		if c.parent.ticketParse != nil {
			ticket = gjson.GetBytes(msg, "ticket").String()
			targets, err = c.parent.ticketParse(c.ns, msgRoom, c.conn.Id(), ticket, targets)
			if err != nil {
				log.Println(err)
				break
			}
		}

		stripped1, err := sjson.DeleteBytes(msg, "targets")
		if err != nil {
			log.Println(err)
			break
		}

		stripped2, err := sjson.DeleteBytes(stripped1, "ticket")
		if err != nil {
			log.Println(err)
			break
		}

		c.parent._ticketMsgFromClient(c.ns, msgRoom, msgFromid, ticket, stripped2, targets...)
	default:
		log.Println("Not allowed Message", msgType)
	}

}

func (c *Conn) close() error {
	if c.closed {
		return nil
	}
	c.closed = true
	func() {
		c.parent.connLock.Lock()
		defer c.parent.connLock.Unlock()
		delete(c.parent.connections, c.ns+c.conn.Id())
	}()

	log.Println("closing conn", c.ns, "#", c.conn.Id())

	return c.conn.Close()
}
