package types

type Search struct {
	Q             string                `json:"q" url:"q,omitempty"`
	CommunityID   Optional[int]         `json:"community_id" url:"community_id,omitempty"`
	CommunityName Optional[string]      `json:"community_name" url:"community_name,omitempty"`
	CreatorID     Optional[int]         `json:"creator_id" url:"creator_id,omitempty"`
	Type          Optional[SearchType]  `json:"type_" url:"type_,omitempty"`
	Sort          Optional[SortType]    `json:"sort" url:"sort,omitempty"`
	ListingType   Optional[ListingType] `json:"listing_type" url:"listing_type,omitempty"`
	Page          Optional[int64]       `json:"page" url:"page,omitempty"`
	Limit         Optional[int64]       `json:"limit" url:"limit,omitempty"`
	Auth          Optional[string]      `json:"auth" url:"auth,omitempty"`
}
type SearchResponse struct {
	Type        string           `json:"type_" url:"type_,omitempty"`
	Comments    []CommentView    `json:"comments" url:"comments,omitempty"`
	Posts       []PostView       `json:"posts" url:"posts,omitempty"`
	Communities []CommunityView  `json:"communities" url:"communities,omitempty"`
	Users       []PersonViewSafe `json:"users" url:"users,omitempty"`
	LemmyResponse
}
type ResolveObject struct {
	Q    string           `json:"q" url:"q,omitempty"`
	Auth Optional[string] `json:"auth" url:"auth,omitempty"`
}
type ResolveObjectResponse struct {
	Comment   Optional[CommentView]    `json:"comment" url:"comment,omitempty"`
	Post      Optional[PostView]       `json:"post" url:"post,omitempty"`
	Community Optional[CommunityView]  `json:"community" url:"community,omitempty"`
	Person    Optional[PersonViewSafe] `json:"person" url:"person,omitempty"`
	LemmyResponse
}
type GetModlog struct {
	ModPersonID Optional[int]    `json:"mod_person_id" url:"mod_person_id,omitempty"`
	CommunityID Optional[int]    `json:"community_id" url:"community_id,omitempty"`
	Page        Optional[int64]  `json:"page" url:"page,omitempty"`
	Limit       Optional[int64]  `json:"limit" url:"limit,omitempty"`
	Auth        Optional[string] `json:"auth" url:"auth,omitempty"`
}
type GetModlogResponse struct {
	RemovedPosts           []ModRemovePostView        `json:"removed_posts" url:"removed_posts,omitempty"`
	LockedPosts            []ModLockPostView          `json:"locked_posts" url:"locked_posts,omitempty"`
	StickiedPosts          []ModStickyPostView        `json:"stickied_posts" url:"stickied_posts,omitempty"`
	RemovedComments        []ModRemoveCommentView     `json:"removed_comments" url:"removed_comments,omitempty"`
	RemovedCommunities     []ModRemoveCommunityView   `json:"removed_communities" url:"removed_communities,omitempty"`
	BannedFromCommunity    []ModBanFromCommunityView  `json:"banned_from_community" url:"banned_from_community,omitempty"`
	Banned                 []ModBanView               `json:"banned" url:"banned,omitempty"`
	AddedToCommunity       []ModAddCommunityView      `json:"added_to_community" url:"added_to_community,omitempty"`
	TransferredToCommunity []ModTransferCommunityView `json:"transferred_to_community" url:"transferred_to_community,omitempty"`
	Added                  []ModAddView               `json:"added" url:"added,omitempty"`
	HiddenCommunities      []ModHideCommunityView     `json:"hidden_communities" url:"hidden_communities,omitempty"`
	LemmyResponse
}
type CreateSite struct {
	Name                       string           `json:"name" url:"name,omitempty"`
	Sidebar                    Optional[string] `json:"sidebar" url:"sidebar,omitempty"`
	Description                Optional[string] `json:"description" url:"description,omitempty"`
	Icon                       Optional[string] `json:"icon" url:"icon,omitempty"`
	Banner                     Optional[string] `json:"banner" url:"banner,omitempty"`
	EnableDownvotes            Optional[bool]   `json:"enable_downvotes" url:"enable_downvotes,omitempty"`
	OpenRegistration           Optional[bool]   `json:"open_registration" url:"open_registration,omitempty"`
	EnableNSFW                 Optional[bool]   `json:"enable_nsfw" url:"enable_nsfw,omitempty"`
	CommunityCreationAdminOnly Optional[bool]   `json:"community_creation_admin_only" url:"community_creation_admin_only,omitempty"`
	RequireEmailVerification   Optional[bool]   `json:"require_email_verification" url:"require_email_verification,omitempty"`
	RequireApplication         Optional[bool]   `json:"require_application" url:"require_application,omitempty"`
	ApplicationQuestion        Optional[string] `json:"application_question" url:"application_question,omitempty"`
	PrivateInstance            Optional[bool]   `json:"private_instance" url:"private_instance,omitempty"`
	DefaultTheme               Optional[string] `json:"default_theme" url:"default_theme,omitempty"`
	DefaultPostListingType     Optional[string] `json:"default_post_listing_type" url:"default_post_listing_type,omitempty"`
	Auth                       string           `json:"auth" url:"auth,omitempty"`
}
type EditSite struct {
	Name                       Optional[string] `json:"name" url:"name,omitempty"`
	Sidebar                    Optional[string] `json:"sidebar" url:"sidebar,omitempty"`
	Description                Optional[string] `json:"description" url:"description,omitempty"`
	Icon                       Optional[string] `json:"icon" url:"icon,omitempty"`
	Banner                     Optional[string] `json:"banner" url:"banner,omitempty"`
	EnableDownvotes            Optional[bool]   `json:"enable_downvotes" url:"enable_downvotes,omitempty"`
	OpenRegistration           Optional[bool]   `json:"open_registration" url:"open_registration,omitempty"`
	EnableNSFW                 Optional[bool]   `json:"enable_nsfw" url:"enable_nsfw,omitempty"`
	CommunityCreationAdminOnly Optional[bool]   `json:"community_creation_admin_only" url:"community_creation_admin_only,omitempty"`
	RequireEmailVerification   Optional[bool]   `json:"require_email_verification" url:"require_email_verification,omitempty"`
	RequireApplication         Optional[bool]   `json:"require_application" url:"require_application,omitempty"`
	ApplicationQuestion        Optional[string] `json:"application_question" url:"application_question,omitempty"`
	PrivateInstance            Optional[bool]   `json:"private_instance" url:"private_instance,omitempty"`
	DefaultTheme               Optional[string] `json:"default_theme" url:"default_theme,omitempty"`
	DefaultPostListingType     Optional[string] `json:"default_post_listing_type" url:"default_post_listing_type,omitempty"`
	LegalInformation           Optional[string] `json:"legal_information" url:"legal_information,omitempty"`
	Auth                       string           `json:"auth" url:"auth,omitempty"`
}
type GetSite struct {
	Auth Optional[string] `json:"auth" url:"auth,omitempty"`
}
type SiteResponse struct {
	SiteView SiteView `json:"site_view" url:"site_view,omitempty"`
	LemmyResponse
}
type GetSiteResponse struct {
	SiteView           Optional[SiteView]           `json:"site_view" url:"site_view,omitempty"`
	Admins             []PersonViewSafe             `json:"admins" url:"admins,omitempty"`
	Online             uint                         `json:"online" url:"online,omitempty"`
	Version            string                       `json:"version" url:"version,omitempty"`
	MyUser             Optional[MyUserInfo]         `json:"my_user" url:"my_user,omitempty"`
	FederatedInstances Optional[FederatedInstances] `json:"federated_instances" url:"federated_instances,omitempty"`
	LemmyResponse
}
type MyUserInfo struct {
	LocalUserView   LocalUserSettingsView    `json:"local_user_view" url:"local_user_view,omitempty"`
	Follows         []CommunityFollowerView  `json:"follows" url:"follows,omitempty"`
	Moderates       []CommunityModeratorView `json:"moderates" url:"moderates,omitempty"`
	CommunityBlocks []CommunityBlockView     `json:"community_blocks" url:"community_blocks,omitempty"`
	PersonBlocks    []PersonBlockView        `json:"person_blocks" url:"person_blocks,omitempty"`
}
type LeaveAdmin struct {
	Auth string `json:"auth" url:"auth,omitempty"`
}
type GetSiteConfig struct {
	Auth string `json:"auth" url:"auth,omitempty"`
}
type GetSiteConfigResponse struct {
	ConfigHjson string `json:"config_hjson" url:"config_hjson,omitempty"`
	LemmyResponse
}
type SaveSiteConfig struct {
	ConfigHjson string `json:"config_hjson" url:"config_hjson,omitempty"`
	Auth        string `json:"auth" url:"auth,omitempty"`
}
type FederatedInstances struct {
	Linked  []string           `json:"linked" url:"linked,omitempty"`
	Allowed Optional[[]string] `json:"allowed" url:"allowed,omitempty"`
	Blocked Optional[[]string] `json:"blocked" url:"blocked,omitempty"`
}
type ListRegistrationApplications struct {
	UnreadOnly Optional[bool]  `json:"unread_only" url:"unread_only,omitempty"`
	Page       Optional[int64] `json:"page" url:"page,omitempty"`
	Limit      Optional[int64] `json:"limit" url:"limit,omitempty"`
	Auth       string          `json:"auth" url:"auth,omitempty"`
}
type ListRegistrationApplicationsResponse struct {
	RegistrationApplications []RegistrationApplicationView `json:"registration_applications" url:"registration_applications,omitempty"`
	LemmyResponse
}
type ApproveRegistrationApplication struct {
	ID         int32            `json:"id" url:"id,omitempty"`
	Approve    bool             `json:"approve" url:"approve,omitempty"`
	DenyReason Optional[string] `json:"deny_reason" url:"deny_reason,omitempty"`
	Auth       string           `json:"auth" url:"auth,omitempty"`
}
type RegistrationApplicationResponse struct {
	RegistrationApplication RegistrationApplicationView `json:"registration_application" url:"registration_application,omitempty"`
	LemmyResponse
}
type GetUnreadRegistrationApplicationCount struct {
	Auth string `json:"auth" url:"auth,omitempty"`
}
type GetUnreadRegistrationApplicationCountResponse struct {
	RegistrationApplications int64 `json:"registration_applications" url:"registration_applications,omitempty"`
	LemmyResponse
}
