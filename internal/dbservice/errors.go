package dbservice

type AdminNotFoundError struct {}

func (e *AdminNotFoundError) Error() string {
	return "admin not found"
}

type UserNotFoundError struct {}

func (e *UserNotFoundError) Error() string {
	return "user not found"
}