package ggwrapper

import (
	"image"

	"github.com/fogleman/gg"
)

// ApplyRoundedCorners 将圆角效果应用于给定图像，并返回带有透明背景的新图像
func ApplyRoundedCorners(src image.Image, radius float64) image.Image {
	width := src.Bounds().Dx()
	height := src.Bounds().Dy()

	// 创建一个新的图形上下文，并确保背景透明
	dc := gg.NewContext(width, height)
	dc.SetRGBA(0, 0, 0, 0) // 完全透明
	dc.Clear()

	// 创建带圆角的矩形路径
	dc.MoveTo(radius, 0)
	dc.LineTo(float64(width)-radius, 0)
	dc.QuadraticTo(float64(width), 0, float64(width), radius)
	dc.LineTo(float64(width), float64(height)-radius)
	dc.QuadraticTo(float64(width), float64(height), float64(width)-radius, float64(height))
	dc.LineTo(radius, float64(height))
	dc.QuadraticTo(0, float64(height), 0, float64(height)-radius)
	dc.LineTo(0, radius)
	dc.QuadraticTo(0, 0, radius, 0)
	dc.ClosePath()

	// 使用该路径剪裁图像
	dc.Clip()

	// 绘制原始图像到上下文
	dc.DrawImage(src, 0, 0)

	// 返回带圆角的新图像
	return dc.Image()
}
