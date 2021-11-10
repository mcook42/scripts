package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    file, err := ioutil.ReadFile("teams.json")
    if err != nil {
        fmt.Println("Error opening teams.json ", err)
        return
    }
    fmt.Println("read in teams.json")

    var teams []Team
    err = json.Unmarshal(file, &teams)
    if err != nil {
        fmt.Println("Error unmarshalling teams.json ", err)
        return
    }

    tfBuilder := strings.Builder{}
    for _, team := range teams {
        outerName := strings.Replace(team.Name, " ", "-", -1)
        resourceHeader := fmt.Sprintf("resource \"github_team\" \"%s\" {\n  name                      = \"%s\" \n privacy = \"%s\" \n description = \"%s\" \n", strings.ToLower(outerName), team.Name, team.Privacy, team.Description)
        tfBuilder.WriteString(resourceHeader)
        tfBuilder.WriteString("}\n\n")
    }
    err = ioutil.WriteFile("out.tf", []byte(tfBuilder.String()), 0777)
    if err != nil {
        fmt.Println("Error writing out file ", err)
        return
    }
    scriptBuilder := strings.Builder{}

    for _, team := range teams {
        outerName := strings.Replace(team.Name, " ", "-", -1)
        command := fmt.Sprintf("terraform import github_team.%s %d \n", strings.ToLower(outerName), team.ID)
        scriptBuilder.WriteString(command)
    }

    err = ioutil.WriteFile("out.sh", []byte(scriptBuilder.String()), 0777)
    if err != nil {
        fmt.Println("Error writing out file ", err)
        return
    }

}

type Team struct {
    Name            string `json:"name,omitempty"`
    ID              int    `json:"id,omitempty"`
    NodeId          string `json:"node_id,omitempty"`
    Slug            string `json:"slug,omitempty"`
    Description     string `json:"description,omitempty"`
    Privacy         string `json:"privacy,omitempty"`
    URL             string `json:"url,omitempty"`
    HTMLUrl         string `json:"html_url,omitempty"`
    MembersUrl      string `json:"members_url,omitempty"`
    RepositoriesUrl string `json:"repositories_url,omitempty"`
    Permission      string `json:"permission,omitempty"`
    Parent          string `json:"parent,omitempty"`
}
