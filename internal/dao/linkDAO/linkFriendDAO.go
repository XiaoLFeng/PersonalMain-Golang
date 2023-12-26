package linkDAO

import (
	"PersonalMain/internal/model/do"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

// GetBlogInformation
//
// 获取友链
func GetBlogInformation() *[]do.BlogDO {
	// 数据库读取信息
	var userDO []do.BlogDO
	result, err := g.Model("xf_blog").All()
	if err == nil {
		if !result.IsEmpty() {
			_ = result.Structs(&userDO)
			g.Log().Cat("Database").Cat("Blog").Notice(context.Background(), "xf_blog 数据表数据提取成功")
			return &userDO
		} else {
			g.Log().Cat("Database").Cat("Blog").Notice(context.Background(), "xf_blog 数据表中没有博客友链相关信息")
			return nil
		}
	} else {
		g.Log().Cat("Database").Cat("Blog").Error(context.Background(), err.Error())
		return nil
	}
}

// GetSort
//
// 获取分类
func GetSort() *[]do.SortDO {
	// 数据库读取信息
	var sortDO []do.SortDO
	result, err := g.Model("xf_blog_sort").OrderAsc("sort").All()
	if err == nil {
		if !result.IsEmpty() {
			_ = result.Structs(&sortDO)
			g.Log().Cat("Database").Cat("Sort").Notice(context.Background(), "xf_sort 数据表数据提取成功")
			return &sortDO
		} else {
			g.Log().Cat("Database").Cat("Sort").Notice(context.Background(), "xf_sort 数据表中没有分类相关信息")
			return nil
		}
	} else {
		g.Log().Cat("Database").Cat("Sort").Error(context.Background(), err.Error())
		return nil
	}
}

// GetBlogForSort
//
// 博客内容归类
func GetBlogForSort(sortDO *[]do.SortDO) *[]do.SortDO {
	// 数据库读取信息
	result, err := g.Model("xf_blog").All()
	if err == nil {
		// 检查数据是否为空
		if result.IsEmpty() {
			g.Log().Cat("Database").Cat("Blog").Notice(context.Background(), "xf_blog 数据表中没有博客相关信息")
			return nil
		}
		// 获取数据
		var blogDO []do.BlogDO
		_ = result.Structs(&blogDO)
		for _, sort := range *sortDO {
			var newBlogsDO *[]do.BlogDO
			for _, blog := range blogDO {
				if *sort.Id == blog.BlogLocation {
					if newBlogsDO == nil {
						newBlogsDO = &[]do.BlogDO{blog}
					} else {
						*newBlogsDO = append(*newBlogsDO, blog)
					}
				}
			}
			(*sortDO)[*sort.Id-1].Blogs = newBlogsDO
		}
		g.Log().Cat("Database").Cat("Blog").Notice(context.Background(), "xf_blog 数据表数据提取成功")
		return sortDO
	} else {
		g.Log().Cat("Database").Cat("Blog").Error(context.Background(), err.Error())
		return nil
	}
}

// GetBlogForName
//
// 检查是否已存在此博客
func GetBlogForName(linkName string) *do.BlogDO {
	// 数据库读取信息
	var blogDO do.BlogDO
	result, err := g.Model("xf_blog").Where("blog_title = ?", linkName).One()
	if err == nil {
		if !result.IsEmpty() {
			_ = result.Struct(&blogDO)
			g.Log().Cat("Database").Cat("Blog").Notice(context.Background(), "xf_blog 数据表成功提取", linkName, "博客信息")
			return &blogDO
		} else {
			g.Log().Cat("Database").Cat("Blog").Notice(context.Background(), "xf_blog 数据表中没有", linkName, "博客友链相关信息")
			return nil
		}
	} else {
		g.Log().Cat("Database").Cat("Blog").Error(context.Background(), err.Error())
		return nil
	}
}

// GetBlogForDomain
//
// 检查是否存在相似链接
func GetBlogForDomain(domain string) *do.BlogDO {
	// 数据库读取信息
	var blogDO do.BlogDO
	result, err := g.Model("xf_blog").Where("blog_url like ?", "%"+domain+"%").One()
	if err == nil {
		if !result.IsEmpty() {
			_ = result.Struct(&blogDO)
			g.Log().Cat("Database").Cat("Blog").Notice(context.Background(), "xf_blog 数据表成功提取", blogDO.BlogTitle, "博客信息")
			return &blogDO
		} else {
			g.Log().Cat("Database").Cat("Blog").Notice(context.Background(), "xf_blog 数据表中没有查找到", domain, "相似的信息")
			return nil
		}
	} else {
		g.Log().Cat("Database").Cat("Blog").Error(context.Background(), err.Error())
		return nil
	}
}

// CreateBlog
//
// 创建博客
func CreateBlog(newBlogDO do.BlogDO) bool {
	// 数据库读取信息
	_, err := g.Model("xf_blog").Data(newBlogDO).Insert()
	if err == nil {
		g.Log().Cat("Database").Cat("Blog").Notice(context.Background(), "xf_blog 数据表成功创建", newBlogDO.BlogTitle, "博客信息")
		return true
	} else {
		g.Log().Cat("Database").Cat("Blog").Error(context.Background(), err.Error())
		return false
	}
}
