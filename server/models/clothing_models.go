package models

// 这个文件用于定义所有衣物相关模型的关联关系，避免循环依赖

// 为了避免循环依赖，我们将在这里重新定义完整的关联关系

// ClothingItemWithRelations 包含完整关联关系的衣物模型
type ClothingItemWithRelations struct {
	ClothingItem

	// 关联关系 - 仅在需要时通过Preload加载
	User               *User               `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Category           *ClothingCategory   `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	Tags               []ClothingTag       `json:"tags,omitempty" gorm:"many2many:clothing_item_tags;"`
	PurchaseRecord     *PurchaseRecord     `json:"purchase_record,omitempty" gorm:"foreignKey:ClothingItemID"`
	MaintenanceRecords []MaintenanceRecord `json:"maintenance_records,omitempty" gorm:"foreignKey:ClothingItemID"`
	WearRecords        []WearRecord        `json:"wear_records,omitempty" gorm:"foreignKey:ClothingItemID"`
	OutfitItems        []OutfitItem        `json:"outfit_items,omitempty" gorm:"foreignKey:ClothingItemID"`
}

// ClothingCategoryWithRelations 包含完整关联关系的分类模型
type ClothingCategoryWithRelations struct {
	ClothingCategory

	ClothingItems []ClothingItem `json:"clothing_items,omitempty" gorm:"foreignKey:CategoryID"`
}

// ClothingTagWithRelations 包含完整关联关系的标签模型
type ClothingTagWithRelations struct {
	ClothingTag

	User          *User          `json:"user,omitempty" gorm:"foreignKey:UserID"`
	ClothingItems []ClothingItem `json:"clothing_items,omitempty" gorm:"many2many:clothing_item_tags;"`
}

// OutfitWithRelations 包含完整关联关系的穿搭模型
type OutfitWithRelations struct {
	Outfit

	User        *User        `json:"user,omitempty" gorm:"foreignKey:UserID"`
	OutfitItems []OutfitItem `json:"outfit_items,omitempty" gorm:"foreignKey:OutfitID"`
}

// OutfitItemWithRelations 包含完整关联关系的穿搭单品模型
type OutfitItemWithRelations struct {
	OutfitItem

	Outfit       *Outfit       `json:"outfit,omitempty" gorm:"foreignKey:OutfitID"`
	ClothingItem *ClothingItem `json:"clothing_item,omitempty" gorm:"foreignKey:ClothingItemID"`
}
