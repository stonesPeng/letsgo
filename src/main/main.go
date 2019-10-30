package main

import "fmt"

func main() {
	client := TelnetClient{}
	connected := client.build("49.235.99.162:3305").dialTo()
	if !connected {
		fmt.Println(client.err.Error())
	}
}
