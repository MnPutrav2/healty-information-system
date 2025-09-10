<script setup lang="ts">
import { getPatient } from '@/lib/api/patient'
import { createRegister, deleteRegister, getCurrentCareNumber, getCurrentRegisterNumber, getRegisters } from '@/lib/api/register';
import { careNumber, regNumber } from '@/lib/careNumber';
import { formatDate, formatDatetime, viewedDate, viewedDateTime } from '@/lib/formatDate';
import { onBeforeMount, reactive, ref, watch } from 'vue'
import type { RegistrationData, RegisterData } from '@/types/register';
import type { Patient } from '@/types/patient';
import { getPoliclinics } from '@/lib/api/policlinic';
import type { Policlinics } from '@/types/policlinic';
import { getDoctors } from '@/lib/api/doctor';
import type { Doctors } from '@/types/doctor';
import type { SearchLimit } from '@/types/response';
import InputDataTime from '../Extras/InputDataTime.vue';
import InputData from '../Extras/InputData.vue';

// Define variabels
const gl = new Date()
const date1 = ref<string>(formatDatetime(gl, "00:00:00"))
const date2 = ref<string>(formatDatetime(gl, "23:59:59"))
const bool = ref<boolean>(false)
const patientDatas = ref<Patient[]>([])
const ambulatoryCare = ref<HTMLElement | null>()
const registerDatas = ref<RegistrationData[]>([])
const menuDataPatient = ref<boolean[]>([])
const doctors = ref<Doctors[] | null>(null)
const policlinics = ref<Policlinics[]>([])
let inter = setInterval(updateTime, 1000)
const payments = ref<Common[]>([
  {
    id: 'UMUM',
    name: 'UMUM'
  },
  {
    id: 'BPJS',
    name: 'BPJS'
  },
])
const registerData: RegisterData = reactive({
  care_number: '',
  register_number: '',
  register_date: '',
  medical_record: '',
  payment_method: '',
  policlinic: '',
  doctor: ''
})
registerData.register_date = formatDatetime(gl, null)
const search = reactive<SearchLimit>({
  search: "",
  limit: 5
})
const search2 = reactive<SearchLimit>({
  search: "",
  limit: 5
})

// External interface
interface Common {
  id: string;
  name: string;
}

// Define function
function clearForm() {
  registerData.medical_record = ''
  registerData.payment_method = ''
  registerData.policlinic = ''
  registerData.doctor = ''
}

function updateTime() {
  registerData.register_date = formatDatetime(new Date(), null)
}

function autoDate(bool: boolean) {
  if (bool) {
    clearInterval(inter)
  } else {
    inter = setInterval(updateTime, 1000)
  }
}

// Handler functions
async function handleReg(mr: string) {
  ambulatoryCare.value?.scrollIntoView({behavior: 'smooth'})

  registerData.medical_record = mr
}


async function handleCreateRegister() {
  const response = await createRegister(registerData)
  const json = await response.json()

  try {
    if (response.status === 201) {
      await handleGetRegister()
      await handleGetCareNumber()
      clearForm()
      alert("Data berhasil dibuat")
    } else {
      alert(json.errors)
    }
  } catch(error) {
    console.log(error)
  }
}

async function handleGetSearchPatient() {
  const response = await getPatient(search)
  const json = await response.json()

  try {
    if (response.status === 200) {
      patientDatas.value = json
    } else {
      console.log(json.errors)
    }
  } catch(error) {
    console.log(error)
  }
}

async function handleGetRegister() {
  const response = await getRegisters(date1.value, date2.value, search2)
  const json = await response.json()

  try {
    if (response.status === 200) {
      registerDatas.value = json
    } else {
      console.log(json.errors)
    }
  } catch(error) {
    console.log(error)
  }
}

async function handleDeleteRegister(care_num: string) {
  const response = await deleteRegister(care_num)
  const json = await response.json()

  try {
    if (response.status === 200) {
      await handleGetRegister()
      await handleGetCareNumber()
      alert("Data berhasil dihapus")
    } else {
      console.log(json.errors)
    }
  } catch(error) {
    console.log(error)
  }
}

