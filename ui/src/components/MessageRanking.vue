<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'

//ユーザー詳細読み込みインターフェース
interface UserDetailWithMessageCount {
  id: string
  displayname: string
  name: string
  groups: string[]
  homechannnel: string
  totalmessagecount: number
}

const route = useRoute()
const rankingDatas = ref<UserDetailWithMessageCount[]>()

//パスパラメタが存在しないときは全ユーザ,存在するときは特定グループのみ
if (route.params.groupid == undefined) {
  onMounted(async () => {
    const res = await fetch('/api/messages')
    if (res.ok) {
      rankingDatas.value = await res.json()
    }
  })
} else {
  onMounted(async () => {
    const res = await fetch('/api/messages/' + route.params.groupid)
    if (res.ok) {
      rankingDatas.value = await res.json()
    }
  })
}
</script>
<template>
  <div>
    <div>
      <h2>{{ $route.params.groupid }}投稿数ランキング</h2>
    </div>
    <ol>
      <li v-for="rankingData in rankingDatas" :key="rankingData.displayname">
        <div>ユーザー:{{ rankingData.displayname }} 投稿数:{{ rankingData.totalmessagecount }}</div>
      </li>
    </ol>
  </div>
</template>
