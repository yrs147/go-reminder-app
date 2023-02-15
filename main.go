package main

import(
	"fmt"
	"os"
	"os/exec"
	"time"
	"github.com/gen2brain/beeep"
	"github.com/oblebedev/when"
	"github.com/oblebedev/when/rules/common"
	"github.com/oblebedev/when/rules/en"
)

const(
	markName="GOLANG_CLI_REMINDER"
	markValue = "1"
)

func main(){
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <hh:mm> <text message\n>",os.Args[0])	
		os.Exit(1)
	}

}