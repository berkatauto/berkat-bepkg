package berkatbepkg

import "time"

type articleID struct {
	randomArticleID int
}

type Article struct {
	Author   string    `json:"author" bson:"author"`
	Title    string    `json:"title" bson:"title"`
	Category string    `json:"category" bson:"category"`
	Tags     Tags      `json:"tags" bson:"tags"`
	Content  Content   `json:"content" bson:"content"`
	Date     time.Time `json:"date" bson:"date"`
}

type Tags struct {
	Tag string `json:"tag" bson:"tag"`
}

type Content struct {
	Paragraph    string `json:"paragraph" bson:"paragraph"`
	Image        string `json:"image" bson:"image"`
	VideoContent string `json:"video_link,omiempty" bson:"video_link,omiempty"` // If available, the video will automatically declared to be embedded.
}

type User struct {
	Fullname      string `json:"fullname" bson:"fullname"`
	Username      string `json:"username" bson:"username"`
	Password      string `json:"password" bson:"password"`
	JournalStatus bool   `json:"journal_bool" bson:"journal_bool"`
	Token         string `json:"token,omitempty" bson:"token,omitempty"`
	Role          string `json:"role,omitempty" bson:"role,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}
