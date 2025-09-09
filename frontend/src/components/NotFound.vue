<script setup lang="ts">
import { useRouter } from 'vue-router';

const router = useRouter()
</script>

<template>
  <section class="anim-slide center" style="height: 100vh;">
    <div>
      <div class="center">
        <h1 style="font-size: 5rem;">404</h1>
        <div style="margin-left: 1rem; margin-right: 1rem; width: 1px; background-color: var(--line-color); height: 5rem;"></div>
        <p>Page Not Found</p>
      </div>
      <div class="center" style="margin-top: 1rem;">
        <button @click="router.push('/')">Back to menu</button>
      </div>
    </div>
  </section>
</template>

<style scoped>
button {
  background-color: transparent;
  color: var(--font-color);
  border: none;
}
</style>
