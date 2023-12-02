package enums

type GroupRole string

const (
	GroupRoleOwner  GroupRole = "owner"
	GroupRoleAdmin  GroupRole = "admin"
	GroupRoleMember GroupRole = "member"
)

type GroupInvitationStatus string

const (
	GroupInvitationStatusPending   GroupInvitationStatus = "pending"
	GroupInvitationStatusAccepted  GroupInvitationStatus = "accepted"
	GroupInvitationStatusRejected  GroupInvitationStatus = "rejected"
	GroupInvitationStatusCancelled GroupInvitationStatus = "canceled"
)
