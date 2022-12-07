package serializer

import "go_ctry/model"

// 定义序列器 Areaservice

type Area struct {
	AreaID    uint    `json:"area_id" form:"area_id"`
	AreaName  string  `json:"area_name" form:"area_name"`
	CityCode  string  `json:"city_code" form:"city_code"`
	CityLevel int     `json:"city_level" form:"city_level"`
	ParentID  uint     `json:"parent_id" form:"parent_id"`
	Children  []*Area `json:"children" form:"children"`
}

// 构造地址树
func BuildAreaTree(list []*Area, parentId uint) []*Area {
	res := make([]*Area, 0)
	for _, v := range list {
		if v.ParentID == parentId {
			v.Children = BuildAreaTree(list, uint(v.AreaID))
			res = append(res, v)
		}
	}
	return res
}

// 定义一个 area 切片
var area_list []*Area

// 序列化 地址
func BuildArea(area_li []*model.Area) []*Area {
	for _, item := range area_li {
		area_list = append(area_list, &Area{
			AreaID:    item.ID,
			AreaName:  item.Name,
			CityCode:  item.City_code,
			CityLevel: item.City_level,
			ParentID:  item.ParentID,
			Children:  nil,
		})
	}

	// 序列化地址数

	return BuildAreaTree(area_list, 0)

}
