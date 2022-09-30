package helpers

func DeReferenceString(s *string) string {
	if s == nil {
		temp := ""
		s = &temp
	}
	val := *s

	return val
}
