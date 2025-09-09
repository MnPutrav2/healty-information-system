<script setup lang="ts">
import { deleteDrugRecipes, getDrugRecipe, getRecipeData, getDrugData, createRecipe, createRecipeCompound, validateRecipe, handoverRecipe } from '@/lib/api/pharmacy';
import { sum } from '@/lib/drugSum';
import { formatDatetime, viewedDateTime } from '@/lib/formatDate';
import type { RecipeType, RecipeData, Drug, DetailRecipe, RecipeForRequest, RecipesForRequest, RecipeCompoundForRequest, RecipeCompound, ValidateTypeRecipe, HandoverTypeRecipe } from '@/types/pharmacy';
import type { SearchLimit } from '@/types/response';
import { nextTick, onBeforeMount, reactive, ref } from 'vue';


// Define variabel
const date = new Date()
const firstDateSearch = ref<string>(formatDatetime(date, "00:00:00"))
const lastDateSearch = ref<string>(formatDatetime(date, "23:59:00"))
const realtimeDate = ref<string>(formatDatetime(date, null))
const recipeData = ref<RecipeData[]>([])
const drugRecipe = ref<RecipeType[]>([])
const pageScroll = ref<HTMLElement | null>()
const pageScrollTop = ref<HTMLElement | null>()
const recipeSelect = ref<string>("")
const careNumber = ref<string>("")
let inter = setInterval(updateTime, 1000)
const bool = ref<boolean>(false)
const open = ref<boolean[]>([])
const openDrugSearch = ref<boolean[]>([])
const addDrugName = reactive<Record<number, string>>({})
const addDrugID = reactive<Record<number, string>>({})
const addDrugUse = reactive<Record<number, string>>({})
const addDrugValue = reactive<Record<number, number>>({})
const addDrugCapacity = reactive<Record<number, number>>({})
const addDrugEmbalming = reactive<Record<number, number>>({})
const addDrugTuslah = reactive<Record<string, number>>({})
const addDrugPrice = reactive<Record<number, number>>({})
const drugSearch = reactive<Record<number, string>>({})
const drugSearchData = ref<Drug[]>([])
const menuDataRecipe = ref<boolean[]>([])
const handover = ref<HandoverTypeRecipe>({
  status: true,
  date: formatDatetime(date, null)
})

// Define function
function scroll() {
  pageScroll.value?.scrollIntoView({behavior: "smooth"})
}

function updateTime() {
  realtimeDate.value = formatDatetime(new Date(), null)
}

function autoDate(bool: boolean) {
  if (bool) {
    clearInterval(inter)
  } else {
    inter = setInterval(updateTime, 1000)
  }
}

function clickedDrugSearch(index: number, drug: Drug) {
  addDrugName[index] = drug.name
  addDrugID[index] = drug.id
  addDrugPrice[index] = drug.price
  addDrugCapacity[index] = drug.capacity

  drugSearch[index] = drug.name
}

function drugRec(capacity: number, fill: number, content: number): number {
  const data: number = content / capacity * fill

  return Math.round(data)
}

