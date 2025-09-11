<script setup lang="ts">
import { UserStatus } from '@/lib/api/user'
import { onBeforeMount, reactive } from 'vue'
import { useRouter, type Router } from 'vue-router'
import { RouterView } from 'vue-router'
import MenuComponents from '@/components/MenusComponents.vue'
import NavbarComponents from '@/components/NavbarComponents.vue'

const router: Router = useRouter()
const token: string | null = localStorage.getItem('token')
interface userData {
  id: number,
  name: string,
  gender: string
}
const userData: userData = reactive({
  id: 0,
  name: '',
  gender: ''
})

async function userStatus() {
  if (!token) {
    router.push('/login')
    return
  }

  try {
    const response = await UserStatus()
    const json = await response.json()

    if (response.status === 200) {
      userData.id = json.employee_id
      userData.name = json.name
      userData.gender = json.gender

      localStorage.setItem("name", userData.name)
    } else {
      console.log(json)
    }
  } catch (error) {
    console.log(error)
  }
}

onBeforeMount(async () => {
  await userStatus()
})
</script>

<template>
  <main class="center">
    <MenuComponents :data="userData">
      <section class="anim-slide cover">
        <NavbarComponents />
        <div class="scroll">
          <RouterView v-slot="{ Component }">
            <component :is="Component" :data="userData" />
          </RouterView>
        </div>
      </section>
    </MenuComponents>
  </main>
</template>

<style scoped>
.cover {
  width: 85%;
  height: 100vh;
}

.scroll {
  width: 100%;
  height: 93%;
}
</style>
