package main


type Time struct {
	EndTime string `json:"endtime"`
	Starttime string `json:"starttime"`
	UpdatedAfter string `json:"updatedafter"`
}

type RectangleLocation struct {
	MinLatitude float32 `json:"minlatitude"`
	MinLongitude float32 `json:"minlongitude"`
	MaxLatitude float32 `json:"maxlatitude"`
	MaxLongitude float32 `json:"maxlongitude"`
}

type CircleLocation struct {
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	MaxRadius float32 `json:"radius"`
	MaxRadiusKm float32 `json:"radiuskm"`
}

type Location struct {
	RectangleLocation RectangleLocation `json:"rectangle"`
	CircleLocation CircleLocation `json:"circle"`

}

type Other struct{
	Catalog string `json:"catalog"`
	Contributor string `json:"contributor"`
	EventId string `json:"eventid"`
	IncludeAllMagnitudes bool `json:"includeallmagnitudes"`
	IncludeAllOrigins bool `json:"includeallorigins"`
	IncludeArrivals bool `json:"includearrivals"`
	IncludeDeleted bool `json:"includedeleted"`
	IncludeSuperSeded bool `json:"includedsuperseded"`
	Limit int8 `json:"limit"`
	MaxDepth float32 `json:"maxdepth"`
	MaxMagnitude float32 `json:"maxmagnitude"`
	MinDepth float32 `json:"mindepth"`
	MinMagnitude float32 `json:"minmagnitude"`
	Offset int8 `json:"offset"`
	OrderBy string `json:"orderby"`
}

type Extensions struct{
	AlertLevel string `json:"alertlevel"`
	Callback string `json:"callback"`
	EventType string `json:"eventtype"`
	JsonError bool `json:"jsonerror"`
	KmlAnimated bool `json:"kmlanimated"`
	KmlColorBy string `json:"kmlcolorby"`
	MaxCDI float32 `json:"maxcdi"`
	Maxgap float32 `json:"maxgap"`
	MaxMmi float32 `json:"maxmmi"`
	MaxSig int16 `json:"maxsig"`
	MinCDI float32 `json:"mincdi"`
	MinFelt int16 `json:"minfelt"`
	Mingap float32 `json:"mingap"`
	MinSig int16 `json:"minsig"`
	NoData int16 `json:"nodata"`
	ProductType string `json:"producttype"`
	ProductCode string `json:"productcode"`
	ReviewStatus string `json:"reviewstatus"`
}


type Methods struct {
	Format map[string]string `json:"format"`
}

type Parameters struct {
	Format string `json:"format"`
	Time string `json:"time"`
	Location string `json:"location"`
}

type Data struct {
	Methods Methods `json:"methods"`
	Parameters Parameters `json:"parameters"`
}