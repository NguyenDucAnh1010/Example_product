package dto

type User struct {
	ID           string `bson:"_id,omitempty"`
	Username     string `bson:"username"`
	PasswordHash string `bson:"password_hash"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
