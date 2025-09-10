<script setup lang="ts">
import { createAmbulatoryCare, deleteAmbulatoryCare, getAmbulatoryCare, updateAmbulatoryCare } from '@/lib/api/ambulatory';
import { getRegisters, updateStatus } from '@/lib/api/register';
import { formatDatetime, viewedDateTime } from '@/lib/formatDate';
import { resetForm } from '@/lib/resetForm';
import { nextTick, onBeforeMount, reactive, ref } from 'vue';
import type { RegistrationData } from '@/types/register';
import type { AmbulatoryCareRequest, AmbulatoryCare, AmbulatoryCareRequestUpdate } from '@/types/ambulatory';
import type { SearchLimit } from '@/types/response';
import router from '@/router';
import InputData from './Extras/InputData.vue';
import InputDataTime from './Extras/InputDataTime.vue';

// Define variabels
const user = defineProps(['data'])
const gl = new Date()
const date1 = ref<string>(formatDatetime(gl, "00:00:00"))
const date2 = ref<string>(formatDatetime(gl, "23:59:59"))
const registerDatas = ref<RegistrationData[]>([])
const pages = ref<HTMLElement | null>()
const searchAmbulatory = ref<string | null>("")
const dateSearch1 = ref<string>(formatDatetime(gl, "00:00:00"))
const dateSearch2 = ref<string>(formatDatetime(gl, "23:59:00"))
const ambulatoryDatas = ref<AmbulatoryCare[]>([])
const editAction = ref<boolean>(false)
const bool = ref<boolean>(false)
const ambInput = ref<boolean>(false)
const menuDataPatient = ref<boolean[]>([])
const menuDataAmbulatory = ref<boolean[]>([])
const dateForEdit = ref<string>('')
let inter = setInterval(updateTime, 1000)
const officers = [
  {
    id: "A00001",
    name: "ADMIN"
  },
  {
    id: "EMP001",
    name: "AGUS"
  },
]
const search = reactive<SearchLimit>({
  search: "",
  limit: 5
})
const ambulatoryData = reactive<AmbulatoryCareRequest>({
  care_number: '',
  medical_record: '',
  date: '',
  body_temperature: 0,
  tension: '',
  pulse: 0,
  respiration: 0,
  height: 0,
  weight: 0,
  spo2: 0,
  gcs: 0,
  awareness: 'Compos Mentis',
  complaint: '',
  examination: '',
  allergy: '',
  followup: '',
  assessment: '',
  instructions: '',
  evaluation: '',
  officer: user.data.id
})

// Define functions
function updateTime() {
  ambulatoryData.date = formatDatetime(new Date(), null)
}

function autoDate(bool: boolean) {
  if (bool) {
    clearInterval(inter)
  } else {
    inter = setInterval(updateTime, 1000)
  }
}

// Handler functions
async function handleCreateAmbulatory() {
  const response = await createAmbulatoryCare(ambulatoryData)
  const json = await response.json()

  try {
    if (response.status === 201) {
      alert("Berhasil simpan data")
      await handleGetAmbulatoryData()
      await handleGetRegister()
      resetForm(ambulatoryData, ambulatoryData.officer)
      ambInput.value = false
    } else {
      alert(json.errors)
    }
  } catch(error) {
    console.log(error)
  }
}

async function handleInputAmbulatory(care: string, mr: string) {
  pages.value?.scrollIntoView({behavior: "smooth"})

  searchAmbulatory.value = care

  ambulatoryData.care_number = care
  ambulatoryData.medical_record = mr
  ambulatoryData.date = ''
  ambulatoryData.body_temperature = 0
  ambulatoryData.tension = ''
  ambulatoryData.pulse = 0
  ambulatoryData.respiration = 0
  ambulatoryData.height = 0
  ambulatoryData.weight = 0
  ambulatoryData.spo2 = 0
  ambulatoryData.gcs = 0
  ambulatoryData.awareness = 'Compos Mentis'
  ambulatoryData.complaint = ''
  ambulatoryData.examination = ''
  ambulatoryData.allergy = '-'
  ambulatoryData.followup = ''
  ambulatoryData.assessment = ''
  ambulatoryData.instructions = ''
  ambulatoryData.evaluation = ''
  ambulatoryData.officer = user.data.id

  await handleGetAmbulatoryData()
}

async function handleGetAmbulatoryData() {
  const response = await getAmbulatoryCare(searchAmbulatory.value, dateSearch1.value, dateSearch2.value)
  const json = await response.json()

  try {
    if (response.status === 200) {
      ambulatoryDatas.value = json
    } else {
      console.log(json.errors)
    }
  } catch(error) {
    console.log(error)
  }
}

