package golang_blog

import (
	"os"
)

func init() {
	Initialize()
}

var (
	BLOG_SRC    = os.Getenv("BLOG_SRC")
	BLOG_TARGET = os.Getenv("BLOG_TARGET")
)

func Initialize() {
	if os.Getenv("BLOG_SRC") == "" {
		os.Setenv("BLOG_SRC", "~/blog_posts")
	}
	if os.Getenv("BLOG_TARGET") == "" {
		os.Setenv("BLOG_TARGET", "$GOPATH/web/posts")
	}
}
