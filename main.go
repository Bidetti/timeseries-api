package main

import (
	"github.com/OpenDataTelemetry/timeseries-api/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default() // Create a new gin router instance
	api := r.Group("/api/timeseries/v0.2/smartcampusmaua")
	{
		api.GET("SmartLights", controller.GetSmartLights)
		api.GET("SmartLights/deviceName/:nodeName", controller.GetSmartLightbyNodeName)
		api.GET("SmartLights/deviceId/:devEUI", controller.GetSmartLightbyDevEUI)
		api.GET("WaterTankLevel", controller.GetWaterTankLevel)
		api.GET("WaterTankLevel/deviceName/:nodeName", controller.GetWaterTankLevelbyNodeName)
		api.GET("WaterTankLevel/deviceId/:devEUI", controller.GetWaterTankLevelbyDevEUI)
		api.GET("LastWaterTankLevel", controller.GetLatestWaterTankLevels)
		api.GET("Hidrometer", controller.GetHidrometer)
		api.GET("Hidrometer/deviceName/:nodeName", controller.GetHidrometerbyNodeName)
		api.GET("Hidrometer/deviceId/:devEUI", controller.GetHidrometerbyDevEUI)
		api.GET("ArtesianWell", controller.GetArtesianWell)
		api.GET("ArtesianWell/deviceName/:nodeName", controller.GetArtesianWellbyNodeName)
		api.GET("ArtesianWell/deviceId/:devEUI", controller.GetArtesianWellbyDevEUI)
	}

	r.Run(":8888")
}
