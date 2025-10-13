package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type IAssetUtil interface {
	UploadBase64Asset(base64Str, folder, name string) (string, error)
	UploadFileAsset(name, folder string, fileHeader *multipart.FileHeader) (string, error)
	DeleteAsset(folder, name string) error
	ProjectRootDir() (string, error)
}

// UploadBase64Asset - upload assets (file: as base64 string)
// example usage:
//
//	savedPath, err := UploadBase64Asset("xxxxxxx", "/products", "new-name")
//
// example with gin:
//
//				engine.POST("/assets-mgmt", func(ctx *gin.Context) {
//	             // Get uploaded file
//	             var form struct {
//	                File string `json:"file"
//	             }
//
//	              // bind user input
//	              if err := ctx.ShouldBind(&form); err != nil {
//	                 ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
//	                 return
//	               }
//
//	                // Call UploadAsset
//	                fn :=strconv.FormatInt(time.Now().UnixMicro(), 10)
//	                savedPath, err := utils.UploadBase64Asset(form.file,
//	                	ctx.Query("folder"), fn)
//	                if err != nil {
//	                	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//	                	return
//	                }
//
//	                scheme := "http"
//	                if ctx.Request.TLS != nil {
//	                    scheme = "https"
//	                }
//	                ctx.JSON(http.StatusOK, gin.H{
//	                    "message": "upload successful",
//	                    "file_path": fmt.Sprintf("%s://%s/assets/%s",
//	                        scheme, ctx.Request.Host, savedPath),
//	                })
//		 })
func UploadBase64Asset(base64Str, folder, name string) (string, error) {
	// Split "data:image/png;base64,..."
	parts := strings.SplitN(base64Str, ",", 2)
	if len(parts) != 2 {
		return "", errors.New("invalid base64 image")
	}

	data, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}

	// Determine extension from MIME
	mimeType := strings.TrimPrefix(strings.SplitN(parts[0], ";", 2)[0], "data:")
	ext := ""
	switch mimeType {
	case "image/png":
		ext = ".png"
	case "image/jpeg":
		ext = ".jpg"
	default:
		return "", errors.New("unsupported MIME type")
	}

	// Ensure upload folder exists
	root, err := ProjectRootDir()
	if err != nil {
		return "", err
	}
	uploadDir := filepath.Join(root, "uploads", folder)
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", err
	}

	// Write temp file
	filePath := filepath.Join(uploadDir, name+ext)
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return "", err
	}

	// return filePath, nil
	finalName := name + ext
	return fmt.Sprintf("%s/%s", folder, finalName), nil
}

// UploadFileAsset - upload assets (file: pdf, excel | image: png, img)
//
// example usage:
//
//	savedPath, err := UploadFileAsset("product-sku-123", "/products", fileHeader)
//
// example with gin:
//
//		    engine.POST("/assets-mgmt", func(ctx *gin.Context) {
//			    // Get uploaded file
//			    fileHeader, err := ctx.FormFile("file")
//			    if err != nil {
//			    	ctx.JSON(http.StatusBadRequest, gin.H{"error": "file not provided"})
//			    	return
//			    }
//
//			    // Call UploadAsset
//			    savedPath, err := utils.UploadFileAsset(ctx.Query("name"),
//			    	ctx.Query("folder"), fileHeader)
//			    if err != nil {
//			    	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//			    	return
//			    }
//
//		        scheme := "http"
//		        if ctx.Request.TLS != nil {
//		            scheme = "https"
//		        }
//		        ctx.JSON(http.StatusOK, gin.H{
//		            "message": "upload successful",
//		            "file_path": fmt.Sprintf("%s://%s/assets/%s",
//		                scheme, ctx.Request.Host, savedPath),
//		        })
//	     })
func UploadFileAsset(name, folder string, fileHeader *multipart.FileHeader) (string, error) {
	// Clean and build full path relative to project root's ./uploads
	root, err := ProjectRootDir()
	if err != nil {
		return "", fmt.Errorf("cannot resolve project root: %w", err)
	}
	basePath := filepath.Join(root, "uploads")
	relativePath := filepath.Clean(folder)
	uploadPath := filepath.Join(basePath, relativePath)

	// Ensure directory exists
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
			return "", fmt.Errorf("failed to create upload directory: %w", err)
		}
	}

	// Open uploaded file
	src, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer func() { _ = src.Close() }()

	// Extract original file extension
	ext := filepath.Ext(fileHeader.Filename)
	if ext == "" {
		return "", fmt.Errorf("uploaded file has no extension")
	}

	// Combine custom name with extension
	finalName := name + ext

	// Create destination file
	dstPath := filepath.Join(uploadPath, finalName)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %w", err)
	}
	defer func() { _ = dst.Close() }()

	// Save file
	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	// Return saved path (relative or full)
	return fmt.Sprintf("%s/%s", folder, finalName), nil
}

// DeleteAsset - remove assets from folder
//
// example with gin:
//
//	    engine.DELETE("/assets-mgmt", func(ctx *gin.Context) {
//		    if err := utils.DeleteAsset(
//		    	ctx.Query("folder"),
//		    	ctx.Query("name"),
//		    ); err != nil {
//		    	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		    }
//		    ctx.JSON(http.StatusOK, gin.H{"message": "delete successful"})
//	    })
func DeleteAsset(folder, name string) error {
	root, err := ProjectRootDir()
	if err != nil {
		return fmt.Errorf("cannot resolve project root: %w", err)
	}
	basePath := filepath.Join(root, "uploads")

	// Prevent absolute paths and directory traversal
	if strings.Contains(folder, "..") || strings.Contains(name, "..") {
		return fmt.Errorf("invalid path: path traversal detected")
	}
	if filepath.IsAbs(folder) || filepath.IsAbs(name) {
		return fmt.Errorf("invalid path: absolute paths not allowed")
	}

	relativePath := filepath.Clean(folder)
	filePath := filepath.Join(basePath, relativePath, name)

	// Double-check it's still inside the uploads directory
	if !strings.HasPrefix(filePath, basePath) {
		return fmt.Errorf("invalid path: outside uploads directory")
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", filePath)
	}

	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

func ProjectRootDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found")
		}
		dir = parent
	}
}
