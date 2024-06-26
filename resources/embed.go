package resources

import "embed"

//go:embed font/* font/*
var F embed.FS

const (
	AlibabaPuHiTi3_Light_TTF = "font/AlibabaPuHuiTi-3/AlibabaPuHuiTi-3-45-Light.ttf"
	AlibabaPuHiTi3_Bold_TTF  = "font/AlibabaPuHuiTi-3/AlibabaPuHuiTi-3-85-Bold.ttf"
	Robot_Bold_TTF           = "font/Robot/Robot-Bold.ttf"
	Robot_Light_TTF          = "font/Robot/Robot-Bold.ttf"
	Robot_Medium_TTF         = "font/Robot/Robot-Bold.ttf"
	Robot_Rgular_TTF         = "font/Robot/Robot-Bold.ttf"
)