async function handleGetRegNumber() {
  const response = await getCurrentRegisterNumber(registerData.register_date, registerData.policlinic)
  const json = await response.json()

  try {
    if (response.status === 200) {
      registerData.register_number = regNumber(json.response)
    } else {
      console.log(json.errors)
    }
  } catch(error) {
    console.log(error)
  }
}

async function handleGetCareNumber() {
  const response = await getCurrentCareNumber(registerData.register_date)
  const json = await response.json()

  try {
    if (response.status === 200) {
      const dat = formatDate(new Date(registerData.register_date))
      registerData.care_number = careNumber(new Date(dat), json.response)
    } else {
      console.log(json.errors)
    }
  } catch(error) {
    console.log(error)
  }
}

async function handleGetDoctors(id: string | null) {
  const response = await getDoctors(id)
  const json = await response.json()

  try {
    if (response.status === 200) {
      doctors.value = json
    } else {
      console.log(json.errors)
    }
  } catch(error) {
    console.log(error)
  }
}

async function handleGetPoliclinics() {
  const response = await getPoliclinics()
  const json = await response.json()

  try {
    if (response.status === 200) {
      policlinics.value = json
    } else {
      console.log(json.errors)
    }
  } catch(error) {
    console.log(error)
  }
}

// Watcher
watch(() => registerData.policlinic, async () => {
    await handleGetRegNumber()
    await handleGetCareNumber()
  }
)

watch(
  () => registerData.policlinic,
  async (newVal) => {
    const selected = policlinics.value.find(poli => poli.id === newVal)
    await handleGetDoctors(selected ? selected.poli_type : null)
  }
)

// Before view page
onBeforeMount(async () => {
  await handleGetRegister()
  await handleGetRegNumber()
  await handleGetCareNumber()
  await handleGetPoliclinics()
})
</script>

