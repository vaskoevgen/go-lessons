/*
    Первая программа 
    на языке Go
*/

package main // определение пакета для текущего файла

import (
	"fmt" // подключение пакета fmt
	"strings"

	"github.com/pborman/uuid"
	"github.com/vaskoevgen/go-lessons/firstmodule"
	"github.com/vaskoevgen/go-modules/secondmodule"
)

func main() {
	firstmodule.Outputtext("test")
	secondmodule.Outputtext1("text2")
	uuid_app()
}

func uuid_app() {
	uuidWithHyphen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)
}