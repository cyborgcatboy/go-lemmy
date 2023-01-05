package types

type Person struct {
	ID              int              `json:"id" url:"id,omitempty"`
	Name            string           `json:"name" url:"name,omitempty"`
	DisplayName     Optional[string] `json:"display_name" url:"display_name,omitempty"`
	Avatar          Optional[string] `json:"avatar" url:"avatar,omitempty"`
	Banned          bool             `json:"banned" url:"banned,omitempty"`
	Published       LemmyTime        `json:"published" url:"published,omitempty"`
	Updated         LemmyTime        `json:"updated" url:"updated,omitempty"`
	ActorID         string           `json:"actor_id" url:"actor_id,omitempty"`
	Bio             Optional[string] `json:"bio" url:"bio,omitempty"`
	Local           bool             `json:"local" url:"local,omitempty"`
	PrivateKey      Optional[string] `json:"private_key" url:"private_key,omitempty"`
	PublicKey       string           `json:"public_key" url:"public_key,omitempty"`
	LastRefreshedAt LemmyTime        `json:"last_refreshed_at" url:"last_refreshed_at,omitempty"`
	Banner          Optional[string] `json:"banner" url:"banner,omitempty"`
	Deleted         bool             `json:"deleted" url:"deleted,omitempty"`
	InboxURL        string           `json:"inbox_url" url:"inbox_url,omitempty"`
	SharedInboxURL  Optional[string] `json:"shared_inbox_url" url:"shared_inbox_url,omitempty"`
	MatrixUserID    Optional[string] `json:"matrix_user_id" url:"matrix_user_id,omitempty"`
	Admin           bool             `json:"admin" url:"admin,omitempty"`
	BotAccount      bool             `json:"bot_account" url:"bot_account,omitempty"`
	BanExpires      LemmyTime        `json:"ban_expires" url:"ban_expires,omitempty"`
}
type PersonSafe struct {
	ID             int              `json:"id" url:"id,omitempty"`
	Name           string           `json:"name" url:"name,omitempty"`
	DisplayName    Optional[string] `json:"display_name" url:"display_name,omitempty"`
	Avatar         Optional[string] `json:"avatar" url:"avatar,omitempty"`
	Banned         bool             `json:"banned" url:"banned,omitempty"`
	Published      LemmyTime        `json:"published" url:"published,omitempty"`
	Updated        LemmyTime        `json:"updated" url:"updated,omitempty"`
	ActorID        string           `json:"actor_id" url:"actor_id,omitempty"`
	Bio            Optional[string] `json:"bio" url:"bio,omitempty"`
	Local          bool             `json:"local" url:"local,omitempty"`
	Banner         Optional[string] `json:"banner" url:"banner,omitempty"`
	Deleted        bool             `json:"deleted" url:"deleted,omitempty"`
	InboxURL       string           `json:"inbox_url" url:"inbox_url,omitempty"`
	SharedInboxURL Optional[string] `json:"shared_inbox_url" url:"shared_inbox_url,omitempty"`
	MatrixUserID   Optional[string] `json:"matrix_user_id" url:"matrix_user_id,omitempty"`
	Admin          bool             `json:"admin" url:"admin,omitempty"`
	BotAccount     bool             `json:"bot_account" url:"bot_account,omitempty"`
	BanExpires     LemmyTime        `json:"ban_expires" url:"ban_expires,omitempty"`
}
type PersonAlias1 struct {
	ID              int              `json:"id" url:"id,omitempty"`
	Name            string           `json:"name" url:"name,omitempty"`
	DisplayName     Optional[string] `json:"display_name" url:"display_name,omitempty"`
	Avatar          Optional[string] `json:"avatar" url:"avatar,omitempty"`
	Banned          bool             `json:"banned" url:"banned,omitempty"`
	Published       LemmyTime        `json:"published" url:"published,omitempty"`
	Updated         LemmyTime        `json:"updated" url:"updated,omitempty"`
	ActorID         string           `json:"actor_id" url:"actor_id,omitempty"`
	Bio             Optional[string] `json:"bio" url:"bio,omitempty"`
	Local           bool             `json:"local" url:"local,omitempty"`
	PrivateKey      Optional[string] `json:"private_key" url:"private_key,omitempty"`
	PublicKey       string           `json:"public_key" url:"public_key,omitempty"`
	LastRefreshedAt LemmyTime        `json:"last_refreshed_at" url:"last_refreshed_at,omitempty"`
	Banner          Optional[string] `json:"banner" url:"banner,omitempty"`
	Deleted         bool             `json:"deleted" url:"deleted,omitempty"`
	InboxURL        string           `json:"inbox_url" url:"inbox_url,omitempty"`
	SharedInboxURL  Optional[string] `json:"shared_inbox_url" url:"shared_inbox_url,omitempty"`
	MatrixUserID    Optional[string] `json:"matrix_user_id" url:"matrix_user_id,omitempty"`
	Admin           bool             `json:"admin" url:"admin,omitempty"`
	BotAccount      bool             `json:"bot_account" url:"bot_account,omitempty"`
	BanExpires      LemmyTime        `json:"ban_expires" url:"ban_expires,omitempty"`
}
type PersonSafeAlias1 struct {
	ID             int              `json:"id" url:"id,omitempty"`
	Name           string           `json:"name" url:"name,omitempty"`
	DisplayName    Optional[string] `json:"display_name" url:"display_name,omitempty"`
	Avatar         Optional[string] `json:"avatar" url:"avatar,omitempty"`
	Banned         bool             `json:"banned" url:"banned,omitempty"`
	Published      LemmyTime        `json:"published" url:"published,omitempty"`
	Updated        LemmyTime        `json:"updated" url:"updated,omitempty"`
	ActorID        string           `json:"actor_id" url:"actor_id,omitempty"`
	Bio            Optional[string] `json:"bio" url:"bio,omitempty"`
	Local          bool             `json:"local" url:"local,omitempty"`
	Banner         Optional[string] `json:"banner" url:"banner,omitempty"`
	Deleted        bool             `json:"deleted" url:"deleted,omitempty"`
	InboxURL       string           `json:"inbox_url" url:"inbox_url,omitempty"`
	SharedInboxURL Optional[string] `json:"shared_inbox_url" url:"shared_inbox_url,omitempty"`
	MatrixUserID   Optional[string] `json:"matrix_user_id" url:"matrix_user_id,omitempty"`
	Admin          bool             `json:"admin" url:"admin,omitempty"`
	BotAccount     bool             `json:"bot_account" url:"bot_account,omitempty"`
	BanExpires     LemmyTime        `json:"ban_expires" url:"ban_expires,omitempty"`
}
type PersonAlias2 struct {
	ID              int              `json:"id" url:"id,omitempty"`
	Name            string           `json:"name" url:"name,omitempty"`
	DisplayName     Optional[string] `json:"display_name" url:"display_name,omitempty"`
	Avatar          Optional[string] `json:"avatar" url:"avatar,omitempty"`
	Banned          bool             `json:"banned" url:"banned,omitempty"`
	Published       LemmyTime        `json:"published" url:"published,omitempty"`
	Updated         LemmyTime        `json:"updated" url:"updated,omitempty"`
	ActorID         string           `json:"actor_id" url:"actor_id,omitempty"`
	Bio             Optional[string] `json:"bio" url:"bio,omitempty"`
	Local           bool             `json:"local" url:"local,omitempty"`
	PrivateKey      Optional[string] `json:"private_key" url:"private_key,omitempty"`
	PublicKey       string           `json:"public_key" url:"public_key,omitempty"`
	LastRefreshedAt LemmyTime        `json:"last_refreshed_at" url:"last_refreshed_at,omitempty"`
	Banner          Optional[string] `json:"banner" url:"banner,omitempty"`
	Deleted         bool             `json:"deleted" url:"deleted,omitempty"`
	InboxURL        string           `json:"inbox_url" url:"inbox_url,omitempty"`
	SharedInboxURL  Optional[string] `json:"shared_inbox_url" url:"shared_inbox_url,omitempty"`
	MatrixUserID    Optional[string] `json:"matrix_user_id" url:"matrix_user_id,omitempty"`
	Admin           bool             `json:"admin" url:"admin,omitempty"`
	BotAccount      bool             `json:"bot_account" url:"bot_account,omitempty"`
	BanExpires      LemmyTime        `json:"ban_expires" url:"ban_expires,omitempty"`
}
type PersonSafeAlias2 struct {
	ID             int              `json:"id" url:"id,omitempty"`
	Name           string           `json:"name" url:"name,omitempty"`
	DisplayName    Optional[string] `json:"display_name" url:"display_name,omitempty"`
	Avatar         Optional[string] `json:"avatar" url:"avatar,omitempty"`
	Banned         bool             `json:"banned" url:"banned,omitempty"`
	Published      LemmyTime        `json:"published" url:"published,omitempty"`
	Updated        LemmyTime        `json:"updated" url:"updated,omitempty"`
	ActorID        string           `json:"actor_id" url:"actor_id,omitempty"`
	Bio            Optional[string] `json:"bio" url:"bio,omitempty"`
	Local          bool             `json:"local" url:"local,omitempty"`
	Banner         Optional[string] `json:"banner" url:"banner,omitempty"`
	Deleted        bool             `json:"deleted" url:"deleted,omitempty"`
	InboxURL       string           `json:"inbox_url" url:"inbox_url,omitempty"`
	SharedInboxURL Optional[string] `json:"shared_inbox_url" url:"shared_inbox_url,omitempty"`
	MatrixUserID   Optional[string] `json:"matrix_user_id" url:"matrix_user_id,omitempty"`
	Admin          bool             `json:"admin" url:"admin,omitempty"`
	BotAccount     bool             `json:"bot_account" url:"bot_account,omitempty"`
	BanExpires     LemmyTime        `json:"ban_expires" url:"ban_expires,omitempty"`
}
type PersonForm struct {
	Name            string                     `json:"name" url:"name,omitempty"`
	DisplayName     Optional[Optional[string]] `json:"display_name" url:"display_name,omitempty"`
	Avatar          Optional[Optional[string]] `json:"avatar" url:"avatar,omitempty"`
	Banned          Optional[bool]             `json:"banned" url:"banned,omitempty"`
	Published       LemmyTime                  `json:"published" url:"published,omitempty"`
	Updated         LemmyTime                  `json:"updated" url:"updated,omitempty"`
	ActorID         Optional[string]           `json:"actor_id" url:"actor_id,omitempty"`
	Bio             Optional[Optional[string]] `json:"bio" url:"bio,omitempty"`
	Local           Optional[bool]             `json:"local" url:"local,omitempty"`
	PrivateKey      Optional[Optional[string]] `json:"private_key" url:"private_key,omitempty"`
	PublicKey       Optional[string]           `json:"public_key" url:"public_key,omitempty"`
	LastRefreshedAt LemmyTime                  `json:"last_refreshed_at" url:"last_refreshed_at,omitempty"`
	Banner          Optional[Optional[string]] `json:"banner" url:"banner,omitempty"`
	Deleted         Optional[bool]             `json:"deleted" url:"deleted,omitempty"`
	InboxURL        Optional[string]           `json:"inbox_url" url:"inbox_url,omitempty"`
	SharedInboxURL  Optional[Optional[string]] `json:"shared_inbox_url" url:"shared_inbox_url,omitempty"`
	MatrixUserID    Optional[Optional[string]] `json:"matrix_user_id" url:"matrix_user_id,omitempty"`
	Admin           Optional[bool]             `json:"admin" url:"admin,omitempty"`
	BotAccount      Optional[bool]             `json:"bot_account" url:"bot_account,omitempty"`
	BanExpires      LemmyTime                  `json:"ban_expires" url:"ban_expires,omitempty"`
}
