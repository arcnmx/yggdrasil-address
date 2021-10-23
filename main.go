// stolen from: https://github.com/yggdrasil-network/yggdrasil-go/issues/856#issuecomment-948752138

package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"

	"github.com/yggdrasil-network/yggdrasil-go/src/address"
)

func main() {
	pk, err := hex.DecodeString(os.Args[1])
	if err != nil {
		panic(err)
	}

	ip := net.IP(address.AddrForKey(pk)[:])
	fmt.Println(ip)
}
