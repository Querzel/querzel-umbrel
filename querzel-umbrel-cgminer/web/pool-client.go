package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
)

func GetPoolInfoFromUser(user string) string {
	btcAddress := btcAddress(user)
	return getPoolInfo(btcAddress)
}

func getPoolInfo(btcAddress string) string {
	client := resty.New()
	if resp, err := client.R().Get("http://solo.ckpool.org/users/" + btcAddress); err != nil {
		panic(err)
	} else {
		return fmt.Sprint(resp)
	}
}

func btcAddress(user string) string {
	idx := strings.Index(user, ".")
	if idx == -1 {
		return user
	} else {
		return user[:idx]
	}
}
