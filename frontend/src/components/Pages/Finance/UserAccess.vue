<script setup lang="ts">
import type { UserReq } from '@/types/hr';
import InputData from '@/components/Extras/InputData.vue';
import { reactive, ref } from 'vue';
import { useRoute } from 'vue-router';
import { addUser } from '@/lib/api/hr';
import type { ResponseError, ResponseSuccess } from '@/types/response';

interface ID {
  id?: string
}

function base64decode(data: string | undefined): string[] {
  if(data == undefined){
    return []
  }

  const res: string = atob(data)
  const arr: string[] = res.split('|')

  return arr
}

const route = useRoute()
const id = ref<ID>(route.query)
const userReq: UserReq = reactive({
  id: "",
  employee_id: base64decode(id.value.id)[0],
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

</script>

<template>
  <section class="anim-slide" ref="pages">
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
  </section>
</template>
