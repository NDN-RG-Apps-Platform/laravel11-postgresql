package error

func ErrMapping(err error) bool {
	// Gabungkan GeneralErrors dan UserErrors menjadi satu slice
	allErrors := append(GeneralErrors, UserErrors...)

	// Loop untuk mengecek apakah err ada di dalam allErrors
	for _, item := range allErrors {
		if err.Error() == item.Error() {
			return true
		}
	}

	return false
}
