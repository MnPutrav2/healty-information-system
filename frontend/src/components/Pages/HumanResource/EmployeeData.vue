<script setup lang="ts">
import { addEmployee, deleteEmployee, getEmployee, updateEmployee } from '@/lib/api/hr';
import { formatDate } from '@/lib/formatDate';
import type { EmployeeData } from '@/types/hr';
import type { ResponseError, ResponseSuccess, SearchLimit } from '@/types/response';
import { nextTick, onBeforeMount, reactive, ref } from 'vue';
import InputData from '@/components/Extras/InputData.vue';
import router from '@/router';
import UserAccessForm from '@/components/Form/UserAccessForm.vue';

const user = defineProps(['data'])
const menuOpen = ref<boolean[]>([])
const open = ref<boolean>(false)
const dropOpenTemplate = ref<boolean[]>([])
const update = ref<boolean>(false)
const updateID = ref<string>("")
const pages = ref<HTMLElement | null>(null)
const employeeData = ref<EmployeeData[]>([])
const employeeReq = reactive<EmployeeData>({
  id: "",
  name: "",
  gender: "Laki - laki",
  birth_place: "",
  birth_date: "",
  wedding: "Menikah",
  address: "",
  nik: "",
  bpjs: "",
  npwp: "",
  phone_number: "",
  email: "",
  status: false
})
const search = reactive<SearchLimit>({
  search: "",
  limit: 5
})

async function handleGetEmployee() {
  try {
    const response: Response = await getEmployee(search)

    if(response.status === 200){
      employeeData.value = await response.json()
    }else{
      const json: ResponseError = await response.json()

      alert(json.errors)
    }
  }catch(e){
    alert(e)
  }
}

async function handleAddEmployee() {
  try {
    const response: Response = await addEmployee(employeeReq)

    if(response.status === 201){
      const json: ResponseSuccess = await response.json()
      await handleGetEmployee()
      resetForm()

      alert(json.response)
    }else{
      const json: ResponseError = await response.json()

      alert(json.errors)
    }
  }catch(e){
    alert(e)
  }

  open.value = false
}

async function handleDeleteEmployee(id: string) {
  try {
    const response: Response = await deleteEmployee(id)

    if(response.status === 200){
      const json: ResponseSuccess = await response.json()
      await handleGetEmployee()

      alert(json.response)
    }else{
      const json: ResponseError = await response.json()

      alert(json.errors)
    }
  }catch(e){
    alert(e)
  }
}

async function handleUpdateEmployee() {
  try {
    const response: Response = await updateEmployee(employeeReq, updateID.value)

    if(response.status === 200){
      const json: ResponseSuccess = await response.json()
      await handleGetEmployee()
      resetForm()

      alert(json.response)
    }else{
      const json: ResponseError = await response.json()

      alert(json.errors)
    }
  }catch(e){
    alert(e)
  }

  open.value = false
  update.value = false
}

function editEmployee(data: EmployeeData, index: number) {
  updateID.value = data.id
  update.value = true
  open.value = true
  dropOpenTemplate.value[index] = false

  Object.assign(employeeReq, {...data, birth_date: formatDate(new Date(data.birth_date))})

  nextTick(() => {
    pages.value?.scrollIntoView({behavior: 'smooth'})
  })
}

function base64encode(id: string, name: string): string{
  const enc: string = btoa(`${id}|${name}`)

  return enc
}

function resetForm() {
  employeeReq.id = ""
  employeeReq.name = ""
  employeeReq.gender = "Laki - laki"
  employeeReq.birth_place = ""
  employeeReq.birth_date = ""
  employeeReq.wedding = "Menikah"
  employeeReq.address = ""
  employeeReq.nik = ""
  employeeReq.bpjs = ""
  employeeReq.npwp = ""
  employeeReq.phone_number = ""
  employeeReq.email = ""
}

onBeforeMount(async () => {
  await handleGetEmployee()
})

</script>

