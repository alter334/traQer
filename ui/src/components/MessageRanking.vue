<script setup lang="ts">
import { ref, onMounted } from 'vue'

interface UserDetailWithMessageCount{
	id: string 
	displayname: string 
	name: string
	groups: string[] 
	homechannnel: string 
	totalmessagecount: number         
}

const rankingDatas = ref<UserDetailWithMessageCount[]>()

onMounted(async () => {
  const res = await fetch('api/messages')
  if (res.ok) {
    rankingDatas.value = await res.json()
  }
})
</script>
<template>
  <div>
    <div><h2>総合投稿数ランキング</h2></div>
    <ol>
      <li v-for="rankingData in rankingDatas" :key="rankingData.displayname">
        <div>ユーザー:{{ rankingData.displayname }} 投稿数:{{ rankingData.totalmessagecount }}</div>
      </li>
    </ol>
  </div>
</template>
