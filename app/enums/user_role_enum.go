package enums

import "fmt"

type UserRoleEnum int8

const (
	RoleAdmin UserRoleEnum = iota + 1
	RoleUser
	RoleGuest
	RoleCustomer
)

func ParseUserRole(role UserRoleEnum) (UserRoleEnum, error) {
	switch role {
	case 1:
		return RoleAdmin, nil
	case 2:
		return RoleUser, nil
	case 3:
		return RoleGuest, nil
	case 4:
		return RoleCustomer, nil
	default:
		return 0, fmt.Errorf("invalid role: %d", role)
	}
}
