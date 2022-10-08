package dao

import (
	"atao-go-blog/entity"
	"atao-go-blog/utils"
	"atao-go-blog/vo"
	"fmt"
	"log"
	"strings"
)

/* 博客数据层*/
// 查询的字段
var searchFeild string = " id ,title ,photo,view ,tags ,catalog ,uploadTime,content,updateTime "

// 查询博客总数
func GetBlogCount() int {
	row := utils.DBConn.QueryRow("select count(1) from blog where del_flag=0")
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var count int
	row.Scan(&count)
	return count
}

/// tags 总数
func GetBlogByTagsCount(tag string) int {
	row := utils.DBConn.QueryRow("SELECT COUNT(1) FROM blog WHERE  del_flag=0 and tags LIKE ?", "%"+tag+"%")
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var count int
	row.Scan(&count)
	return count
}

// 分类查询文章总数
func GetBlogByCategoryCount(catelog string) int {
	row := utils.DBConn.QueryRow("SELECT COUNT(1) FROM blog WHERE del_flag=0 and catalog = ? ", catelog)
	if row.Err() != nil {
		log.Println("GetBlogByCategoryCount: ", row.Err())
	}
	var count int
	row.Scan(&count)
	return count
}

// // 查询博客所有内容
func GetAllBlog(page, pageSize int) ([]entity.Blog, error) {

	rows, err := utils.DBConn.Query("SELECT"+searchFeild+"FROM blog where  del_flag=0  ORDER BY uploadTime DESC  LIMIT ?, ?", page, pageSize)
	if err != nil {
		log.Println("查询出错", err)
		return nil, err
	}

	var blogs []entity.Blog
	for rows.Next() {
		var blog entity.Blog

		err := rows.Scan(&blog.Id, &blog.Title, &blog.Photo, &blog.View, &blog.Tags, &blog.Catalog, &blog.UpLoadTime, &blog.Content, &blog.UpdateTime)
		if err != nil {
			log.Println("blog取值出错", err)
			return nil, err
		}

		blogs = append(blogs, blog)
	}
	return blogs, nil
}

//
// 分页查询 -查询博客所有内容-条件查询
func GetAllBlogByTag(page, pageSize int, condition string) ([]entity.Blog, error) {
	rows, err := utils.DBConn.Query("SELECT"+searchFeild+"FROM blog WHERE del_flag= 0 and tags LIKE ? ORDER BY uploadTime DESC  LIMIT ?,?", "%"+condition+"%", page, pageSize)
	if err != nil {
		log.Println("查询出错", err.Error())
		return nil, err
	}

	var blogs []entity.Blog
	for rows.Next() {
		var blog entity.Blog

		err := rows.Scan(&blog.Id, &blog.Title, &blog.Photo, &blog.View, &blog.Tags, &blog.Catalog, &blog.UpLoadTime, &blog.Content, &blog.UpdateTime)
		if err != nil {
			log.Println("blog取值出错", err)
			return nil, err
		}

		blogs = append(blogs, blog)
	}
	return blogs, nil
}

// 分类查询
func GetAllBlogByCategory(page, pageSize int, condition string) ([]entity.Blog, error) {
	rows, err := utils.DBConn.Query("SELECT "+searchFeild+" FROM blog WHERE del_flag=0 and catalog = ? ORDER BY uploadTime DESC  LIMIT ?, ? ", condition, page, pageSize)
	if err != nil {
		log.Println("GetAllBlogByCategory查询出错:", err)
		return nil, err
	}

	var blogs []entity.Blog
	for rows.Next() {
		var blog entity.Blog

		err := rows.Scan(&blog.Id, &blog.Title, &blog.Photo, &blog.View, &blog.Tags, &blog.Catalog, &blog.UpLoadTime, &blog.Content, &blog.UpdateTime)
		if err != nil {
			log.Println("blog取值出错", err)
			return nil, err
		}

		blogs = append(blogs, blog)
	}
	return blogs, nil
}

