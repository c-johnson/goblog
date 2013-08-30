package golang_blog

import (
	"encoding/json"
	"fmt"
	"github.com/russross/blackfriday"
	"hash/crc64"
	"io/ioutil"
	"path"
	"time"
)

type post struct {
	Title       string
	Shortname   string
	DateDisplay time.Time
	DateCreated time.Time
	DateUpdated time.Time
	Checksum    int
	Tags        []string
}

type Generator struct {
	rootPath string
	list     []post
}

// func Sync(g Generator) {
// 	srcPath := path.Join(g.rootPath, "src/")
// 	files, err := ioutil.ReadDir(srcPath)
// 	if err == nil {
// 		for _, fileInfo := range files {

// 		}
// 	}
// }

type test struct {
	date time.Time
}

func testTime() {

}

func saveList() {

	newPost := post{Title: "The Man Who Saved My Life", Shortname: "cigarettes", DateCreated: time.Now()}
	jsonPost, err := json.Marshal(newPost)
	if err == nil {
		fmt.Println("rawjson = " + string(jsonPost))
	}
}

// func loadList(g *Generator) {
// 	listPath := path.Join(g.rootPath, "list.json")
// 	file, err := ioutil.ReadFile(g.rootPath)
// 	if err == nil {

// 	}
// }

func Compile(root string, shortname string) string {
	path := path.Join(root, "src/"+shortname+".md")

	file, err := ioutil.ReadFile(path)

	checksum := crc64.Checksum(file, crc64.MakeTable(crc64.ISO))

	saveList()

	fmt.Println(checksum)

	if err == nil {
		return string(blackfriday.MarkdownBasic(file))
	} else {
		return ""
	}
}
