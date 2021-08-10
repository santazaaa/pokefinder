package api

func IsNotFoundError(err error) bool {
	return err != nil && err.Error() == "NOT_FOUND"
}
