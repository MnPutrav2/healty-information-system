<script setup lang="ts">
import { getDoctorsAll } from '@/lib/api/doctor';
import { addExamination, deleteExaminationData, getExaminationData } from '@/lib/api/examination';
import { getExamination } from '@/lib/api/finance';
import { formatDatetime, viewedDateTime } from '@/lib/formatDate';
import type { Doctors } from '@/types/doctor';
import type { ExaminationDetail, ExaminationReq } from '@/types/examination';
import type { ExaminationCost } from '@/types/finance';
import type { ResponseError, ResponseSuccess, SearchLimit } from '@/types/response';
import { onBeforeMount, reactive, ref } from 'vue';
import { useRoute } from 'vue-router';
import InputData from '../Extras/InputData.vue';

const doctors = ref<Doctors[]>([])
const route = useRoute()
const date = new Date()
const examinationAdd = reactive<ExaminationReq>({
  care_number: String(route.query.careNumber),
  examination: "",
  doctor_id: "",
  nurse_id: "",
  service_type: "Ralan",
  date: formatDatetime(date, null)
})

const search = reactive<SearchLimit>({
  search: "",
  limit: 10
})

const examData = ref<ExaminationCost[]>([])
const examinationData = ref<ExaminationDetail[]>([])

async function handleGetDoctorAll() {
  try{
    const response: Response = await getDoctorsAll();

    if(response.status === 200){
      doctors.value = await response.json()
    }else{
      const json: ResponseError = await response.json()

      alert(json.errors)
    }
  }catch(e){
    alert(e)
  }
}

async function handleDeleteExaminationData(id: string) {
  try{
    const response: Response = await deleteExaminationData(id);

    if(response.status === 200){
      const json: ResponseSuccess = await response.json()
      await handleGetExaminationData()

      alert(json.response)
    }else{
      const json: ResponseError = await response.json()

      alert(json.errors)
    }
  }catch(e){
    alert(e)
  }
}

async function handleAddExamination(id: string) {
  const data: ExaminationReq = {...examinationAdd, examination: id}

  try{
    const response: Response = await addExamination(data);

    if(response.status === 201){
      await handleGetExaminationData()
    }else{
      const json: ResponseError = await response.json()

      alert(json.errors)
    }
  }catch(e){
    alert(e)
  }
}

async function handleGetExamination() {
  try {
    const response: Response = await getExamination(search)

    if(response.status === 200){
      examData.value = await response.json()
    }else{
      const json: ResponseError = await response.json()

      alert(json.errors)
    }
  }catch(e){
    alert(`Error : ${e}`)
  }
}

async function handleGetExaminationData() {
  try {
    const response: Response = await getExaminationData(String(route.query.careNumber))

    if(response.status === 200){
      examinationData.value = await response.json()
    }else{
      const json: ResponseError = await response.json()

      alert(json.errors)
    }
  }catch(e){
    alert(`Error : ${e}`)
  }
}

onBeforeMount(async () => {
  await handleGetDoctorAll()
  await handleGetExaminationData()
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
          <form class="form-data-custom" v-on:submit.prevent="">
            <div class="center" style="justify-content: flex-start; align-items: flex-end; padding-left: 1rem;">
              <div style="padding: 0.5rem;">
                <div style="padding: 0.5rem;">
                  <label style="font-size: var(--font-size);" for="dr">Dokter</label>
                </div>
                <slot></slot>
              </div>
              <InputData :props="{ id: 'dr', name: 'Dokter' }">
                <select id="dr" v-model="examinationAdd.doctor_id">
                  <option v-for="doctor in doctors" :key="doctor.id" :value="doctor.id">{{ doctor.name }}</option>
                </select>
              </InputData>
              <button>Save</button>
            </div>
          </form>
        </div>

        <h4 style="margin-bottom: 1rem; margin-top: 1rem;">Tindakan</h4>
        <hr>

        <div style="width: 100%; height: 15rem; margin-top: 1rem" class="scroll">
          <table class="table-custom" style="width: 100%">
            <thead>
              <tr>
                <td>Action</td>
                <td>Nama</td>
                <td>Dokter</td>
                <td>Perawat</td>
                <td>Tanggal</td>
                <td>Biaya</td>
              </tr>
            </thead>
            <tbody>
              <tr v-for="data in examinationData" :key="data.id">
                <td>
                  <button class="button-action" @click="handleDeleteExaminationData(data.id)">Delete</button>
                </td>
                <td>{{ data.examination }}</td>
                <td>{{ data.doctor_name }}</td>
                <td>{{ data.nurse_name }}</td>
                <td>{{ viewedDateTime(data.date) }}</td>
                <td>{{ data.total }}</td>
              </tr>
            </tbody>
          </table>
        </div>

        <div style="padding: 0.5rem; padding-top: 2rem; padding-bottom: 2rem;">
          <form class="form-search-input" v-on:submit.prevent="handleGetExamination">
            <div class="box-search">
              <div class="box">
                <input type="number" v-model="search.limit" id="lm" placeholder="limit">
                <input type="text" id="sc" v-model="search.search" placeholder="Tindakan">
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

        <hr>

        <div style="width: 100%; height: 15rem; margin-top: 1rem" class="scroll">
          <table class="table-custom" style="width: 100%;">
            <thead>
              <tr>
                <td>Action</td>
                <td>Nama pemeriksaan</td>
                <td>Cara bayar</td>
                <td>Biaya dokter</td>
                <td>Biaya perawat</td>
                <td>Biaya management</td>
              </tr>
            </thead>
            <tbody>
              <tr v-for="data in examData" :key="data.id">
                <td>
                  <button class="button-action" @click="handleAddExamination(data.id)">Add</button>
                </td>
                <td>{{ data.examination_name }}</td>
                <td>{{ data.payment_method }}</td>
                <td>{{ data.doctor_cost }}</td>
                <td>{{ data.nurse_cost }}</td>
                <td>{{ data.management_cost }}</td>
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
  border: none;
  padding: 0.5rem;
  border-radius: 1rem;
  color: var(--font-color);
}
</style>
