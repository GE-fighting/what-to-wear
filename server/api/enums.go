package api

// AttachmentType 附件类型枚举
type AttachmentType string

const (
	AttachmentTypeImage AttachmentType = "image" // 图片
	AttachmentTypeVideo AttachmentType = "video" // 视频
	AttachmentTypeFile  AttachmentType = "file"  // 其他文件
)

// IsValid 检查附件类型是否有效
func (t AttachmentType) IsValid() bool {
	switch t {
	case AttachmentTypeImage, AttachmentTypeVideo, AttachmentTypeFile:
		return true
	default:
		return false
	}
}

// IsImage 检查是否为图片类型
func (t AttachmentType) IsImage() bool {
	return t == AttachmentTypeImage
}

// IsVideo 检查是否为视频类型
func (t AttachmentType) IsVideo() bool {
	return t == AttachmentTypeVideo
}

// EntityType 关联实体类型枚举
type EntityType string

const (
	EntityTypeClothingItem EntityType = "clothing_item" // 衣物
	EntityTypeOutfit       EntityType = "outfit"        // 穿搭
	EntityTypeUser         EntityType = "user"          // 用户
	EntityTypeMaintenance  EntityType = "maintenance"   // 保养记录
	EntityTypeWearRecord   EntityType = "wear_record"   // 穿着记录
	EntityTypePurchase     EntityType = "purchase"      // 购买记录
)

// IsValid 检查实体类型是否有效
func (t EntityType) IsValid() bool {
	switch t {
	case EntityTypeClothingItem, EntityTypeOutfit, EntityTypeUser,
		EntityTypeMaintenance, EntityTypeWearRecord, EntityTypePurchase:
		return true
	default:
		return false
	}
}

// ClothingStatus 衣物状态枚举
type ClothingStatus string

const (
	ClothingStatusActive   ClothingStatus = "active"   // 在用
	ClothingStatusInactive ClothingStatus = "inactive" // 闲置
	ClothingStatusDonated  ClothingStatus = "donated"  // 已捐赠
	ClothingStatusSold     ClothingStatus = "sold"     // 已出售
	ClothingStatusLost     ClothingStatus = "lost"     // 丢失
	ClothingStatusDamaged  ClothingStatus = "damaged"  // 损坏
)

// IsValid 检查衣物状态是否有效
func (s ClothingStatus) IsValid() bool {
	switch s {
	case ClothingStatusActive, ClothingStatusInactive, ClothingStatusDonated,
		ClothingStatusSold, ClothingStatusLost, ClothingStatusDamaged:
		return true
	default:
		return false
	}
}

// WeatherType 天气类型枚举
type WeatherType string

const (
	WeatherTypeSunny  WeatherType = "sunny"  // 晴天
	WeatherTypeRainy  WeatherType = "rainy"  // 雨天
	WeatherTypeCloudy WeatherType = "cloudy" // 多云
	WeatherTypeSnowy  WeatherType = "snowy"  // 雪天
	WeatherTypeFoggy  WeatherType = "foggy"  // 雾天
	WeatherTypeWindy  WeatherType = "windy"  // 大风
)

// IsValid 检查天气类型是否有效
func (w WeatherType) IsValid() bool {
	switch w {
	case WeatherTypeSunny, WeatherTypeRainy, WeatherTypeCloudy,
		WeatherTypeSnowy, WeatherTypeFoggy, WeatherTypeWindy:
		return true
	default:
		return false
	}
}

// Gender 性别枚举
type Gender string

const (
	GenderMale   Gender = "male"   // 男性
	GenderFemale Gender = "female" // 女性
	GenderOther  Gender = "other"  // 其他
)

// IsValid 检查性别是否有效
func (g Gender) IsValid() bool {
	switch g {
	case GenderMale, GenderFemale, GenderOther:
		return true
	default:
		return false
	}
}

// OutfitRating 穿搭评分枚举
type OutfitRating int

const (
	OutfitRatingPoor      OutfitRating = 1 // 很差
	OutfitRatingFair      OutfitRating = 2 // 一般
	OutfitRatingGood      OutfitRating = 3 // 良好
	OutfitRatingVeryGood  OutfitRating = 4 // 很好
	OutfitRatingExcellent OutfitRating = 5 // 优秀
)

// IsValid 检查评分是否有效
func (r OutfitRating) IsValid() bool {
	return r >= OutfitRatingPoor && r <= OutfitRatingExcellent
}

// TagType 标签类型枚举
type TagType string

const (
	TagTypeSeason   TagType = "season"   // 季节标签
	TagTypeOccasion TagType = "occasion" // 场合标签
	TagTypeStyle    TagType = "style"    // 风格标签
	TagTypeCustom   TagType = "custom"   // 自定义标签
)

// IsValidTagType 检查标签类型是否有效
func IsValidTagType(tagType string) bool {
	validTypes := []TagType{
		TagTypeSeason, TagTypeOccasion, TagTypeStyle,
		TagTypeCustom,
	}

	for _, validType := range validTypes {
		if TagType(tagType) == validType {
			return true
		}
	}
	return false
}

// MaintenanceType 保养类型枚举
type MaintenanceType string

const (
	MaintenanceWashing     MaintenanceType = "washing"      // 清洗
	MaintenanceDryCleaning MaintenanceType = "dry_cleaning" // 干洗
	MaintenanceRepair      MaintenanceType = "repair"       // 修补
	MaintenancePolishing   MaintenanceType = "polishing"    // 抛光
	MaintenanceWaterproof  MaintenanceType = "waterproof"   // 防水处理
	MaintenanceStorage     MaintenanceType = "storage"      // 收纳保存
	MaintenanceOther       MaintenanceType = "other"        // 其他
)

