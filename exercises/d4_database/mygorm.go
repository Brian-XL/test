package d4_database

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Age       uint8
	Posts     []Post `gorm:"foreignKey:UserID; references:ID"`
	PostCount uint
}

type Post struct {
	gorm.Model
	Title      string
	Author     string
	Content    string
	UserID     *uint
	Comments   []Comment `gorm:"foreignKey:PostID; references:ID"`
	User       *User     `gorm:"foreignKey:UserID; references:ID"`
	HasComment bool
}

type Comment struct {
	gorm.Model
	Author   string
	Content  string
	PostID   *uint
	Parent   *Comment `gorm:"foreignKey:ParentID; references:ID"`
	ParentID *uint
	Post     *Post `gorm:"foreignKey:PostID; references:ID"`
}

func InitializeGorm(db *gorm.DB) {
	err := db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatal("Failed to Automigrate", err)
	}
	comments := []Comment{
		{Author: "William", Content: "Thank you sir! What a great article."},
		{Author: "Jack", Content: "Is it really useful?"},
	}
	posts := []Post{
		{Title: "The thesis of the new gas oil", Author: "Francis", Content: "", Comments: comments},
		{Title: "Nice to meet you", Author: "Francis", Content: "I am here to make friends", Comments: []Comment{
			{Author: "Wendy", Content: "Welcome to our place!"},
			{Author: "Janne", Content: "Hello budy."},
		}},
		{Title: "Hello everyone", Author: "Francis", Content: "This is my new blog"},
	}
	user := User{Name: "Francis", Age: 32, Posts: posts}

	db.Create(&user)
	fmt.Println(user)

	comments = []Comment{
		{Author: "Francis", Content: "Nice to meet you."},
		{Author: "John", Content: "Welcome :)"},
		{Author: "Judy", Content: "Welcome :)"},
	}
	posts = []Post{
		{Title: "Story A", Author: "Jack", Content: "", Comments: comments},
		{Title: "Story B", Author: "Jack", Content: "I am here to make friends"},
	}
	user = User{Name: "Jack", Age: 22, Posts: posts}

	db.Create(&user)
}

func SearchAssociation(db *gorm.DB) {
	var user User
	db.Preload("Posts.Comments").First(&user, "name = ?", "Francis")
	count := db.Model(&user).Association("Posts").Count()
	fmt.Printf("User %s has %d articles.\n", user.Name, count)
	posts := user.Posts
	for _, v := range posts {
		fmt.Println(v.Title, " ", v.Author, " ", v.Content)
		fmt.Printf("   The Post received %d comments.\n", db.Model(&v).Association("Comments").Count())
		fmt.Println("   comments:")
		for _, comment := range v.Comments {
			fmt.Println("     ", comment.Author, "says:", comment.Content)
		}
	}

}

func MostCommentsPost(db *gorm.DB) {
	// type Result struct {
	// 	PostID uint
	// 	Count  int
	// }
	// var result Result
	// db.Table("comments").Select("post_id, COUNT(*) AS count").Group("post_id").Order("count DESC").Limit(1).Scan(&result)
	// fmt.Println("Result:", result)

	// var id uint
	// db.Table("comments").Select("post_id").Group("post_id").Order("COUNT(*) DESC").Limit(1).Scan(&id)
	// fmt.Println("id:", id)

	// var post Post
	// db.Model(post).
	// 	Select("posts.*, COUNT(comments.id) AS comments_count").
	// 	Joins("LEFT JOIN comments ON comments.post_id = posts.id").
	// 	Group("posts.id").
	// 	Order("comments_count DESC").
	// 	Limit(1).
	// 	Preload("User").
	// 	First(&post)

	// fmt.Println(post)

	var post Post
	db.Model(&post).Select("posts.*").Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").Order("COUNT(comments.id) DESC").Limit(1).First(&post)

	fmt.Println(post.Title)
}

// 钩子函数，文章创建时自动更新用户的文章数量统计字段
func (post *Post) AfterCreate(tx *gorm.DB) (err error) {
	var user User
	tx.Model(&user).First(&user, post.UserID)
	//tx.Model(&user).Where("id = ?", post.UserID).First(&user)

	user.PostCount += 1
	tx.Model(&user).Where("id = ?", user.ID).Updates(&user)
	return
}

func CreatePost(db *gorm.DB) {
	var uid uint = 2
	post := Post{Title: "调查", Content: "太平洋", Author: "William", UserID: &uid}
	db.Create(&post)
}

// 钩子函数，评论删除时
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// var post Post
	// tx.Preload("Comments").First(&post, c.PostID)
	// if len(post.Comments) == 0 {
	// 	fmt.Println("无评论")
	// 	tx.Model(&Post{}).Where("id = ?", post.ID).Update("has_comment", false)
	// }

	var count int64
	tx.Model(&c).Where("post_id = ?", c.PostID).Count(&count)
	if count == 0 {
		tx.Model(&Post{}).Where("id = ?", c.PostID).Update("has_comment", false)
	}

	return
}

func DeleteComment(db *gorm.DB) {
	var c []Comment
	db.Model(&Comment{}).Where("post_id = ?", 2).Find(&c)
	for _, v := range c {
		db.Delete(&v)
	}
}
