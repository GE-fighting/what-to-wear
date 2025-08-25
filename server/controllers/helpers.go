package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"what-to-wear/server/api"
	"what-to-wear/server/api/errors"
)

// getUserID 从上下文获取用户ID
func getUserID(c *gin.Context) uint {
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		return 0
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		return 0
	}

	return userID
}

// getUserIDRequired 获取用户ID，如果不存在则返回错误响应
func getUserIDRequired(c *gin.Context) (uint, bool) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, api.Unauthorized("未授权访问"))
		return 0, false
	}
	return userID, true
}

// parseUintParam 解析URL参数为uint类型
func parseUintParam(c *gin.Context, paramName string) (uint, error) {
	paramStr := c.Param(paramName)
	if paramStr == "" {
		return 0, errors.ErrInvalidRequest("缺少参数: " + paramName)
	}

	value, err := strconv.ParseUint(paramStr, 10, 32)
	if err != nil {
		return 0, errors.ErrInvalidRequest("无效的参数: " + paramName)
	}

	return uint(value), nil
}

// parseUintParamRequired 解析必需的URL参数，如果解析失败则返回错误响应
func parseUintParamRequired(c *gin.Context, paramName string) (uint, bool) {
	value, err := parseUintParam(c, paramName)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.BadRequest(err.Error()))
		return 0, false
	}
	return value, true
}

// parseIntQuery 解析查询参数为int类型，带默认值
func parseIntQuery(c *gin.Context, queryName string, defaultValue int) int {
	queryStr := c.DefaultQuery(queryName, strconv.Itoa(defaultValue))
	value, err := strconv.Atoi(queryStr)
	if err != nil {
		return defaultValue
	}
	return value
}

// parseFloatQuery 解析查询参数为float64类型，带默认值
func parseFloatQuery(c *gin.Context, queryName string, defaultValue float64) float64 {
	queryStr := c.DefaultQuery(queryName, strconv.FormatFloat(defaultValue, 'f', -1, 64))
	value, err := strconv.ParseFloat(queryStr, 64)
	if err != nil {
		return defaultValue
	}
	return value
}

// parseBoolQuery 解析查询参数为bool类型，带默认值
func parseBoolQuery(c *gin.Context, queryName string, defaultValue bool) bool {
	queryStr := c.DefaultQuery(queryName, strconv.FormatBool(defaultValue))
	value, err := strconv.ParseBool(queryStr)
	if err != nil {
		return defaultValue
	}
	return value
}

// validatePagination 验证并设置分页参数的默认值
func validatePagination(page, pageSize *int) {
	if *page <= 0 {
		*page = 1
	}
	if *pageSize <= 0 {
		*pageSize = 20
	}
	if *pageSize > 100 {
		*pageSize = 100
	}
}

// isValidTagType 验证标签类型
func isValidTagType(tagType string) bool {
	validTypes := []string{"season", "occasion", "style", "color", "material", "brand", "custom"}
	for _, validType := range validTypes {
		if tagType == validType {
			return true
		}
	}
	return false
}

// isValidSortOrder 验证排序方向
func isValidSortOrder(sortOrder string) bool {
	return sortOrder == "asc" || sortOrder == "desc"
}

// isValidClothingSortBy 验证衣物排序字段
func isValidClothingSortBy(sortBy string) bool {
	validFields := []string{"name", "price", "wear_count", "durability_score", "created_at", "updated_at", "brand", "color"}
	for _, field := range validFields {
		if sortBy == field {
			return true
		}
	}
	return false
}
