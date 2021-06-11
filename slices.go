package slices

func FindString(sl *[]string, i *string) (*string, bool) {
	for _, c := range *sl {
		if c == *i {
			return &c, true
		}
	}

	return nil, false
}

// returns -1 if unable to find element
func IndexOf(sl *[]interface{}, s *interface{}) (int, bool) {
	var i int
	for _, el := range *sl {
		if el == *s {
			return i, true
		}
		i++
	}

	return -1, false
}
