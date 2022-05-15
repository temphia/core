package binder

type ipcBinder struct {
	core   *Binder
	loaded bool
	slots  map[string]interface{}
}

func (ib *ipcBinder) clear() {
	ib.loaded = false
	for k := range ib.slots {
		delete(ib.slots, k)
	}
}

func (ib *ipcBinder) getSlot(slotname string) (interface{}, error) {
	// if ib.loaded {
	// 	ib.slots[slotname]

	// }

	// ib.core.filterResource(entities.ResSlot, func(key string, res *entities.Resource) {
	// 	// sres, err := res.Slot()
	// 	// if err != nil {
	// 	// 	panic(err)
	// 	// }

	// 	// sres.Invoker
	// })

	return nil, nil

}
