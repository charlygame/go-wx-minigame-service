package user

type WXAuth struct {
	OpenID     string `json:"openid,omitempty" bson:"open_id,omitempty"`
	SessionKey string `json:"session_key,omitempty" bson:"session_key,omitempty"`
}

type User struct {
	ID           string `json:"id,omitempty" bson:"_id,omitempty"`
	Username     string `json:"username,omitempty" bson:"username,omitempty"`
	Score        int    `json:"score,omitempty" bson:"score,omitempty"`
	WxOpenId     string `json:"wx_open_id,omitempty" bson:"wx_open_id,omitempty"`
	WxSessionKey string `json:"wx_session_key,omitempty" bson:"wx_session_key,omitempty"`
}
