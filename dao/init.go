package dao

import (
	"log"

	"github.com/twsouza/tax-packet-shipping/common"
)

var config *common.Config

func init() {
	var err error
	config, err = common.GetConfig()
	log.Println("config err:", err)

	if err != nil {
		log.Println("Err db.connect->common.GetConfig", err)
	}
}
