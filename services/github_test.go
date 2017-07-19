package services

import (
	"testing"
	"context"
	"github.com/google/go-github/github"
	"github.com/stretchr/testify/assert"
	"time"
)



func TestGithubService_GetIssuesWithLabels(t *testing.T) {
	ghs := NewGithubService()

	issues, _, err := ghs.client.Issues.ListByRepo(context.Background(), "facebook", "react", &github.IssueListByRepoOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	})

	//var dat map[string]interface{}
	//bytes, _ := ioutil.ReadAll(resp.Body)
	//if err := json.Unmarshal(bytes, &dat); err != nil {
	//	panic(err)
	//}

	//fmt.Println(dat)

	assert.Nil(t, err)
	assert.NotEmpty(t, issues)

	first := issues[0]
	time.Sleep(2 * time.Second)

	assert.NotNil(t, first.Repository)
}