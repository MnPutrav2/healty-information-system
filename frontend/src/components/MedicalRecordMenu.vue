<script setup lang="ts">
import { ref, watch } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute()
const open = ref<boolean>(false)

watch(
  () => route.query.careNum,
  () => {
    open.value = true
  }
)
</script>

<template>
  <section style="position: fixed; right: 2.4rem; transition: all 0.2s linear;" class="center" :style="!open ? 'transform: translateX(14.5rem)' : 'transform: translateX(0rem)'">
    <div style="width: 2rem; height: 10rem; border-radius: 1rem 0 0 1rem; background-color: var(--background-color); border: 1px solid var(--line-color-transparent);">
      <button style="width: 100%; height: 100%; background-color: transparent; border: none;" @click="open = !open" class="center">
        <p style="rotate: 90deg; color: var(--font-color);">Menu</p>
      </button>
    </div>
    <div style="height: 80vh; min-width: 10rem; max-width: 10rem; background-color: var(--background-color); border: 1px solid var(--line-color-transparent); border-radius: 1rem; padding: 1rem;">
      <!-- isi -->
       <slot></slot>
    </div>
  </section>
</template>
