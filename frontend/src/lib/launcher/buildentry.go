package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

func main() {
	path := build()

	out, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	buildIframe(string(out))
	buildSubOrigin(string(out))

}

func buildSubOrigin(artifact string) {
	final := fmt.Sprintf("var __dirname = ''; var module = {}; module['exports']={};%s", string(artifact))
	cdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	newPath := path.Join(cdir, "../../../../data/templates/suborigin_loader.js")

	err = ioutil.WriteFile(newPath, []byte(final), 0777)
	if err != nil {
		panic(err)
	}

}

func buildIframe(artifact string) {
	out := []byte(artifact)

	out = bytes.ReplaceAll(out, []byte("`"), []byte("\\`"))
	out = bytes.ReplaceAll(out, []byte("$"), []byte("\\$"))

	final := fmt.Sprintf("export default `var __dirname = ''; var module = {}; module['exports']={};%s`", string(out))
	err := ioutil.WriteFile("entry/entry_generated.js", []byte(final), 0777)
	if err != nil {
		panic(err)
	}

}

func build() string {
	build := path.Join(os.TempDir(), "entry_compiled")
	cmd := exec.Command("ncc", "build", "entry/entry.ts", "--out", build)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	filePath := path.Join(build, "index.js")

	return filePath
}

/*

✦ ❯ make generate_launcher_templates
cd code/frontend/lib/launcher && go run buildentry.go
ncc: Version 0.33.1
ncc: Compiling file index.js into CJS
ncc: Using typescript@3.9.10 (local user-provided)
24kB  ../../../../../../../../../../tmp/entry_compiled/index.js
24kB  [2541ms] - ncc 0.33.1


*/
