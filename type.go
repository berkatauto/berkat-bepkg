package berkatbepkg

import "fmt"

type Article struct {
	Title string `json:"title" bson:"title"`
	Tags  string `json:"tags" bson:"tags"`
}

// NewArticle adalah fungsi pembuat untuk membuat instance baru dari Article
func NewArticle(title, tags string) *Article {
    return &Article{
        Title: title,
        Tags: tags,
    }
}

// PrintInfo adalah metode untuk mencetak informasi artikel ke layar
func (a *Article) PrintInfo() {
    fmt.Printf("Title: %s\nContent: %s\n", a.Title, a.Content)
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
