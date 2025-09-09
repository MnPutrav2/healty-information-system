<script setup lang="ts">
import { createDrugData, deleteDrugData, getDistributor, getDrugData, updateDrugData } from '@/lib/api/pharmacy';
import { onBeforeMount, reactive, ref } from 'vue';
import type { DrugData, Distributor, ResponseDrugData, RequestBodyDrugDataUpdate } from '@/types/pharmacy';
import { viewedDateTime } from '@/lib/formatDate';
import type { SearchLimit } from '@/types/response';

// Define variabels
const search = reactive<SearchLimit>({
  search: "",
  limit: 5
})
const drugDatas = ref<ResponseDrugData[]>([])
const pages = ref<HTMLElement | null>()
const bool = ref<boolean>(false)
const distributors = ref<Distributor[]>([])
const menuDataDrug = ref<boolean[]>([])
const drug = reactive<DrugData>({
  id: "",
  name: "",
  distributor: "",
  capacity: 0,
  fill: 0,
  unit: "",
  composition: "",
  category: "",
  price: 0,
  expired_date: "",
})

// Define functions
// Handler functions
async function handleCreateDrugData() {
  const response = await createDrugData(drug)
  const json = await response.json()

  try {

    if (response.status === 201) {
      await handleGetDrugData()
      alert("simpan berhasil")
    }else{
      alert(json.errors)
    }

  } catch(error) {
    console.log(error)
  }
}

async function handleGetDrugData() {
  const response = await getDrugData(search)
  const json = await response.json()

  try {

    if (response.status === 200) {
      drugDatas.value = json
    }else{
      alert(json.errors)
    }

  } catch(error) {
    console.log(error)
  }
}

function handleEditDrugData(data: ResponseDrugData) {
  pages.value?.scrollIntoView({behavior: "smooth"})

  Object.assign(drug);
  drug.distributor = data.distributor_id
}

async function handleUpdateDrugData(id: string) {
  const dt = ref<RequestBodyDrugDataUpdate>({
    id: id,
    data: drug
  })

  const response = await updateDrugData(dt.value)
  const json = await response.json()

  try {

    if (response.status === 200) {
      alert("berhasil update data")
      await handleGetDrugData()
      bool.value = false
    }else{
      alert(json.errors)
    }

  } catch(error) {
    console.log(error)
  }
}

async function handleDeleteDrugData(id: string) {
  const response = await deleteDrugData(id)
  const json = await response.json()

  try {

    if (response.status === 200) {
      alert("berhasil hapus data")
      await handleGetDrugData()
    }else{
      alert(json.errors)
    }

  } catch(error) {
    console.log(error)
  }
}

async function handleGetDistributor() {
  const response = await getDistributor()
  const json = await response.json()

  try {

    if (response.status === 200) {
      distributors.value = json
    }else{
      alert(json.errors)
    }

  } catch(error) {
    console.log(error)
  }
}

// Before view page
onBeforeMount(async ()=> {
  await handleGetDrugData()
  await handleGetDistributor()
})
</script>

<template>
  <section class="anim-slide" ref="pages">
    <h3 style="margin: 0.5rem;">Data Obat</h3>
    <div style="margin-top: 2rem; margin-bottom: 2rem;" class="bottom-line">
      <form class="form-data-custom" v-on:submit.prevent="handleCreateDrugData">

        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Input Data Obat</h4>
        <div style="display: grid; grid-template-columns: auto auto auto; padding-left: 1rem;">
          <div class="input-field">
            <div class="cover">
              <label for="id">Kode Obat</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <input type="text" id="id" placeholder="Kode obat" v-model="drug.id" maxlength="6" required>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="nm">Nama obat</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <input type="text" id="nm" placeholder="Nama obat" v-model="drug.name" maxlength="20" required>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="distributor">Distributor</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <select id="distributor" placeholder="Distributor" v-model="drug.distributor">
              <option v-for="dis in distributors" :key="dis.id" :value="dis.id">{{ dis.name }}</option>
            </select>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="cp">Kapasitas</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <input type="number" id="cp" placeholder="Kapasitas" v-model="drug.capacity" required>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="unit">Unit</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <select id="unit" v-model="drug.unit">
              <option value="mg">MG</option>
              <option value="ml">ML</option>
            </select>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="com">Komposisi</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <input type="text" id="com" placeholder="komposisi" v-model="drug.composition" required>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="fill">Isi/Box</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <input type="number" id="fill" placeholder="isi" v-model="drug.fill" required>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="ucategorynit">Kategori</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <select id="ucategorynit" v-model="drug.category">
              <option value="Obat">Obat</option>
              <option value="BHP">BHP</option>
            </select>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="price">Harga</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <input type="number" id="price" placeholder="harga" v-model="drug.price" required>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="exp">Tanggal kadaluarsa</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <input type="date" id="exp" placeholder="harga" v-model="drug.expired_date" required>
          </div>
        </div>

        <div>
          <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Action</h4>
          <button type="submit">Save</button>
          <button type="button" v-if="bool" @click="handleUpdateDrugData(drug.id)">Update</button>
        </div>

      </form>
    </div>

    <div style="padding-bottom: 2rem;">
      <form class="form-data-custom" v-on:submit.prevent="handleGetDrugData">
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Data Obat</h4>
        <div class="center" style="justify-content: flex-start; align-items: flex-end; padding-left: 1rem;">
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="sc">Cari</label>
            </div>
            <input type="text" id="sc" v-model="search.search" placeholder="Nama obat">
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="lm">Limit</label>
            </div>
            <input type="number" id="lm" v-model="search.limit" placeholder="limit">
          </div>
          <button type="submit">Cari</button>
        </div>
      </form>
    </div>

    <div style="width: 100%; overflow-x: scroll; overflow-y: scroll; height: 20rem; scrollbar-width: thin;">
      <table class="table-custom">
        <thead>
          <tr>
            <td>No</td>
            <td>Kode</td>
            <td>Nama</td>
            <td>Komposisi</td>
            <td>Distributor</td>
            <td>kapasitas</td>
            <td>Isi/box</td>
            <td>Unit</td>
            <td>Kategory</td>
            <td>Harga per tablet</td>
            <td>Tanggal Kadaluarsa</td>
            <td>Action</td>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(drug, index) in drugDatas" :key="drug.id">
            <td>{{ index + 1 }}</td>
            <td>{{ drug.id }}</td>
            <td>{{ drug.name }}</td>
            <td>{{ drug.composition }}</td>
            <td>{{ drug.distributor }}</td>
            <td>{{ drug.capacity }}</td>
            <td>{{ drug.fill }}</td>
            <td>{{ drug.unit }}</td>
            <td>{{ drug.category }}</td>
            <td>{{ drug.price }}</td>
            <td>{{ viewedDateTime(drug.expired_date) }}</td>
            <td style="position: relative;">
                <div class="drop-down">
                    <button class="act" @click="menuDataDrug[index] = !menuDataDrug[index]">:</button>
                    <div class="menu" v-if="menuDataDrug[index]">
                        <ul>
                            <li>
                              <button class="button-action" @click="handleEditDrugData(drug); bool = true; menuDataDrug[index] = false">Edit</button>
                            </li>
                            <li>
                                <button class="button-action" @click="handleDeleteDrugData(drug.id); menuDataDrug[index] = false">Delete</button>
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
