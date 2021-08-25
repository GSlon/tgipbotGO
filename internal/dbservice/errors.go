package dbservice

type AdminNotFoundError struct {}

func (e *AdminNotFoundError) Error() string {
	return "admin not found"
}

type AdminAlreadyExistsError struct {}

func (e *AdminAlreadyExistsError) Error() string {
	return "admin already exists"
}

type UserNotFoundError struct {}

func (e *UserNotFoundError) Error() string {
	return "user not found"
}

type UserAlreadyExistsError struct {}

func (e *UserAlreadyExistsError) Error() string {
	return "user already exists"
}

type LogNotFoundError struct {}

func (e *LogNotFoundError) Error() string {
	return "log not found"
}

type LogAlreadyExistsError struct {}

func (e *LogAlreadyExistsError) Error() string {
	return "log already exists"
}