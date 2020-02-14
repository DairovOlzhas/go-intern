package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"github.com/fatih/color"
	"os"
)

var showFiles, showDirs *bool

func rec(path, pic string) {


	filesAndDirs, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	// ╠═╚

	var files, dirs []os.FileInfo

	for _, fd := range filesAndDirs {
		if fd.IsDir() {
			dirs = append(dirs, fd)
		} else {
			files = append(files, fd)
		}
	}


	if *showFiles {

		for i, f := range files {
			fmt.Printf(pic)
			if i == len(files) - 1 && (!(*showDirs) || len(dirs) == 0){
				fmt.Printf("╚═══")
			} else {
				fmt.Printf("╠═══")
			}
			fmt.Println(f.Name())
		}
	}
	if *showDirs {

		for i, d := range dirs {
			fmt.Printf(pic)
			if i == len(dirs) - 1 {
				fmt.Printf("╚═══")
			} else {
				fmt.Printf("╠═══")
			}
			color.Set(color.FgYellow)
			fmt.Println(d.Name())
			color.Unset()

			if i == len(dirs) - 1 {
				rec(path+"/"+d.Name(), pic+"    ")
			} else {
				rec(path+"/"+d.Name(), pic+"║   ")
			}

		}
	}
}

func main(){
	var path *string
	//path = "."
	path = flag.String("path", ".", "a string")
	showFiles = flag.Bool("f", false, "a bool")
	showDirs = flag.Bool("d", false, "a bool")


	flag.Parse()

	if *showDirs == false && *showFiles == false {
		*showDirs = true
		*showFiles = true
	}
	//fmt.Println(*showFiles)
	//fmt.Println(*showDirs)

	//fmt.Println(*path)
	rec(*path,"")
	//rec(".", "")

}

