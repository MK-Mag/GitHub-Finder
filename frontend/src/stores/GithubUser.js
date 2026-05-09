import { defineStore } from 'pinia'
import { FetchingRepoItems , FetchingUser } from '../../wailsjs/go/main/App' // تأكد من صحة المسار

export const useGithubStore = defineStore('github', {
  state: () => ({
    user: null,
    repositories: [],
    loading: false,
    error: null,
  }),

  getters: {
    // تجهيز بيانات لغات البرمجة (Pie Chart)
    languageChartData(state) {
      const counts = {}
      state.repositories.forEach(repo => {
        if (repo.language) {
          counts[repo.language] = (counts[repo.language] || 0) + 1
        }
      })
      return {
        series: Object.values(counts),
        labels: Object.keys(counts)
      }
    },

    // تجهيز بيانات أشهر 5 مشاريع (Bar Chart)
    topReposChartData(state) {
      const sorted = [...state.repositories]
        .sort((a, b) => b.stargazers_count - a.stargazers_count) // انتبه للصيغة الصغيرة
        .slice(0, 5)

      return {
        series: [{
          name: 'Stars',
          data: sorted.map(r => r.stargazers_count)
        }],
        categories: sorted.map(r => r.name)
      }
    }
  },

  actions: {
    async loadUserStats(username) {
      this.loading = true
      this.error = null
      try {
        const data = await FetchingRepoItems(username)
        this.repositories = data || []
      } catch (err) {
        this.error = "Failed to get the data"
        console.error(err)
      } finally {
        this.loading = false
      }
    },
    async loadUserinfo(username){
      this.loading = true
      this.error = null
      try{
        const profile = await FetchingUser(username)
        this.user = profile
      } catch(err){
        this.error = "Failed to get the profile data"
        console.error(err)
      } finally {
        this.loading = false
      }
    }
  }
})