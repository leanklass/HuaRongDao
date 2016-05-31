package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"os"

	"github.com/ghj1976/HuaRongDao/level"
	"github.com/ghj1976/HuaRongDao/textures"
)

func main() {
	// 待处理的数据
	levelArr := level.InitData()

	// 准备字体
	textures.LoadGameFont("../assets/f1.ttf")

	// 绘图区域创建
	d := 2
	ff := 30
	_, _, areaWidth, areaHeight := textures.GetBounds(d)
	mmm := image.NewRGBA(image.Rect(0, 0, 10*areaWidth+11*ff, 4*areaHeight+5*ff))
	bg := color.RGBA{171, 190, 62, 255}
	draw.Draw(mmm, mmm.Bounds(), &image.Uniform{bg}, image.ZP, draw.Src)

	i := 0
	for _, le := range levelArr {
		if i > 40 {
			break
		}
		a := i % 10
		b := i / 10
		// log.Println(a, "-", b)

		le.MapArray = level.Layout2Map(le.Layout)
		le.ChessMans = level.ChessManArray2Map(le.MapArray)
		//log.Println(le.Layout)
		r := rand.Intn(3)
		if r == 0 {
			le.LevelStatus = level.LevelNotPass
		} else if r == 1 {
			le.LevelStatus = level.LevelPass
		} else {
			le.LevelStatus = level.LevelBestPass
		}

		m := textures.LevelRGBA(2, &le)

		// 绘制在大背景图上
		draw.Draw(mmm,
			image.Rect(ff+a*(areaWidth+ff),
				ff+b*(areaHeight+ff),
				ff+a*(areaWidth+ff)+areaWidth,
				ff+b*(areaHeight+ff)+areaHeight),
			m, image.ZP, draw.Src)

		i++
	}

	//	// 绘图
	//	draw.Draw(m, image.Rect(100, 100, 200, 300), &image.Uniform{blue}, image.ZP, draw.Src)

	// 保存到文件
	imgfile, _ := os.OpenFile("001.png", os.O_RDWR|os.O_CREATE, 0660)
	defer imgfile.Close()
	png.Encode(imgfile, mmm)

	textures.ReleaseFont()
	log.Println("finish")
}
