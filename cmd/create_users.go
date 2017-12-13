// Script to create 200 users and associate them to a team

package main

import (
  "fmt"
  "os/exec"
)

func create_command(username string, email string) {
  cmd := exec.Command("./platform", "user", "create", "--email", email, "--username", username, "--password", "Password1")
  out, err := cmd.CombinedOutput()
  if err != nil {
    fmt.Println("cmd.run failed")
  }
  fmt.Printf("combined out:\n%s\n", string(out))
}

func team_command(username string) {
  cmd := exec.Command("./platform", "team", "add", "test-team", username)
  out, err := cmd.CombinedOutput()
  if err != nil {
    fmt.Println("cmd.run failed")
    
  }
  fmt.Printf("combined out:\n%s\n", string(out))
}

func create() {
  for i := 200; i < 1000; i++ {
    var username string = fmt.Sprintf("user%d", i)
    var email string = fmt.Sprintf("%s@example.com", username)
    
    fmt.Println(username)
    fmt.Println(email)
    create_command(username, email)
  }
}

func team() {
  for i := 200; i < 1000; i++ {
    var username string = fmt.Sprintf("user%d", i)
    team_command(username)
  }
}

func main() {
  // create_command("usernamewow", "fffoo@example.com")
  // team_command("user50")
  create()
  team()
}
