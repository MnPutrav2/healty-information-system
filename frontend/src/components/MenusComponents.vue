<script setup lang="ts">
import { UserLogout, UserPages } from '@/lib/api/user'
import { onBeforeMount, ref, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()
const query = ref<Query>(route.query)

// Define variabels
const getUserData = defineProps(['data', 'menu'])
const userData = ref<userData>(getUserData.data)
const openNav = ref<boolean>(getUserData.menu)

// External type
interface userData {
  id: number,
  name: string,
  gender: string
}

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

onBeforeMount(async () => {
  await userPages()
})

watch(route, () => {
  query.value = route.query
})

watch(() => getUserData.menu, (newVal) => {
  openNav.value = newVal
})
</script>

<template>
  <section class="anim-slide responsive-navbar phone" :style="openNav ? 'width: 100%' : 'width: 0%'">
    <div style="padding: 0.5rem;">
      <div class="profil-card">
        <div style="height: 90%;">
          <div class="center" style="padding: 0.5rem; justify-content: flex-start">
            <div class="icon"></div>
            <div style="margin-left: 0.5rem;">
              <p style="font-weight: bold;">{{ userData.name }}</p>
              <button @click="logout" class="logout">Logout</button>
            </div>
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
  <section class="anim-slide responsive-navbar desktop">
    <div style="padding: 0.5rem;">
      <div class="profil-card">
        <div style=" height: 90%">
          <div class="center" style="padding: 0.5rem; justify-content: flex-start">
            <div class="icon"></div>
            <div style="margin-left: 0.5rem;">
              <p style="font-weight: bold;">{{ userData.name }}</p>
              <button @click="logout" class="logout">Logout</button>
            </div>
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

  <slot></slot>
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
