package algorithm

import (
	"fmt"
)

func Permute(str []rune,l int,r int)  {
	if l==r {
		fmt.Println(string(str))
	}else {
		for i:=l;i<=r ; i++ {
			str[l], str[i] = str[i], str[l]
			Permute(str,l+1,r)
			str[l], str[i] = str[i], str[l]
		}
	}
}