package main

import (
        "bufio"
        "fmt"
        "os"
        "path/filepath"
        "strings"
)

func main() {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Voer het pad in: ")
        path, _ := reader.ReadString('\n')
        path = strings.TrimSpace(path)

        totalSize := getDocumentsSize(path)
        fmt.Printf("Totale grootte van documentbestanden in %s: %s\n", path, formatBytes(totalSize))

        fmt.Println("\nDruk op Enter om af te sluiten...")
        bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func getDocumentsSize(drive string) int64 {
        var totalSize int64
        err := filepath.Walk(drive, func(path string, info os.FileInfo, err error) error {
                if err != nil {
                        fmt.Println("Fout bij het ophalen van:", path, err)
                        return err
                }
                if !info.IsDir() {
                        fmt.Println("Bestand gevonden:", path, "extensie:", filepath.Ext(path)) // Log bezochte bestanden
                        if isDocumentFile(path) {
                                fmt.Println("Document bestand:", path)
                                totalSize += info.Size()
                        }
                }
                return nil
        })
        if err != nil {
                fmt.Println("Fout bij het ophalen van de schijf:", err)
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
