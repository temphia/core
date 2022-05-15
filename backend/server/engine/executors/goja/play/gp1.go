package main

import (
	"github.com/dop251/goja"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/backend/server/btypes/easyerr"
	"github.com/temphia/temphia/backend/server/btypes/rtypes/event"
)

func main1() {
	const s1 = `
	function f(param) {

		throw "this is bullshit"
		   
	return {
	 	Meta: { aaa: 11}
	 }

}
`

	const s2 = `
function f(param) {
	throw "this is bullshit"
}
`

	vm := goja.New()
	_, err := vm.RunString(s2)
	if err != nil {
		panic(err)
	}

	var fn func(ev *event.Request) (*event.Response, error)
	err = vm.ExportTo(vm.Get("f"), &fn)
	if err != nil {
		panic(err)
	}

	pp.Println(fn(&event.Request{
		Id:   "hahah",
		Type: "xyz",
		Name: "who",
	}))
}

// 1. if return contain error then it will translate to exception inside js vm
// 2. but if same retuntype is interface{} but return go err (errors.New("")) then it will look like object type
// 3. if func has multiple rerturn types (interface{}, interface{}) then it will be [return_type1, return_type2]
// 		but if it is (interface{}, error) then it will throw error if error is not nil ortherwise single return type

func main2() {

	const script = `

	try {
		const resp = __call_example()
		__log(typeof resp) 
		__log(JSON.stringify(resp))
	}catch (err) {
		__log(err)
	}
	`

	vm := goja.New()

	vm.Set("__log", func(msg interface{}) {
		pp.Println(msg)
	})

	vm.Set("__call_example", func() (interface{}, error) {
		return 1, nil //  easyerr.Error("aaa")
	})

	_, err := vm.RunString(script)
	if err != nil {
		panic(err)
	}

}

func mainErrorHandeling() {

	const script = `

	try {
		const resp = __call_example()
		__log(typeof resp) 
		__log(JSON.stringify(resp))
	}catch (err) {
		__log(err)
	}
	`

	vm := goja.New()

	vm.Set("__log", func(msg interface{}) {
		pp.Println(msg)
	})

	vm.Set("__call_example", func() error {
		return easyerr.Error("aaa")
	})

	_, err := vm.RunString(script)
	if err != nil {
		panic(err)
	}

}
