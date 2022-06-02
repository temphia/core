package sockdhub

import (
	"fmt"

	"github.com/temphia/core/backend/server/btypes/service"
)

type PlugConnOptions struct {
	TenantId string
	UserId   string
	GroupId  string
	DeviceId string
	Plug     string
	Conn     service.Conn
}

type UserConnOptions struct {
	TenantId string
	UserId   string
	GroupId  string
	DeviceId string
	Conn     service.Conn
}

type UpdateDynRoomTagsOptions struct {
	TenantId  string
	DynSource string
	DynGroup  string
	ConnId    string
}

func (s *SockdHub) AddPlugConnection(opts PlugConnOptions) error {
	return s.sockd.NewConnection(&service.ConnOptions{
		NameSpace: opts.TenantId,
		Conn:      opts.Conn,
		Expiry:    10000,
		PreJoinRooms: map[string][]string{
			ROOM_PLUG_DEV: {
				fmt.Sprintf("plug_%s", opts.Plug),
			},
		},
	})
}

func (s *SockdHub) AddUserConnOptions(opts UserConnOptions) error {
	connTags := []string{
		fmt.Sprint("sys.user_", opts.UserId),
		fmt.Sprint("sys.ugroup_", opts.GroupId),
		fmt.Sprint("sys.device_", opts.DeviceId),
		TAG_REALUSER,
	}

	return s.sockd.NewConnection(&service.ConnOptions{
		NameSpace: opts.TenantId,
		Conn:      opts.Conn,
		Expiry:    10000,
		PreJoinRooms: map[string][]string{
			ROOM_SYS_USERS: connTags,
			ROOM_SYSTABLE:  connTags,
		},
	})

}

func (s *SockdHub) UpdateDynRoomTags(opts UpdateDynRoomTagsOptions) error {

	return s.sockd.RoomUpdateTags(
		opts.TenantId,
		ROOM_SYSTABLE,
		opts.ConnId,
		&service.UpdateTagOptions{
			AddTags:    []string{fmt.Sprintf("dgroup.%s.%s", opts.DynSource, opts.DynGroup)},
			ClearOld:   true,
			RemoveTags: []string{},
		})
}

func (s *SockdHub) UpdateRoomTags(tenantId, room, connid string, opts *service.UpdateTagOptions) error {
	return s.sockd.RoomUpdateTags(
		tenantId,
		room,
		connid,
		opts,
	)
}
