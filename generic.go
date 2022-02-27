package yavp

func GenericValidator(key string, errs ...error) error {
	errors := make(Errors)
	for _, err := range errs {
		if err != nil {
			errors[key] = err
			return errors
		}
	}
	return nil
}