// IsValidMaintenanceType 检查保养类型是否有效
func IsValidMaintenanceType(maintenanceType string) bool {
	validTypes := []MaintenanceType{
		MaintenanceWashing, MaintenanceDryCleaning, MaintenanceRepair,
		MaintenancePolishing, MaintenanceWaterproof, MaintenanceStorage, MaintenanceOther,
	}

	for _, validType := range validTypes {
		if MaintenanceType(maintenanceType) == validType {
			return true
		}
	}
	return false
}

// ItemRole 单品角色枚举
type ItemRole string

const (
	ItemRoleMain      ItemRole = "main"      // 主要单品
	ItemRoleAccent    ItemRole = "accent"    // 点缀单品
	ItemRoleBase      ItemRole = "base"      // 基础单品
	ItemRoleOuter     ItemRole = "outer"     // 外层单品
	ItemRoleInner     ItemRole = "inner"     // 内层单品
	ItemRoleBottom    ItemRole = "bottom"    // 下装
	ItemRoleShoes     ItemRole = "shoes"     // 鞋子
	ItemRoleAccessory ItemRole = "accessory" // 配饰
)

// IsValidItemRole 检查单品角色是否有效
func IsValidItemRole(role string) bool {
	validRoles := []ItemRole{
		ItemRoleMain, ItemRoleAccent, ItemRoleBase, ItemRoleOuter,
		ItemRoleInner, ItemRoleBottom, ItemRoleShoes, ItemRoleAccessory,
	}

	for _, validRole := range validRoles {
		if ItemRole(role) == validRole {
			return true
		}
	}
	return false
}

// 系统标签枚举定义
// SystemTag 系统标签信息
type SystemTag struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// 季节标签枚举
const (
	SeasonSpring = "春季"
	SeasonSummer = "夏季"
	SeasonAutumn = "秋季"
	SeasonWinter = "冬季"
)

// 场合标签枚举
const (
	OccasionFormal  = "正式"
	OccasionLeisure = "休闲"
	OccasionSports  = "运动"
	OccasionDate    = "约会"
	OccasionWork    = "工作"
	OccasionParty   = "聚会"
	OccasionHome    = "居家"
)

// 风格标签枚举
const (
	StyleSimple   = "简约"
	StyleSweet    = "甜美"
	StyleCool     = "酷帅"
	StyleElegant  = "优雅"
	StyleRetro    = "复古"
	StyleStreet   = "街头"
	StyleLiterary = "文艺"
)

// 系统标签数据映射
var SystemTags = map[TagType][]SystemTag{
	TagTypeSeason: {
		{ID: 1, Name: SeasonSpring, Type: "season", Description: "适合春季穿着", Icon: ""},
		{ID: 2, Name: SeasonSummer, Type: "season", Description: "适合夏季穿着", Icon: ""},
		{ID: 3, Name: SeasonAutumn, Type: "season", Description: "适合秋季穿着", Icon: ""},
		{ID: 4, Name: SeasonWinter, Type: "season", Description: "适合冬季穿着", Icon: ""},
	},
	TagTypeOccasion: {
		{ID: 5, Name: OccasionFormal, Type: "occasion", Description: "正式场合", Icon: ""},
		{ID: 6, Name: OccasionLeisure, Type: "occasion", Description: "休闲场合", Icon: ""},
		{ID: 7, Name: OccasionSports, Type: "occasion", Description: "运动健身", Icon: ""},
		{ID: 8, Name: OccasionDate, Type: "occasion", Description: "约会场合", Icon: ""},
		{ID: 9, Name: OccasionWork, Type: "occasion", Description: "工作场合", Icon: ""},
		{ID: 10, Name: OccasionParty, Type: "occasion", Description: "聚会场合", Icon: ""},
		{ID: 11, Name: OccasionHome, Type: "occasion", Description: "居家休息", Icon: ""},
	},
	TagTypeStyle: {
		{ID: 12, Name: StyleSimple, Type: "style", Description: "简约风格", Icon: ""},
		{ID: 13, Name: StyleSweet, Type: "style", Description: "甜美风格", Icon: ""},
		{ID: 14, Name: StyleCool, Type: "style", Description: "酷帅风格", Icon: ""},
		{ID: 15, Name: StyleElegant, Type: "style", Description: "优雅风格", Icon: ""},
		{ID: 16, Name: StyleRetro, Type: "style", Description: "复古风格", Icon: ""},
		{ID: 17, Name: StyleStreet, Type: "style", Description: "街头风格", Icon: ""},
		{ID: 18, Name: StyleLiterary, Type: "style", Description: "文艺风格", Icon: ""},
	},
}

// GetSystemTagsByType 根据标签类型获取系统标签枚举
func GetSystemTagsByType(tagType TagType) []SystemTag {
	if tags, exists := SystemTags[tagType]; exists {
		return tags
	}
	return []SystemTag{}
}

// GetAllSystemTags 获取所有系统标签
func GetAllSystemTags() map[TagType][]SystemTag {
	return SystemTags
}

// IsSystemTag 检查是否为系统标签
func IsSystemTag(tagName string, tagType TagType) bool {
	if tags, exists := SystemTags[tagType]; exists {
		for _, tag := range tags {
			if tag.Name == tagName {
				return true
			}
		}
	}
	return false
}

// GetSystemTagByName 根据名称和类型获取系统标签信息
func GetSystemTagByName(tagName string, tagType TagType) (*SystemTag, bool) {
	if tags, exists := SystemTags[tagType]; exists {
		for _, tag := range tags {
			if tag.Name == tagName {
				return &tag, true
			}
		}
	}
	return nil, false
}
