package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("123456"))
}
func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$h1ObbRdK8GHKZGhd22QwNetNyiUdI8/myJotPQWUAfzkeXgZvPzoO", "123456"))
}
