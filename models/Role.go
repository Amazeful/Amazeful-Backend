package models

type UserRole int

const (
	UserRoleGlobal UserRole = 1 << iota
	UserRoleSubscriber
	UserRoleVIP
	UserRoleEditor
	UserRoleModerator
	UserRoleSuperMod
	UserRoleBroadcaster
	UserRoleAdmin
)
