package main

import (
	"encoding/json"
	"os"
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
	for i := range msg.Env {
		if obj.Env[i] != msg.Env[i] {
			panic(i)
		}
	}
}
 