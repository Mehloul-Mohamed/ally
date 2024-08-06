package app

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/Mehloul-Mohamed/ally/api"
	"github.com/Mehloul-Mohamed/ally/styles"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
	"github.com/charmbracelet/lipgloss/tree"
)

// Function to produce a fancy tree to display challenges
func buildTree(categories []string, challMap map[string][]api.CtfdChall) *tree.Tree {
	// Build & Render Tree
	ind := func(_ tree.Children, _ int) string { return "    " }
	rootTree := tree.New().
		Root(styles.Header.Render("Challenges")).
		EnumeratorStyle(styles.Category).
		ItemStyle(styles.Category).
		Enumerator(func(_ tree.Children, _ int) string { return "•" }).
		Indenter(ind)

	for _, v := range categories {
		categoryTree := tree.New().
			Root(v).
			Indenter(ind)
		// Add challenges to category
		for _, c := range challMap[v] {
			// This should be done in a simpler manner, but lipgloss is acting weirdly so I'm stuck with this(Probably I'm the issue but whatever)
			var style lipgloss.Style
			if c.SolvedByMe {
				style = styles.Solved
			} else {
				style = styles.Unsolved
			}
			categoryTree.Child(
				fmt.Sprintf(
					"%s %s", styles.Id.Render("#"+strconv.Itoa(c.ID)),
					style.Render(c.Name),
				),
			)
		}
		rootTree.Child(categoryTree)
	}
	return rootTree
}

func StartCtf(name string, url string, token string) error {
	// Setup directory
	home, _ := os.UserHomeDir()
	d := home + "/ctf/"
	_, err := os.Stat(d)
	if errors.Is(err, os.ErrNotExist) {
		log.Println("~/ctf missing. Creating it...")
		err = os.Mkdir(d, 0777)
	}

	if err != nil {
		return err
	}

	err = os.Mkdir(d+name, 0777)
	if err != nil {
		return err
	}
	os.Chdir(d + name)

	// Store credentials
	f, err := os.Create("credentials.txt")
	if err != nil {
		return err
	}

	defer f.Close()
	f.WriteString(url + "\n" + token)
	return nil
}

func DisplayChallList(url, token string) error {
	challs, err := api.GetChallList(url, token)
	if err != nil {
		return err
	}

	// Build Challenge Map & Categories Slice
	// We have to maintain a seperate categories slice so we can sort them and have a consitant way to loop over the map
	var categories []string
	challMap := make(map[string][]api.CtfdChall)
	t := ""
	for _, c := range challs.Data {
		t = c.Category
		if t == "" {
			t = "Uncategorized"
			challMap[t] = append(challMap[t], c)
		} else {
			challMap[t] = append(challMap[t], c)
		}
		if !slices.Contains(categories, t) {
			categories = append(categories, t)
		}
	}
	sort.Strings(categories)

	rootTree := buildTree(categories, challMap)
	fmt.Println(rootTree)
	return nil
}

func Attempt(id int, url, token string) error {
	wd, _ := os.Getwd()
	chall, err := api.GetChallenge(id, url, token)
	if err != nil {
		return err
	}
	os.Mkdir(wd+"/"+chall.Data.Name, 0777)
	os.Chdir(wd + "/" + chall.Data.Name)

	// Download Files
	for _, file := range chall.Data.Files {
		resp, err := http.Get(url + file)
		if err != nil {
			return err
		}
		fileBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		f, err := os.Create(strings.Split(strings.Split(file, "/")[3], "?")[0])
		if err != nil {
			return err
		}
		defer f.Close()
		f.Write(fileBytes)
	}
	return nil
}

func DisplayTeamInfo(url, token string) error {
	topThree, err := api.GetTopThree(url, token)
	if err != nil {
		return err
	}
	team, err := api.GetTeamInfo(url, token)
	if err != nil {
		return err
	}

	scoreboard := list.New(
		styles.Header.MarginLeft(0).Render("ScoreBoard: "),
		list.New(
			"🥇 "+styles.First.Render(topThree.Data.Num1.Name)+styles.First.Render(strconv.Itoa(topThree.Data.Num1.Score), "points"),
			"🥈 "+styles.Second.Render(topThree.Data.Num2.Name)+styles.Second.Render(strconv.Itoa(topThree.Data.Num2.Score), "points"),
			"🥉 "+styles.Third.Render(topThree.Data.Num3.Name)+styles.Third.Render(strconv.Itoa(topThree.Data.Num3.Score), "points"),
		).Enumerator(list.Roman),
	)
	stats := list.New(
		styles.Header.MarginLeft(0).Render("Team Stats:"),
		list.New(fmt.Sprintf("ID: %d", team.Data.ID)),
		list.New(fmt.Sprintf("Score: %d", team.Data.Score)),
		list.New(fmt.Sprintf("Place: %s", team.Data.Place)),
	)
	fmt.Print(scoreboard, "\n\n")
	fmt.Println(stats)
	return nil
}
