<script setup lang="ts">
import { createRequestLabolatorium, getLabolatoriumData, getLabolatoriumNumber } from '@/lib/api/labolatorium'
import { formatDatetime } from '@/lib/formatDate'
import { labNumber } from '@/lib/labNumber'
import type { LabsType, LabTypeRequest } from '@/types/labolatorium'
import { onBeforeMount, reactive, ref, watch } from 'vue'
import InputData from '../Extras/InputData.vue'

// Define variabel

const props = defineProps(['data'])
const date = new Date()
const search = ref<string>("")
const labsDataRequest = ref<LabsType[]>([])
const labDataForRequest = reactive<LabTypeRequest>({
  care_number: props.data,
  date: formatDatetime(date, null),
  lab_id:"0001000",
  officer_id: "A00001",
  referral_id: "S00001",
  data: labsDataRequest.value
})
const labDatas = ref<LabsType[]>([])

// Define function

async function handleCreateLabRequest() {
  const data: LabTypeRequest = {...labDataForRequest, data:labsDataRequest.value}

  const response = await createRequestLabolatorium(data)
  const json = await response.json()

  if (response.status === 201) {
    alert("success")
    labsDataRequest.value = []
  } else {
    alert(json.errors)
  }
}

function handleAddLabData(lab: LabsType) {
  labsDataRequest.value.push(lab)
}

function handleDeleteLabData(id: string){
  labsDataRequest.value = labsDataRequest.value.filter(d => d.id !== id)
}

async function handleGetLabolatoriumData() {
    const response = await getLabolatoriumData(5, search.value)
    const json = await response.json()

    if (response.status === 200) {
        labDatas.value = json
    } else {
        alert(json.errors)
    }
}

async function handleGetLabolatoriumNumber(date: string): Promise<number> {
    const response = await getLabolatoriumNumber(date)
    const json = await response.json()

    if (response.status === 200) {
        return parseInt(json.response)
    } else {
        return 0
    }
}

watch(() => labDataForRequest.date, (newDate) => {
  (async () => {
    const number = await handleGetLabolatoriumNumber(newDate)
    labDataForRequest.lab_id = labNumber(new Date(newDate), number)
  })()
})

// mount
onBeforeMount(async () => {
  await handleGetLabolatoriumData()
  const number = await handleGetLabolatoriumNumber(formatDatetime(date, null))
  labDataForRequest.lab_id = labNumber(date, number)
})

</script>

<template>
  <teleport to='body'>
    <section class="container center" >
      <div class="cover2 scroll">
        <div class="close-btn center">
          <slot name="btn-close"></slot>
        </div>

        <div style="padding-top: 2rem; padding-bottom: 2rem;">
          <form class="form-data-custom" v-on:submit.prevent="handleCreateLabRequest">
            <div class="center" style="justify-content: flex-start; align-items: flex-end; padding-left: 1rem;">
              <InputData :props="{ id: 'mr', name: 'Nomor rawat' }">
                <input type="text" id="mr" v-model="labDataForRequest.care_number" readonly placeholder="no rawat">
              </InputData>
              <InputData :props="{ id: 'ss', name: 'Nomor permintaan' }">
                <input type="text" id="ss" v-model="labDataForRequest.lab_id" placeholder="permintaan">
              </InputData>
              <InputData :props="{ id: 'da', name: 'Tanggal permintaan' }">
                <input type="datetime-local" v-model="labDataForRequest.date" id="da" step="1" placeholder="resep">
              </InputData>
              <button>Save</button>
            </div>
          </form>
        </div>

        <h4 style="margin-bottom: 1rem;">Pemeriksaan lab</h4>
        <hr>

        <div style="width: 100%; height: 15rem; margin-top: 1rem" class="scroll">
          <table class="table-custom" style="width: 100%">
            <thead>
              <tr>
                <td>Action</td>
                <td>Nama</td>
                <td>Harga</td>
              </tr>
            </thead>
            <tbody>
              <tr v-for="lab in labsDataRequest" :key="lab.id">
                <td>
                  <button class="button-action" @click="handleDeleteLabData(lab.id)">Delete</button>
                </td>
                <td>{{ lab.name }}</td>
                <td>{{ lab.total }}</td>
              </tr>
            </tbody>
          </table>
        </div>

        <div style="padding-top: 2rem; padding-bottom: 2rem;">
          <form class="form-data-custom">
            <div class="center" style="justify-content: flex-start; align-items: flex-end; padding-left: 1rem;">
              <InputData :props="{ id: 'nr', name: 'Cari nama pemeriksaan' }">
                <input type="text" v-model="search" id="nr" placeholder="cari pemeriksaan">
              </InputData>
              <button>Find</button>
            </div>
          </form>
        </div>

        <hr>

        <div style="width: 100%; height: 15rem; margin-top: 1rem" class="scroll">
          <table class="table-custom" style="width: 100%">
            <thead>
              <tr>
                <td>Action</td>
                <td>Nama</td>
                <td>Harga</td>
              </tr>
            </thead>
            <tbody>
              <tr v-for="lab in labDatas" :key="lab.id">
                <td>
                  <button class="button-action" @click="handleAddLabData(lab)">Add</button>
                </td>
                <td>{{ lab.name }}</td>
                <td>{{ lab.total }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </section>
  </teleport>
</template>

<style scoped>
.container {
  inset: 0;
  position: fixed;
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
