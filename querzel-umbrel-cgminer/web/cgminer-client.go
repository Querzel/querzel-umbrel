package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

type RawResponse []byte

func cgminerCommand(command Command) (*RawResponse, error) {
	return cgminerCommandWithServer(command, *flagCgminerHost, *flagCgminerPort)
}

func cgminerCommandWithServer(command Command, host string, port int) (*RawResponse, error) {
	if conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port)); err != nil {
		return nil, err
	} else {
		defer conn.Close()

		request := command.ToJsonString()
		if _, err = fmt.Fprint(conn, request); err != nil {
			panic(err)
		}

		if response, err := ioutil.ReadAll(conn); err != nil {
			return nil, err
		} else {
			result := RawResponse(response)
			return &result, nil
		}
	}
}
