package types

type ModAddCommunityView struct {
	ModAddCommunity ModAddCommunity  `json:"mod_add_community" url:"mod_add_community,omitempty"`
	Moderator       PersonSafe       `json:"moderator" url:"moderator,omitempty"`
	Community       CommunitySafe    `json:"community" url:"community,omitempty"`
	ModdedPerson    PersonSafeAlias1 `json:"modded_person" url:"modded_person,omitempty"`
}
type ModAddView struct {
	ModAdd       ModAdd           `json:"mod_add" url:"mod_add,omitempty"`
	Moderator    PersonSafe       `json:"moderator" url:"moderator,omitempty"`
	ModdedPerson PersonSafeAlias1 `json:"modded_person" url:"modded_person,omitempty"`
}
type ModBanFromCommunityView struct {
	ModBanFromCommunity ModBanFromCommunity `json:"mod_ban_from_community" url:"mod_ban_from_community,omitempty"`
	Moderator           PersonSafe          `json:"moderator" url:"moderator,omitempty"`
	Community           CommunitySafe       `json:"community" url:"community,omitempty"`
	BannedPerson        PersonSafeAlias1    `json:"banned_person" url:"banned_person,omitempty"`
}
type ModBanView struct {
	ModBan       ModBan           `json:"mod_ban" url:"mod_ban,omitempty"`
	Moderator    PersonSafe       `json:"moderator" url:"moderator,omitempty"`
	BannedPerson PersonSafeAlias1 `json:"banned_person" url:"banned_person,omitempty"`
}
type ModHideCommunityView struct {
	ModHideCommunity ModHideCommunity `json:"mod_hide_community" url:"mod_hide_community,omitempty"`
	Admin            PersonSafe       `json:"admin" url:"admin,omitempty"`
	Community        CommunitySafe    `json:"community" url:"community,omitempty"`
}
type ModLockPostView struct {
	ModLockPost ModLockPost   `json:"mod_lock_post" url:"mod_lock_post,omitempty"`
	Moderator   PersonSafe    `json:"moderator" url:"moderator,omitempty"`
	Post        Post          `json:"post" url:"post,omitempty"`
	Community   CommunitySafe `json:"community" url:"community,omitempty"`
}
type ModRemoveCommentView struct {
	ModRemoveComment ModRemoveComment `json:"mod_remove_comment" url:"mod_remove_comment,omitempty"`
	Moderator        PersonSafe       `json:"moderator" url:"moderator,omitempty"`
	Comment          Comment          `json:"comment" url:"comment,omitempty"`
	Commenter        PersonSafeAlias1 `json:"commenter" url:"commenter,omitempty"`
	Post             Post             `json:"post" url:"post,omitempty"`
	Community        CommunitySafe    `json:"community" url:"community,omitempty"`
}
type ModRemoveCommunityView struct {
	ModRemoveCommunity ModRemoveCommunity `json:"mod_remove_community" url:"mod_remove_community,omitempty"`
	Moderator          PersonSafe         `json:"moderator" url:"moderator,omitempty"`
	Community          CommunitySafe      `json:"community" url:"community,omitempty"`
}
type ModRemovePostView struct {
	ModRemovePost ModRemovePost `json:"mod_remove_post" url:"mod_remove_post,omitempty"`
	Moderator     PersonSafe    `json:"moderator" url:"moderator,omitempty"`
	Post          Post          `json:"post" url:"post,omitempty"`
	Community     CommunitySafe `json:"community" url:"community,omitempty"`
}
type ModStickyPostView struct {
	ModStickyPost ModStickyPost `json:"mod_sticky_post" url:"mod_sticky_post,omitempty"`
	Moderator     PersonSafe    `json:"moderator" url:"moderator,omitempty"`
	Post          Post          `json:"post" url:"post,omitempty"`
	Community     CommunitySafe `json:"community" url:"community,omitempty"`
}
type ModTransferCommunityView struct {
	ModTransferCommunity ModTransferCommunity `json:"mod_transfer_community" url:"mod_transfer_community,omitempty"`
	Moderator            PersonSafe           `json:"moderator" url:"moderator,omitempty"`
	Community            CommunitySafe        `json:"community" url:"community,omitempty"`
	ModdedPerson         PersonSafeAlias1     `json:"modded_person" url:"modded_person,omitempty"`
}