async function handleUpdateStatus(careNum: string, status: string) {
  const response = await updateStatus(careNum, status)
  const json = await response.json()

  try {
    if (response.status === 200) {
      alert("success")
      await handleGetRegister()
    } else {
      alert(json.errors)
    }
  } catch(error) {
    alert('error :' + error)
  }
}

async function handleGetRegister() {
  const response = await getRegisters(date1.value, date2.value, search)
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

function handleEditAmbulatory(amb: AmbulatoryCare) {
  ambInput.value = true

  nextTick(() => {
    pages.value?.scrollIntoView({behavior: "smooth"})
  })
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const {officer_name, date, ...filter} = amb
  Object.assign(ambulatoryData, filter)
  ambulatoryData.date = formatDatetime(new Date(amb.date), null)
}

async function handleDeleteAmbulatory(careNum: string, date: string) {
  const response = await deleteAmbulatoryCare(careNum, date)
  const json = await response.json()

  try {
    if (response.status === 200) {
      await handleGetAmbulatoryData()
      alert("hapus data berhasil")
    } else {
      alert(json.errors)
    }
  } catch(error) {
    console.log(error)
  }
}

async function handleUpdateAmbulatory(careNum: string, date: string) {
  const data = ref<AmbulatoryCareRequestUpdate>({
    care_number: careNum,
    date: date,
    data: ambulatoryData
  })

  const response = await updateAmbulatoryCare(data.value)
  const json = await response.json()

  try {
    if (response.status === 200) {
      await handleGetAmbulatoryData()
      alert("update data berhasil")
      ambInput.value = false
    } else {
      alert(json.errors)
    }
  } catch(error) {
    console.log(error)
  }
}

// Before page view
onBeforeMount(async () => {
  await handleGetRegister()
})
</script>

<template>
  <section class="anim-slide" ref="pages">
    <div v-if="ambInput">
      <h3 style="margin: 0.5rem;">Rawat Jalan</h3>
      <div style="margin-top: 2rem; margin-bottom: 2rem;" class="bottom-line">
        <form class="form-data-custom" v-on:submit.prevent="handleCreateAmbulatory">

          <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Input Pemeriksaan</h4>
          <div style="display: grid; grid-template-columns: auto auto auto; padding-left: 1rem;">
            <InputData :props="{ id: 'cm', name: 'Nomor rawat' }">
              <input type="text" id="cm" v-model="ambulatoryData.care_number" placeholder="Nomor rawat" required>
            </InputData>
            <InputData :props="{ id: 'awa', name: 'Kesadaran' }">
              <select id="awa" v-model="ambulatoryData.awareness">
                <option value="Compos Mentis">Compos Mentis</option>
                <option value="Somnolence">Somnolence</option>
                <option value="Sopor">Sopor</option>
                <option value="Coma">Coma</option>
                <option value="Alert">Alert</option>
                <option value="Confusion">Confusion</option>
                <option value="Voice">Voice</option>
                <option value="Pain">Pain</option>
                <option value="Unresponsive">Unresponsive</option>
              </select>
            </InputData>
            <InputData :props="{ id: 'al', name: 'Alergi' }">
              <input type="text" id="al" v-model="ambulatoryData.allergy" placeholder="Alergi" required>
            </InputData>
            <InputDataTime :props="{ id: 'date', name: 'Tanggal' }">
              <input type="datetime-local" step="1" id="date" v-model="ambulatoryData.date" placeholder="tanggal">
              <div :class="bool ? 'clock-inactive' : 'clock-active'" class="button-clock" @click="autoDate(bool = !bool)" role="button" tabindex="0"></div>
            </InputDataTime>
            <InputData :props="{ id: 'ofc', name: 'Petugas' }">
              <select id="ofc" v-model="ambulatoryData.officer">
                <option v-for="ofc in officers" :value="ofc.id" :key="ofc.id">{{ ofc.name }}</option>
              </select>
            </InputData>
          </div>

          <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Pemeriksaan</h4>
          <div style="display: grid; grid-template-columns: auto auto auto; padding-left: 1rem;">
            <InputData :props="{ id: 'complaint', name: 'Subject' }">
              <textarea id="complaint" v-model="ambulatoryData.complaint" placeholder="keluhan pasien"></textarea>
            </InputData>
            <InputData :props="{ id: 'exam', name: 'Objek' }">
              <textarea id="exam" v-model="ambulatoryData.examination" placeholder="Pemeriksaan"></textarea>
            </InputData>
            <InputData :props="{ id: 'as', name: 'Assesment' }">
              <textarea id="as" v-model="ambulatoryData.assessment" placeholder="Assesment"></textarea>
            </InputData>
            <InputData :props="{ id: 'as', name: 'Plan' }">
              <textarea id="fol" v-model="ambulatoryData.followup" placeholder="Tindak lanjut"></textarea>
            </InputData>
            <InputData :props="{ id: 'int', name: 'Intruksi' }">
              <textarea id="int" v-model="ambulatoryData.instructions" placeholder="Intruksi"></textarea>
            </InputData>
            <InputData :props="{ id: 'evl', name: 'Evaluasi' }">
              <textarea id="evl" v-model="ambulatoryData.evaluation" placeholder="Evaluasi"></textarea>
            </InputData>
          </div>

          <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Pemeriksaan</h4>
          <div style="display: grid; grid-template-columns: auto auto auto; padding-left: 1rem;">
            <InputData :props="{ id: 'bt', name: 'Suhu Badan' }">
              <input type="number" id="bt" v-model="ambulatoryData.body_temperature" placeholder="suhu badan">
            </InputData>
            <InputData :props="{ id: 'hg', name: 'Tinggi Badan (Cm)' }">
              <input type="number" id="bt" v-model="ambulatoryData.height" placeholder="tinggi badan">
            </InputData>
            <InputData :props="{ id: 'wg', name: 'Berat Badan (Kg)' }">
              <input type="number" id="wg" v-model="ambulatoryData.weight" placeholder="berat badan">
            </InputData>
            <InputData :props="{ id: 'ts', name: 'Tensi (mmHg)' }">
              <input type="text" id="ts" v-model="ambulatoryData.tension" placeholder="Tensi">
            </InputData>
            <InputData :props="{ id: 'pl', name: 'Nadi' }">
              <input type="number" id="pl" v-model="ambulatoryData.pulse" placeholder="nadi">
            </InputData>
            <InputData :props="{ id: 'spo', name: 'SPO' }">
              <input type="number" id="spo" v-model="ambulatoryData.spo2" placeholder="spo2">
            </InputData>
            <InputData :props="{ id: 'gcs', name: 'GCS' }">
              <input type="number" id="gcs" v-model="ambulatoryData.gcs" placeholder="GCS">
            </InputData>
            <InputData :props="{ id: 'rs', name: 'Respirasi' }">
              <input type="number" id="rs" v-model="ambulatoryData.respiration" placeholder="Respirasi">
            </InputData>
          </div>

          <div>
            <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Action</h4>
            <button type="submit">Save</button>
            <button type="button" v-if="editAction" @click="handleUpdateAmbulatory(ambulatoryData.care_number, dateForEdit)">Update</button>
          </div>

        </form>
      </div>
    </div>

    <div style="padding-top: 2rem; padding-bottom: 2rem;">
      <form class="form-data-custom" v-on:submit.prevent="handleGetAmbulatoryData">
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Data Pemeriksaan Pasien</h4>
        <div class="center" style="justify-content: flex-start; align-items: flex-end; padding-left: 1rem;">
          <InputData :props="{ id: 'dt1', name: 'Tanggal awal' }">
            <input type="datetime-local" step="1" id="dt1" v-model="dateSearch1" placeholder="date">
          </InputData>
          <InputData :props="{ id: 'dt2', name: 'Tanggal akhir' }">
            <input type="datetime-local" step="1" id="dt2" v-model="dateSearch2" placeholder="date">
          </InputData>
          <InputData :props="{ id: 'sc', name: 'Cari' }">
            <input type="text" id="sc" v-model="searchAmbulatory" placeholder="No Rawat/Nama Pasien">
          </InputData>
          <button>Cari</button>
        </div>
      </form>
    </div>

    <div style="width: 100%; overflow-x: scroll; overflow-y: scroll; height: 20rem; scrollbar-width: thin;" class="bottom-line">
      <table class="table-custom">
        <thead>
          <tr>
            <td>No Rawat</td>
            <td>Nama</td>
            <td>Tanggal Rawat</td>
            <td>No Rekam Medis</td>
            <td>Petugas</td>
            <td>Suhu</td>
            <td>Tensi</td>
            <td>Nadi</td>
            <td>Berat Badan</td>
            <td>Tinggi Badan</td>
            <td>SPO2</td>
            <td>GCS</td>
            <td>Respirasi</td>
            <td>Keluhan</td>
            <td>Assesment</td>
            <td>Intruksi</td>
            <td>Pemeriksaan</td>
            <td>Tindak Lanjut</td>
            <td>Evaluasi</td>
            <td>Action</td>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(amb, index) in ambulatoryDatas" :key="index">
            <td>{{ amb.care_number }}</td>
            <td>{{ amb.name }}</td>
            <td>{{ viewedDateTime(amb.date) }}</td>
            <td>{{ amb.medical_record }}</td>
            <td>{{ amb.officer_name }}</td>
            <td>{{ amb.body_temperature }}</td>
            <td>{{ amb.tension }}</td>
            <td>{{ amb.pulse }}</td>
            <td>{{ amb.weight }}</td>
            <td>{{ amb.height }}</td>
            <td>{{ amb.spo2 }}</td>
            <td>{{ amb.gcs }}</td>
            <td>{{ amb.respiration }}</td>
            <td>{{ amb.complaint }}</td>
            <td>{{ amb.assessment }}</td>
            <td>{{ amb.instructions }}</td>
            <td>{{ amb.examination }}</td>
            <td>{{ amb.followup }}</td>
            <td>{{ amb.evaluation }}</td>
            <td style="position: relative;">
                <div class="drop-down">
                    <button class="act" @click="menuDataAmbulatory[index] = !menuDataAmbulatory[index]">:</button>
                    <div class="menu" v-if="menuDataAmbulatory[index]">
                        <ul>
                            <li>
                              <button class="button-action" @click="handleEditAmbulatory(amb); editAction = true; dateForEdit = amb.date; menuDataAmbulatory[index] = false">Edit</button>
                            </li>
                            <li>
                                <button class="button-action" @click="handleDeleteAmbulatory(amb.care_number, amb.date); menuDataAmbulatory[index] = false">Delete</button>
                            </li>
                        </ul>
                    </div>
                </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div style="padding-top: 2rem; padding-bottom: 2rem;">
      <form class="form-data-custom" v-on:submit.prevent="handleGetRegister">
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Cari pasien Rawat Jalan</h4>
        <div class="center" style="justify-content: flex-start; align-items: flex-end; padding-left: 1rem;">
          <InputData :props="{ id: 'dtd1', name: 'Tanggal awal' }">
            <input type="datetime-local" step="1" id="dt1" v-model="date1" placeholder="date">
          </InputData>
          <InputData :props="{ id: 'dtd2', name: 'Tanggal akhir' }">
            <input type="datetime-local" step="1" id="dt2" v-model="date2" placeholder="date">
          </InputData>
          <InputData :props="{ id: 'sc', name: 'Cari' }">
            <input type="text" id="sc" v-model="search.search" placeholder="No Rawat/Nama Pasien">
          </InputData>
          <InputData :props="{ id: 'lm', name: 'Limit' }">
            <input type="number" v-model="search.limit" id="lm" placeholder="limit">
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
            <td>No rekam medis</td>
            <td>Nama</td>
            <td>Tanggal registrasi</td>
            <td>Jenis Kelamin</td>
            <td>Cara Bayar</td>
            <td>Poliklinik</td>
            <td>Dokter</td>
            <td>Status</td>
            <td>Action</td>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(reg, index) in registerDatas" :key="reg.care_number" :class="reg.status == 'Sudah' ? 'status' : ''">
            <td>{{ reg.care_number }}</td>
            <td>{{ reg.medical_record }}</td>
            <td>{{ reg.name }}</td>
            <td>{{ viewedDateTime(reg.register_date) }}</td>
            <td>{{ reg.gender }}</td>
            <td>{{ reg.payment_method }}</td>
            <td>{{ reg.policlinic_name }}</td>
            <td>{{ reg.doctor_name }}</td>
            <td>{{ reg.status }}</td>
            <td style="position: relative;">
              <div class="drop-down">
                <button class="act" @click="menuDataPatient[index] = !menuDataPatient[index]">:</button>
                <div class="menu" v-if="menuDataPatient[index]">
                  <ul>
                    <li>
                      <button class="button-action" @click="handleInputAmbulatory(reg.care_number, reg.medical_record); menuDataPatient[index] = false; ambInput = true">Input</button>
                    </li>
                    <li>
                        <button class="button-action" @click="router.push(`?careNum=${reg.care_number}`);menuDataPatient[index] = false">Menu</button>
                    </li>
                    <li>
                      <button class="button-action" @click="handleUpdateStatus(reg.care_number, 'Sudah'); menuDataPatient[index] = false">Sudah periksa</button>
                    </li>
                    <li>
                      <button class="button-action" @click="handleUpdateStatus(reg.care_number, 'Belum'); menuDataPatient[index] = false">Belum periksa</button>
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

<style scoped>
.status > td {
  background-color: rgba(0, 255, 0, 0.107);
}
</style>
