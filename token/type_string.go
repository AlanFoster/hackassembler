// Code generated by "stringer -type=Type"; DO NOT EDIT.

package token

import "fmt"

const _Type_name = "VALUELEFT_BRACKETRIGHT_BRACKETATEQUALSOPERATORSEMICOLONINVALIDEOF"

var _Type_index = [...]uint8{0, 5, 17, 30, 32, 38, 46, 55, 62, 65}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
