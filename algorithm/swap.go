package algorithm


func swap(a *int,b *int){
	*a += *b
	*b = *a - *b
	*a -= *b
}
