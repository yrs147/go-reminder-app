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

	now := time.Now()
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)
	
	t, err := w.Parse(os.Args[1], now)

	if err != nil {
		fmt.Printf(err)
		os.Exit(2)
	}

	if t==nil {
	fmt.Println("Unablr to parse time!")
	os.Exit(2)
	}
	if now.After(t.Time){
		fmt.Println("set a future time!")
		os.Exit(3)
	}

	diff := t.Time.Sub(now)


}