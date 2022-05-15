package sockcore

import (
	"fmt"
	"time"

	"github.com/k0kubun/pp"
)

var Debug = false

func (s *SockCore) debug() {
	if !Debug {
		return
	}

	for {
		time.Sleep(time.Second * 20)

		connIds := make([]string, 0, len(s.connections))
		for k := range s.connections {
			connIds = append(connIds, k)
		}

		roomIds := make([]string, 0, len(s.rooms))
		for k, v := range s.rooms {

			s.SendBroadcast(v.ns, v.name, []string{}, []byte(`{"message": "hello from server"}`))

			roomIds = append(roomIds, k)
		}

		pp.Println("####################################################")
		fmt.Println("No of conns ", connIds)
		fmt.Println("No of Rooms ", roomIds)
		pp.Println("=======================")
		fmt.Println("@=> ROOMS", s.rooms)
		fmt.Println("@=> CONNS", s.connections)
		pp.Println("####################################################")

	}
}
