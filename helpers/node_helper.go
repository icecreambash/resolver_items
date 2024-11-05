package helpers

import (
	"AinedIndexChessCLI/db_models"
	"AinedIndexChessCLI/db_models/interfaces"
	"encoding/json"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

var (
	once     = sync.Once{}
	BuildCat map[string]string
)

func GetBuildingCategories() map[string]string {
	once.Do(func() {
		var categories = make(map[string]string)
		categories["mkd"] = "МКД"
		categories["stock"] = "Склад"
		categories["private_house"] = "Частный дом"
		categories["cottage"] = "Коттедж"
		categories["townhouse"] = "Таунхаус"
		categories["business"] = "Бизнес-центр"
		categories["mall"] = "Торговый центр"
		categories["parking"] = "Паркинг"
		BuildCat = categories
	})
	return BuildCat
}

func GetTitle(item db_models.Item, room db_models.Room) string {
	title := ""
	switch item.Category {
	case "flat":
		if room.Rooms != 0 {
			title = strconv.Itoa(room.Rooms) + "-к квартира"
		} else {
			title = "Квартира"
		}
		break
	case "office":
		title = "Офис"
		break
	case "parking_space":
		title = "Машиноместо"
		break
	}

	if room.AreaFull != 0 {
		title += " " + strconv.FormatFloat(room.AreaFull, 'f', 2, 64) + "м²"
	}
	if room.NumberObject != "" {
		title += " №" + room.NumberObject
	}
	return title
}

func GetSubtitle(item db_models.Item, DB *gorm.DB) string {
	subtitle := ""
	buildCat := GetBuildingCategories()
	var parent db_models.Tree
	DB.Table("trees").First(&parent, item.NodeID)
	if (parent != db_models.Tree{}) {
		if parent.ModelType == "Item" {

			jk := FindFirstComplex(parent, DB)
			if (jk != db_models.Complex{}) {
				subtitle = "ЖК " + jk.Name
			}
			if parent.ModelID.UUIDValue.String() != "" {
				var itemDB db_models.Item
				DB.Table("items").First(&itemDB, parent.ModelID.UUIDValue)
				if (itemDB != db_models.Item{}) {
					category := buildCat[itemDB.Category]
					if category != "" {
						if subtitle != "" {
							subtitle += " / "
						}
						subtitle += category
					}
					var build db_models.Building
					DB.Table("buildings").First(&build, itemDB.ModelID)
					if (build != db_models.Building{}) {
						if build.NumberObject != "" {
							subtitle += " №" + build.NumberObject
						}
					}
				}

			}

		}

	}

	return subtitle
}

func GetParents(item db_models.Item) {

}

func GetModelsByNode(nodes []db_models.Tree, DB *gorm.DB) []interfaces.Model {
	var models []interfaces.Model
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for _, node := range nodes {
		wg.Add(1)
		go func(node db_models.Tree) {
			defer wg.Done()
			model := GetModelByNode(node, DB)
			if model != nil {
				mu.Lock()
				models = append(models, model)
				mu.Unlock()
			}
		}(node)
	}
	wg.Wait()
	return models
}

func GetModelByNode(node db_models.Tree, DB *gorm.DB) interfaces.Model {
	switch node.ModelType {
	case "Item":
		var item db_models.Item
		DB.Table("items").First(&item, node.ModelID.UUIDValue)
		if (item != db_models.Item{}) {
			var build db_models.Building
			DB.Table("buildings").First(&build, item.ModelID)
			if (build != db_models.Building{}) {
				return build
			}
		}
		break
	case "BaseGroup":
		var item db_models.BaseGroup
		DB.Table("base_groups").First(&item, node.ModelID.IntValue)
		if (item != db_models.BaseGroup{}) {
			return item
		}
		break
	case "ComplexGroup":
		var item db_models.Complex
		DB.Table("complex_groups").First(&item, node.ModelID.IntValue)
		if (item != db_models.Complex{}) {
			return item
		}
		break
	}
	return nil
}

func FindFirstComplex(node db_models.Tree, DB *gorm.DB) db_models.Complex {
	var complexGroup interfaces.Model
	if node.ModelType != "ComplexGroup" {
		var nodes []db_models.Tree
		if node.ParentID != 0 {
			if err := DB.Find(&nodes, node.ParentID).Error; err != nil {
				panic(err)
			}
			var wg sync.WaitGroup
			for _, item := range nodes {
				wg.Add(1)
				go func(item db_models.Tree) {
					defer wg.Done()
					if item.ModelType != "ComplexGroup" {
						FindFirstComplex(item, DB)
					} else {
						complexGroup = GetModelByNode(item, DB)
						return
					}
				}(item)
			}
			wg.Wait()
		}

	} else {
		complexGroup = GetModelByNode(node, DB)
	}
	return ConvertToComplex(complexGroup)
}

func ConvertToComplex(jk interface{}) db_models.Complex {
	complexByte, err := json.Marshal(jk)
	if err != nil {
		panic(err)
	}
	var comGroup db_models.Complex
	if err := json.Unmarshal(complexByte, &comGroup); err != nil {
		panic(err)
	}
	return comGroup
}
