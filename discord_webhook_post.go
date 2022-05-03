package main

import (
    "io/ioutil"
    "log"
    "fmt"
    "net/http"
    "strings"
    "os"
)

func post_to_webhook(message string, webhook_url string) {

    // Standard net Client
    client := &http.Client{}

    // Using Sprintlin to easily concatenate strings.
    message_str := fmt.Sprintln(`{"content":"`, message,`"}`)
    fmt.Printf(message_str)

    var data = strings.NewReader(message_str)

    // Sending the request
    req, err := http.NewRequest("POST", webhook_url, data)
    if err != nil {
        log.Fatal(err)
    }

    req.Header.Set("Accept", "application/json")
    req.Header.Set("Content-Type", "application/json")
    resp, err := client.Do(req)

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()
    bodyText, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    log.Printf("%s\n", bodyText)
}

func main () {
    // We expect just one string argument ( the user's message )
    args := os.Args
    if len(args) != 2 {
        log.Fatal("Please provide your message in quotations.")
        return
    }
    post_to_webhook(args[1], os.Getenv("DISCORD_WEBHOOK"))
}