// Handler fuunction
function handleAddDrugRecipe(index: number, type: string, recipe: string, drugcom: DetailRecipe[]) {
  open.value[index] =! open.value[index]

  const name = addDrugName[index] || ""
  const id = addDrugID[index] || ""
  const value = addDrugValue[index] || 0
  const use = addDrugUse[index] || ""
  const embalming = addDrugEmbalming[index] || 0
  const tuslah = addDrugTuslah[index] || 0
  const price = addDrugPrice[index] || 0
  const capacity = addDrugCapacity[index] || 0

  const recipesForRequest = ref<RecipesForRequest[]>([])

  if(type == "common") {
    const recipeRes: RecipesForRequest = {
      drug_id: id,
      value,
      use,
      embalming,
      tuslah,
      total_price: sum(value, price, embalming, tuslah),
    }

    recipesForRequest.value.push(recipeRes)

    const recipeRequest = reactive<RecipeForRequest>({
      care_number: careNumber.value,
      recipe_number: recipeSelect.value,
      date: formatDatetime(date, null),
      validate: formatDatetime(date, "00:00:00"),
      handover: formatDatetime(date, "00:00:00"),
      type: "add",
      drug: recipesForRequest.value
    })

    handleAddRecipe(recipeRequest)
  } else {
    const recipeCompounds = ref<RecipeCompound[]>([])

    const recipeRequest = reactive<RecipeCompoundForRequest>({
      care_number: careNumber.value,
      recipe_number: recipeSelect.value,
      date: formatDatetime(date, null),
      validate: formatDatetime(date, "00:00:00"),
      handover: formatDatetime(date, "00:00:00"),
      type: "add",
      recipes: recipeCompounds.value
    })

    const recipeCompound = ref<RecipeCompound>({
      recipe_name: recipe,
      value: drugcom[0].value,
      use: drugcom[0].use,
      drug: []
    })

    const val: number = drugRec(capacity, drugcom[0].compound_value, value)

    recipeCompound.value.drug.push({
      name: name,
      drug_id: id,
      value: val,
      embalming,
      tuslah,
      price: sum(val, price, embalming, tuslah),
    })

    recipeCompounds.value.push(recipeCompound.value)

    handleAddRecipeCompound(recipeRequest)
  }

  drugSearch[index] = ""
  addDrugName[index] = ""
  addDrugID[index] = ""
  addDrugValue[index] = 0
  addDrugUse[index] = ""
  addDrugEmbalming[index] = 0
  addDrugTuslah[index] = 0
  addDrugPrice[index] = 0
  addDrugCapacity[index] = 0
}

async function handleAddRecipe(recipe: RecipeForRequest) {
  const response = await createRecipe(recipe)
  const json = await response.json()

  try {

    if (response.status === 201) {
      alert('Resep berhasil dibuat!')
      handleGetDrugRecipe(recipeSelect.value)
    }else{
      alert(json.errors)
    }

  } catch(error) {
    console.log(error)
  }
}

async function handleAddRecipeCompound(recipe: RecipeCompoundForRequest) {
  const response = await createRecipeCompound(recipe)
  const json = await response.json()

  try {

    if (response.status === 201) {
      alert("Berhasil membuat resep!")
      handleGetDrugRecipe(recipeSelect.value)
    }else{
      alert(json.errors)
    }

  } catch(error) {
    console.log(error)
  }
}

async function handleGetRecipeData() {
  const response = await getRecipeData(firstDateSearch.value, lastDateSearch.value)
  const json = await response.json()

  try {

    if (response.status === 200) {
      recipeData.value = json
    }else{
      alert(json.errors)
    }

  } catch(error) {
    console.log(error)
  }
}

async function handleGetDrugRecipe(recipe: string) {
  const response = await getDrugRecipe(recipe)
  const json = await response.json()

  try {

    if (response.status === 200) {
      drugRecipe.value = json
      open.value = []
    }else{
      alert(json.errors)
    }

  } catch(error) {
    console.log(error)
  }

  nextTick(() => {
    scroll()
  })
}

async function handleValidateRecipe() {
  pageScrollTop.value?.scrollIntoView({behavior: "smooth"})

  const val = ref<ValidateTypeRecipe>({
    validate_status: true,
    validate_date: realtimeDate.value
  })

  const response = await validateRecipe(recipeSelect.value, val.value)
  const json = await response.json()

  try {

    if (response.status === 200) {
      alert('Validasi berhasil')
      await handleGetRecipeData()
      drugRecipe.value = []
    }else{
      alert(json.errors)
    }

  } catch(error) {
    console.log(error)
  }
}

async function handleDeleteDrugRecipe(recipe: string, drugID: string, comname: string) {
  const response = await deleteDrugRecipes(recipe, drugID, comname)
  const json = await response.json()

  try {

    if (response.status === 200) {
      alert('Berhasil hapus data')
      await handleGetRecipeData()
      await handleGetDrugRecipe(recipeSelect.value)
    }else{
      alert(json.errors)
    }

  } catch(error) {
    console.log(error)
  }

  nextTick(() => {
    scroll()
  })
}

async function handleGetDrugData(index: number) {
  const search = reactive<SearchLimit>({
    search: drugSearch[index] || "",
    limit: 10
  })
  openDrugSearch.value[index] = true

  const response = await getDrugData(search)
  const json = await response.json()

  try {

    if (response.status === 200) {
      drugSearchData.value = json
    }else{
      alert(json.errors)
    }

  } catch(error) {
    console.log(error)
  }
}

