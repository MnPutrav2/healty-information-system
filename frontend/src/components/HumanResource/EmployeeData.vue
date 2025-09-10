<script setup lang="ts">
import { addEmployee, deleteEmployee, getEmployee, updateEmployee } from '@/lib/api/hr';
import { formatDate } from '@/lib/formatDate';
import type { EmployeeData } from '@/types/hr';
import type { ResponseError, ResponseSuccess, SearchLimit } from '@/types/response';
import { onBeforeMount, reactive, ref } from 'vue';

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
  email: ""
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
  pages.value?.scrollIntoView({behavior: 'smooth'})

  updateID.value = data.id
  update.value = true
  open.value = true
  dropOpenTemplate.value[index] = false

  employeeReq.id = data.id
  employeeReq.name = data.name
  employeeReq.gender = data.gender
  employeeReq.birth_place = data.birth_place
  employeeReq.birth_date = data.birth_date
  employeeReq.wedding = data.wedding
  employeeReq.address = data.address
  employeeReq.nik = data.nik
  employeeReq.bpjs = data.bpjs
  employeeReq.npwp = data.npwp
  employeeReq.phone_number = data.phone_number
  employeeReq.email = data.email
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
    <h3 style="margin: 0.5rem;">Data Pegawai</h3>
    <div style="padding-top: 0.5rem; padding-bottom: 1rem;" class="bottom-line">
      <button class="button-action" style="margin: 0.5rem;" @click="open = !open; update = false" v-if="!open">Add pegawai</button>
      <form class="form-data-custom" v-else v-on:submit.prevent="">
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Tambah pegawai</h4>
        <div style="display: grid; grid-template-columns: auto auto auto; padding-left: 1rem;">
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="id">ID pegawai</label>
            </div>
            <input type="text" v-model="employeeReq.id" id="id" placeholder="ID pemeriksaan" required>
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="nm">Nama</label>
            </div>
            <input type="text" id="nm" v-model="employeeReq.name" placeholder="Nama" required>
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="jk">Jenis kelamin</label>
            </div>
            <select id="jk" v-model="employeeReq.gender" required>
              <option value="Laki - laki">Laki - laki</option>
              <option value="Perempuan">Perempuan</option>
            </select>
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="pn">Status pernikahan</label>
            </div>
            <select id="pn" v-model="employeeReq.wedding" required>
              <option value="Menikah">Menikah</option>
              <option value="Belum menikah">Belum menikah</option>
              <option value="Cerai">Cerai</option>
              <option value="Cerai mati">Cerai mati</option>
            </select>
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="bp">Tempat lahir</label>
            </div>
            <input type="text" id="bp" v-model="employeeReq.birth_place" placeholder="Tempat lahir" required>
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="tl">Tanggal lahir</label>
            </div>
            <input type="date" id="tl" v-model="employeeReq.birth_date" placeholder="Tempat lahir" required>
          </div>
          <div style="padding: 0.5rem;">
            <div>
              <label for="al">Alamat</label>
            </div>
            <textarea id="al" v-model="employeeReq.address" placeholder="Alamat" required></textarea>
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="nik">NIK</label>
            </div>
            <input type="text" v-model="employeeReq.nik" id="nik" placeholder="NIK" required>
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="bpj">BPJS</label>
            </div>
            <input type="text" v-model="employeeReq.bpjs" id="bpj" placeholder="BPJS" required>
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="npw">NPWP</label>
            </div>
            <input type="text" id="npw" v-model="employeeReq.npwp" placeholder="NPWP" required>
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="tlp">Nomor telpon</label>
            </div>
            <input type="text" id="tlp" v-model="employeeReq.phone_number" placeholder="Nomor telpon" required>
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="em">Email</label>
            </div>
            <input type="email" v-model="employeeReq.email" id="em" placeholder="Email" required>
          </div>
          <div>
            <button type="button" @click="resetForm">Reset</button>
            <button type="button" v-if="update" @click="handleUpdateEmployee">Update</button>
            <button type="submit" v-else @click="handleAddEmployee">Add</button>
            <button type="button" @click="open = !open">Close</button>
          </div>
        </div>
      </form>
    </div>

    <div>
      <form class="form-data-custom" v-on:submit.prevent="handleGetEmployee">
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Cari data pegawai</h4>
        <div class="center" style="justify-content: flex-start; align-items: flex-end; padding-left: 1rem;">
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="nv">Nama pegawai</label>
            </div>
            <input type="text" id="nv" v-model="search.search" placeholder="Nama pegawai">
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="l">Limit</label>
            </div>
            <input type="number" id="l" v-model="search.limit" placeholder="Limit">
          </div>
          <button type="submit">Cari</button>
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
            <td>Action</td>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(data, index) in employeeData" :key="data.id">
            <td>{{ data.id }}</td>
            <td>{{ data.name }}</td>
            <td>{{ data.gender }}</td>
            <td>{{ data.birth_place + ", " + formatDate(new Date(data.birth_date)) }}</td>
            <td>{{ data.wedding }}</td>
            <td>{{ data.address }}</td>
            <td>{{ data.nik }}</td>
            <td>{{ data.bpjs }}</td>
            <td>{{ data.npwp }}</td>
            <td>{{ data.phone_number }}</td>
            <td>{{ data.email }}</td>
            <td style="position: relative;">
              <div class="drop-down">
                  <button class="act" @click="dropOpenTemplate[index] = !dropOpenTemplate[index]">:</button>
                  <div class="menu" v-if="dropOpenTemplate[index]">
                      <ul>
                          <li>
                              <button @click="handleDeleteEmployee(data.id)">Delete</button>
                          </li>
                          <li>
                            <button @click="editEmployee(data, index)">Update</button>
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
</template>
