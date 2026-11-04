# GitHub Finder

A desktop application for exploring public GitHub profiles and repositories. Built with **Go**, **Wails**, **Vue 3**, and **ApexCharts**.

## Features

- Look up a public GitHub user by username.
- View profile details, bio, avatar, and account statistics.
- Browse public repositories and their primary languages.
- Visualize language distribution across repositories.
- Compare the five most-starred repositories.
- Open the original GitHub profile and repositories.

## Tech Stack

| Layer | Technology |
| --- | --- |
| Desktop runtime | Wails v2 |
| Backend | Go |
| Frontend | Vue 3 |
| State management | Pinia |
| Charts | ApexCharts |
| Frontend tooling | Vite |

## Requirements

- Go (compatible with the version required by Wails v2)
- Node.js and npm
- Wails CLI v2
- Platform dependencies required by [Wails](https://wails.io/docs/gettingstarted/installation)

## Getting Started

Clone the repository and install frontend dependencies:

```bash
git clone https://github.com/MK-Mag/GitHub-Finder.git
cd GitHub-Finder
cd frontend
npm install
cd ..
```

Start the desktop app in development mode:

```bash
wails dev
```

Build a production application:

```bash
wails build
```

## Project Structure

```text
.
├── app.go                 # GitHub API client and data models
├── main.go                # Wails application setup
├── frontend/
│   ├── src/
│   │   ├── App.vue        # Main application interface
│   │   ├── stores/
│   │   │   └── GithubUser.js # Profile and repository state
│   │   └── style.css      # Global application styles
│   └── package.json
├── build/                 # Wails build assets
├── go.mod
└── wails.json
```

## How It Works

The Vue interface calls Go methods exposed by Wails. The Go backend requests public profile and repository data from the GitHub REST API, then returns typed data to the frontend. Pinia keeps the profile, repositories, loading state, and errors in one place; ApexCharts presents repository statistics.

## API and Limits

GitHub Finder reads public GitHub data and does not require a personal access token. Unauthenticated GitHub API requests are subject to GitHub's rate limits. The application does not access private repositories.

## License

No license has been specified yet.
