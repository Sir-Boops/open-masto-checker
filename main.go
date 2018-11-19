package main

import "fmt"
import "strings"

func main() {

  // Download the list from git
  git_url := "https://git.sergal.org/Sir-Boops/yiff-social-instances/raw/branch/master/list.txt"
  url_list := strings.Split(fetch_remote_text(git_url), "\n"); // Convert the string to a []string
  url_list = url_list[:len(url_list) - 1] // Remove the last line

  for i := 0; i < len(url_list); i++ {
    if check_if_open(url_list[i]) {
      fmt.Println("<a href='https://" + url_list[i] + "/about'>" + url_list[i] + "</a><br>");
    }
  }
}
