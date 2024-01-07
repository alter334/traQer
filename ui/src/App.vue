<template>
  <main>
    <div :class="$style.container">
      <header :class="$style.header">
        <router-link to="/ping">Ping|</router-link>
        <router-link to="/message">Message|</router-link>
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
