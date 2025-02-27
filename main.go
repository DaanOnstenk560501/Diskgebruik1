package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	drive := "C:\\Users\\Daan\\Documents"
	totalSize := getDocumentsSize(drive)
	fmt.Printf("Totale grootte van documentbestanden in %s: %s\n", drive, formatBytes(totalSize))
}

func getDocumentsSize(drive string) int64 {
	var totalSize int64
	err := filepath.Walk(drive, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Fout bij het bezoeken van:", path, err)
			return err
		}
		if !info.IsDir() && !strings.HasPrefix(filepath.Base(path), ".") {
			fmt.Println("Bezocht bestand:", path)
			if isDocumentFile(path) {
				fmt.Println("Document bestand:", path)
				totalSize += info.Size()
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Fout bij het doorlopen van de schijf:", err)
	}
	return totalSize
}

func isDocumentFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".doc" || ext == ".docx" || ext == ".pdf" || ext == ".txt" || ext == ".odt" || ext == ".go"
}

func formatBytes(bytes int64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	switch {
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/GB)
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/MB)
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/KB)
	default:
		return fmt.Sprintf("%d bytes", bytes)
	}
}
