package berkatbepkg

type Article struct {
	Author   string `json:"author" bson:"author"`
	Title    string `json:"title" bson:"title"`
	Category string `json:"category" bson:"category"`
	Tags     string `json:"tags" bson:"tags"`
	Content  string `json:"contect" bson:"content"`
}

type VideoArticle struct {
	Author   string `json:"author" bson:"author"`
	Title    string `json:"title" bson:"title"`
	Category string `json:"category" bson:"category"`
	Tags     string `json:"tags" bson:"tags"`
	Video    string `json:"video" bson:"video"`
}

type User struct {
	Fullname      string `json:"fullname" bson:"fullname"`
	Username      string `json:"username" bson:"username"`
	Email         string `json:"email" bson:"email"`
	Password      string `json:"password" bson:"password"`
	JournalStatus string `json:"journal_bool" bson:"journal_bool"`
	Role          string `json:"role,omitempty" bson:"role,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

/*
package peda

type GeometryPolygon struct {
	Coordinates [][][]float64 `json:"coordinates" bson:"coordinates"`
	Type        string        `json:"type" bson:"type"`
}

type GeometryLineString struct {
	Coordinates [][]float64 `json:"coordinates" bson:"coordinates"`
	Type        string      `json:"type" bson:"type"`
}

type GeometryPoint struct {
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
	Type        string    `json:"type" bson:"type"`
}

type GeoJsonLineString struct {
	Type       string             `json:"type" bson:"type"`
	Properties Properties         `json:"properties" bson:"properties"`
	Geometry   GeometryLineString `json:"geometry" bson:"geometry"`
}

type GeoJsonPolygon struct {
	Type       string          `json:"type" bson:"type"`
	Properties Properties      `json:"properties" bson:"properties"`
	Geometry   GeometryPolygon `json:"geometry" bson:"geometry"`
}

type Geometry struct {
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
	Type        string      `json:"type" bson:"type"`
}
type GeoJson struct {
	Type       string     `json:"type" bson:"type"`
	Properties Properties `json:"properties" bson:"properties"`
	Geometry   Geometry   `json:"geometry" bson:"geometry"`
}

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
