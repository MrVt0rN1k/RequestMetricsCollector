package main

var (
	size, allSizeFile, less1mb, between1mband5mb, between5mband10mb, more10bm int
)

func fileSize(length int) {
	if length < 1024 {
		size = length
	} else {
		size = length / 1024
	}

	switch {
	case size > 0 && size <= 1000:
		less1mb++
	case size > 1000 && size <= 5000:
		between1mband5mb++
	case size > 5000 && size <= 10000:
		between5mband10mb++
	case size >= 10000:
		more10bm++
	}
	allSizeFile += size
}
