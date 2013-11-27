package golang_blog

import (
	"encoding/json"
	"fmt"
	"github.com/russross/blackfriday"
	"hash/crc64"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

type post struct {
	Title       string
	Shortname   string
	DateDisplay time.Time
	DateCreated time.Time
	DateUpdated time.Time
	Public      bool
	Checksum    int
	Tags        []string
}

type Generator struct {
	rootPath string
	list     []post
}

type test struct {
	date time.Time
}

func Generate() {
	initialize()
	manifest := GetManifest(os.Getenv("BLOG_SRC"))
	SaveManifest(os.Getenv("BLOG_POSTS"), strings.Join(manifest, "\n"))
	GenerateFiles(os.Getenv("BLOG_SRC"), os.Getenv("BLOG_POSTS"))
}

func initialize() {
	if os.Getenv("BLOG_SRC") == "" {
		os.Setenv("BLOG_SRC", "~/blog_posts")
	}
	if os.Getenv("BLOG_POSTS") == "" {
		os.Setenv("BLOG_POSTS", "$GOPATH/web/posts")
	}
}

func SaveManifest(posts_target string, manifestJson string) {
	os.MkdirAll(posts_target, os.ModePerm)
	manifestFile, _ := os.Create(path.Join(posts_target, "manifest.json"))
	manifestFile.WriteString(manifestJson)
}

func GetManifest(posts_src string) []string {
	postsData, _ := ioutil.ReadDir(posts_src)

	manifestJson := make([]string, 1)

	for _, postData := range postsData {
		newPost := post{Shortname: Shortname(postData.Name()), DateCreated: time.Now(), Public: false}

		jsonPost, err := json.Marshal(newPost)
		if err == nil {
			manifestJson = append(manifestJson, string(jsonPost))
		}
	}

	return manifestJson
}

func GenerateFiles(posts_src string, posts_target string) {
	postsData, _ := ioutil.ReadDir(posts_src)

	outDir := path.Join(posts_target, "out")

	os.MkdirAll(outDir, os.ModePerm)
	fmt.Println("outdir = ", outDir)

	for _, postData := range postsData {
		fileName := path.Join(posts_src, postData.Name())
		file, err := ioutil.ReadFile(fileName)

		if err == nil {
			fileHtml := string(blackfriday.MarkdownBasic(file))
			outPath := path.Join(outDir, Shortname(postData.Name())+".html")
			err := ioutil.WriteFile(outPath, []byte(fileHtml), os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	}
}

func Shortname(str string) string {
	return strings.Split(str, ".")[0]
}

func Compile(root string, shortname string) string {
	path := path.Join(root, shortname+".md")

	file, err := ioutil.ReadFile(path)

	checksum := crc64.Checksum(file, crc64.MakeTable(crc64.ISO))

	fmt.Println(checksum)

	if err == nil {
		return string(blackfriday.MarkdownBasic(file))
	} else {
		return ""
	}
}
