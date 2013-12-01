package golang_blog

import (
	"fmt"
	"os"
	"path"
)

// var _ = fmt.Println

func init() {
	Initialize()
	fmt.Println(BLOG_SRC)
	fmt.Println(BLOG_WEB)
	fmt.Println(BLOG_TARGET)
}

var (
	BLOG_SRC    = os.Getenv("BLOG_SRC")
	GOPATH      = os.Getenv("GOPATH")
	BLOG_WEB    = path.Join(GOPATH, "web")
	BLOG_TARGET = path.Join(BLOG_WEB, "posts")
)

func Initialize() {
	if os.Getenv("BLOG_SRC") == "" {
		os.Setenv("BLOG_SRC", "~/blog_posts")
	}
}
