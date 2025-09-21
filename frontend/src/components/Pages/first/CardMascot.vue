<script setup lang="ts">
import { ref } from 'vue';

interface Day {
  id: number
  name: string
}

const date = new Date()
const days = ref<Day[]>([
  {
    id: 1,
    name: "Senin"
  },
  {
    id: 2,
    name: "Selasa"
  },
  {
    id: 3,
    name: "Rabu"
  },
  {
    id: 4,
    name: "Kamis"
  },
  {
    id: 5,
    name: "Jumat"
  },
  {
    id: 6,
    name: "Sabtu"
  },
  {
    id: 7,
    name: "Minggu"
  }
])
</script>

<template>
  <div style="width: 100%; padding-top: 1.5rem;" class="center">
    <div class="card">
      <div class="responsive-flex">
        <div class="responsive-width">
          <h1><slot></slot></h1>
          <h3 style="margin-top: 1rem;">Medika Permata Bunda</h3>
          <div class="mini-card">
            <div class="center">
              <p class="day" v-for="day in days" :key="day.id" :style="date.getDay() == day.id ? 'color:white; border:1px solid white' : ''">{{ day.name }}</p>
            </div>
            <h2></h2>
          </div>
        </div>
        <div class="center responsive-width">
          <img class="mascout-img" src="../../../assets/images/mascout.png" alt="">
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
h1 {
  font-size: 1.5rem;
}

.day {
  margin: 0.5rem;
  padding: 0.3rem;
  font-size: 0.5rem;
  border-radius: 0.3rem;
  border: 1px solid rgba(255, 255, 255, 0.245);
  color: var(--font-color-sec);
}

h3 {
  color: var(--font-color-sec);
}

.mini-card {
  background-image: linear-gradient(50deg, rgba(0, 102, 122, 0.214), rgba(0, 113, 144, 0.364));
  border: 1px solid rgb(55, 188, 255);
  margin-top: 1.5rem;
  transition: all 0.2s;
  border-radius: 1rem;
}

.mini-card:hover {
  box-shadow: 0px 0px 10px 1px rgba(0, 208, 255, 0.346);
}

.mascout-img {
  width: 20rem;
}
</style>
