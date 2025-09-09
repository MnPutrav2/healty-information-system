<script setup lang="ts">
import { getLabolatoriumDetailData, getLabolatoriumRequest, updateLabolatoriumSample, updateLabolatoriumValue } from '@/lib/api/labolatorium'
import { nextTick, onBeforeMount, ref } from 'vue'
import type { LabolatoriumRequestData,  LabReq, LabolatoriumValue, LabUpdateSample, LabolatoriumDetailData } from '@/types/labolatorium'
import { formatDatetime, viewedDateTime } from '@/lib/formatDate'

// Define variabel
const date = new Date()
const date1 = ref<string>(formatDatetime(date, "00:00:00"))
const date2 = ref<string>(formatDatetime(date, "23:59:00"))
const labDatas = ref<LabolatoriumRequestData[]>([])
const dropOpen = ref<boolean[]>([])
const bool = ref<boolean>(false)
const sampleDate = ref<string>("")
const scrollPage = ref<HTMLElement | null>(null)
const scrollPageUp = ref<HTMLElement | null>(null)
const open = ref<boolean>(false)
const detailData = ref<LabolatoriumDetailData[]>([])
let inter = setInterval(updateTime, 1000)
const id = ref<string>("")
const templ = ref<LabReq[]>([])

// Define function
function updateTime() {
  sampleDate.value = formatDatetime(new Date(), null)
}

function addValue(lab: string, index: string, event: Event) {
  const target = event.target as HTMLInputElement | null

  templ.value.forEach((item) => {
    if(item.template_id == index) {
      item.id = lab
      item.value = target == null ? "" : target.value.toString()
    }
  })
}

async function scroll(index: number, ex: string) {
  open.value = true
  dropOpen.value[index] = false
  id.value = ex

  await handleGetLabolatoriumDetailData(ex)

  nextTick(() => {
    scrollPage.value?.scrollIntoView({behavior: 'smooth'})
  })
}

function autoDate(bool: boolean) {
  if (bool) {
    clearInterval(inter)
  } else {
    inter = setInterval(updateTime, 1000)
  }
}

async function handleGetLabolatoriumRequestData() {
  const response = await getLabolatoriumRequest(date1.value, date2.value)
  const json = await response.json()

  try {
    if (response.status === 200) {
      labDatas.value = json
    } else {
      alert(json.errors)
    }
  }catch(e){
    alert(`Error : ${e}`)
  }
}

async function handleUpdateLabolatoriumVal() {
  const dt: LabolatoriumValue = {
    date: sampleDate.value,
    value: templ.value
  }

  try {
    const response = await updateLabolatoriumValue(dt, id.value)
    const json = await response.json()

    if (response.status === 200) {
      alert("sukses")
      open.value = false
      await handleGetLabolatoriumRequestData()

      nextTick(() => {
        scrollPageUp.value?.scrollIntoView({behavior: 'smooth'})
      })
    } else {
      alert(json.errors)
    }
  }catch(e){
    alert(`Error : ${e}`)
  }
}

async function handleGetLabolatoriumDetailData(id: string) {
  const response = await getLabolatoriumDetailData(id)
  const json = await response.json()
  const dataz = ref<LabolatoriumDetailData[]>([])

  try {
    if (response.status === 200) {
      detailData.value = json
      dataz.value = json
    } else {
      alert(json.errors)
    }

    dataz.value.forEach((item) => {
      item.template.forEach((dt) => {
        templ.value.push({template_id: dt.template_id, id: "", value: ""})
      })
    })
  }catch(e){
    alert(`Error : ${e}`)
  }
}

async function handleUpdateLabolatoriumSample(id: string) {
  const data = ref<LabUpdateSample>({status: true, value: sampleDate.value})
  const response = await updateLabolatoriumSample(data.value, id)
  const json = await response.json()

  try {
    if (response.status === 200) {
      alert("sample")
      await handleGetLabolatoriumRequestData()
    } else {
      alert(json.errors)
    }
  }catch(e){
    alert(`Error : ${e}`)
  }
}

// on before
onBeforeMount(async () => {
  await handleGetLabolatoriumRequestData()
})
</script>

