package main

import (
	"io/ioutil"

	"github.com/k0kubun/pp"
	"github.com/rs/xid"

	"github.com/temphia/temphia/backend/server/btypes/rtypes"
	"github.com/temphia/temphia/backend/server/btypes/rtypes/event"
	"github.com/temphia/temphia/backend/server/engine/executors/goja"
	"github.com/temphia/temphia/backend/tests/mbindings"
)

func main() {

	server, err := ioutil.ReadFile("build/bundle.js")
	handleErr(err)

	mbind := &mbindings.MockedBindings{
		SelfFiles: map[string][]byte{"server.js": server},
	}

	pp.Println(mbind)

	executor, err := goja.NewBuilder(rtypes.ExecutorOption{
		Binder:   nil,
		PlugId:   "",
		AgentId:  "",
		Slug:     "",
		ExecType: "",
	})
	handleErr(err)

	eresp, err := executor.Process(&event.Request{
		Id:   xid.New().String(),
		Type: "creq",
		Name: "main",
		Vars: map[string]interface{}{
			"ctx_var_1": 128,
		},
		Data: []byte(`{"hello": "zzzzz"}`),
	})

	handleErr(err)

	pp.Println(eresp)

}

func handleErr(err error) {
	if err != nil {
		pp.Println(err)
		panic(err)
	}

}
