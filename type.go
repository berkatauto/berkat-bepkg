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
		Tags:  tags,
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
