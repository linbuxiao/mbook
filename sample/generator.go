package sample

import "mbook/proto"

func NewKeyBoard() *proto.KeyBoard {
	keyBoard := &proto.KeyBoard{
		Layout:  randomKeyBoardLayout(),
		Backlit: randomBool(),
	}

	return keyBoard
}

//func NewCPU() *proto.CPU  {
//
//}
