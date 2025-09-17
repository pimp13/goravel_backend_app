package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/goravel/framework/contracts/filesystem"
	"github.com/goravel/framework/facades"
)

type UploadService interface {
	UploadImage(file filesystem.File) (string, error)
}

type uploadService struct{}

func NewUploadService() UploadService {
	return &uploadService{}
}

func (s *uploadService) UploadImage(file filesystem.File) (string, error) {
	uploadPath := "uploads"
	fullPath := facades.App().PublicPath(uploadPath)

	uniqueFileName := fmt.Sprintf("%s_%d", uuid.NewString(), time.Now().Unix())

	filePath, err := facades.Storage().PutFileAs(fullPath, file, uniqueFileName)
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	return filePath, nil
}

/*
func (s *uploadService) UploadImage(file *multipart.FileHeader, destDir string) (string, error) {
	// ساخت پوشه اگر وجود نداشت
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
			return "", err
		}
	}

	// تولید نام یونیک برای فایل
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s_%d%s", uuid.NewString(), time.Now().Unix(), ext)
	filePath := filepath.Join(destDir, filename)

	// ذخیره فایل
	if err := saveUploadedFile(file, filePath); err != nil {
		return "", err
	}

	return filename, nil
}

// تابع کمکی برای ذخیره فایل
func saveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.ReadFrom(src)
	return err
}
*/