async function handleRecipeHandover(recipe: string) {
  const hand: HandoverTypeRecipe = {...handover.value, date: realtimeDate.value}

  const response = await handoverRecipe(recipe, hand)
  const json = await response.json()

  try {

    if (response.status === 200) {
      alert("penyerahan resep")
      await handleGetRecipeData()
    }else{
      alert(json.errors)
    }

  } catch(error) {
    console.log(error)
  }
}

// Before page view
onBeforeMount(async () => {
  await handleGetRecipeData()
})
</script>

<template>
  <section class="anim-slide" ref="pageScrollTop">
    <h3 style="margin: 0.5rem;">Resep masuk</h3>
    <div style="padding-top: 2rem; padding-bottom: 2rem;" class="bottom-line">
      <form class="form-data-custom" v-on:submit.prevent="handleGetRecipeData">
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Cari resep</h4>
        <div class="center" style="justify-content: flex-start; align-items: flex-end; padding-left: 1rem;">
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="dt1">Tanggal awal</label>
            </div>
            <input type="datetime-local" step="1" id="dt1" v-model="firstDateSearch" placeholder="tanggal">
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="dt2">Tanggal akhir</label>
            </div>
            <input type="datetime-local" step="1" id="dt2" v-model="lastDateSearch" placeholder="tanggal">
          </div>
          <button type="submit">Cari</button>
        </div>
      </form>
    </div>

    <div style="width: 100%; overflow-x: scroll; overflow-y: scroll; height: 30rem; scrollbar-width: thin;">
      <table class="table-custom">
        <thead>
          <tr>
            <td>No resep</td>
            <td>No rawat</td>
            <td>Validasi</td>
            <td>Nama</td>
            <td>Tanggal resep</td>
            <td>Tanggal validasi</td>
            <td>Tanggal penyerahan</td>
            <td>Action</td>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(rec, index) in recipeData" :key="rec.recipe_id">
            <td>{{ rec.recipe_id }}</td>
            <td>{{ rec.care_number }}</td>
            <td :class="rec.validate_status ? 'validate' : 'invalidate'">{{ rec.validate_status ? 'Sudah' : 'Belum' }}</td>
            <td>{{ rec.name }}</td>
            <td>{{ viewedDateTime(rec.date) }}</td>
            <td>{{ viewedDateTime(rec.validate) }}</td>
            <td>{{ viewedDateTime(rec.handover) }}</td>
            <td style="position: relative;">
                <div class="drop-down">
                    <button class="act" @click="menuDataRecipe[index] = !menuDataRecipe[index]">:</button>
                    <div class="menu" v-if="menuDataRecipe[index]">
                        <ul>
                            <li>
                              <button class="button-action" @click="handleGetDrugRecipe(rec.recipe_id); recipeSelect = rec.recipe_id; careNumber = rec.care_number; menuDataRecipe[index] = false" :disabled="rec.validate_status">Validasi</button>
                            </li>
                            <li>
                              <button class="button-action" @click="handleRecipeHandover(rec.recipe_id); menuDataRecipe[index] = false">Penyerahan</button>
                            </li>
                        </ul>
                    </div>
                </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div ref="pageScroll" v-if="drugRecipe && drugRecipe.length > 0">
      <hr>
      <div style="padding-top: 2rem; padding-bottom: 2rem;">
        <form class="form-data-custom" v-on:submit.prevent="handleValidateRecipe">
          <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Validasi resep</h4>
          <div class="center" style="justify-content: flex-start; align-items: flex-end; padding-left: 1rem;">
            <div style="padding: 0.5rem;">
              <div style="margin-bottom: 0.5rem;">
                <label for="dt1">Tanggal validasi</label>
              </div>
              <div class="center">
                <input type="datetime-local" step="1" id="dt1" v-model="realtimeDate" placeholder="tanggal">
                <div :class="bool ? 'clock-inactive' : 'clock-active'" class="button-clock" @click="autoDate(bool = !bool)" role="button" tabindex="0"></div>
              </div>
            </div>
            <button type="submit">Validasi</button>
          </div>
        </form>
      </div>

      <div style="width: 100%; overflow-x: scroll; overflow-y: scroll; height: 30rem; scrollbar-width: thin;">
        <table class="table-custom">
          <thead>
            <tr>
              <td>Jenis resep</td>
              <td>Data</td>
            </tr>
          </thead>
          <tbody>
            <template v-for="(drug, index) in drugRecipe" :key="index">
              <tr>
                <td colspan="1">{{ drug.compound_name == 'com' ? 'Biasa' : `Racikan : ${drug.compound_name}` }}</td>
                <td colspan="5">
                  <button @click="open[index] = !open[index]" v-if="!open[index]" class="button-action">Tambah obat</button>
                  <form class="form-data-custom" v-on:submit.prevent="handleAddDrugRecipe(index, drug.recipe_type, drug.compound_name, drug.data)" v-if="open[index]">
                    <div class="center" style="justify-content: flex-start; align-items: flex-end; padding-left: 1rem;">
                      <button type="submit">Add</button>
                      <div class="input-field" style="position: relative;">
                        <div>
                          <div style="margin-bottom: 0.5rem;">
                            <label for="nmo">Nama obat</label>
                          </div>
                          <input type="text" v-model="drugSearch[index]" @change="handleGetDrugData(index)" id="nmo">
                          <div class="find-input" v-if="drugSearchData && drugSearchData.length > 0 && openDrugSearch[index]">
                            <button type="button" @click="openDrugSearch[index] = false">close</button>
                            <ul>
                              <li v-for="drug in drugSearchData" :key="drug.id" @click="clickedDrugSearch(index, drug)">{{ drug.name }}</li>
                            </ul>
                          </div>
                        </div>
                      </div>
                      <div style="padding: 0.5rem;" v-if="drug.recipe_type != 'compound'">
                        <div style="margin-bottom: 0.5rem;">
                          <label for="atp">Aturan pakai</label>
                        </div>
                        <input type="text" id="atp" v-model="addDrugUse[index]" placeholder="Aturan pakai" required>
                      </div>
                      <div style="padding: 0.5rem;">
                        <div style="margin-bottom: 0.5rem;">
                          <label for="jml">{{ drug.recipe_type == 'compound' ? 'Kandungan' : 'Jumlah' }}</label>
                        </div>
                        <input type="number" id="jml" v-model="addDrugValue[index]" :placeholder="drug.recipe_type == 'compound' ? 'Kandungan' : 'jumlah'" required>
                      </div>
                      <div style="padding: 0.5rem;">
                        <div style="margin-bottom: 0.5rem;">
                          <label for="emb">Embalase</label>
                        </div>
                        <input type="number" id="emb" v-model="addDrugEmbalming[index]" placeholder="embalase">
                      </div>
                      <div style="padding: 0.5rem;">
                        <div style="margin-bottom: 0.5rem;">
                          <label for="tsl">Tuslah</label>
                        </div>
                        <input type="number" id="tsl" v-model="addDrugTuslah[index]" placeholder="tuslah">
                      </div>
                    </div>
                  </form>
                </td>
              </tr>
              <tr>
                <td>
                </td>
                <td>
                  <table>
                    <thead>
                      <tr>
                        <td>Action</td>
                        <td>Nama obat</td>
                        <td>Jumlah</td>
                        <td>Aturan pakai</td>
                        <td>Embalase</td>
                        <td>Tuslah</td>
                        <td>Total</td>
                      </tr>
                    </thead>
                    <tbody>
                      <tr v-for="dt in drug.data" :key="dt.drug_id">
                        <td>
                          <button class="button-action" @click="handleDeleteDrugRecipe(dt.recipe_id, dt.drug_id, dt.compound_name)">Delete</button>
                        </td>
                        <td>{{ dt.drug_name }}</td>
                        <td>{{ dt.value }}</td>
                        <td>{{ dt.use }}</td>
                        <td>{{ dt.embalming }}</td>
                        <td>{{ dt.tuslah }}</td>
                        <td>{{ dt.total_price }}</td>
                      </tr>
                    </tbody>
                  </table>
                </td>
              </tr>
            </template>
          </tbody>
        </table>
      </div>
    </div>
  </section>
</template>
