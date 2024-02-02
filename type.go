package berkatbepkg

import "time"

type Article struct {
	ID       string    `json:"article_id,omitempty" bson:"article_id,omitempty"`
	Author   string    `json:"author" bson:"author"`
	Title    string    `json:"title" bson:"title"`
	Category string    `json:"category" bson:"category"`
	Tags     string    `json:"tags" bson:"tags"`
	Content  Content   `json:"content" bson:"content"`
	Date     time.Time `json:"date" bson:"date"`
}

type Content struct {
	Paragraph string `json:"paragraph" bson:"paragraph"`
	Image     string `json:"image" bson:"image"`
}

type User struct {
	Fullname      string `json:"fullname" bson:"fullname"`
	Username      string `json:"username" bson:"username"`
	Password      string `json:"password" bson:"password"`
	NoHP          string `json:"no_hp,omitempty" bson:"no_hp,omitempty"`
	JournalStatus string `json:"journal_bool" bson:"journal_bool"`
	Token         string `json:"token,omitempty" bson:"token,omitempty"`
	Role          string `json:"role,omitempty" bson:"role,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type RegisterInfo struct {
	Message       string `json:"message,omitempty" bson:"message,omitempty"`
	Status        bool   `json:"status" bson:"status"`
	Fullname      string `json:"fullname" bson:"fullname"`
	Username      string `json:"username" bson:"username"`
	Password      string `json:"password" bson:"password"`
	JournalStatus string `json:"journal_bool" bson:"journal_bool"`
	Role          string `json:"role,omitempty" bson:"role,omitempty"`
	NoHP          string `json:"no_hp,omitempty" bson:"no_hp,omitempty"`
}

type ArticleInfo struct {
	Message  string    `json:"message,omitempty" bson:"message,omitempty"`
	Author   string    `json:"author" bson:"author"`
	Title    string    `json:"title" bson:"title"`
	Category string    `json:"category" bson:"category"`
	Tags     string    `json:"tags" bson:"tags"`
	Content  []Content `json:"content" bson:"content"`
	Date     string    `json:"date" bson:"date"`
}

type whatsauth struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
