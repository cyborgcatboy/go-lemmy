//  Source: lemmy/crates/db_schema/src/source/email_verification.rs
// Code generated by go.arsenm.dev/go-lemmy/cmd/gen (struct generator). DO NOT EDIT.

package types

type EmailVerification struct {
	ID               int32     `json:"id" url:"id,omitempty"`
	LocalUserID      int       `json:"local_user_id" url:"local_user_id,omitempty"`
	Email            string    `json:"email" url:"email,omitempty"`
	VerificationCode string    `json:"verification_code" url:"verification_code,omitempty"`
	Published        LemmyTime `json:"published" url:"published,omitempty"`
}
type EmailVerificationForm struct {
	LocalUserID       int    `json:"local_user_id" url:"local_user_id,omitempty"`
	Email             string `json:"email" url:"email,omitempty"`
	VerificationToken string `json:"verification_token" url:"verification_token,omitempty"`
}
