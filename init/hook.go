package hook

import (
	"fmt"
	ps "github.com/mitchellh/go-ps"
)

type process struct {
	name    string
	id      int
	address *int
}

func GetProcessesList() []process {

	p, err := ps.Processes()
	if err != nil {
		fmt.Println(err)
	}

	var processList []process

	for _, v := range p {

		var process = process{}

		process.name = v.Executable()
		process.id = v.Pid()
		process.address = &process.id

		processList = append(processList, process)
	}
	return processList
}

func FindWoWProcess(processList []process) process {

	var wowstruct process

	for _, wow := range processList {
		if wow.name == "Wow.exe" {
			wowstruct = wow
		}
	}
	return wowstruct
}