package main

import (
	"errors"
	"fmt"
	"github.com/reiver/go-telnet"
)

type TelnetClient struct {
	remoteAddr string //telnet remote addr  example ip:port
	status     int    // 0 can't connect | 1 connected
	err        error  // error message
}

func (telnetClient *TelnetClient) build(addr string) *TelnetClient {
	telnetClient.remoteAddr = addr
	if !telnetClient.addrCheck() {
		telnetClient.result(0, errors.New("远端地址格式有误"))
	}
	telnetClient.result(1, errors.New(""))
	return telnetClient
}

func (telnetClient *TelnetClient) dialTo() bool {
	if 1 != telnetClient.status {
		return false
	}
	conn, e := telnet.DialTo(telnetClient.remoteAddr)
	if nil != e {
		telnetClient.result(0, errors.New("connect "+telnetClient.remoteAddr+" timeout..."))
		return false
	}
	fmt.Println("connected to " + telnetClient.remoteAddr + "...")
	conn.Close()
	return true
}

func (telnetClient TelnetClient) addrCheck() bool {
	//@TODO regex judge ,maybe in space time
	return true
}

func (telnetClient *TelnetClient) result(status int, err error) {
	telnetClient.status = status
	telnetClient.err = err
}
