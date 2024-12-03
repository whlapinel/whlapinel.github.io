package data

import "gh_static_portfolio/cmd/domain"

type blogRepo struct {
}

func NewBlogRepo() domain.BlogRepository {
	return &blogRepo{}
}

func (r *blogRepo) GetAll() ([]*domain.Blog, error) {
	return blogs, nil
}

var blogs = []*domain.Blog{
	domain.NewBlog("Blog 1", "This is the first blog. It is a very long blog. It is so long that it will wrap around to the next line."),
	domain.NewBlog("Blog 1", "This is the first blog. It is a very long blog. It is so long that it will wrap around to the next line."),
	domain.NewBlog("Blog 2", "This is the second blog. It is a very long blog. It is so long that it will wrap around to the next line."),
	domain.NewBlog("Blog 3", "This is the third blog. It is a very long blog. It is so long that it will wrap around to the next line."),
	domain.NewBlog("Blog 4", "This is the fourth blog. It is a very long blog. It is so long that it will wrap around to the next line."),
	domain.NewBlog("Blog 5", "This is the fifth blog. It is a very long blog. It is so long that it will wrap around to the next line."),
}
