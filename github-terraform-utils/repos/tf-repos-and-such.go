package main

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/google/go-github/github"
)

func main() {
    ts := github.BasicAuthTransport{
        Username:  "",
        Password:  "",
        Transport: nil,
    }

    teamIds := make([]int64, 4)
    teamIds[0] = 4947840
    teamIds[1] = 5494169
    teamIds[2] = 3291385
    teamIds[3] = 4769796

    client := github.NewClient(ts.Client())
    ctx := context.Background()
    for _, id := range teamIds {
        team, _, err := client.Teams.GetTeam(ctx, id)
        if err != nil {
            fmt.Println(err)
        }
        marshalled, _ := json.MarshalIndent(team, "", "\t")
        fmt.Println("team: ", string(marshalled))
    }
    //repositories, _, err := client.Repositories.List(context.Background(), "", &github.RepositoryListOptions{
    //    ListOptions: github.ListOptions{
    //        PerPage: 100,
    //    },
    //})
    //if err != nil {
    //    fmt.Println(err)
    //    return
    //}

    //file, err := os.OpenFile("main.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, fs.ModeAppend)
    //if err != nil {
    //    fmt.Println("error opening main.tf", err)
    //    return
    //}

    //builder := strings.Builder{}
    //for _, repository := range repositories {
    //    //sprintf := fmt.Sprintf("module \"%s\" {\n  source = \"./modules/repository\"\n\n  repository_name = \"%s\" \n  writer_teams    = []", *repository.Name, *repository.Name)
    //    //builder.WriteString(sprintf)
    //    //builder.WriteString("\n } \n\n")
    //    //file.WriteString(builder.String())
    //
    //    //sprintf := fmt.Sprintf(
    //    //   "terraform import module.%s.github_repository.repository %s \n",
    //    //   *repository.Name,
    //    //   *repository.Name,
    //    //)
    //
    //    sprintf := fmt.Sprintf("terraform import module.%s.github_branch_default.default-branch %s \n", *repository.Name, *repository.Name)
    //    //sprintf := fmt.Sprintf("terraform import module.%s.github_branch_protection_v3.approver-teams %s:%s \n", repository.GetName(), repository.GetName(), repository.GetDefaultBranch())
    //    builder.WriteString(sprintf)
    //}
    //
    //err = ioutil.WriteFile("import.sh", []byte(builder.String()), fs.ModePerm)
    //if err != nil {
    //    fmt.Println(err)
    //    return
    //}
}
