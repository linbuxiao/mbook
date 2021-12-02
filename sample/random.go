package sample

import (
	"math/rand"
	"mbook/proto"
)

func randomKeyBoardLayout() proto.KeyBoard_Layout {
	switch rand.Intn(3) {
	case 1:
		return proto.KeyBoard_AZERTY
	case 2:
		return proto.KeyBoard_QWERTY
	case 3:
		return proto.KeyBoard_QWERTZ
	default:
		return proto.KeyBoard_UNKNOWN
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}
