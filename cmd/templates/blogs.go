package templates

import (
	"gh_static_portfolio/cmd/domain"

	"github.com/a-h/templ"
)

type blogsPage struct {
	blogs []*domain.Blog
}

func NewBlogsPage(blogs []*domain.Blog) Templifier {
	return &blogsPage{
		blogs: blogs,
	}
}

func (p *blogsPage) Templify() templ.Component {
	return BlogsListComponent(p.blogs)
}

func (p *blogsPage) Title() string {
	return "Blogs"
}

func (p *blogsPage) Directory() string {
	return rootDir
}

type blogPage struct {
	blog *domain.Blog
}

func NewBlogPage(blog *domain.Blog) Templifier {
	return &blogPage{
		blog: blog,
	}
}

func (p *blogPage) Templify() templ.Component {
	return BlogComponent(p.blog)
}

func (p *blogPage) Title() string {
	return p.blog.Title()
}

func (p *blogPage) Directory() string {
	return blogsDir
}
