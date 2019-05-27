//Enter the path you wish to crawl as an argument when you're executing go
package main

import (
	"fmt"
	"os/exec"
	"os"
	"runtime"
)

func main() {
	if(runtime.GOOS == "windows") {
		fmt.Println("Using windows os")
		path := os.Args[1]
		outputFile := os.Args[2]
		cmd := exec.Command("cmd.exe","/C","C:\\Users\\viseshini.palle\\Desktop\\InfoGov\\drive.bat",path,outputFile)
		if err := cmd.Run(); err != nil { 
			fmt.Println("Error: ", err)
			return
		}
	} else {
		fmt.Println("Using linux os") 
		cmd := exec.Command("/bin/sh","drive.sh")
				if err := cmd.Run(); err != nil { 
			fmt.Println("Error: ", err)
			return
		}
		}
}