package sense

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

type Accessory struct {
	*accessory.Accessory

	TemperatureSensor *service.TemperatureSensor
	HumiditySensor    *service.HumiditySensor
	LightSensor       *service.LightSensor
	AirQualitySensor  *service.AirQualitySensor
}

func NewAccessory(name string) *Accessory {
	info := accessory.Info{
		Name:         name,
		Manufacturer: "Hello",
		Model:        "Sense",
	}
	acc := Accessory{}
	acc.Accessory = accessory.New(info, accessory.TypeThermostat)

	acc.TemperatureSensor = service.NewTemperatureSensor()
	acc.AddService(acc.TemperatureSensor.Service)

	acc.HumiditySensor = service.NewHumiditySensor()
	acc.AddService(acc.HumiditySensor.Service)

	acc.LightSensor = service.NewLightSensor()
	acc.AddService(acc.LightSensor.Service)

	acc.AirQualitySensor = service.NewAirQualitySensor()
	acc.AddService(acc.AirQualitySensor.Service)

	return &acc
}

func (acc *Accessory) UpdateValues(room *RoomMeasurement) {
	acc.TemperatureSensor.CurrentTemperature.SetValue(room.Temperature.Value)
	acc.HumiditySensor.CurrentRelativeHumidity.SetValue(room.Humidity.Value)
	acc.LightSensor.CurrentAmbientLightLevel.SetValue(room.Light.Value)
	acc.AirQualitySensor.AirQuality.SetValue(airQualityValue(room.Particulates.Value))
}

func airQualityValue(aqi float64) int {
	switch {
	case aqi < 50:
		return characteristic.AirQualityExcellent
	case aqi < 100:
		return characteristic.AirQualityGood
	case aqi < 150:
		return characteristic.AirQualityFair
	case aqi < 200:
		return characteristic.AirQualityInferior
	default:
		return characteristic.AirQualityPoor
	}
}
