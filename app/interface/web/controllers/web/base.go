package web

func CheckPage(page, pageSize int) (int32, int32) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 15
	}

	return int32(page), int32(pageSize)
}
