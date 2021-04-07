package yavp

func When(condition bool, validator error) error {
	if condition {
		return validator
	} else {
		return nil
	}
}
