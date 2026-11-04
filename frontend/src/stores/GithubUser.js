import { defineStore } from 'pinia'
import { FetchingRepoItems, FetchingUser } from '../../wailsjs/go/main/App'

export const useGithubStore = defineStore('github', {
  state: () => ({
    user: null,
    repositories: [],
    loading: false,
    error: null,
  }),

  getters: {
    languageChartData(state) {
      const counts = {}

      for (const repository of state.repositories) {
        if (repository.language) {
          counts[repository.language] = (counts[repository.language] || 0) + 1
        }
      }

      return {
        labels: Object.keys(counts),
        series: Object.values(counts),
      }
    },

    topReposChartData(state) {
      const topRepositories = [...state.repositories]
        .sort((a, b) => b.stargazers_count - a.stargazers_count)
        .slice(0, 5)

      return {
        categories: topRepositories.map(repository => repository.name),
        series: [{
          name: 'Stars',
          data: topRepositories.map(repository => repository.stargazers_count),
        }],
      }
    },
  },

  actions: {
    async loadUser(username) {
      this.loading = true
      this.error = null
      this.user = null
      this.repositories = []

      try {
        const [profile, repositories] = await Promise.all([
          FetchingUser(username),
          FetchingRepoItems(username),
        ])

        this.user = profile
        this.repositories = repositories || []
      } catch (error) {
        this.error = error?.message || 'Unable to retrieve data from GitHub.'
        console.error('GitHub Finder:', error)
      } finally {
        this.loading = false
      }
    },
  },
})
