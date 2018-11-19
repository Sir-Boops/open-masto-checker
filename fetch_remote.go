package main

import "net/http"
import "io/ioutil"

func fetch_remote_text(url string) string {

  resp, err := http.Get(url); // Fetch the remote text

  ans := "";

  if err == nil {
    raw_body, err := ioutil.ReadAll(resp.Body); // Convert the body to bytes
    if err == nil {
      ans = string(raw_body); // Convert bytes to char
    }
    defer resp.Body.Close();
  }
  return ans;
}