// 归档查询
func GetArchIvesBlog(page, pageSize int) ([]entity.Blog, error) {
	rows, err := utils.DBConn.Query("SELECT "+searchFeild+" , DATE_FORMAT(b.uploadTime, '%Y-%c')AS YEAR   FROM blog b where del_flag=0  ORDER BY DATE_FORMAT(b.uploadTime, '%Y-%m-%d') DESC LIMIT ?,?", page, pageSize)
	if err != nil {
		log.Println("查询出错", err)
		return nil, err
	}

	var blogs []entity.Blog
	for rows.Next() {
		var blog entity.Blog

		err := rows.Scan(&blog.Id, &blog.Title, &blog.Photo, &blog.View, &blog.Tags, &blog.Catalog, &blog.UpLoadTime, &blog.Content, &blog.UpdateTime, &blog.YearMonth)
		if err != nil {
			log.Println("blog取值出错", err)
			return nil, err
		}

		blogs = append(blogs, blog)
	}
	return blogs, nil
}

func GetArchIveYearMonth() ([]string, error) {
	rows, err := utils.DBConn.Query("SELECT DATE_FORMAT(uploadTime,'%Y-%c')  FROM blog where del_flag=0  GROUP BY DATE_FORMAT(uploadTime,'%Y-%m') ORDER BY DATE_FORMAT(uploadTime, '%Y-%m-%d') DESC")
	if err != nil {
		log.Println("查询出错", err)
		return nil, err
	}

	var result []string
	for rows.Next() {

		var yearMonth string
		err := rows.Scan(&yearMonth)
		if err != nil {
			log.Println("blog取值出错", err)
			return nil, err
		}

		result = append(result, yearMonth)
	}
	return result, nil
}

/// end

// 根据id 查询博客
func GetBlogById(id int) (entity.Blog, error) {
	rows := utils.DBConn.QueryRow("SELECT "+searchFeild+" FROM blog where del_flag=0 and id= ?", id)
	if rows.Err() != nil {
		log.Println("查询出错", rows.Err())
		return entity.Blog{}, rows.Err()
	}

	var blog entity.Blog
	err := rows.Scan(&blog.Id, &blog.Title, &blog.Photo, &blog.View, &blog.Tags, &blog.Catalog, &blog.UpLoadTime, &blog.Content, &blog.UpdateTime)
	if err != nil {
		log.Println("blog取值出错", err)
		return entity.Blog{}, err
	}
	return blog, nil
}

// 根据id 批量修改blog 访问数据
func UpdateBlogViewbyId(fields, ids string) (int, error) {
	r, err := utils.DBConn.Exec("update blog set view = case id " + fields + " end " + "where id in(" + ids + ")")

	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	i, _ := r.RowsAffected()
	return int(i), nil
}

// 搜索文章根据文字
func GetBlogBytext(text string) ([]vo.SearchBlogVo, error) {

	fmt.Println(len(text))
	s := strings.TrimSpace(text)
	if len(s) > 0 {
		text = text + "*"
	}
	sql := "SELECT id,title ,content FROM blog WHERE MATCH(title,content) AGAINST( '" + text + "' IN BOOLEAN MODE)"
	rows, err := utils.DBConn.Query(sql)

	if err != nil {
		return nil, err
	}
	var blogItems []vo.SearchBlogVo
	for rows.Next() {
		var blogItem vo.SearchBlogVo

		err2 := rows.Scan(&blogItem.Id, &blogItem.Title, &blogItem.Content)
		if err2 != nil {
			return nil, err2
		}

		if len(blogItem.Content) > 150 {
			blogItem.Content = blogItem.Content[0:150]
			blogItem.Content = strings.Replace(string(blogItem.Content)+".....", "#", "", -1)
		}
		// 补充正则替换

		blogItems = append(blogItems, blogItem)
	}
	return blogItems, nil

}

func GetHotArticle() []vo.HotArticle {
	// SELECT id , title ,uploadTime FROM blog ORDER BY VIEW  DESC LIMIT 0,3
	rows, err := utils.DBConn.Query("SELECT id , title , photo,uploadTime FROM blog ORDER BY VIEW  DESC LIMIT 0,3")

	if err != nil {
		panic(err)
	}

	var hotArticles []vo.HotArticle
	for rows.Next() {
		var hotArticle vo.HotArticle
		rows.Scan(&hotArticle.Id, &hotArticle.Title, &hotArticle.Photo, &hotArticle.CreateDate)
		hotArticles = append(hotArticles, hotArticle)
	}

	return hotArticles

}
