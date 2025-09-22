<script setup lang="ts">
import { createRecipe, getCurrentRecipeNumber, getDrugData, getRecipeNumber } from '@/lib/api/pharmacy';
import { onBeforeMount, reactive, ref, watch } from 'vue';
import { sum, loopSum } from '@/lib/drugSum';
import { formatDatetime } from '@/lib/formatDate';
import type { Recipe, RecipeForRequest, Drug, RecipesForRequest } from '@/types/pharmacy';
import { recipeNumber } from '@/lib/careNumber';
import type { SearchLimit } from '@/types/response';
import InputData from '../Extras/InputData.vue';
import { useRoute } from 'vue-router';

// Define variabels
const route = useRoute()
const recipes = ref<Recipe[]>([])
const recipesForRequest = ref<RecipesForRequest[]>([])
const recipeType = ref(false)
const drugs = ref<Drug[]>([])
const valMap = reactive<Record<string, number>>({})
const useMap = reactive<Record<string, string>>({})
const embalmingMap = reactive<Record<string, number>>({})
const tuslahMap = reactive<Record<string, number>>({})
const date = new Date()
const searchDrug = ref<string>("")
const recipeRequest = reactive<RecipeForRequest>({
  care_number: String(route.query.careNumber),
  recipe_number: "",
  date: formatDatetime(date, null),
  validate: formatDatetime(date, "00:00:00"),
  handover: formatDatetime(date, "00:00:00"),
  type: "create",
  drug: []
})

// Define functions
function removeRecipe(id: string) {
  recipes.value = recipes.value.filter(r => r.id !== id)
  recipesForRequest.value = recipesForRequest.value.filter(r => r.drug_id !== id)
}

// Handler functions
function handleAddDrug(data: Drug) {
  const value = valMap[data.id] || 0
  const use = useMap[data.id] || ''
  const embalming = embalmingMap[data.id] || 0
  const tuslah = tuslahMap[data.id] || 0

  const recipe: Recipe = {
    id: data.id,
    name: data.name,
    value,
    use,
    embalming,
    tuslah,
    price: data.price,
  }

  const recipeRes: RecipesForRequest = {
    drug_id: data.id,
    value,
    use,
    embalming,
    tuslah,
    total_price: sum(value, data.price, embalming, tuslah),
  }

  recipes.value.push(recipe)
  recipesForRequest.value.push(recipeRes)

  valMap[data.id] = 0
  useMap[data.id] = ''
  embalmingMap[data.id] = 0
  tuslahMap[data.id] = 0
}

async function handleCreateRecipe() {
  const res = {...recipeRequest, drug: recipesForRequest.value}
  const response = await createRecipe(res)
  const json = await response.json()

  try {

    if (response.status === 201) {
      alert('Resep berhasil dibuat!')
    }else{
      alert(json.errors)
    }

  } catch(error) {
    console.log(error)
  }
}

async function handleGetDrugData() {
  const search: SearchLimit = {
    search: searchDrug.value,
    limit: 10
  }

  const response = await getDrugData(search)
  const json = await response.json()

  try {

    if (response.status === 200) {
      drugs.value = json
    }else{
      alert(json.errors)
    }

  } catch(error) {
    console.log(error)
  }
}

async function handleGetCurrentRecipeNumber() {
  const response = await getCurrentRecipeNumber(recipeRequest.date)
  const json = await response.json()

  try {

    if (response.status === 200) {
      recipeRequest.recipe_number = recipeNumber(new Date(recipeRequest.date), json.response)
    }else{
      alert(json.errors)
    }

  } catch(error) {
    console.log(error)
  }
}

async function handleGetRecipeNumber() {
  const response = await getRecipeNumber(recipeRequest.care_number)
  const json = await response.json()

  try {

    if (response.status === 200) {
      recipeRequest.recipe_number = json.response
    }else{
      alert(json.errors)
    }

  } catch(error) {
    console.log(error)
  }
}

// Wacher
watch(recipeType, async (val) => {
  await (val ? handleGetRecipeNumber() : handleGetCurrentRecipeNumber());
  recipeRequest.type = val ? "add" : "create";
});

