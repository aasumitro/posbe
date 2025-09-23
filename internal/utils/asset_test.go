package utils_test

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"

	"github.com/aasumitro/posbe/internal/utils"
)

func TestUploadAndDeleteAssetsUtils(t *testing.T) {
	// Prepare dummy file content
	fileContent := []byte("test file content for UploadAsset")
	tempFile, err := os.CreateTemp("", "test-upload-*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer func() { _ = os.Remove(tempFile.Name()) }()
	if _, err := tempFile.Write(fileContent); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	_ = tempFile.Close()

	// Reopen for multipart file header simulation
	srcFile, err := os.Open(tempFile.Name())
	if err != nil {
		t.Fatalf("failed to reopen temp file: %v", err)
	}
	defer func() { _ = srcFile.Close() }()

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	part, err := writer.CreateFormFile("file", filepath.Base(tempFile.Name()))
	if err != nil {
		t.Fatalf("failed to create form file: %v", err)
	}
	if _, err := io.Copy(part, srcFile); err != nil {
		t.Fatalf("failed to copy file content: %v", err)
	}
	_ = writer.Close()

	reader := multipart.NewReader(&b, writer.Boundary())
	form, err := reader.ReadForm(10 << 20)
	if err != nil {
		t.Fatalf("failed to parse multipart form: %v", err)
	}
	fileHeader := form.File["file"][0]

	// Upload
	name := "test-upload-renamed"
	folder := "test-products"
	savedPath, err := utils.UploadAsset(name, folder, fileHeader)
	if err != nil {
		t.Fatalf("UploadAsset failed: %v", err)
	}

	// Check file exists
	root, err := utils.ProjectRootDir()
	if err != nil {
		t.Fatalf("cannot resolve project root: %v", err)
	}
	absPath := filepath.Join(root, "uploads", savedPath)
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		t.Fatalf("uploaded file does not exist at path: %s", absPath)
	}

	// Check file content
	uploadedContent, err := os.ReadFile(absPath)
	if err != nil {
		t.Fatalf("failed to read uploaded file: %v", err)
	}
	if !bytes.Equal(uploadedContent, fileContent) {
		t.Errorf("uploaded file content does not match original")
	}

	// Delete
	err = utils.DeleteAsset(folder, name+filepath.Ext(fileHeader.Filename))
	if err != nil {
		t.Fatalf("DeleteAsset failed: %v", err)
	}

	// Verify file is deleted
	if _, err := os.Stat(savedPath); err == nil || !os.IsNotExist(err) {
		t.Errorf("file was not deleted properly")
	}

	_ = os.RemoveAll(filepath.Join(root, "uploads", folder))
}
