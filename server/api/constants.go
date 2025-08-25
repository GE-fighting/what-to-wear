package api

// HTTP 状态码常量
const (
	StatusOK                  = 200
	StatusCreated             = 201
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusUnprocessableEntity = 422
	StatusInternalServerError = 500
)

// 分页相关常量
const (
	DefaultPageSize = 20
	MaxPageSize     = 100
	DefaultPage     = 1
)

// 文件上传相关常量
const (
	MaxFileSize     = 10 * 1024 * 1024 // 10MB
	MaxImageSize    = 5 * 1024 * 1024  // 5MB
	MaxVideoSize    = 50 * 1024 * 1024 // 50MB
	AllowedImageExt = "jpg,jpeg,png,gif,webp"
	AllowedVideoExt = "mp4,avi,mov,wmv"
)

// 缓存相关常量
const (
	CacheKeyPrefix     = "wtw:"
	CacheExpireDefault = 3600 // 1小时
	CacheExpireShort   = 300  // 5分钟
	CacheExpireLong    = 86400 // 24小时
)

// JWT相关常量
const (
	JWTExpireHours = 24 * 7 // 7天
	JWTIssuer      = "what-to-wear"
)

// 业务规则常量
const (
	MaxOutfitItems    = 20  // 一套穿搭最多20件衣物
	MaxTagsPerItem    = 10  // 每件衣物最多10个标签
	MaxUserClothings  = 1000 // 用户最多1000件衣物
)