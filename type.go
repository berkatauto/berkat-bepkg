package berkatbepkg

type RandomNumber struct {
	random int
}

type articleID struct {
	randomArticleID int
}

type Article struct {
	Author   string  `json:"author" bson:"author"`
	Title    string  `json:"title" bson:"title"`
	Category string  `json:"category" bson:"category"`
	Tags     Tags    `json:"tags" bson:"tags"`
	Content  Content `json:"contect" bson:"content"`
	// Date     time.Time `json:"date" bson:"date"`
}

type Tags struct {
	Tag string `json:"tag" bson:"tag"`
}

type Content struct {
	// ImageHeader  base64.Encoding `bson:"image_encode"`
	Paragraph    string `json:"paragraph" bson:"paragraph"`
	VideoContent string `json:"video_link" bson:"video_link"` // If available, the video will automatically declared to be embedded.
}

// type VideoArticle struct {
// 	Author   string `json:"author" bson:"author"`
// 	Title    string `json:"title" bson:"title"`
// 	Category string `json:"category" bson:"category"`
// 	Tags     string `json:"tags" bson:"tags"`
// 	Video    string `json:"video" bson:"video"`
// }

type User struct {
	UserID        RandomNumber `json:"user_id" bson:"user_id"`
	Fullname      string       `json:"fullname" bson:"fullname"`
	Username      string       `json:"username" bson:"username"`
	Password      string       `json:"password" bson:"password"`
	JournalStatus bool         `json:"journal_bool" bson:"journal_bool"`
	Token         string       `json:"token,omitempty" bson:"token,omitempty"`
	Role          string       `json:"role,omitempty" bson:"role,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

/*
type Properties struct {
	Name string `json:"name" bson:"name"`
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
*/
