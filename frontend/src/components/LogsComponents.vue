<script setup lang="ts">
import { logsApi } from '@/lib/api/logs';
import { formatDatetime, viewedDateTime } from '@/lib/formatDate';
import { onBeforeMount, ref } from 'vue';
import InputData from './Extras/InputData.vue';

// Define variabels
const date = new Date()
const date1 = ref<string>(formatDatetime(date, "01:00:00"))
const date2 = ref<string>(formatDatetime(date, "23:59:59"))
const logs = ref<Logs[]>([])

// External type
interface Logs {
  id: string,
  user: string,
  level: string,
  message: string,
  path: string,
  date: string,
}

// Define functions
async function getLogs() {
  const response = await logsApi(date1.value, date2.value)
  const json = await response.json()

  try {
    if (response.status === 200) {
      logs.value = json
    } else {
      console.log(json.errors)
    }
  } catch(error) {
    console.log(error)
  }
}

// Before page view
onBeforeMount(async () => {
  await getLogs()
})
</script>

<template>
  <section class="anim-slide" ref="pages">
    <h3 style="margin: 0.5rem;">Logs Activity</h3>
    <div style="margin-top: 2rem; margin-bottom: 2rem;">
      <form class="form-data-custom" v-on:submit.prevent="getLogs">
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Cari</h4>
        <div class="responsive-grid" style="padding-left: 1rem;">
          <InputData :props="{ id: 'tgl1', name: 'Tanggal awal' }">
            <input type="datetime-local" step="1" id="tgl1" v-model="date1" placeholder="tanggal">
          </InputData>
          <InputData :props="{ id: 'tgl2', name: 'Tanggal akhir' }">
            <input type="datetime-local" step="1" id="tgl2" v-model="date2" placeholder="tanggal">
          </InputData>
        </div>
        <div style="margin: 1rem;">
          <button type="submit" style="width: 2rem;">Cari</button>
        </div>
      </form>

      <table class="table-custom" style="width: 100%; margin-top: 2rem;">
        <thead>
          <tr>
            <td>Date</td>
            <td>User</td>
            <td>Level</td>
            <td>Message</td>
            <td>Path</td>
          </tr>
        </thead>
        <tbody>
          <tr v-for="log in logs" :key="log.id" :class="log.level == 'WARN' ? 'warn' : log.level == 'ERROR' ? 'error' : 'info'">
            <td>{{ viewedDateTime(log.date) }}</td>
            <td>{{ log.user == "" || log.user == "0" ? "SYSTEM" : log.user }}</td>
            <td>{{ log.level }}</td>
            <td>{{ log.message }}</td>
            <td>{{ log.path }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </section>
</template>

<style scoped>
.warn {
  color: yellow;
}

.warn > td {
  background-color: rgba(255, 255, 0, 0.255);
}

.error {
  color: rgb(255, 0, 0);
}

.error > td {
  background-color: rgba(255, 0, 0, 0.255);
}

.info {
  color: rgb(255, 255, 255);
}
</style>