watch(() => recipeRequest.date, async () => {
  await handleGetCurrentRecipeNumber()
})

// Before page view
onBeforeMount(async () => {
  await handleGetCurrentRecipeNumber()
})
</script>

<template>
  <teleport to='body' >
    <section class="container center">
      <div class="cover2 scroll">
        <div class="close-btn center">
          <slot name="btn-close"></slot>
        </div>

        <div style="padding-top: 2rem; padding-bottom: 2rem;">
          <form class="form-data-custom" v-on:submit.prevent="handleCreateRecipe">
            <div class="center" style="justify-content: flex-start; align-items: flex-end; padding-left: 1rem;">
              <InputData :props="{ id: 'mr', name: 'Nomor rawat' }">
                <input type="text" id="mr" v-model="recipeRequest.care_number" readonly placeholder="no rawat">
              </InputData>
              <InputData :props="{ id: 'ss', name: 'Nomor resep' }">
                <input type="text" id="ss" v-model="recipeRequest.recipe_number" placeholder="resep">
              </InputData>
              <InputData :props="{ id: 'ck', name: 'Timpa resep' }">
                <input type="checkbox" id="ck" v-model="recipeType" placeholder="resep">
              </InputData>
              <InputData :props="{ id: 'da', name: 'Tanggal resep' }">
                <input type="datetime-local" id="da" step="1" v-model="recipeRequest.date" placeholder="resep">
              </InputData>
              <button>Save</button>
            </div>
          </form>
        </div>

        <h4 style="margin-bottom: 1rem;">Resep Obat</h4>
        <hr>

        <div style="width: 100%; height: 15rem; margin-top: 1rem;" class="scroll">
          <table class="table-custom">
            <thead>
              <tr>
                <td>Action</td>
                <td>Nama</td>
                <td>Jumlah</td>
                <td>Harga/tablet</td>
                <td>Aturan pakai</td>
                <td>Embalase</td>
                <td>Tuslah</td>
                <td>Total</td>
              </tr>
            </thead>
            <tbody>
              <tr v-for="rec in recipes" :key="rec.id">
                <td>
                  <button class="button-action" @click="removeRecipe(rec.id)">Delete</button>
                </td>
                <td>{{ rec.name }}</td>
                <td>{{ rec.value }}</td>
                <td>{{ rec.price }}</td>
                <td>{{ rec.use }}</td>
                <td>{{ rec.embalming }}</td>
                <td>{{ rec.tuslah }}</td>
                <td>{{ sum(rec.price, rec.value, rec.embalming, rec.tuslah) }}</td>
              </tr>
            </tbody>
          </table>
        </div>

        <div style="margin-top: 2rem; margin-left: 2rem;">Total = {{ loopSum(recipes) }}</div>

        <div style="padding: 0.5rem; padding-top: 2rem; padding-bottom: 2rem;">
          <form class="form-search-input" v-on:submit.prevent="handleGetDrugData">
            <div class="box-search">
              <div class="box">
                <input class="search-input" type="text" id="sc" v-model="searchDrug" placeholder="Nama obat">
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

        <div style="width: 100%; height: 15rem" class="scroll">
          <table class="table-custom">
            <thead>
              <tr>
                <td>Action</td>
                <td>Nama</td>
                <td>Jumlah</td>
                <td>Aturan pakai</td>
                <td>Embalase</td>
                <td>Tuslah</td>
                <td>Komposisi</td>
                <td>Harga/tablet</td>
              </tr>
            </thead>
            <tbody>
              <tr v-for="rec in drugs" :key="rec.id">
                <td>
                  <button class="button-action" @click="handleAddDrug(rec)">Add</button>
                </td>
                <td>{{ rec.name }}</td>
                <td>
                  <input type="number" v-model.number="valMap[rec.id]" />
                </td>
                <td>
                  <input type="text" v-model="useMap[rec.id]" />
                </td>
                <td>
                  <input type="number" v-model="embalmingMap[rec.id]" />
                </td>
                <td>
                  <input type="number" v-model="tuslahMap[rec.id]" />
                </td>
                <td>{{ rec.composition }}</td>
                <td>{{ rec.price }}</td>
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
