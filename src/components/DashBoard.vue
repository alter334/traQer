<script setup lang="ts">
import { onMounted, ref } from 'vue'
const users = ref<string>('no data')

onMounted(async () => {
  const res = await fetch('https://q.trap.jp/api/v3/users')
  if (res.ok) {
    users.value = await res.json()
  } else {
    users.value = await res.statusText
  }
})
</script>

<template>
  <div>{{ users }}</div>
</template>
