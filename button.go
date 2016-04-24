package main

const (
	BtnPress  = iota // 按钮被按下状态
	BtnNormal        // 按钮正常状态
)

type BtnStatus byte // 按钮的状态枚举

// 游戏中的按钮类
type GameBtn struct {
	status        BtnStatus // 按钮的状态， 一共2种，按下、正常
	GameRectangle           // 按钮所在位置（长方形）
}