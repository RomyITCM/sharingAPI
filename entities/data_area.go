package entities

type DataArea struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type DataMasterRegion struct {
	RegionId   string `json:"region_id"`
	RegionName string `json:"region_name"`
}

type DataMasterArea struct {
	AreaId   string `json:"area_id"`
	AreaName string `json:"area_name"`
}
