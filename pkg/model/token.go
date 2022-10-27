package model

type Token struct {
	Symbol          string
	Name            string
	Logo            string
	ContractAddress string
	Platform        string
	// PlatformSpecificData store platform specific data as a json string
	PlatformSpecificData string
	UsdPrice             float64
	UsdMarketCap         float64
	Usd24HourChange      float64
}
