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

const rankingData = ref<UserDetailWithMessageCount[]>()

onMounted(async () => {
  const res = await fetch('api/messages')
  if (res.ok) {
    rankingData.value = await res.json()
  }
})
</script>
<template>
  {{ rankingData }}
</template>
