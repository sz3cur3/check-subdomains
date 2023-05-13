package main

import (
    "bufio"
    "flag"
    "fmt"
    "net"
    "os"
)

func main() {
    // Define command-line flags
    inputFilePath := flag.String("i", "", "input file path")
    outputFilePath := flag.String("o", "", "output file path")
    flag.Parse()

    // Open input file for reading
    inputFile, err := os.Open(*inputFilePath)
    if err != nil {
        fmt.Printf("error opening input file: %v\n", err)
        os.Exit(1)
    }
    defer inputFile.Close()

    // Open output file for writing
    outputFile, err := os.Create(*outputFilePath)
    if err != nil {
        fmt.Printf("error creating output file: %v\n", err)
        os.Exit(1)
    }
    defer outputFile.Close()

    // Read input file line by line
    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        subdomain := scanner.Text()

        // Resolve subdomain using DNS
        ips, err := net.LookupIP(subdomain)
        if err != nil {
            fmt.Fprintf(outputFile, "%s cannot be resolved\n", subdomain)
        } else {
            for _, ip := range ips {
                fmt.Fprintf(outputFile, "%s resolves to %s\n", subdomain, ip)
            }
        }
    }
    if err := scanner.Err(); err != nil {
        fmt.Printf("error reading input file: %v\n", err)
        os.Exit(1)
    }
}