<template>
  <section class="anim-slide" ref="scrollPageUp">
    <h3 style="margin: 0.5rem;">Permintaan Labolatorium</h3>
    <div style="padding-top: 2rem; padding-bottom: 2rem;" class="bottom-line">
      <form class="form-data-custom" v-on:submit.prevent="handleGetLabolatoriumRequestData">
        <div class="center" style="justify-content: flex-start; align-items: flex-end; padding-left: 1rem;">
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="dt1">Tanggal awal</label>
            </div>
            <input type="datetime-local" step="1" id="dt1" v-model="date1" placeholder="tanggal">
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="dt2">Tanggal akhir</label>
            </div>
            <input type="datetime-local" step="1" id="dt2" v-model="date2" placeholder="tanggal">
          </div>
          <button type="submit">Find</button>
        </div>
      </form>
      <form class="form-data-custom">
        <div class="center" style="justify-content: flex-start; align-items: flex-end; padding-left: 1rem;">
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="dt1">Sample</label>
            </div>
             <div class="center">
              <div class="center">
                <input type="datetime-local" step="1" id="date" v-model="sampleDate" placeholder="tanggal">
                <div :class="bool ? 'clock-inactive' : 'clock-active'" class="button-clock" @click="autoDate(bool = !bool)" role="button" tabindex="0"></div>
              </div>
            </div>
          </div>
        </div>
      </form>
    </div>

    <div style="width: 100%; overflow-x: scroll; overflow-y: scroll; height: 20rem; scrollbar-width: thin;">
      <table class="table-custom">
        <thead>
          <tr>
            <td>No</td>
            <td>Nomor permintaan</td>
            <td>Nama pasien</td>
            <td>Perujuk</td>
            <td>Sampel</td>
            <td>Hasil</td>
            <td>Action</td>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(lab, index) in labDatas" :key="index">
            <td>{{ index + 1 }}</td>
            <td>{{ lab.lab_id }}</td>
            <td>{{ lab.patient_name }}</td>
            <td>{{ lab.referral_name }}</td>
            <td>{{ lab.sample == null ? lab.sample : viewedDateTime(lab.sample) }}</td>
            <td>{{ lab.validate_date == null ? lab.validate_date : viewedDateTime(lab.validate_date) }}</td>
            <td style="position: relative;">
              <div class="drop-down">
                <button class="act" @click="dropOpen[index] = !dropOpen[index]">:</button>
                <div class="menu" v-if="dropOpen[index]">
                  <ul>
                    <li>
                        <button @click="handleUpdateLabolatoriumSample(lab.lab_id); dropOpen[index] = false" :disabled="lab.sample != null">Sample</button>
                    </li>
                  </ul>
                  <ul>
                    <li>
                        <button @click="scroll(index, lab.lab_id)" :disabled="lab.sample == null || lab.validate_date != null">Hasil</button>
                    </li>
                  </ul>
                </div>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div style="padding-top: 2rem; padding-bottom: 2rem;" ref="scrollPage" v-if="open">
      <form class="form-data-custom" v-on:submit.prevent="handleUpdateLabolatoriumVal">
        <div class="center" style="justify-content: flex-start; align-items: flex-end; padding-left: 1rem;">
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="dt1">Hasil</label>
            </div>
             <div class="center">
              <div class="center">
                <input type="datetime-local" step="1" id="date" v-model="sampleDate" placeholder="tanggal">
                <div :class="bool ? 'clock-inactive' : 'clock-active'" class="button-clock" @click="autoDate(bool = !bool)" role="button" tabindex="0"></div>
              </div>
            </div>
          </div>
          <button type="submit">Save</button>
        </div>
      </form>
    </div>

    <div style="width: 100%; overflow-x: scroll; overflow-y: scroll; height: 20rem; scrollbar-width: thin;" v-if="open">
      <table class="table-custom">
        <thead>
          <tr>
            <td>No</td>
            <td>Nama Pemeriksaan</td>
            <td>Template</td>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(lab, index) in detailData" :key="lab.id">
            <td>{{ index + 1 }}</td>
            <td>{{ lab.name }}</td>
            <td>
              <table>
                <thead>
                  <tr>
                    <td>Input</td>
                    <td>Nama Template</td>
                    <td>Satuan</td>
                    <td>Nilai Normal</td>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="(temp, index) in lab.template" :key="index">
                    <td>
                      <input class="normal-input" @input="addValue(lab.id, temp.template_id, $event)" type="text">
                    </td>
                    <td>{{ temp.name }}</td>
                    <td>{{ temp.unit }}</td>
                    <td>{{ temp.normal_value }}</td>
                  </tr>
                </tbody>
              </table>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </section>
</template>
