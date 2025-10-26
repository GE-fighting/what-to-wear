package services

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"what-to-wear/server/config"
	"what-to-wear/server/logger"
)

// OSSService OSS服务接口
type OSSService interface {
	// 生成预签名上传URL
	GeneratePresignedUploadURL(ctx context.Context, bucketName, objectKey string, expires time.Duration) (string, error)
	// 生成预签名下载URL
	GeneratePresignedDownloadURL(ctx context.Context, bucketName, objectKey string, expires time.Duration) (string, error)
	// 生成文件上传预签名URL（简化版本）
	GeneratePresignedURL(fileName, fileType string) (string, error)
	// 生成文件下载预签名URL
	GenerateDownloadURL(fileName string) (string, error)
}

// ossService OSS服务实现
type ossService struct {
	client *oss.Client
	config *config.Config
}

// NewOSSService 创建OSS服务实例
func NewOSSService(cfg *config.Config) (OSSService, error) {
	// 创建配置
	ossConfig := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.OSS.AccessKeyID, cfg.OSS.AccessKeySecret)).
		WithRegion(cfg.OSS.Region)

	// 创建OSS客户端
	client := oss.NewClient(ossConfig)

	return &ossService{
		client: client,
		config: cfg,
	}, nil
}

// GeneratePresignedUploadURL 生成预签名上传URL
func (s *ossService) GeneratePresignedUploadURL(ctx context.Context, bucketName, objectKey string, expires time.Duration) (string, error) {
	log := logger.GetLogger()
	log.Info("Generating presigned upload URL", logger.Fields{
		"bucket": bucketName,
		"key":    objectKey,
	})

	// 默认15分钟过期，最大7天
	if expires == 0 {
		expires = 15 * time.Minute
	}

	// 构建请求
	putObjectRequest := &oss.PutObjectRequest{
		Bucket: oss.Ptr(bucketName),
		Key:    oss.Ptr(objectKey),
	}

	// 生成预签名URL
	result, err := s.client.Presign(ctx, putObjectRequest, oss.PresignExpires(expires))
	if err != nil {
		log.ErrorWithErr(err, "Failed to generate presigned upload URL", logger.Fields{
			"bucket": bucketName,
			"key":    objectKey,
		})
		return "", fmt.Errorf("failed to generate presigned upload URL: %w", err)
	}

	log.Info("Presigned upload URL generated successfully", logger.Fields{
		"bucket":    bucketName,
		"key":       objectKey,
		"expiresIn": expires.String(),
	})

	return result.URL, nil
}

// GeneratePresignedDownloadURL 生成预签名下载URL
func (s *ossService) GeneratePresignedDownloadURL(ctx context.Context, bucketName, objectKey string, expires time.Duration) (string, error) {
	log := logger.GetLogger()
	log.Info("Generating presigned download URL", logger.Fields{
		"bucket": bucketName,
		"key":    objectKey,
	})

	// 默认15分钟过期，最大7天
	if expires == 0 {
		expires = 15 * time.Minute
	}

	// 构建请求
	getObjectRequest := &oss.GetObjectRequest{
		Bucket: oss.Ptr(bucketName),
		Key:    oss.Ptr(objectKey),
	}

	// 生成预签名URL
	result, err := s.client.Presign(ctx, getObjectRequest, oss.PresignExpires(expires))
	if err != nil {
		log.ErrorWithErr(err, "Failed to generate presigned download URL", logger.Fields{
			"bucket": bucketName,
			"key":    objectKey,
		})
		return "", fmt.Errorf("failed to generate presigned download URL: %w", err)
	}

	log.Info("Presigned download URL generated successfully", logger.Fields{
		"bucket":    bucketName,
		"key":       objectKey,
		"expiresIn": expires.String(),
	})

	return result.URL, nil
}

// GeneratePresignedURL 生成文件上传预签名URL（简化版本）
func (s *ossService) GeneratePresignedURL(fileName, fileType string) (string, error) {
	log := logger.GetLogger()
	log.Info("Generating presigned URL for file upload", logger.Fields{
		"fileName": fileName,
		"fileType": fileType,
	})

	// 生成唯一的文件名（可选，防止文件名冲突）
	timestamp := time.Now().Unix()
	ext := filepath.Ext(fileName)
	uniqueFileName := fmt.Sprintf("%d_%s", timestamp, fileName)
	if ext == "" && fileType != "" {
		uniqueFileName = fmt.Sprintf("%s.%s", uniqueFileName, fileType)
	}

	bucketName := s.config.OSS.BucketName

	// 使用配置中的过期时间
	expires := time.Duration(s.config.OSS.Expires) * time.Second

	result, err := s.GeneratePresignedUploadURL(context.Background(), bucketName, uniqueFileName, expires)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	log.Info("Presigned URL generated successfully", logger.Fields{
		"fileName":     fileName,
		"uniqueName":   uniqueFileName,
		"expiresInSec": s.config.OSS.Expires,
	})

	return result, nil
}

// GenerateDownloadURL 生成文件下载预签名URL
func (s *ossService) GenerateDownloadURL(fileName string) (string, error) {
	log := logger.GetLogger()
	log.Info("Generating download URL for file", logger.Fields{
		"fileName": fileName,
	})

	bucketName := s.config.OSS.BucketName

	// 使用配置中的过期时间
	expires := time.Duration(s.config.OSS.Expires) * time.Second

	result, err := s.GeneratePresignedDownloadURL(context.Background(), bucketName, fileName, expires)
	if err != nil {
		return "", fmt.Errorf("failed to generate download URL: %w", err)
	}

	log.Info("Download URL generated successfully", logger.Fields{
		"fileName":     fileName,
		"expiresInSec": s.config.OSS.Expires,
	})

	return result, nil
}
