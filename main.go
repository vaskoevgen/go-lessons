package main

import (
	"fmt"
	"strings"

	"github.com/pborman/uuid"
	"github.com/vaskoevgen/go-lessons/modules/firstmodule"
)

func main() {
	firstmodule.outputtext("test")
	uuid_app()
}

func uuid_app() {
	uuidWithHyphen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)
}