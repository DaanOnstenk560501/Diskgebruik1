package main

import (
        "bufio"
        "fmt"
        "log"
        "os"
        "path/filepath"
        "strings"
)

func main() {
        logFile, err := os.OpenFile("folder_gebruik.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
        if err != nil {
                log.Fatal("Fout bij het openen van logbestand:", err)
        }
        defer logFile.Close()

        log.SetOutput(logFile)

        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Voer het pad in: ")
        path, _ := reader.ReadString('\n')
        path = strings.TrimSpace(path)

        log.Println("\nScan 1:") // Nieuwe scan markeren
        totalSize := getDocumentsSize(path, logFile)
        log.Println("Totale grootte van documentbestanden in", path, ":", formatBytes(totalSize))

        fmt.Println("Totale grootte van documentbestanden in", path, ":", formatBytes(totalSize))
        fmt.Println("\nDruk op Enter om af te sluiten...")
        bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func getDocumentsSize(drive string, logFile *os.File) int64 {
        var totalSize int64
        err := filepath.Walk(drive, func(path string, info os.FileInfo, err error) error {
                if err != nil {
                        log.Println("Fout bij het bezoeken van:", path, err)
                        return err
                }
                if !info.IsDir() {
                        log.Println("Bezocht bestand:", path, "extensie:", filepath.Ext(path))
                        if isDocumentFile(path) {
                                log.Println("Document bestand:", path)
                                totalSize += info.Size()
                        }
                }
                return nil
        })
        if err != nil {
                log.Println("Fout bij het doorlopen van de schijf:", err)
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
