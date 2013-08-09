package main

import (
	"encoding/json"
	"os"
	"os/exec"

)

type Msg struct {
	Env []string `json:"env"`
}

func main() {
	env := os.Environ()
	msg := &Msg{env}
	j, e := json.Marshal(&msg)
	if e != nil {
		panic(e)
	}
	obj := new(Msg)
	e = json.Unmarshal(j, obj)
	if e != nil {
		panic(e)
	}
	cmd := exec.Command("ls", "/home/agonopolskiy/")
	cmd.Env = obj.Env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	e = cmd.Start()
	if e != nil {
		panic(e)
	}
	e = cmd.Wait()
	if e != nil {
		panic(e)
	}
}
 


















