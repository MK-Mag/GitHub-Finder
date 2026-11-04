package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const githubAPIBaseURL = "https://api.github.com"

type GitHubUser struct {
	Login       string `json:"login"`
	ID          int    `json:"id"`
	AvatarURL   string `json:"avatar_url"`
	HTMLURL     string `json:"html_url"`
	Name        string `json:"name"`
	Company     string `json:"company"`
	Blog        string `json:"blog"`
	Location    string `json:"location"`
	Bio         string `json:"bio"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
	CreatedAt   string `json:"created_at"`
}

type Repository struct {
	Name            string `json:"name"`
	FullName        string `json:"full_name"`
	Description     string `json:"description"`
	HTMLURL         string `json:"html_url"`
	StargazersCount int    `json:"stargazers_count"`
	ForksCount      int    `json:"forks_count"`
	Language        string `json:"language"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type App struct {
	ctx    context.Context
	client *http.Client
}

func NewApp() *App {
	return &App{
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) requestGitHub(endpoint string, target any) error {
	ctx := a.ctx
	if ctx == nil {
		ctx = context.Background()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, githubAPIBaseURL+endpoint, nil)
	if err != nil {
		return fmt.Errorf("create GitHub request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := a.client
	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request GitHub API: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		switch res.StatusCode {
		case http.StatusNotFound:
			return fmt.Errorf("GitHub user or resource was not found")
		case http.StatusForbidden:
			return fmt.Errorf("GitHub API rate limit reached or request forbidden")
		default:
			return fmt.Errorf("GitHub API returned status %d", res.StatusCode)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return fmt.Errorf("decode GitHub response: %w", err)
	}

	return nil
}

// FetchingUser returns the public profile for a GitHub username.
func (a *App) FetchingUser(username string) (*GitHubUser, error) {
	username = url.PathEscape(username)
	var user GitHubUser

	if err := a.requestGitHub("/users/"+username, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// FetchingRepoItems returns up to 100 public repositories for a GitHub user.
func (a *App) FetchingRepoItems(username string) ([]Repository, error) {
	username = url.PathEscape(username)
	endpoint := fmt.Sprintf("/users/%s/repos?per_page=100&sort=updated", username)

	var repositories []Repository
	if err := a.requestGitHub(endpoint, &repositories); err != nil {
		return nil, err
	}

	return repositories, nil
}
