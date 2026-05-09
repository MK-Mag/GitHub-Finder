package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type GitHubUser struct {
	Login       string `json:"login"`        // اسم المستخدم الفريد
	ID          int    `json:"id"`           // المعرف الرقمي
	AvatarURL   string `json:"avatar_url"`   // رابط الصورة الشخصية
	HTMLURL     string `json:"html_url"`     // رابط الملف الشخصي على المتصفح
	Name        string `json:"name"`         // الاسم الظاهر (قد يكون فارغاً)
	Company     string `json:"company"`      // الشركة
	Blog        string `json:"blog"`         // الموقع الشخصي
	Location    string `json:"location"`     // الموقع الجغرافي
	Bio         string `json:"bio"`          // السيرة الذاتية
	PublicRepos int    `json:"public_repos"` // عدد المستودعات العامة
	Followers   int    `json:"followers"`    // عدد المتابعين
	Following   int    `json:"following"`    // من يتابعهم
	CreatedAt   string `json:"created_at"`   // تاريخ إنشاء الحساب
}

// الهياكل البيانية (Structs)
type Repository struct {
	Name            string `json:"name"`
	StargazersCount int    `json:"stargazers_count"`
	Language        string `json:"language"`
	CreatedAt       string `json:"created_at"`
}

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// دالة جلب المستودعات
func (a *App) FetchingRepoItems(username string) ([]Repository, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API error: %d", res.StatusCode)
	}

	var repos []Repository
	err = json.NewDecoder(res.Body).Decode(&repos)
	if err != nil {
		return nil, err
	}

	return repos, nil
}

func (a *App) FetchingUser(Username string) (*GitHubUser, error) {
	Url := fmt.Sprintf("https://api.github.com/users/%s", Username)

	res, err := http.Get(Url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("there an error with the api :%d", res.StatusCode)
	}

	var User GitHubUser
	err = json.NewDecoder(res.Body).Decode(&User)
	if err != nil {
		return nil, err
	}
	return &User, err
}