<template>
  <section class="anim-slide" ref="ambulatoryCare">
    <h3 style="margin: 0.5rem;">Registrasi Rawat Jalan</h3>
    <div style="margin-top: 2rem; margin-bottom: 2rem;" class="bottom-line">
      <form class="form-data-custom" v-on:submit.prevent="handleCreateRegister()">

        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Rawat Jalan</h4>
        <div style="display: grid; grid-template-columns: auto auto auto; padding-left: 1rem;">
          <InputData :props="{ id: 'regnum', name: 'Nomor registrasi' }">
            <input type="text" id="regnum" v-model="registerData.register_number" placeholder="Nomor registrasi" required>
          </InputData>
          <InputDataTime :props="{ id: 'dt1', name: 'Tanggal reg' }">
            <input type="datetime-local" step="1" id="date" v-model="registerData.register_date" placeholder="tanggal">
            <div :class="bool ? 'clock-inactive' : 'clock-active'" class="button-clock" @click="autoDate(bool = !bool)" role="button" tabindex="0"></div>
          </InputDataTime>
          <InputData :props="{ id: 'care_num', name: 'Nomor Rawat' }">
            <input type="text" id="care_num" v-model="registerData.care_number" placeholder="Nomor rawat" required>
          </InputData>
          <InputData :props="{ id: 'mr', name: 'Nomor RM' }">
            <input type="text" id="mr" v-model="registerData.medical_record" placeholder="Nomor RM" required>
          </InputData>
          <InputData :props="{ id: 'pay', name: 'Cara bayar' }">
            <select id="pay" v-model="registerData.payment_method">
              <option v-for="pay in payments" :key="pay.id" :value="pay.id">{{ pay.name }}</option>
            </select>
          </InputData>
          <InputData :props="{ id: 'poli', name: 'Poliklinik' }">
            <select id="poli" v-model="registerData.policlinic">
              <option v-for="poli in policlinics" :key="poli.id" :value="poli.id">{{ poli.name }}</option>
            </select>
          </InputData>
          <InputData :props="{ id: 'dr', name: 'Dokter' }" v-if="doctors">
            <select id="dr" v-model="registerData.doctor">
              <option v-for="doc in doctors" :key="doc.id" :value="doc.id">{{ doc.name }}</option>
            </select>
          </InputData>

        </div>

        <div>
          <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Action</h4>
          <button type="submit">Save</button>
        </div>

      </form>
    </div>

    <div style="padding-top: 2rem; padding-bottom: 2rem;">
      <form class="form-data-custom" v-on:submit.prevent="handleGetSearchPatient()">
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Cari pasien</h4>
        <div class="center" style="justify-content: flex-start; align-items: flex-end; padding-left: 1rem;">
          <InputData :props="{ id: 'sc', name: 'Cari' }">
            <input type="text" id="sc" v-model="search.search" placeholder="Nama pemeriksaan">
          </InputData>
          <InputData :props="{ id: 'lm', name: 'Limit' }">
            <input type="number" v-model="search.limit" id="lm" placeholder="limit">
          </InputData>
          <button>Cari</button>
        </div>
      </form>
    </div>

    <div style="width: 100%; overflow-x: scroll; overflow-y: scroll; height: 20rem; scrollbar-width: thin;" class="bottom-line">
      <table class="table-custom">
        <thead>
          <tr>
            <td>Action</td>
            <td>No rekam medis</td>
            <td>Nama</td>
            <td>Tempat/tanggal lahir</td>
            <td>Jenis Kelamin</td>
            <td>Alamat</td>
            <td>NIK</td>
            <td>BPJS</td>
            <td>No telepon</td>
          </tr>
        </thead>
        <tbody>
          <tr v-for="patient in patientDatas" :key="patient.medical_record">
            <td>
              <button class="button-action" @click="handleReg(patient.medical_record)">Register</button>
            </td>
            <td>{{ patient.medical_record }}</td>
            <td>{{ patient.name }}</td>
            <td>{{ `${patient.birth_place}, ${viewedDate(patient.birth_date)}` }}</td>
            <td>{{ patient.gender }}</td>
            <td>{{ patient.address }}</td>
            <td>{{ patient.nik }}</td>
            <td>{{ patient.bpjs }}</td>
            <td>{{ patient.phone_number }}</td>
          </tr>
        </tbody>
      </table>
    </div>

    <div style="padding-top: 2rem; padding-bottom: 2rem;">
      <form class="form-data-custom" v-on:submit.prevent="handleGetRegister">
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Cari pasien Rawat Jalan</h4>
        <div class="center" style="justify-content: flex-start; align-items: flex-end; padding-left: 1rem;">
          <InputData :props="{ id: 'dt1', name: 'Tanggal awal' }">
            <input type="datetime-local" step="1" id="dt1" v-model="date1" placeholder="tanggal">
          </InputData>
          <InputData :props="{ id: 'dt2', name: 'Tanggal akhir' }">
            <input type="datetime-local" step="1" id="dt2" v-model="date2" placeholder="tanggal">
          </InputData>
          <InputData :props="{ id: 'sc1', name: 'Cari' }">
            <input type="text" id="sc" v-model="search2.search" placeholder="No Rawat/Nama pasien">
          </InputData>
          <InputData :props="{ id: 'lm1', name: 'Limit' }">
            <input type="number" v-model="search2.limit" id="lm1" placeholder="limit">
          </InputData>
          <button>Cari</button>
        </div>
      </form>
    </div>

    <div style="width: 100%; overflow-x: scroll; overflow-y: scroll; height: 20rem; scrollbar-width: thin;">
      <table class="table-custom">
        <thead>
          <tr>
            <td>No Rawat</td>
            <td>No Reg</td>
            <td>No rekam medis</td>
            <td>Nama</td>
            <td>Tanggal registrasi</td>
            <td>Jenis Kelamin</td>
            <td>Cara Bayar</td>
            <td>Poliklinik</td>
            <td>Dokter</td>
            <td>Action</td>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(reg, index) in registerDatas" :key="reg.care_number">
            <td>{{ reg.care_number }}</td>
            <td>{{ reg.register_number }}</td>
            <td>{{ reg.medical_record }}</td>
            <td>{{ reg.name }}</td>
            <td>{{ viewedDateTime(reg.register_date) }}</td>
            <td>{{ reg.gender }}</td>
            <td>{{ reg.payment_method }}</td>
            <td>{{ reg.policlinic_name }}</td>
            <td>{{ reg.doctor_name }}</td>
            <td style="position: relative;">
                <div class="drop-down">
                    <button class="act" @click="menuDataPatient[index] = !menuDataPatient[index]">:</button>
                    <div class="menu" v-if="menuDataPatient[index]">
                        <ul>
                            <li>
                              <button class="button-action" @click="handleDeleteRegister(reg.care_number); menuDataPatient[index] = false">Delete</button>
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
