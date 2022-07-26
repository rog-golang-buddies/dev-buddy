package githubBroker

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/rog-golang-buddies/dev-buddy/internal/pkg/constants"
	"golang.org/x/oauth2"
)

func CreateBroker(ctx context.Context) (*github.Client, error) {
	// Get PAT from ctx.
	personalAccessToken := ctx.Value(constants.GHPATHeader)

	// Set a Static Token Source and create Client.
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: fmt.Sprint(personalAccessToken)},
	)
	tc := oauth2.NewClient(ctx, ts)

	// create github client.
	client := github.NewClient(tc)

	return client, nil
}

func GetAllIssueNames(ctx context.Context, client *github.Client, repositoryName string) (string, error) {
	// create search options
	opts := &github.SearchOptions{Sort: "created", Order: "asc"}

	// get the githubOwner
	githubOwner := fmt.Sprint(ctx.Value(constants.OwnerKey))

	// create search string
	// docs.github.com/en/search-github/searching-on-github/searching-issues-and-pull-requests
	searchString := "is:issue is:open repo:" + githubOwner + "/" + repositoryName

	// get the github issues
	githubIssues, _, err := client.Search.Issues(ctx, searchString, opts)
	if err != nil {
		return "", err
	}

	issuesList := githubIssues.Issues

	var returnIssuesList string

	// add all the issues with formatting to be returned
	for position, issue := range issuesList {
		returnIssuesList += fmt.Sprintln(position, *issue.Title)
	}
	return returnIssuesList, nil
}

func CreateOrganizationInvite(ctx context.Context, client *github.Client, username string) (string, error) {
	// var response string
	// create search options
	opts := &github.SearchOptions{Sort: "created", Order: "asc"}

	githubUsername, _, err := client.Search.Users(ctx, username, opts)
	if err != nil {
		return "", err
	}
	githubUser := githubUsername.Users

	// if len(githubUser) > 1 {
	// 	return "more than one user found, please recheck", nil
	// }
	githubUserId := githubUser[0].ID
	teamIDList := []int64{6379585}
	roleString := "direct_member"
	inviteOpts := github.CreateOrgInvitationOptions{
		InviteeID: githubUserId,
		Role:      &roleString,
		TeamID:    teamIDList,
	}

	organizationName := fmt.Sprint(ctx.Value(constants.OwnerKey))
	invitation, _, err := client.Organizations.CreateOrgInvitation(ctx, organizationName, &inviteOpts)
	if err != nil {
		return "", err
	}
	response := "Invite created for the user " + username + " at " + invitation.CreatedAt.String()
	return response, nil
}
