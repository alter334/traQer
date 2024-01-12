<template>
  <main>
    <div :class="$style.container">
      <header :class="$style.header">
        <router-link to="/ping">Ping|</router-link>
        <router-link to="/messages">Message|</router-link>
        <router-link to="/messages/1d88436d-a6d9-4290-8702-05509116f921">23B</router-link>
      </header>

      <router-view />
    </div>
  </main>
</template>

<style module>
.container {
  max-width: fit-content;
  margin: auto;
}
.header {
  display: flex;
  justify-content: center;
}
</style>
