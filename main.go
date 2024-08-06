package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Mehloul-Mohamed/ally/app"
)

var Url, Token string

func help(f string) {
	switch f {
	case "start":
		fmt.Println("usage: ally start name url token\n\n" +
			"positional arguments:\n" +
			"  name\t\t\tCTF name\n" +
			"  url\t\t\tCTF url\n" +
			"  token\t\t\tYour API token")
	case "list":
		fmt.Println("usage: ally list")
	case "attempt":
		fmt.Println("usage: ally attempt id\n\n" +
			"positional arguments:\n" +
			"  id\t\t\tChallenge id")
	case "fetch":
		fmt.Println("usage: ally fetch")
	case "info":
		fmt.Println("usage: ally info")
	case "main":
		fmt.Println("usage: ally {start,list,attempt,info}\n" +
			"positional arguments:\n" +
			"  {start,list,attempt}\n" +
			"\tstart\t\t\tStart a CTF\n" +
			"\tlist\t\t\tShow challenge list\n" +
			"\tattempt\t\t\tAttempt a challenge\n" +
			"\tinfo\t\t\tShow scoreboard & team stats\n" +
			"\tfetch\t\t\tFetch all challenges")
	}
}

func main() {
	if len(os.Args) < 2 {
		help("main")
		return
	}

	if os.Args[1] == "start" {
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
		if errors.Is(err, fs.ErrExist) {
			log.Fatalln("Ctf already started")
		}

		if err != nil {
			panic(err)
		}
		return
	}

	// Other Options
	// Read Credentials
	wd, _ := os.Getwd()
	bytes, err := os.ReadFile(wd + "/credentials.txt")

	if errors.Is(err, os.ErrNotExist) {
		log.Fatalln("Credentials file not found")
	}

	if err != nil {
		panic(err)
	}

	slice := strings.Split(string(bytes), "\n")
	Url = slice[0]
	Token = slice[1]

	switch os.Args[1] {
	case "list":
		if len(os.Args) != 2 {
			help("list")
			return
		}
		err := app.DisplayChallList(Url, Token)
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
			log.Fatalln("id must be an integer")
		}
		err = app.Attempt(id, Url, Token)
		if err != nil {
			panic(err)
		}
	case "fetch":
		if len(os.Args) != 2 {
			help("fetch")
			return
		}
		err := app.FetchAll(Url, Token)
		if err != nil {
			panic(err)
		}
	// Credit to shadow1004 on GitHub for the idea
	case "info":
		if len(os.Args) != 2 {
			help("info")
			return
		}
		err := app.DisplayTeamInfo(Url, Token)
		if err != nil {
			panic(err)
		}
	default:
		help("main")
	}
}
