package main

import "time"
import "strings"
import "net/http"
import "io/ioutil"

// Check if instance is open or not
// True: open
// False: closed
func check_if_open(url string) bool {

  client := http.Client{Timeout: time.Duration(5 * time.Second)} // Create a custom client
  resp, err := client.Get("https://" + url + "/about"); // Request the instance with the custom client

  ans := false; // Always assume it's closed

  if err == nil {
    raw_body, err := ioutil.ReadAll(resp.Body); // Parse the body to bytes
    if err == nil {
      if !strings.Contains(string(raw_body), "closed-registrations-message") {
        ans = true;
      }
    }
    defer resp.Body.Close(); // Close the connection
  }
  return ans;
}
