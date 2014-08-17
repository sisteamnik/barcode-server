package bcserver

import (
	"bufio"
	"fmt"
	"os"
)

func Run() {
	f, err := os.OpenFile("/dev/ttyUSB0", os.O_RDONLY|os.O_SYNC, os.ModeCharDevice)
	if err != nil {
		panic(err)
	}

	for {
		r := bufio.NewReader(f)
		ln, err := r.ReadString('\n')
		if err == nil {
			fmt.Print(string(ln))
		} else {
			fmt.Print("error")
		}
	}

}

func init() {
	Run()
}
