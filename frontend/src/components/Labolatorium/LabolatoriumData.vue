<script setup lang="ts">
import { createLabolatoriumData, createLabolatoriumTemplateData, deleteLabolatoriumData, deleteLabolatoriumTemplate, getLabolatoriumData, getLabolatoriumTemplate, updateLabolatoriumData } from '@/lib/api/labolatorium';
import type { LabolatoriumCreate, LabolatoriumDatas, LabolatoriumTemplate, LabolatoriumTemplateCreate } from '@/types/labolatorium';
import { nextTick, onBeforeMount, reactive, ref } from 'vue';

// Define variabel
const pageScroll = ref<HTMLElement | null>()
const labolatoriumData = ref<LabolatoriumDatas[]>([])
const labolatoriumTemplate = ref<LabolatoriumTemplate[]>([])
const open = ref<boolean>(false)
const dropOpen = ref<boolean[]>([])
const dropOpenTemplate = ref<boolean[]>([])
const idNotEdit = ref<string>("")
const idForTemplate = ref<string>("")
const createLabolatorium = reactive<LabolatoriumCreate>({
    id: "",
    name: "",
    referral_fee: 0,
    officer_fee: 0,
    management: 0,
})
const createTemplate = reactive<LabolatoriumTemplateCreate>({
    id: "",
    value: {
      id: "",
      name: "",
      unit: "",
      normal_value: "",
      template_id: ''
    }
})

// Define function
function resetForm() {
    open.value = false

    createLabolatorium.id = ""
    createLabolatorium.name = ""
    createLabolatorium.referral_fee = 0
    createLabolatorium.officer_fee = 0
    createLabolatorium.management = 0
}

function editLab(index: number,data: LabolatoriumDatas) {
    open.value = true
    dropOpen.value[index] = false
    labolatoriumTemplate.value = []

    createLabolatorium.id = data.id
    createLabolatorium.name = data.name
    createLabolatorium.referral_fee = data.referral_fee
    createLabolatorium.officer_fee = data.officer_fee
    createLabolatorium.management = data.management
}

// Handler function
async function handleCreateLabolatoriumData() {
  try {
    const response = await createLabolatoriumData(createLabolatorium)
    const json = await response.json()

    if (response.status === 201) {
        alert("Berhasil")
        await handleGetLabolatoriumData()
        resetForm()
    } else {
        alert(json.errors)
    }
  }catch(e){
    alert(`Error : ${e}`)
  }
}

async function handleGetLabolatoriumData() {
  try {
    const response = await getLabolatoriumData(5, "")
    const json = await response.json()

    if (response.status === 200) {
        labolatoriumData.value = json
        resetForm()
    } else {
        alert(json.errors)
    }
  }catch(e){
    alert(`Error : ${e}`)
  }
}

async function handleGetLabolatoriumTemplate(id: string) {
  try {
    const response = await getLabolatoriumTemplate(id)
    const json = await response.json()

    if (response.status === 200) {
        labolatoriumTemplate.value = json
        resetForm()
    } else {
        alert(json.errors)
    }
  }catch(e){
    alert(`Error : ${e}`)
  }
}

async function handleDeleteLabData(index: number, id: string) {
    dropOpen.value[index] = false
    labolatoriumTemplate.value = []

  try {
    const response = await deleteLabolatoriumData(id)
    const json = await response.json()

    if (response.status === 200) {
        alert("Berhasil")
        await handleGetLabolatoriumData()
    } else {
        alert(json.errors)
    }
  }catch(e){
    alert(`Error : ${e}`)
  }
}

async function handleCreateLabolatoriumTemplate() {
  const data: LabolatoriumTemplateCreate = {...createTemplate, id: idForTemplate.value}

  try {
    const response = await createLabolatoriumTemplateData(data)
    const json = await response.json()

    if (response.status === 201) {
        alert("Berhasil")
        await handleGetLabolatoriumTemplate(idForTemplate.value)
    } else {
        alert(json.errors)
    }
  }catch(e){
    alert(`Error : ${e}`)
  }
}

async function handleUpdateLabData() {
  const dataToUpdate = {
    id: createLabolatorium.id,
    name: createLabolatorium.name,
    referral_fee: createLabolatorium.referral_fee,
    officer_fee: createLabolatorium.officer_fee,
    management: createLabolatorium.management,
  }

  try {
    const response = await updateLabolatoriumData(idNotEdit.value, dataToUpdate)
    const json = await response.json()

    if (response.status === 200) {
        alert("Berhasil")
        await handleGetLabolatoriumData()
        resetForm()
    } else {
        alert(json.errors)
    }
  }catch(e){
    alert(`Error : ${e}`)
  }
}

async function handleDeleteTemplate(data: LabolatoriumTemplate) {
  try {
    const response = await deleteLabolatoriumTemplate(data.id, data.unit)
    const json = await response.json()

    if (response.status === 200) {
        alert("Berhasil")
        await handleGetLabolatoriumTemplate(data.id)
    } else {
        alert(json.errors)
    }
  }catch(e){
    alert(`Error : ${e}`)
  }
}

// function handleAddTemplateData() {
//     template.value.push(createTemplate)
// }

