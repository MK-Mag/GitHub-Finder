<script setup>
import { ref, computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useGithubStore } from './stores/GithubUser'

const searchTerm = ref('')
const githubStore = useGithubStore()
const { user, repositories, loading, error, languageChartData, topReposChartData } =
  storeToRefs(githubStore)

const hasRepositories = computed(() => repositories.value.length > 0)

async function handleSearch() {
  const username = searchTerm.value.trim()
  if (!username || loading.value) return
  await githubStore.loadUser(username)
}
</script>

<template>
  <main class="app-shell">
    <header class="topbar">
      <a class="brand" href="#" aria-label="GitHub Finder home">
        <span class="brand-mark">GH</span>
        <span>GitHub Finder</span>
      </a>
      <span class="topbar-caption">A closer look at public GitHub profiles</span>
    </header>

    <section class="hero">
      <p class="eyebrow">DEVELOPER PROFILE EXPLORER</p>
      <h1>Explore the work<br /><span>behind a GitHub profile.</span></h1>
      <p class="hero-copy">
        Search for a GitHub username to discover repositories, languages, and activity at a glance.
      </p>

      <form class="search-form" @submit.prevent="handleSearch">
        <label class="sr-only" for="username">GitHub username</label>
        <span class="search-icon" aria-hidden="true">⌕</span>
        <input
          id="username"
          v-model="searchTerm"
          type="text"
          placeholder="Enter a GitHub username"
          autocomplete="off"
          spellcheck="false"
          :disabled="loading"
        />
        <button type="submit" :disabled="loading || !searchTerm.trim()">
          {{ loading ? 'Searching…' : 'Search profile' }}
        </button>
      </form>
      <p class="search-hint">Public profiles only · No access token required</p>
    </section>

    <section v-if="loading" class="state-panel" aria-live="polite">
      <span class="loader"></span>
      <p>Fetching profile and repositories…</p>
    </section>

    <section v-else-if="error" class="state-panel error-panel" role="alert">
      <span class="state-icon">!</span>
      <h2>We couldn't load this profile</h2>
      <p>{{ error }}</p>
      <span class="muted">Check the username or try again in a moment.</span>
    </section>

    <template v-else-if="user">
      <section class="profile-card">
        <img class="avatar" :src="user.avatar_url" :alt="user.login + ' avatar'" />
        <div class="profile-main">
          <p class="eyebrow">GITHUB PROFILE</p>
          <h2>{{ user.name || user.login }}</h2>
          <a class="username" :href="user.html_url" target="_blank" rel="noreferrer">
            @{{ user.login }} ↗
          </a>
          <p v-if="user.bio" class="bio">{{ user.bio }}</p>
          <p v-else class="bio muted">No bio provided.</p>
          <div class="profile-meta">
            <span v-if="user.location">⌖ {{ user.location }}</span>
            <a v-if="user.blog" :href="user.blog.startsWith('http') ? user.blog : 'https://' + user.blog" target="_blank" rel="noreferrer">
              ↗ {{ user.blog }}
            </a>
            <span v-if="user.company">▤ {{ user.company }}</span>
          </div>
        </div>
        <div class="profile-stats">
          <div><strong>{{ user.public_repos }}</strong><span>Repositories</span></div>
          <div><strong>{{ user.followers }}</strong><span>Followers</span></div>
          <div><strong>{{ user.following }}</strong><span>Following</span></div>
        </div>
      </section>

      <section class="section-heading">
        <div>
          <p class="eyebrow">PROFILE OVERVIEW</p>
          <h2>Repository insights</h2>
        </div>
        <span class="result-count">{{ repositories.length }} repositories loaded</span>
      </section>

      <section v-if="hasRepositories" class="charts-grid">
        <article class="panel chart-panel">
          <div class="panel-heading">
            <div><h3>Languages</h3><p>Primary language per repository</p></div>
          </div>
          <apexchart
            v-if="languageChartData.series.length"
            type="donut"
            height="280"
            :options="{
              labels: languageChartData.labels,
              theme: { mode: 'dark' },
              chart: { background: 'transparent', toolbar: { show: false } },
              legend: { position: 'bottom', labels: { colors: '#a5afc1' } },
              dataLabels: { enabled: false },
              stroke: { colors: ['#151c2b'] },
              plotOptions: { pie: { donut: { size: '66%' } } }
            }"
            :series="languageChartData.series"
          />
          <p v-else class="empty-chart">No language data is available for these repositories.</p>
        </article>

        <article class="panel chart-panel">
          <div class="panel-heading">
            <div><h3>Most-starred repositories</h3><p>Top five by star count</p></div>
          </div>
          <apexchart
            type="bar"
            height="280"
            :options="{
              theme: { mode: 'dark' },
              chart: { background: 'transparent', toolbar: { show: false } },
              plotOptions: { bar: { borderRadius: 5, horizontal: true, barHeight: '55%' } },
              dataLabels: { enabled: false },
              xaxis: { categories: topReposChartData.categories, labels: { style: { colors: '#8995aa' } } },
              yaxis: { labels: { style: { colors: '#cbd5e1' } } },
              grid: { borderColor: '#283247' },
              colors: ['#58a6ff']
            }"
            :series="topReposChartData.series"
          />
        </article>
      </section>

      <section v-else class="state-panel">
        <h3>No public repositories yet</h3>
        <p>This profile doesn't have any public repositories to analyze.</p>
      </section>

      <section class="repositories-section">
        <div class="section-heading">
          <div><p class="eyebrow">PUBLIC WORK</p><h2>Repositories</h2></div>
          <span class="result-count">{{ repositories.length }} total</span>
        </div>
        <div v-if="hasRepositories" class="repo-list">
          <a
            v-for="repo in repositories"
            :key="repo.full_name"
            class="repo-item"
            :href="repo.html_url"
            target="_blank"
            rel="noreferrer"
          >
            <div class="repo-copy">
              <h3>{{ repo.name }} <span aria-hidden="true">↗</span></h3>
              <p>{{ repo.description || 'No description provided.' }}</p>
              <div class="repo-meta">
                <span v-if="repo.language" class="language-dot">{{ repo.language }}</span>
                <span>★ {{ repo.stargazers_count }}</span>
                <span>⑂ {{ repo.forks_count }}</span>
              </div>
            </div>
          </a>
        </div>
        <p v-else class="muted">Nothing to show here yet.</p>
      </section>
    </template>

    <section v-else class="welcome-note">
      <div class="welcome-symbol">⌘</div>
      <h2>Your next discovery starts here.</h2>
      <p>Enter a username above to see profile details and repository insights.</p>
    </section>

    <footer class="footer">
      <span>GitHub Finder</span>
      <span>Built with Go · Wails · Vue</span>
      <a href="https://github.com/MK-Mag/GitHub-Finder" target="_blank" rel="noreferrer">Source code ↗</a>
    </footer>
  </main>
</template>
