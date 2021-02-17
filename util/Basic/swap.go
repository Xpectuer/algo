package basic

func swap(a, b *int) {
	*a ^= *b
	*b ^= *a
	*a ^= *b
}