async function handleAddTemplate(index: number, id: string) {
  dropOpen.value[index] = false
  createTemplate.id = id
  await handleGetLabolatoriumTemplate(id)

  nextTick(() => {
    pageScroll.value?.scrollIntoView({behavior:  "smooth"})
  })
}

// on mount
onBeforeMount(async () => {
  await handleGetLabolatoriumData()
})
</script>

<template>
  <section class="anim-slide" ref="pages">
    <h3 style="margin: 0.5rem;">Data Labolatorium</h3>
    <div style="padding-top: 2rem; padding-bottom: 2rem;" class="bottom-line">
      <form class="form-data-custom" v-on:submit.prevent="handleCreateLabolatoriumData">
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Buat data pemeriksaan labolatorium</h4>
        <div style="display: grid; grid-template-columns: auto auto auto; padding-left: 1rem;">
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="id">ID pemeriksaan</label>
            </div>
            <input type="text" id="id" v-model="createLabolatorium.id" placeholder="ID pemeriksaan">
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="nm">Nama pemeriksaan</label>
            </div>
            <input type="text" id="nm" v-model="createLabolatorium.name" placeholder="Nama pemeriksaan">
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="prd">Biaya perujuk</label>
            </div>
            <input type="number" id="prd" v-model="createLabolatorium.referral_fee" placeholder="Biaya perujuk">
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="pr">Biaya petugas</label>
            </div>
            <input type="number" id="pr" v-model="createLabolatorium.officer_fee" placeholder="Biaya petugas">
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="ma">Manajement</label>
            </div>
            <input type="number" id="ma" v-model="createLabolatorium.management" placeholder="Biaya manajement">
          </div>
          <div>
            <button type="button" @click="resetForm">Reset</button>
            <button type="button" @click="handleUpdateLabData" v-if="open">Update</button>
            <button type="submit" v-if="!open">Create</button>
          </div>
        </div>
      </form>
    </div>

    <div style="width: 100%; overflow-x: scroll; overflow-y: scroll; height: 20rem; scrollbar-width: thin;" class="bottom-line">
      <table class="table-custom">
        <thead>
          <tr>
            <td>No</td>
            <td>ID</td>
            <td>Nama Pemeriksaan</td>
            <td>Tarif perujuk</td>
            <td>Tarif petugas</td>
            <td>Management</td>
            <td>Total</td>
            <td>Action</td>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(data, index) in labolatoriumData" :key="data.id">
            <td>{{ index + 1 }}</td>
            <td>{{ data.id }}</td>
            <td>{{ data.name }}</td>
            <td>{{ data.referral_fee }}</td>
            <td>{{ data.officer_fee }}</td>
            <td>{{ data.management }}</td>
            <td>{{ data.total }}</td>
            <td style="position: relative;">
                <div class="drop-down">
                    <button class="act" @click="dropOpen[index] = !dropOpen[index]">:</button>
                    <div class="menu" v-if="dropOpen[index]">
                        <ul>
                            <li>
                                <button @click="editLab(index, data); idNotEdit = data.id">Edit</button>
                            </li>
                            <li>
                                <button @click="handleDeleteLabData(index, data.id)">Delete</button>
                            </li>
                            <li>
                                <button @click="handleAddTemplate(index, data.id); idForTemplate = data.id">Template</button>
                            </li>
                        </ul>
                    </div>
                </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div ref="pageScroll" v-if="labolatoriumTemplate == null || labolatoriumTemplate.length > 0">
      <form class="form-data-custom" v-on:submit.prevent="handleCreateLabolatoriumTemplate">
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Template hasil</h4>
        <div class="center" style="justify-content: flex-start; align-items: flex-end; padding-left: 1rem;">
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="tn">Nama template pemeriksaan</label>
            </div>
            <input type="text" id="tn" v-model="createTemplate.value.name" placeholder="Nama pemeriksaan">
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="un">Satuan</label>
            </div>
            <input type="text" id="un" v-model="createTemplate.value.unit" placeholder="Satuan">
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="nv">Nilai normal</label>
            </div>
            <input type="text" id="nv" v-model="createTemplate.value.normal_value" placeholder="Nilai normal">
          </div>
          <button type="submit">Create</button>
        </div>
      </form>
    </div>

    <div style="width: 100%; overflow-x: scroll; overflow-y: scroll; height: 20rem; scrollbar-width: thin;" v-if="labolatoriumTemplate == null || labolatoriumTemplate.length > 0">
      <table class="table-custom">
        <thead>
          <tr>
            <td>Nama template</td>
            <td>Satuan</td>
            <td>Nilai normal</td>
            <td>Action</td>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(data, index) in labolatoriumTemplate" :key="index">
            <td>{{ data.name }}</td>
            <td>{{ data.unit }}</td>
            <td>{{ data.normal_value }}</td>
            <td style="position: relative;">
                <div class="drop-down">
                    <button class="act" @click="dropOpenTemplate[index] = !dropOpenTemplate[index]">:</button>
                    <div class="menu" v-if="dropOpenTemplate[index]">
                        <ul>
                            <li>
                                <button @click="handleDeleteTemplate(data)">Delete</button>
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
