package constant

type UserStatus int

const (
	UserStatusInactive UserStatus = iota
	UserStatusActive
	UserStatusBanned
)
