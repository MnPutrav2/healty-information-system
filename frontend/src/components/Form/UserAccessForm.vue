<script setup lang="ts">
import type { UserReq } from '@/types/hr';
import InputData from '@/components/Extras/InputData.vue';
import { reactive, watch } from 'vue';
import { useRoute } from 'vue-router';
import { addUser } from '@/lib/api/hr';
import type { ResponseError, ResponseSuccess } from '@/types/response';

function base64decode(data: string | undefined): string[] {
  if(data == undefined){
    return []
  }

  const res: string = atob(data)
  const arr: string[] = res.split('|')

  return arr
}

const route = useRoute()
const userReq: UserReq = reactive({
  id: "",
  employee_id: "",
  username: "",
  password: ""
})

async function handleAddUser() {
  try{
    const response: Response = await addUser(userReq)

    if(response.status === 201){
      const json: ResponseSuccess = await response.json()

      alert(json.response)
    }else{
      const json: ResponseError = await response.json()

      alert(json.errors)
    }
  }catch(e){
    alert(e)
  }
}

watch(route, () => {
  const id = route.query
  userReq.employee_id = base64decode(String(id.id))[0]
})
</script>

<template>
  <section class="container center" style="position: fixed; z-index: 99; inset: 0;">
    <div class="cover2 scroll">
      <div class="close-btn center">
        <slot name="btn-close"></slot>
      </div>

      <h3 style="margin: 0.5rem;">Data Pegawai</h3>
      <div style="padding-top: 0.5rem; padding-bottom: 1rem;" class="bottom-line">
        <form class="form-data-custom" v-on:submit.prevent="handleAddUser">
          <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Buat akses user</h4>
          <div class="responsive-grid" style="padding-left: 1rem;">
            <InputData :props="{id: 'usid', name: 'User ID'}">
              <input type="text" id="usid" v-model="userReq.id" maxlength="6" placeholder="User ID">
            </InputData>
            <InputData :props="{id: 'emid', name: 'ID pegawai'}">
              <input type="text" id="emid" v-model="userReq.employee_id" maxlength="6" placeholder="ID pegawai" readonly>
            </InputData>
            <InputData :props="{id: 'us', name: 'Username'}">
              <input type="text" id="us" v-model="userReq.username" maxlength="12" placeholder="Username">
            </InputData>
            <InputData :props="{id: 'pass', name: 'Password'}">
              <input type="text" id="pass" v-model="userReq.password" maxlength="8"  placeholder="Password">
            </InputData>
          </div>
          <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Action</h4>
          <div style="margin: 1rem;">
            <button type="submit">Add</button>
          </div>
        </form>
      </div>
    </div>
  </section>
</template>

<style scoped>
.container {
  position: fixed;
  inset: 0;
  width: 100%;
  height: 100%;
  backdrop-filter: blur(2px);
  z-index: 3;
}

.close-btn {
  width: 2rem;
  height: 2rem;
  background-color: var(--background-color);
  border: 1px solid var(--line-color);
  position: absolute;
  right: 1rem;
  margin-top: -1.5rem;
  border-radius: 2rem;
}

.cover2 {
  width: 94%;
  height: 89%;
  padding: 0.5rem;
  border-radius: 1rem;
  background-color: var(--background-color);
  border: 1px solid var(--line-color);
  z-index: 3;
}

input {
  background-color: var(--background-color);
  border: 1px solid var(--line-color);
  padding: 0.5rem;
  border-radius: 0.5rem;
  color: var(--font-color);
}
</style>
