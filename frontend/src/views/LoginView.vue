<script setup lang="ts">
import { GetAuthToken, UserStatus } from '@/lib/api/user'
import router from '@/router'
import type { ResponseError, ResponseToken } from '@/types/response'
import { onBeforeMount, reactive } from 'vue'

const userAccount: userAccount = reactive({
  username: "",
  password: ""
})

interface userAccount {
  username: string,
  password: string
}

async function userStatus() {
  const token: string | null = localStorage.getItem("token")

  if (token != null) {
    try {
      const response: Response = await UserStatus()

      if (response.status === 200) {
        alert("Session login masih ada")

        router.push("/")
      }else{
        localStorage.removeItem("token")
      }
    } catch (error) {
      console.log(error)
    }
  }
}

async function login() {
  try {
    const response: Response = await GetAuthToken(userAccount)

    if (response.status === 201) {
      const json: ResponseToken = await response.json()
      localStorage.setItem('token', json.token)

      router.push('/')
      console.log(json.status)
    }else{
      const json: ResponseError = await response.json()
      alert(json.errors)
    }
  } catch(error) {
    console.log(error)
  }
}

onBeforeMount(async () => {
  await userStatus()
})
</script>

<template>
  <main class="center">
    <section class="form-cover center" style="width: 30%;">
      <div class="anim-slide">
        <h3 style="margin: 0.5rem;">LOGIN</h3>
        <form v-on:submit.prevent="login" class="form-custom">
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="username">Username</label>
            </div>
            <input type="text" id="username" v-model="userAccount.username" required autocomplete="username" placeholder="Username">
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="password">Password</label>
            </div>
            <input type="password" id="password" v-model="userAccount.password" required autocomplete="current-password" placeholder="Password">
          </div>
          <div class="center">
            <button type="submit">login</button>
          </div>
        </form>
      </div>
    </section>
    <div role="separator" aria-orientation="vertical" style="width: 1px; height: 95vh; background-color: var(--line-color);"></div>
    <section class="form-cover center" style="width: 70%">
      <div style="padding: 1rem;" class="anim-slide center">
        <div>
          <h1 style="margin-bottom: 1rem; font-size: 2rem;">Healty Information System</h1>
          <h1 style="margin-bottom: 1rem; opacity: 0.6; font-size: 2rem;">Healty Information System</h1>
          <h1 style="margin-bottom: 1rem; opacity: 0.2; font-size: 2rem;">Healty Information System</h1>
        </div>
      </div>
    </section>
  </main>
</template>
