package main

import (
	"Go_Work/models/res"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
	"os"
)

const file = "models/res/err_code.json"

type ErrMap map[res.ErrCode]string

func testmain() {
	byteDate, err := os.ReadFile(file)
	if err != nil {
		logrus.Error(err)
		return
	}
	var errMap = ErrMap{}
	err = json.Unmarshal(byteDate, &errMap)
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(errMap[res.SettingsError])
}
