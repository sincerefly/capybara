package resources

import "embed"

//go:embed font/* font/*
var F embed.FS

const (
	AlibabaPuHiTi3LightTTF = "font/AlibabaPuHuiTi-3/AlibabaPuHuiTi-3-45-Light.ttf"
	AlibabaPuHiTi3BoldTTF  = "font/AlibabaPuHuiTi-3/AlibabaPuHuiTi-3-85-Bold.ttf"
	RobotBoldTTF           = "font/Robot/Robot-Bold.ttf"
	RobotLightTTF          = "font/Robot/Robot-Light.ttf"
	RobotMediumTTF         = "font/Robot/Robot-Medium.ttf"
	RobotRegularTTF        = "font/Robot/Robot-Regular.ttf"
)
