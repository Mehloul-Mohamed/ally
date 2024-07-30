package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Mehloul-Mohamed/ally/app"
)

func help(f string) {
	switch f {
	case "start":
		fmt.Print("usage: ally start name url token\n\n")
		fmt.Println("positional arguments:")
		fmt.Println("  name\t\t\tCTF name")
		fmt.Println("  url\t\t\tCTF url")
		fmt.Println("  token\t\t\tYour API token")
	case "list":
		fmt.Println("usage: ally list")
	case "attempt":
		fmt.Print("usage: ally attempt id\n\n")
		fmt.Println("positional arguments:")
		fmt.Println("  id\t\t\tChallenge id")
	case "info":
		fmt.Println("usage: ally info")
	case "main":
		fmt.Println("usage: ally {start,list,attempt,info}")
		fmt.Println("positional arguments:")
		fmt.Println("  {start,list,attempt}")
		fmt.Println("\tstart\t\t\tStart a CTF")
		fmt.Println("\tlist\t\t\tShow challenge list")
		fmt.Println("\tattempt\t\t\tAttempt a challenge")
		fmt.Println("\tinfo\t\t\tShow scoreboard & team stats")
	}
}

func main() {
	if len(os.Args) < 2 {
		help("main")
		return
	}
	switch os.Args[1] {
	case "start":
		if len(os.Args) != 5 {
			help("start")
			return
		}
		name := os.Args[2]
		url := strings.TrimSuffix(os.Args[3], "/")
		token := "token " + os.Args[4]
		if name == "" || url == "" || token == "" {
			help("main")
		}
		err := app.StartCtf(name, url, token)
		if err != nil {
			panic(err)
		}
	case "list":
		if len(os.Args) != 2 {
			help("list")
			return
		}
		err := app.DisplayChallList()
		if err != nil {
			panic(err)
		}
	case "attempt":
		if len(os.Args) != 3 || os.Args[2] == "-h" {
			help("attempt")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic("id must be an integer")
		}
		err = app.Attempt(id)
		if err != nil {
			panic(err)
		}
	// Credit to shadow1004 on GitHub for the idea
	case "info":
		if len(os.Args) != 2 {
			help("info")
			return
		}
		err := app.DisplayTeamInfo()
		if err != nil {
			panic(err)
		}
	default:
		help("main")
	}
}
