package dto

// GeneratePresignedURLRequest 生成预签名URL请求
type GeneratePresignedURLRequest struct {
	FileName string `json:"file_name" binding:"required"` // 文件名
	FileType string `json:"file_type" binding:"required"` // 文件类型
}

// GeneratePresignedURLResponse 生成预签名URL响应
type GeneratePresignedURLResponse struct {
	URL string `json:"url"` // 预签名URL
}