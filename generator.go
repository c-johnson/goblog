package golang_blog

import (
	"bytes"
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

var _ = crc64.MakeTable

func init() {
	if os.Getenv("BLOG_SRC") == "" {

		os.Setenv("BLOG_SRC", "~/blog_posts")
	}
	if os.Getenv("BLOG_POSTS") == "" {
		os.Setenv("BLOG_POSTS", "$GOPATH/web/posts")
	}
}

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

func (p post) String() string {
	return "Post shortname = " + p.Shortname
}

type Generator struct {
	rootPath string
	list     []post
}

type test struct {
	date time.Time
}

func Generate(manifest bool) {
	if manifest {
		SaveManifest()
	}
	GenerateHtml()
}

func PublicPosts() []post {
	manifest := Manifest()
	ret := make([]post, 0)

	for _, post := range manifest {
		if post.Public {
			ret = append(ret, post)
		}
	}
	return ret
}

func GenerateHtml() {
	posts_src := path.Join(os.Getenv("BLOG_SRC"))
	posts_target := path.Join(os.Getenv("BLOG_POSTS"), "out")
	os.MkdirAll(posts_target, os.ModePerm)
	posts := PublicPosts()
	for _, post := range posts {
		if post.Public {
			html := Compile(posts_src, post.Shortname)
			outPath := path.Join(posts_target, post.Shortname+".html")
			err := ioutil.WriteFile(outPath, []byte(html), os.ModePerm)
			if err != nil {
				fmt.Println("Fuck, an err ", err)
			}
		}
	}
}

func ManifestBytes() []byte {
	posts_target := os.Getenv("BLOG_POSTS")
	fullpath := path.Join(posts_target, "manifest.json")
	file, err := ioutil.ReadFile(fullpath)

	if err == nil {
		return file
	}

	return nil
}

func Manifest() []post {
	posts_target := os.Getenv("BLOG_POSTS")
	fullpath := path.Join(posts_target, "manifest.json")
	file, err := ioutil.ReadFile(fullpath)
	var posts = make([]post, 1)

	if err == nil {
		err = json.Unmarshal(file, &posts)
		if err == nil {
			return posts
		} else {
			fmt.Println("err = ", err)
		}
	}

	return nil
}

func WriteManifest(posts_target string, manifest []post) {
	os.MkdirAll(posts_target, os.ModePerm)
	manifestPath := path.Join(posts_target, "manifest.json")

	file, err := os.Create(manifestPath)

	if err == nil {
		manifestJson, err := json.Marshal(manifest)
		if err == nil {
			var b bytes.Buffer
			json.Indent(&b, manifestJson, "", "\t")
			b.WriteTo(file)
		}
	} else {
		fmt.Println(err)
	}
}

func SaveManifest() {
	posts_src := os.Getenv("BLOG_SRC")
	posts_target := os.Getenv("BLOG_POSTS")

	postsData, _ := ioutil.ReadDir(posts_src)

	manifest := make([]post, 1)

	for _, postData := range postsData {
		newPost := post{Shortname: Shortname(postData.Name()), DateCreated: time.Now(), Public: false}
		manifest = append(manifest, newPost)
	}

	WriteManifest(posts_target, manifest)
}

func Shortname(str string) string {
	return strings.Split(str, ".")[0]
}

func Compile(root string, shortname string) string {
	path := path.Join(root, shortname+".md")
	file, err := ioutil.ReadFile(path)

	// checksum := crc64.Checksum(file, crc64.MakeTable(crc64.ISO))
	// fmt.Println(checksum, "awef")

	if err == nil {
		return string(blackfriday.MarkdownBasic(file))
	} else {
		return ""
	}
}
