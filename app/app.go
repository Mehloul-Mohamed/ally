package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/Mehloul-Mohamed/ally/api"
	"github.com/Mehloul-Mohamed/ally/styles"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
)

func ParseChallJson(bytes []byte) (*api.CtfdChallListResponse, error) {
	var challs api.CtfdChallListResponse
	err := json.Unmarshal(bytes, &challs)
	if err != nil {
		return nil, err
	}
	return &challs, nil
}

func FetchChallList(url string, token string) ([]byte, error) {
	// if no credentials are provided, get them from the credentials file
	if url == "" && token == "" {
		wd, _ := os.Getwd()
		bytes, err := os.ReadFile(wd + "/credentials.txt")
		if err != nil {
			return nil, err
		}
		slice := strings.Split(string(bytes), "\n")
		url = slice[0]
		token = slice[1]
	}
	bytes, err := api.GetChallList(url, token)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

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
			// This should be done in a simpler manner, but lipgloss is acting weirdly so I'm stuck with this
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
	d := home + "/ctf/" + name
	err := os.Mkdir(d, 0777)
	if err != nil {
		return err
	}
	os.Chdir(d)

	// Fetch challenge list
	_, err = FetchChallList(url, token)
	if err != nil {
		return err
	}

	// Store credentials
	f, err := os.Create("credentials.txt")
	if err != nil {
		return err
	}

	defer f.Close()
	f.WriteString(url + "\n" + token)
	return nil
}

func DisplayChallList() error {
	bytes, err := FetchChallList("", "")
	if err != nil {
		return err
	}
	challs, err := ParseChallJson(bytes)
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

func Attempt(id int) error {
	wd, _ := os.Getwd()
	bytes, err := os.ReadFile(wd + "/credentials.txt")
	if err != nil {
		return err
	}

	slice := strings.Split(string(bytes), "\n")
	url := slice[0]
	token := slice[1]

	chall, err := api.GetChallenge(id, url, token)
	if err != nil {
		return err
	}
	os.Mkdir(wd+"/"+chall.Data.Name, 0777)
	os.Chdir(wd + "/" + chall.Data.Name)

	for _, file := range chall.Data.Files {
		resp, err := http.Get(url + file)
		if err != nil {
			return err
		}
		fileBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		f, err := os.Create(fmt.Sprintf("%s", strings.Split(strings.Split(file, "/")[3], "?")[0]))
		if err != nil {
			return err
		}
		defer f.Close()
		f.Write(fileBytes)
	}
	return nil
}
