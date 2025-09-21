<script setup lang="ts">
import { logsApi } from '@/lib/api/logs';
import { formatDatetime, viewedDateTime } from '@/lib/formatDate';
import { onBeforeMount, ref } from 'vue';

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

    <div style="padding: 0.5rem; padding-top: 2rem; padding-bottom: 2rem;">
      <form class="form-search-input" v-on:submit.prevent="getLogs">
        <div class="box-search">
          <div class="box">
            <input class="search-input" type="datetime-local" step="1" id="dt1" v-model="date1" placeholder="date">
            <input class="search-input" type="datetime-local" step="1" id="dt2" v-model="date2" placeholder="date">
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
