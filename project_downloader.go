package main

import "fmt"
import "os"
import "io"
import "encoding/csv"
import "./projectManager"

// Download Drupal's project from a CSV file.
func main() {

    csvfile, err := os.Open("module-examples.csv")

    if err != nil {
        fmt.Println(err)
        return
    }

    defer csvfile.Close()

    reader := csv.NewReader(csvfile)

    for {
        record, err := reader.Read()

        if err == io.EOF {
            break
        } else if err!= nil {
            fmt.Println(err)
            return
        }

        p := projectManager.NewProject(record[0])
        p.DownloadProject("/tmp")
        break
    }
}
