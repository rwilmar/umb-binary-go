package umbBinary

type UmbChannels struct {
	Channel           uint16 //channel number
	Description       string
	Unit              string // symbol for printing
	UnitOfMeasurement string // see: http://www.foodvoc.org/vocabularies/om-1.8/Unit_of_measure
}

var umbChannelsMap = map[uint16]UmbChannels{
	100: {100, "temperature °C", "°C", "om:Celsius_temperature_unit"},
	105: {105, "temperature °F", "°F", "om:Fahrenheit_temperature_unit"},

	110: {110, "dewpoint °C", "°C", "om:Celsius_temperature_unit"},
	115: {115, "dewpoint °F", "°F", "om:Fahrenheit_temperature_unit"},

	113: {113, "dome heater temp", "°C", "om:Celsius_temperature_unit"},
	118: {118, "dome heater temp", "°F", "om:Fahrenheit_temperature_unit"},

	200: {200, "relative humidity", "%", "om:Relative_humidity_unit"},
	205: {205, "absolute humidity", "g/m³", "om:Current_density_unit"},
	210: {210, "mixing ratio ", "g/kg", "om:Current_density_unit"},

	300: {300, "absolute air pressure", "hPa", "om:Pressure_unit"},
	305: {305, "rel. air pressure", "hPa", "om:Pressure_unit"},

	400: {400, "wind speed", "m/s", "om:speed_unit"},
	405: {405, "wind speed", "km/h", "om:speed_unit"},
	410: {410, "wind speed", "mph", "om:speed_unit"},
	415: {415, "wind speed", "kts", "om:speed_unit"},

	500: {500, "wind direction", "°", "Celsius_temperature_unit"},
	502: {502, "wind direction compass", "°", "Celsius_temperature_unit"},
	510: {510, "wind direction heading", "°", "Celsius_temperature_unit"},

	700: {700, "Precipitation intens (0 = No, 60 = Rain, 70 = Snow)", "", ""},

	600: {600, "Precipitation Liter / m²", "l/m²", "om:Amount_of_substance_flow_unit"},
	601: {601, "Daily Precipitation Liter / m²", "l/m²", "om:Amount_of_substance_flow_unit"},
	620: {620, "Precip Water film height in mm", "mm", "om:Amount_of_substance_flow_unit"},
	621: {621, "Daily Precipitation mm", "mm", "om:Amount_of_substance_flow_unit"},
	640: {640, "Water film height in inches", "inch", "om:Amount_of_substance_flow_unit"},
	660: {660, "Water film height in mils ", "mil", "om:Amount_of_substance_flow_unit"},
	605: {605, "Liter/m² since last request", "l/m²", "om:Amount_of_substance_flow_unit"},
	625: {625, "Water film height in mm since last request", "mm", "om:Amount_of_substance_flow_unit"},
	645: {645, "Water film height in inches since last request", "inch", "om:Amount_of_substance_flow_unit"},
	665: {665, "Water film height in mils since last request", "mil", "om:Amount_of_substance_flow_unit"},

	900: {900, "Global Radiation", "W/m²", "om:Radiant_intensity_unit"},
	902: {902, "UV-index", "digits", "om:Radiant_intensity_unit"},
	903: {903, "ambient light level", "klx", "om:Luminance_unit"},
	904: {904, "twilight", "lx", "om:Luminance_unit"},
	910: {910, "sun direction azimuth", "°", "om:Plane_angle_unit"},
	911: {911, "sun direction elevation", "°", "om:Plane_angle_unit"},

	3900: {3900, "latitude", "°", "om:Plane_angle_unit"},
	3901: {3901, "longitude", "°", "om:Plane_angle_unit"},
	3902: {3902, "altitude", "m", "om:Length_unit"},

	4060: {4060, "wifi status", "", "singular unit"},
	4061: {4061, "wifi signal", "dBm", "om:Energy_unit"},
	4071: {4071, "gps num satellites", "", "singular unit"},
	4640: {4640, "R2S heater status (on/off)", "", "singular unit"},
	4702: {4702, "boot count", "", "singular unit"},
	4703: {4703, "system time - UTC", "", "om:Time_unit"},
	4704: {4704, "system time - local", "", "om:Time_unit"},
	4705: {4705, "CPU load", "%", "om:Percentage_unit"},

	10000: {10000, "voltage supply", "V", "om:Electric_potential_unit"},
}
