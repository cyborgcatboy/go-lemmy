package types

type PasswordResetRequest struct {
	ID             int32     `json:"id" url:"id,omitempty"`
	TokenEncrypted string    `json:"token_encrypted" url:"token_encrypted,omitempty"`
	Published      LemmyTime `json:"published" url:"published,omitempty"`
	LocalUserID    int       `json:"local_user_id" url:"local_user_id,omitempty"`
}
type PasswordResetRequestForm struct {
	LocalUserID    int    `json:"local_user_id" url:"local_user_id,omitempty"`
	TokenEncrypted string `json:"token_encrypted" url:"token_encrypted,omitempty"`
}
