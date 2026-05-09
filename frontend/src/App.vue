<script setup>
import { ref } from 'vue'
import { useGithubStore } from './stores/githubUser'
import { storeToRefs } from 'pinia'

const searchTerm = ref('')
const isVisible = ref(false)
const githubStore = useGithubStore()

// استخراج البيانات من الـ Store مع الحفاظ على التفاعل (Reactivity)
const { languageChartData, topReposChartData, loading , user , repositories } = storeToRefs(githubStore)

const handleSearch = async () => {
  if (!searchTerm.value.trim()) return
  isVisible.value = true
  await githubStore.loadUserStats(searchTerm.value)
  await githubStore.loadUserinfo(searchTerm.value)
}
</script>

<template>
  <div class="min-h-screen bg-gray-900 text-white px-8 pt-10 pb-8">
    
    <!-- العنوان ومنطقة البحث -->
    <div class="flex flex-col items-center gap-6 mb-12">
      <h1 class="text-5xl font-bold">GitHub Finder</h1>
      
      <div class="flex gap-3">
        <input 
          v-model="searchTerm"
          @keyup.enter="handleSearch"
          placeholder="Enter Username"
          class="bg-gray-800 p-3 rounded-lg border border-gray-700 w-64 outline-none focus:border-blue-500"
        />
        <button @click="handleSearch" class="bg-blue-600 px-6 py-3 rounded-lg font-bold hover:bg-blue-700 transition">
          Search
        </button>
      </div>
    </div>

    <!-- حالة التحميل -->
    <div v-if="loading" class="text-center text-blue-400 text-xl">trying to get the data</div>

    <div v-else-if ="isVisible">

      <!--User Card-->
    <div class="flex flex-col md:flex-row gap-6 mb-8">
      <img
       :src="user.avatar_url" alt="avatar" class="w-32 h-32 rounded-full border-2 border-gray-700 object-cover">
    </div>

    <div class="flex-1">
      <h2 class="text-2xl font-bold">{{user.name || user.login}}</h2>
      <p class="text-gray-400">@{{ user.login }}</p>
      <p class="mt-2">{{ user.bio || 'Bio is not found' }}</p>
    </div>

    <div class="flex flex-wrap gap-4 mt-4 text-sm"></div>

    <!-- عرض الرسوم البيانية -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-8 max-w-6xl mx-auto">
            
      <!-- مخطط اللغات -->
      <div class="bg-gray-800 p-6 rounded-2xl border border-gray-700">
        <h3 class="text-lg font-semibold mb-4 text-gray-400">Languages Used</h3>
        <apexchart 
          type="pie" 
          :options="{ labels: languageChartData.labels, theme: {mode: 'dark'} }" 
          :series="languageChartData.series" 
        />
      </div>

      <!-- مخطط المشاريع -->
      <div class="bg-gray-800 p-6 rounded-2xl border border-gray-700">
        <h3 class="text-lg font-semibold mb-4 text-gray-400">Top 5 Projects (Stars)</h3>
        <apexchart 
          type="bar" 
          :options="{ xaxis: { categories: topReposChartData.categories }, theme: {mode: 'dark'} }" 
          :series="topReposChartData.series" 
        />
      </div>
</div>
    </div>
  </div>
</template>