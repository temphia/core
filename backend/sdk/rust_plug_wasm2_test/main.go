package main

import (
	"io/ioutil"
	"os"

	"github.com/k0kubun/pp"
	"github.com/rs/xid"
	"github.com/temphia/temphia/backend/server/btypes/rtypes"
	"github.com/temphia/temphia/backend/server/btypes/rtypes/event"
	"github.com/temphia/temphia/backend/server/engine/executors/wasmer2"
	"github.com/temphia/temphia/backend/tests/mbindings"
)

func main() {

	wasmFile, err := ioutil.ReadFile(os.Args[1])
	handleErr(err)

	mbind := &mbindings.MockedBindings{
		SelfFiles: map[string][]byte{"server.wasm": wasmFile},
	}

	pp.Println(mbind)

	builder := wasmer2.NewBuilder()
	executor, err := builder(rtypes.ExecutorOption{
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
		Name: "hello_wasm",
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