<template>
  <section class="anim-slide" ref="pages">
    <h3 style="margin: 0.5rem;">Data Pegawai {{ user.data.name }}</h3>
    <div style="padding-top: 0.5rem; padding-bottom: 1rem;" class="bottom-line">
      <button class="button-add" style="margin: 0.5rem;" @click="open = !open; update = false" v-if="!open">Add pegawai</button>
      <form class="form-data-custom" v-else v-on:submit.prevent="">
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Tambah pegawai</h4>
        <div class="responsive-grid" style="padding-left: 1rem;">
          <InputData :props="{id: 'id', name: 'ID pegawai'}">
            <input type="text" v-model="employeeReq.id" id="id" placeholder="ID pemeriksaan" required>
          </InputData>
          <InputData :props="{id: 'nm', name: 'Nama'}">
            <input type="text" v-model="employeeReq.id" id="nm" placeholder="Nama" required>
          </InputData>
          <InputData :props="{id: 'jk', name: 'Jenis kelamin'}">
            <select id="jk" v-model="employeeReq.gender" required>
              <option value="Laki - laki">Laki - laki</option>
              <option value="Perempuan">Perempuan</option>
            </select>
          </InputData>
          <InputData :props="{id: 'pn', name: 'Status pernikahan'}">
            <select id="pn" v-model="employeeReq.wedding" required>
              <option value="Menikah">Menikah</option>
              <option value="Belum menikah">Belum menikah</option>
              <option value="Cerai">Cerai</option>
              <option value="Cerai mati">Cerai mati</option>
            </select>
          </InputData>
          <InputData :props="{id: 'bp', name: 'Tempat lahir'}">
            <input type="text" id="bp" v-model="employeeReq.birth_place" placeholder="Tempat lahir" required>
          </InputData>
          <InputData :props="{id: 'tl', name: 'Tempat lahir'}">
            <input type="date" id="tl" v-model="employeeReq.birth_date" placeholder="Tempat lahir" required>
          </InputData>
          <InputData :props="{id: 'al', name: 'Alamat'}">
            <textarea id="al" v-model="employeeReq.address" placeholder="Alamat" required></textarea>
          </InputData>
          <InputData :props="{id: 'nik', name: 'NIK'}">
            <input type="text" v-model="employeeReq.nik" id="nik" placeholder="NIK" required>
          </InputData>
          <InputData :props="{id: 'bpj', name: 'BPJS'}">
            <input type="text" v-model="employeeReq.bpjs" id="bpj" placeholder="BPJS" required>
          </InputData>
          <InputData :props="{id: 'npw', name: 'NPWP'}">
            <input type="text" id="npw" v-model="employeeReq.npwp" placeholder="NPWP" required>
          </InputData>
          <InputData :props="{id: 'npw', name: 'Nomor telpon'}">
            <input type="text" id="tlp" v-model="employeeReq.phone_number" placeholder="Nomor telpon" required>
          </InputData>
          <InputData :props="{id: 'em', name: 'Email'}">
            <input type="email" v-model="employeeReq.email" id="em" placeholder="Email" required>
          </InputData>
        </div>
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Action</h4>
        <div>
          <button type="button" @click="resetForm">Reset</button>
          <button type="button" v-if="update" @click="handleUpdateEmployee">Update</button>
          <button type="submit" v-else @click="handleAddEmployee">Add</button>
          <button type="button" @click="open = !open">Close</button>
        </div>
      </form>
    </div>

    <div style="padding: 0.5rem; padding-top: 2rem; padding-bottom: 2rem;">
      <form class="form-search-input" v-on:submit.prevent="handleGetEmployee">
        <div class="box-search">
          <div class="box">
            <input class="search-input" type="number" id="lm" v-model="search.limit" placeholder="limit">
            <input class="search-input" type="text" id="sc" v-model="search.search" placeholder="Nama pegawai">
          </div>
          <button type="submit">
            <p>Search</p>
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="384px" height="384px">
              <path d="M 9 2 C 5.1458514 2 2 5.1458514 2 9 C 2 12.854149 5.1458514 16 9 16 C 10.747998 16 12.345009 15.348024 13.574219 14.28125 L 14 14.707031 L 14 16 L 20 22 L 22 20 L 16 14 L 14.707031 14 L 14.28125 13.574219 C 15.348024 12.345009 16 10.747998 16 9 C 16 5.1458514 12.854149 2 9 2 z M 9 4 C 11.773268 4 14 6.2267316 14 9 C 14 11.773268 11.773268 14 9 14 C 6.2267316 14 4 11.773268 4 9 C 4 6.2267316 6.2267316 4 9 4 z"/>
            </svg>
          </button>
        </div>
      </form>
    </div>

    <div style="width: 100%; overflow-x: scroll; overflow-y: scroll; height: 20rem; scrollbar-width: thin;" class="bottom-line">
      <table class="table-custom">
        <thead>
          <tr>
            <td>ID pegawai</td>
            <td>Nama</td>
            <td>Jenis kelamin</td>
            <td>Tempat, tanggal lahir</td>
            <td>Status menikah</td>
            <td>Alamat</td>
            <td>NIK</td>
            <td>No BPJS</td>
            <td>No NPWP</td>
            <td>Nomor telepon</td>
            <td>Email</td>
            <td>Status</td>
            <td>Action</td>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(employee, index) in employeeData" :key="employee.id">
            <td>{{ employee.id }}</td>
            <td>{{ employee.name }}</td>
            <td>{{ employee.gender }}</td>
            <td>{{ employee.birth_place + ", " + formatDate(new Date(employee.birth_date)) }}</td>
            <td>{{ employee.wedding }}</td>
            <td>{{ employee.address }}</td>
            <td>{{ employee.nik }}</td>
            <td>{{ employee.bpjs }}</td>
            <td>{{ employee.npwp }}</td>
            <td>{{ employee.phone_number }}</td>
            <td>{{ employee.email }}</td>
            <td :style="employee.status ? 'color: lightgreen' : 'color: red'">{{ employee.status ? 'Aktif' : 'Non aktif' }}</td>
            <td style="position: relative;">
              <div class="drop-down">
                  <button class="act" @click="dropOpenTemplate[index] = !dropOpenTemplate[index]">:</button>
                  <div class="menu" v-if="dropOpenTemplate[index]">
                      <ul>
                          <li>
                              <button class="button-action" @click="handleDeleteEmployee(employee.id)">Delete</button>
                          </li>
                          <li>
                            <button class="button-action" @click="editEmployee(employee, index)">Update</button>
                          </li>
                          <li>
                            <button class="button-action" @click="router.push(`?id=${base64encode(employee.id, employee.name)}`); menuOpen[0] = true; dropOpenTemplate[index] = false">Add user access</button>
                          </li>
                      </ul>
                  </div>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </section>

  <teleport to='body' v-if="menuOpen[0]">
    <UserAccessForm>
      <template #btn-close>
        <button class="act" @click="menuOpen[0] = false">X</button>
      </template>
    </UserAccessForm>
  </teleport>
</template>
