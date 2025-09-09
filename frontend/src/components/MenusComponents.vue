<script setup lang="ts">
import { UserLogout, UserPages } from '@/lib/api/user'
import { onBeforeMount, ref, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import MedicalRecordMenu from './MedicalRecordMenu.vue'
import RecipeInput from './Form/RecipeInput.vue'
import RecipeCompoundInput from './Form/RecipeCompoundInput.vue'
import LabolatoriumForm from './Form/LabolatoriumForm.vue'
import ExaminationForm from './Form/ExaminationForm.vue'

const router = useRouter()
const route = useRoute()
const query = ref<Query>(route.query)
const name: string | null = localStorage.getItem("name")

interface pathData {
  group: string
  path: pathGroup[]
}

interface Query {
  careNum?: string
}

interface pathGroup {
  name: string
  path: string
}

const pathData = ref<pathData[]>([])
const token: string | null = localStorage.getItem('token')

async function userPages() {
  if (!token) {
    return
  }

  try {
    const response = await UserPages()
    const json = await response.json()

    if (response.status === 200) {
      pathData.value = json
    } else {
      router.push('/login')
      console.log(json)
    }
  } catch (error) {
    router.push('/login')
    console.log(error)
  }
}

function routing(path: string) {
  router.push(path)
}

async function logout() {
  const response = await UserLogout()
  const json = await response.json()

  try {
    if (response.status === 200) {
      localStorage.setItem('token', '')
      localStorage.setItem('user', '')

      router.push({
        path: '/login'
      })

      console.log(json.status)
    }else{
      localStorage.setItem('token', '')

      router.push({
        path: '/login'
      })

      console.log(json.errors)
    }
  } catch(error) {
    localStorage.setItem('token', '')
    localStorage.setItem('user', '')

    router.push({
      path: '/login'
    })

    console.log(error)
  }
}

const activeMenu = ref<string | null>(null)

function toggleMenu(menuName: string) {
  activeMenu.value = activeMenu.value === menuName ? null : menuName
}

onBeforeMount(async () => {
  await userPages()
})

watch(route, () => {
  query.value = route.query
})
</script>

<template>
  <section class="anim-slide" style="width: 15%; height: 100vh; border-right: 1px solid var(--line-color-transparent);">
    <div class="profil-card">
      <div style=" height: 90%; border-bottom: 1px solid var(--line-color-transparent)">
        <div class="center" style="padding: 0.5rem; justify-content: flex-start">
          <div class="icon"></div>
          <div style="margin-left: 0.5rem;">
            <p style="font-weight: bold;">{{ name }}</p>
            <button @click="logout" class="logout">Logout</button>
          </div>
        </div>
      </div>
    </div>
    <div style="width: 100%; height: 80%;" class="scroll">
      <ul>
        <li>
          <button class="menu-button" :class="route.path == '/' ? 'menu-button-active' : ''" @click="routing('/')">Home</button>
        </li>
        <li v-for="(path, index) in pathData" :key="index" class="path">
          <details v-if="path.group != '-'">
            <summary>{{ path.group }}</summary>
            <button v-for="r in path.path" :key="r.name" class="menu-button" :class="route.path == r.path ? 'menu-button-active' : ''" @click="routing(r.path)">{{ r.name }}</button>
          </details>
          <button v-else v-for="r in path.path" :key="r.name" class="menu-button" :class="route.path == r.path ? 'menu-button-active' : ''" @click="routing(r.path)">{{ r.name }}</button>
        </li>
      </ul>
    </div>
  </section>

  <!-- Menu area -->
   <RecipeInput style="position: fixed; transition: all 0.2s linear;" :data="query.careNum" v-if="activeMenu === 'resep'">
      <template #btn-close>
        <button class="act" @click="toggleMenu('resep')">X</button>
      </template>
    </RecipeInput>

    <RecipeCompoundInput style="position: fixed; transition: all 0.2s linear;" :data="query.careNum" v-if="activeMenu === 'resep-racik'">
      <template #btn-close>
        <button class="act" @click="toggleMenu('resep-racik')">X</button>
      </template>
    </RecipeCompoundInput>

    <LabolatoriumForm style="position: fixed; transition: all 0.2s linear;" :data="query.careNum" v-if="activeMenu === 'lab-input'">
      <template #btn-close>
        <button class="act" @click="toggleMenu('lab-input')">X</button>
      </template>
    </LabolatoriumForm>

    <ExaminationForm style="position: fixed; transition: all 0.2s linear;" :data="query.careNum" v-if="activeMenu === 'exam-input'">
      <template #btn-close>
        <button class="act" @click="toggleMenu('exam-input')">X</button>
      </template>
    </ExaminationForm>
  <!-- Menu area -->

  <slot></slot>
  <MedicalRecordMenu v-if="route.path == '/ambulatory-care'">
    <div style="margin-bottom: 1rem;">
      <p style="margin-bottom: 0.5rem;">{{ query.careNum == '' ? 'Pilih pasien' : query.careNum }}</p>
      <hr>
    </div>
    <template v-if="query.careNum != null">
      <div>
        <button style="background-color: transparent; color: var(--font-color); border: none;" @click="toggleMenu('resep')">Input Resep</button>
      </div>
      <div style="margin-top: 1rem;">
        <button style="background-color: transparent; color: var(--font-color); border: none;" @click="toggleMenu('resep-racik')">Input Resep Racikan</button>
      </div>
      <div style="margin-top: 1rem;">
        <button style="background-color: transparent; color: var(--font-color); border: none;" @click="toggleMenu('lab-input')">Input Permintaan LAB</button>
      </div>
      <div style="margin-top: 1rem;">
        <button style="background-color: transparent; color: var(--font-color); border: none;" @click="toggleMenu('exam-input')">Input Tindakan</button>
      </div>
    </template>
  </MedicalRecordMenu>
</template>

<style scoped>
.menu-button {
  padding: 0.5rem;
  text-align: start;
  width: 100%;
  background-color: transparent;
  cursor: pointer;
  color: var(--font-color-sec);
  border: none;
}

.act {
  width: 100%;
  height: 100%;
  background-color: transparent;
  border: none;
  color: var(--font-color);
}

.menu-button-active {
  border-left: 3px solid rgb(0, 174, 255);
  padding: 0.5rem;
  color: var(--font-color);
  background-image: linear-gradient(to right, rgba(0, 174, 255, 0.233), transparent);
  transition: all 0.2s;
}

details {
  margin: 0.5rem 0rem;
  cursor: pointer;
  transition: all 0.3s ease;
  margin-left: 0.5rem;
  color: var(--font-color-sec);
}

summary::marker,
summary::-webkit-details-marker {
  display: none;
}

/* Tambahkan panah custom */
summary {
  position: relative;
  margin-bottom: 0.5rem;
  cursor: pointer;
  list-style: none;
}

summary::before {
  content: "›"; /* panah ke kanan */
  position: absolute;
  right: 0;
  transition: transform 0.2s ease;
}

/* Saat <details> dibuka, putar panah */
details[open] summary::before {
  transform: rotate(90deg); /* jadikan panah ke bawah */
}

.icon {
  width: 2rem;
  height: 2rem;
  background-image: url('../assets/icons/profile.png');
  background-size: cover;
  border-radius: 2rem;
}
</style>
