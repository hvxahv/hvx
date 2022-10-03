package activitypub

// https://www.w3.org/TR/activitystreams-vocabulary/

const (
	FollowType          = "Follow"
	AcceptType          = "Accept"
	RejectType          = "Reject"
	UndoType            = "Undo"
	AddType             = "Add"
	AnnounceType        = "Announce"
	ArriveType          = "Arrive"
	BlockType           = "Block"
	CreateType          = "Create"
	DeleteType          = "Delete"
	DislikeType         = "Dislike"
	FlagType            = "Flag"
	IgnoreType          = "Ignore"
	InviteType          = "Invite"
	JoinType            = "Join"
	LeaveType           = "Leave"
	LikeType            = "Like"
	ListenType          = "Listen"
	MoveType            = "Move"
	OfferType           = "Offer"
	QuestionType        = "Question"
	ReadType            = "Read"
	RemoveType          = "Remove"
	TentativeRejectType = "TentativeReject"
	TentativeAcceptType = "TentativeAccept"
	TravelType          = "Travel"
	UpdateType          = "Update"
	ViewType            = "View"
	NoteType            = "Note"
)

// Actor https://www.w3.org/TR/activitystreams-vocabulary/#actor-types.
const (
	ApplicationType  = "Application"
	GroupType        = "Group"
	OrganizationType = "Organization"
	PersonType       = "Person"
	ServiceType      = "Service"
	ChannelType      = "Channel"
)
