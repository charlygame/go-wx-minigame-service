package user

type User struct {
	ID         string `json:"id" bson:"_id, omitempty"`
	Username   string `json:"username" bson:"username, omitempty"`
	SessionKey string `json:"session_key" bson:"session_key, omitempty"`
	OpenID     string `json:"open_id" bson:"open_id, omitempty"`
	Score      int    `json:"score" bson:"score, omitempty"`
}
