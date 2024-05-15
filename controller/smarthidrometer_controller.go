package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/OpenDataTelemetry/timeseries-api/database"
	"github.com/gin-gonic/gin"
)

func GetHidrometer(c *gin.Context) {
	intervalStr := c.DefaultQuery("interval", "15")

	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 400 {
		c.JSON(400, gin.H{"error": "Interval must be less than 400"})
		return
	}

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer influxDB.Close()

	query := `
		SELECT *
		FROM "Hidrometer"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		ORDER BY time DESC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() {
		value := iterator.Value()
		obj := gin.H{
			"fields": gin.H{
				"data_boardVoltage":            value["data_boardVoltage"],
				"data_counter":                 value["data_counter"],
				"fCnt":                         value["fCnt"],
				"rxInfo_altitude_0":            value["rxInfo_altitude_0"],
				"rxInfo_altitude_1":            value["rxInfo_altitude_1"],
				"rxInfo_latitude_0":            value["rxInfo_latitude_0"],
				"rxInfo_latitude_1":            value["rxInfo_latitude_1"],
				"rxInfo_loRaSNR_0":             value["rxInfo_loRaSNR_0"],
				"rxInfo_loRaSNR_1":             value["rxInfo_loRaSNR_1"],
				"rxInfo_longitude_0":           value["rxInfo_longitude_0"],
				"rxInfo_longitude_1":           value["rxInfo_longitude_1"],
				"rxInfo_rssi_0":                value["rxInfo_rssi_0"],
				"rxInfo_rssi_1":                value["rxInfo_rssi_1"],
				"rxInfo_dataRate_spreadFactor": value["rxInfo_dataRate_spreadFactor"],
				"txInfo_frequency":             value["txInfo_frequency"],
			},
			"name": "Hidrometer",
			"tags": gin.H{
				"applicationID":              value["applicationID"],
				"applicationName":            value["applicationName"],
				"devEUI":                     value["devEUI"],
				"fPort":                      value["fPort"],
				"host":                       value["host"],
				"nodeName":                   value["nodeName"],
				"rxInfo_mac_0":               value["rxInfo_mac_0"],
				"rxInfo_mac_1":               value["rxInfo_mac_1"],
				"rxInfo_name_0":              value["rxInfo_name_0"],
				"rxInfo_name_1":              value["rxInfo_name_1"],
				"txInfo_adr":                 value["txInfo_adr"],
				"txInfo_codeRate":            value["txInfo_codeRate"],
				"txInfo_dataRate_bandwidth":  value["txInfo_dataRate_bandwidth"],
				"txInfo_dataRate_modulation": value["txInfo_dataRate_modulation"],
			},
			"timestamp": value["time"],
		}
		objs = append(objs, obj)
	}
	fmt.Println(len(objs))
	if len(objs) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"message": "No content"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs)
}

func GetHidrometerbyNodeName(c *gin.Context) {
	nodeName := c.Param("nodeName")
	intervalStr := c.DefaultQuery("interval", "15")

	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 400 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Interval must be less than 400"})
		return
	}

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer influxDB.Close()

	query := fmt.Sprintf(`
		SELECT *
		FROM "Hidrometer"
		WHERE "nodeName" = '%s' AND "time" >= now() - interval '%s' minutes
		ORDER BY time DESC;
	`, nodeName, intervalStr)

	iterator, err := influxDB.Query(context.Background(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
				"data_boardVoltage":            value["data_boardVoltage"],
				"data_counter":                 value["data_counter"],
				"fCnt":                         value["fCnt"],
				"rxInfo_altitude_0":            value["rxInfo_altitude_0"],
				"rxInfo_altitude_1":            value["rxInfo_altitude_1"],
				"rxInfo_latitude_0":            value["rxInfo_latitude_0"],
				"rxInfo_latitude_1":            value["rxInfo_latitude_1"],
				"rxInfo_loRaSNR_0":             value["rxInfo_loRaSNR_0"],
				"rxInfo_loRaSNR_1":             value["rxInfo_loRaSNR_1"],
				"rxInfo_longitude_0":           value["rxInfo_longitude_0"],
				"rxInfo_longitude_1":           value["rxInfo_longitude_1"],
				"rxInfo_rssi_0":                value["rxInfo_rssi_0"],
				"rxInfo_rssi_1":                value["rxInfo_rssi_1"],
				"rxInfo_dataRate_spreadFactor": value["rxInfo_dataRate_spreadFactor"],
				"txInfo_frequency":             value["txInfo_frequency"],
			},
			"name": "Hidrometer",
			"tags": gin.H{
				"applicationID":              value["applicationID"],
				"applicationName":            value["applicationName"],
				"devEUI":                     value["devEUI"],
				"fPort":                      value["fPort"],
				"host":                       value["host"],
				"nodeName":                   value["nodeName"],
				"rxInfo_mac_0":               value["rxInfo_mac_0"],
				"rxInfo_mac_1":               value["rxInfo_mac_1"],
				"rxInfo_name_0":              value["rxInfo_name_0"],
				"rxInfo_name_1":              value["rxInfo_name_1"],
				"txInfo_adr":                 value["txInfo_adr"],
				"txInfo_codeRate":            value["txInfo_codeRate"],
				"txInfo_dataRate_bandwidth":  value["txInfo_dataRate_bandwidth"],
				"txInfo_dataRate_modulation": value["txInfo_dataRate_modulation"],
			},
			"timestamp": value["time"],
		}
		objs = append(objs, obj)
	}
	fmt.Println(len(objs))
	if len(objs) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"message": "No content"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs)
}

func GetHidrometerbyDevEUI(c *gin.Context) {
	devEUI := c.Param("devEUI")
	intervalStr := c.DefaultQuery("interval", "15")

	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 400 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Interval must be less than 400"})
		return
	}

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer influxDB.Close()

	query := fmt.Sprintf(`
		SELECT *
		FROM "Hidrometer"
		WHERE "devEUI" = '%s' AND "time" >= now() - interval '%s' minutes
		ORDER BY time DESC;
	`, devEUI, intervalStr)

	iterator, err := influxDB.Query(context.Background(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for iterator.Next() {
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
				"data_boardVoltage":            value["data_boardVoltage"],
				"data_counter":                 value["data_counter"],
				"fCnt":                         value["fCnt"],
				"rxInfo_altitude_0":            value["rxInfo_altitude_0"],
				"rxInfo_altitude_1":            value["rxInfo_altitude_1"],
				"rxInfo_latitude_0":            value["rxInfo_latitude_0"],
				"rxInfo_latitude_1":            value["rxInfo_latitude_1"],
				"rxInfo_loRaSNR_0":             value["rxInfo_loRaSNR_0"],
				"rxInfo_loRaSNR_1":             value["rxInfo_loRaSNR_1"],
				"rxInfo_longitude_0":           value["rxInfo_longitude_0"],
				"rxInfo_longitude_1":           value["rxInfo_longitude_1"],
				"rxInfo_rssi_0":                value["rxInfo_rssi_0"],
				"rxInfo_rssi_1":                value["rxInfo_rssi_1"],
				"rxInfo_dataRate_spreadFactor": value["rxInfo_dataRate_spreadFactor"],
				"txInfo_frequency":             value["txInfo_frequency"],
			},
			"name": "Hidrometer",
			"tags": gin.H{
				"applicationID":              value["applicationID"],
				"applicationName":            value["applicationName"],
				"devEUI":                     value["devEUI"],
				"fPort":                      value["fPort"],
				"host":                       value["host"],
				"nodeName":                   value["nodeName"],
				"rxInfo_mac_0":               value["rxInfo_mac_0"],
				"rxInfo_mac_1":               value["rxInfo_mac_1"],
				"rxInfo_name_0":              value["rxInfo_name_0"],
				"rxInfo_name_1":              value["rxInfo_name_1"],
				"txInfo_adr":                 value["txInfo_adr"],
				"txInfo_codeRate":            value["txInfo_codeRate"],
				"txInfo_dataRate_bandwidth":  value["txInfo_dataRate_bandwidth"],
				"txInfo_dataRate_modulation": value["txInfo_dataRate_modulation"],
			},
			"timestamp": value["time"],
		}
		objs = append(objs, obj)
	}
	fmt.Println(len(objs))
	if len(objs) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"message": "No content"})
		return
	}
	c.IndentedJSON(http.StatusOK, objs)
}
