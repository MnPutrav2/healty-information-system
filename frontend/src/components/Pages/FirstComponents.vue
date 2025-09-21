<script setup lang="ts">
import { ref } from 'vue';
import CardMascot from '@/components/Pages/first/CardMascot.vue';

// Define variabels
const getUserData = defineProps(['data'])
const userData = ref<userData>(getUserData.data)

// External type
interface userData {
  id: number,
  name: string,
  gender: string
}

</script>

<template>
  <section class="anim-slide" style="width: 100%; height: 100%;">
    <CardMascot>Selamat datang, {{ userData.name }}</CardMascot>
  </section>
</template>
