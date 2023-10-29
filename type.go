package berkatbepkg

import (
	"encoding/json"
	"fmt"
)

type Article struct {
	Title string `json:"title" bson:"title"`
	Tags  string `json:"tags" bson:"tags"`
}

type VideoArticle struct {
	Title string `json:"title" bson:"title"`
	Tags  string `json:"tags" bson:"tags"`
	Video string `json:"video" bson:"video"`
}

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role,omitempty" bson:"role,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

// NewArticle adalah fungsi pembuat untuk membuat instance baru dari Article
func NewArticle(title, tags string) *Article {
	return &Article{
		Title: title,
		Tags:  tags,
	}
}

func NewVideoArticle(title, tags, video string) *VideoArticle {
	return &VideoArticle{
		Title: title,
		Tags:  tags,
		Video: video,
	}
}

// ToJSON mengonversi objek Article ke format JSON
func (a *Article) ToJSON() (string, error) {
	jsonData, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// PrintInfo mencetak informasi artikel ke layar
func (a *Article) PrintInfo() {
	fmt.Printf("Title: %s\nTags: %s\n", a.Title, a.Tags)
}
