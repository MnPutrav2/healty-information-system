<script setup lang="ts">
import { createExamination, deleteExamination, getExamination, updateExamination } from '@/lib/api/finance';
import type { ExaminationCost } from '@/types/finance';
import type { ResponseError, ResponseSuccess, SearchLimit } from '@/types/response';
import { reactive, ref } from 'vue';
import InputData from '@/components/Extras/InputData.vue';

const examinationData = reactive<ExaminationCost>({
  id: "",
  examination_name: "",
  payment_method: "",
  doctor_cost: 0,
  nurse_cost: 0,
  management_cost: 0
})

const search = reactive<SearchLimit>({
  search: "",
  limit: 10
})

const datas = ref<ExaminationCost[]>()
const dropOpen = ref<boolean[]>([])
const update = ref<boolean>(false)
const idUpdate = ref<string>("")
const pages = ref<HTMLElement | null>()

async function handleCreateExaminationData() {
  try {
    const response = await createExamination(examinationData)

    if(response.status === 201){
      const json: ResponseSuccess = await response.json()
      await handleGetExamination()

      alert(json.response)
    }else{
      const json: ResponseError = await response.json()

      alert(json.errors)
    }
  }catch(e){
    alert(`Error : ${e}`)
  }
}

async function handleGetExamination() {
  try {
    const response: Response = await getExamination(search)

    if(response.status === 200){
      datas.value = await response.json()
    }else{
      const json: ResponseError = await response.json()

      alert(json.errors)
    }
  }catch(e){
    alert(`Error : ${e}`)
  }
}

async function handleDeleteExamination(id: string) {
  try {
    const response: Response = await deleteExamination(id)

    if(response.status === 200){
      const json: ResponseSuccess = await response.json()
      await handleGetExamination()

      alert(json.response)
    }else{
      const json: ResponseError = await response.json()

      alert(json.errors)
    }
  }catch(e){
    alert(`Error : ${e}`)
  }
}

async function handleUpdateExamination() {
  try {
    const response: Response = await updateExamination(examinationData, idUpdate.value)

    if(response.status === 200){
      const json: ResponseSuccess = await response.json()
      await handleGetExamination()

      alert(json.response)
    }else{
      const json: ResponseError = await response.json()

      alert(json.errors)
    }

    update.value = false
  }catch(e){
    update.value = false

    alert(`Error : ${e}`)
  }
}

function editExamination(data: ExaminationCost) {
  idUpdate.value = data.id

  pages.value?.scrollIntoView({behavior: "smooth"})

  examinationData.id = data.id
  examinationData.examination_name = data.examination_name
  examinationData.payment_method = data.payment_method
  examinationData.doctor_cost = data.doctor_cost
  examinationData.nurse_cost = data.nurse_cost
  examinationData.management_cost = data.management_cost
}
</script>

<template>
  <section class="anim-slide" ref="pages">
    <h3 style="margin: 0.5rem;">Data Pemeriksaan</h3>
    <div style="padding-top: 2rem; padding-bottom: 2rem;" class="bottom-line">
      <form class="form-data-custom" v-on:submit.prevent="handleCreateExaminationData">
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Buat data pemeriksaan</h4>
        <div class="responsive-grid" style="padding-left: 1rem;">
          <InputData :props="{ id: 'id', name: 'ID' }">
            <input type="text" id="id" v-model="examinationData.id" placeholder="ID">
          </InputData>
          <InputData :props="{ id: 'nm', name: 'Nama pemeriksaan' }">
            <input type="text" id="nm" v-model="examinationData.examination_name" placeholder="Nama pemeriksaan">
          </InputData>
          <InputData :props="{ id: 'by', name: 'Cara bayar' }">
            <select id="by" v-model="examinationData.payment_method">
              <option value="UMUM">UMUM</option>
              <option value="BPJS">BPJS</option>
            </select>
          </InputData>
          <InputData :props="{ id: 'dr', name: 'Biaya dokter' }">
            <input type="number" id="dr" v-model="examinationData.doctor_cost" placeholder="Biaya dokter">
          </InputData>
          <InputData :props="{ id: 'pr', name: 'Biaya perawat' }">
            <input type="number" id="pr" v-model="examinationData.nurse_cost" placeholder="Biaya perawat">
          </InputData>
          <InputData :props="{ id: 'ma', name: 'Manajement' }">
            <input type="number" id="ma" v-model="examinationData.management_cost" placeholder="Biaya manajement">
          </InputData>
        </div>
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Action</h4>
        <div style="padding-left: 1rem;">
          <button type="submit">Create</button>
          <button type="button" v-if="update" @click="handleUpdateExamination">Update</button>
        </div>
      </form>
    </div>

    <div style="padding: 0.5rem; padding-top: 2rem; padding-bottom: 2rem;">
      <form class="form-search-input" v-on:submit.prevent="handleGetExamination">
        <div class="box-search">
          <div class="box">
            <input class="search-input" type="number" id="lm" v-model="search.limit" placeholder="limit">
            <input class="search-input" type="text" id="sc" v-model="search.search" placeholder="Nama pemeriksaan">
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
            <td>ID</td>
            <td>Nama pemeriksaan</td>
            <td>Cara bayar</td>
            <td>Biaya dokter</td>
            <td>Biaya perawat</td>
            <td>Biaya management</td>
            <td>Action</td>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(data, index) in datas" :key="data.id">
            <td>{{ data.id }}</td>
            <td>{{ data.examination_name }}</td>
            <td>{{ data.payment_method }}</td>
            <td>{{ data.doctor_cost }}</td>
            <td>{{ data.nurse_cost }}</td>
            <td>{{ data.management_cost }}</td>
            <td style="position: relative;">
              <div class="drop-down">
                <button class="act" @click="dropOpen[index] = !dropOpen[index]">:</button>
                <div class="menu" v-if="dropOpen[index]">
                  <ul>
                    <li>
                        <button class="button-action" @click="handleDeleteExamination(data.id); dropOpen[index] = false">Delete</button>
                    </li>
                    <li>
                        <button class="button-action" @click="editExamination(data), dropOpen[index] = false, update = true">Update</button>
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
