<script setup lang="ts">
import { district, province, regencie, village } from '@/lib/api/address';
import { createPatient, deletePatient, getCurrentMedicalRecord, getPatient, updatePatient } from '@/lib/api/patient';
import { medicalRecord } from '@/lib/medicalRecordFormat';
import { onBeforeMount, reactive, ref, watch } from 'vue'
import type { Patient } from '@/types/patient';
import { viewedDate } from '@/lib/formatDate';
import type { Address } from '@/types/register';
import type { SearchLimit } from '@/types/response';

// Define variabels
const search = reactive<SearchLimit>({
  search: "",
  limit: 5
})
const patientDatas = ref<Patient[]>([])
const mr = ref<string>("")
const edits = ref<HTMLElement | null>()
const editAction = ref<boolean>(false)
const provincesDataSearch = ref<Address[]>([])
const findProvince = ref<string>("")
const provinceOpen = ref<boolean>(false)
const provinces = ref<Address[]>([])
const regencieDataSearch = ref<Address[]>([])
const findRegencie = ref<string>("")
const regencieOpen = ref<boolean>(false)
const regencies = ref<Address[]>([])
const districtDataSearch = ref<Address[]>([])
const findDistrict = ref<string>("")
const districtOpen = ref<boolean>(false)
const districts = ref<Address[]>([])
const villageDataSearch = ref<Address[]>([])
const findVillage = ref<string>("")
const villageOpen = ref<boolean>(false)
const villages = ref<Address[]>([])
const menuDataPatient = ref<boolean[]>([])
const patientData: Patient = reactive({
  medical_record: '',
  name: '',
  gender: '',
  wedding: '',
  religion: '',
  education: '',
  birth_place: '',
  birth_date: '',
  work: '',
  address: '',
  village: 0,
  district: 0,
  regencie: 0,
  province: 0,
  nik: '',
  bpjs: '',
  phone_number: '',
  parent_name: '',
  relationship: '',
  parent_gender: ''
});

// Define functions
function sanitizePatientData(data: Patient): Patient {
  return {
    ...data,
    province: Number(data.province),
    regencie: Number(data.regencie),
    district: Number(data.district),
    village: Number(data.village),
  };
}

async function edit(patient: Patient) {
  edits.value?.scrollIntoView({behavior: 'smooth'})

  Object.assign(patientData, patient)
}

function saveProvince(prov: Address) {
  patientData.province = prov.id
  findProvince.value = prov.name
  provinceOpen.value = false
  provincesDataSearch.value = []
}

function saveRegencie(reg: Address) {
  patientData.regencie = reg.id
  findRegencie.value = reg.name
  regencieOpen.value = false
  regencieDataSearch.value = []
}

function saveDistrict(dis: Address) {
  patientData.district = dis.id
  findDistrict.value = dis.name
  districtOpen.value = false
  districtDataSearch.value = []
}

function saveVillage(vil: Address) {
  patientData.village = vil.id
  findVillage.value = vil.name
  villageOpen.value = false
  villageDataSearch.value = []
}

async function resetForm() {
  await handleGetmedicalRecord()

  editAction.value = false

  patientData.name = ''
  patientData.gender = ''
  patientData.wedding = ''
  patientData.religion = ''
  patientData.education = ''
  patientData.birth_place = ''
  patientData.birth_date = ''
  patientData.work = ''
  patientData.address = ''
  patientData.village = 0
  patientData.district = 0
  patientData.regencie = 0
  patientData.province = 0
  patientData.nik = ''
  patientData.bpjs = ''
  patientData.phone_number = ''
  patientData.parent_name = ''
  patientData.relationship = ''
  patientData.parent_gender = ''
}

// Handler functions
async function handleCreatePatientData() {
  const response = await createPatient(sanitizePatientData(patientData))
  const json = await response.json()

  try {

    if (response.status === 201) {
      alert("simpan berhasil")
      await handleGetPatient()
      await resetForm()
      await handleGetmedicalRecord()
    }else{
      alert(json.errors)
    }

  } catch(error) {
    console.log(error)
  }
}

async function handleGetPatient() {
  const response = await getPatient(search)
  const json = await response.json()

  try {
    if (response.status === 200) {
      patientDatas.value = json
    } else {
      console.log(json.errors)
    }
  } catch(error) {
    console.log(error)
  }
}

async function handleGetSearchPatient() {
  const response = await getPatient(search)
  const json = await response.json()

  try {
    if (response.status === 200) {
      patientDatas.value = json
    } else {
      console.log(json.errors)
    }
  } catch(error) {
    console.log(error)
  }
}

async function handleDeletePatient(mr: string) {
  const response = await deletePatient(mr)
  const json = await response.json()

  try {
    if (response.status === 200) {
      await handleGetPatient()
      alert("hapus data berhasil")
    } else {
      alert(json.errors)
    }
  } catch(error) {
    console.log(error)
  }
}

async function handleUpdatePatient() {
  const response = await updatePatient(sanitizePatientData(patientData), mr.value)
  const json = await response.json()

  try {
    if (response.status === 200) {
      await handleGetPatient()
      await resetForm()
      await handleGetmedicalRecord()
      editAction.value = false
      alert("update data berhasil")
    } else {
      alert(json.errors)
    }
  } catch(error) {
    console.log(error)
  }
}

async function handleGetmedicalRecord() {
  const response = await getCurrentMedicalRecord()
  const json = await response.json()
  const mr = ref<string>()

  try {
    if (response.status === 200) {
      patientData.medical_record = medicalRecord(json.response)
    } else {
      console.log(response.status)

      return
    }
  } catch(error) {
    console.log(error)

    return
  }

  return mr.value
}

// Watcher
watch(() => findProvince.value, () => {
  provinceOpen.value = true
  const query = findProvince.value.toLowerCase();
  const results = provinces.value.filter(u =>
    u.name.toLowerCase().includes(query)
  );

  provincesDataSearch.value = results;
});

watch(() => patientData.address, async () => {
  const response = await province()
  const json = await response.json()

  try {
    if (response.status === 200) {
      provinces.value = json
    } else {
      console.log(response.status)
    }
  } catch(error) {
    console.log(error)
  }
})

watch(() => findRegencie.value, () => {
  regencieOpen.value = true
  const query = findRegencie.value.toLowerCase();
  const results = regencies.value.filter(u =>
    u.name.toLowerCase().includes(query)
  );

  regencieDataSearch.value = results;
});

watch(() => patientData.province, async () => {
  const response = await regencie(patientData.province)
  const json = await response.json()

  try {
    if (response.status === 200) {
      regencies.value = json
    } else {
      console.log(response.status)
    }
  } catch(error) {
    console.log(error)
  }
})

watch(() => findDistrict.value, () => {
  districtOpen.value = true
  const query = findDistrict.value.toLowerCase();
  const results = districts.value.filter(u =>
    u.name.toLowerCase().includes(query)
  );

  districtDataSearch.value = results;
});

watch(() => patientData.regencie, async () => {
  const response = await district(patientData.regencie)
  const json = await response.json()

  try {
    if (response.status === 200) {
      districts.value = json
    } else {
      console.log(response.status)
    }
  } catch(error) {
    console.log(error)
  }
})

watch(() => findVillage.value, () => {
  villageOpen.value = true
  const query = findVillage.value.toLowerCase();
  const results = villages.value.filter(u =>
    u.name.toLowerCase().includes(query)
  );

  villageDataSearch.value = results;
});

watch(() => patientData.district, async () => {
  const response = await village(patientData.district)
  const json = await response.json()

  try {
    if (response.status === 200) {
      villages.value = json
    } else {
      console.log(response.status)
    }
  } catch(error) {
    console.log(error)
  }
})

// Before page view
onBeforeMount(async () => {
  await handleGetPatient()
  await handleGetmedicalRecord()
})
</script>

<template>
  <section class="anim-slide" ref="edits">
    <h3 style="margin: 0.5rem;">Registrasi Pasien</h3>
    <div style="margin-top: 2rem; margin-bottom: 2rem;" class="bottom-line">
      <form class="form-data-custom" v-on:submit.prevent="handleCreatePatientData">

        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Data pasien</h4>
        <div style="display: grid; grid-template-columns: auto auto auto; padding-left: 1rem;">
          <div class="input-field">
            <div class="cover">
              <label for="mr">Nomor rekam medis</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <input type="text" id="mr" v-model="patientData.medical_record" placeholder="No RM" maxlength="6" required>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="name">Nama pasien</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <input type="text" id="name" v-model="patientData.name" placeholder="Nama pasien" maxlength="50" required>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="gender">Jenis kelamin</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <select id="gender" v-model="patientData.gender">
              <option value="Laki - laki">Laki - laki</option>
              <option value="Perempuan">Perempuan</option>
            </select>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="wedding">Status pernikahan</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <select id="wedding" v-model="patientData.wedding">
              <option value="Menikah">Menikah</option>
              <option value="Belum menikah">Belum menikah</option>
              <option value="Cerai">Cerai</option>
              <option value="Cerai mati">Cerai mati</option>
            </select>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="religion">Agama</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <select id="religion" v-model="patientData.religion">
              <option value="Islam">Islam</option>
              <option value="Kristen">Kristen</option>
              <option value="Katholik">Katholik</option>
              <option value="Hindu">Hindu</option>
              <option value="Budha">Budha</option>
              <option value="Konghuchu">Konghuchu</option>
            </select>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="education">Pendidikan</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <select id="education" v-model="patientData.education">
              <option value="Tidak/Belum sekolah">Tidak/Belum sekolah</option>
              <option value="SD/Sederajat">SD/Sederajat</option>
              <option value="SMP/Sederajat">SMP/Sederajat</option>
              <option value="SMA/Sederajat">SMA/Sederajat</option>
              <option value="Diploma (D1 - D4)">Diploma (D1 - D4)</option>
              <option value="Sarjana (S1)">Sarjana (S1)</option>
              <option value="Magister (S2)">Magister (S2)</option>
              <option value="Doctor (S3)">Doctor (S3)</option>
            </select>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="birthplace">Tempat lahir</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <input type="text" id="birthplace" v-model="patientData.birth_place" placeholder="Tempat lahir" maxlength="20" required>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="birthdate">Tanggal lahir</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <input type="date" id="birthdate" v-model="patientData.birth_date" placeholder="Tanggal lahir" required>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="work">Pekerjaan</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <input type="text" id="work" v-model="patientData.work" placeholder="Pekerjaan" maxlength="20" required>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="nik">NIK</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <input type="text" id="nik" v-model="patientData.nik" placeholder="320*****" maxlength="16" required>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="bpjs">BPJS</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <input type="text" id="bpjs" v-model="patientData.bpjs" placeholder="000****" maxlength="20" required>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="phone_number">No. HP</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <input type="text" id="phone_number" v-model="patientData.phone_number" placeholder="08*****" maxlength="13" required>
          </div>
        </div>

        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Alamat</h4>
        <div style="display: grid; grid-template-columns: auto auto auto; padding-left: 1rem;">
          <div class="input-field">
            <div class="cover">
              <label for="address">Alamat</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <input type="text" id="address" v-model="patientData.address" placeholder="Alamat" required>
          </div>

          <div class="input-field" style="position: relative;">
            <div class="cover">
              <label>Provinsi</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <div>
              <input type="text" v-model="findProvince" placeholder="Provinsi">
              <div class="find-input" :style="{ display: !provinceOpen ? 'none' : '' }">
                <button type="button" @click="provinceOpen = false">close</button>
                <ul>
                  <li v-for="prov in provincesDataSearch" :key="prov.id" @click="saveProvince(prov)">{{ prov.name }}</li>
                </ul>
              </div>
            </div>
          </div>

          <div class="input-field" style="position: relative;">
            <div class="cover">
              <label>Kabupaten</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <div>
              <input type="text" v-model="findRegencie" placeholder="Kabupaten">
              <div class="find-input" :style="{ display: !regencieOpen ? 'none' : '' }">
                <button type="button" @click="regencieOpen = false">close</button>
                <ul>
                  <li v-for="reg in regencieDataSearch" :key="reg.id" @click="saveRegencie(reg)">{{ reg.name }}</li>
                </ul>
              </div>
            </div>
          </div>

          <div class="input-field" style="position: relative;">
            <div class="cover">
              <label>Kecamatan</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <div>
              <input type="text" v-model="findDistrict" placeholder="Kecamatan">
              <div class="find-input" :style="{ display: !districtOpen ? 'none' : '' }">
                <button type="button" @click="districtOpen = false">close</button>
                <ul>
                  <li v-for="dis in districtDataSearch" :key="dis.id" @click="saveDistrict(dis)">{{ dis.name }}</li>
                </ul>
              </div>
            </div>
          </div>

          <div class="input-field" style="position: relative;">
            <div class="cover">
              <label>Kelurahan</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <div>
              <input type="text" v-model="findVillage" placeholder="Kelurahan">
              <div class="find-input" :style="{ display: !villageOpen ? 'none' : '' }">
                <button type="button" @click="villageOpen = false">close</button>
                <ul>
                  <li v-for="vil in villageDataSearch" :key="vil.id" @click="saveVillage(vil)">{{ vil.name }}</li>
                </ul>
              </div>
            </div>
          </div>

        </div>

        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Data penanggung jawab</h4>
        <div style="display: grid; grid-template-columns: auto auto auto; padding-left: 1rem;">
          <div class="input-field">
            <div class="cover">
              <label for="parent_name">Nama</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <input type="text" id="parent_name" v-model="patientData.parent_name" placeholder="Nama Penanggung jawab" required>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="relationship">Hubungan</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <select id="relationship" v-model="patientData.relationship">
              <option value="Ibu">Ibu</option>
              <option value="Ayah">Ayah</option>
              <option value="Saudara">Saudara</option>
              <option value="Teman">Teman</option>
            </select>
          </div>

          <div class="input-field">
            <div class="cover">
              <label for="parent_gender">Jenis Kelamin</label>
            </div>
            <span style="padding-right: 0.5rem;">:</span>
            <select id="parent_gender" v-model="patientData.parent_gender">
              <option value="Laki - laki">Laki - laki</option>
              <option value="Perempuan">Perempuan</option>
            </select>
          </div>

        </div>

        <div>
          <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Action</h4>
          <button type="submit">Save</button>
          <button type="button" @click="resetForm">Reset</button>
          <button type="button" v-if="editAction" @click="handleUpdatePatient()">Update</button>
        </div>
      </form>
    </div>

    <div style="padding-top: 2rem; padding-bottom: 2rem;">
      <form class="form-data-custom" v-on:submit.prevent="handleGetSearchPatient">
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Cari pasien</h4>
        <div class="center" style="justify-content: flex-start; align-items: flex-end; padding-left: 1rem;">
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="parent_name">Cari</label>
            </div>
            <input type="text" id="parent_name" placeholder="No RM/Nama Pasien" v-model="search.search">
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="limit1">Limit</label>
            </div>
            <input type="number" id="limit1" placeholder="limit data" v-model="search.limit">
          </div>
          <button>Cari</button>
        </div>
      </form>
    </div>

    <div style="width: 100%; height: 20rem; overflow-x: scroll; overflow-y: scroll; scrollbar-width: thin;">
      <table class="table-custom">
        <thead>
          <tr>
            <td>No rekam medis</td>
            <td>Nama</td>
            <td>Tempat/tanggal lahir</td>
            <td>Jenis Kelamin</td>
            <td>Alamat</td>
            <td>NIK</td>
            <td>BPJS</td>
            <td>No telepon</td>
            <td>Action</td>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(patient, index) in patientDatas" :key="patient.medical_record">
            <td>{{ patient.medical_record }}</td>
            <td>{{ patient.name }}</td>
            <td>{{ `${patient.birth_place}, ${viewedDate(patient.birth_date)}` }}</td>
            <td>{{ patient.gender }}</td>
            <td>{{ patient.address }}</td>
            <td>{{ patient.nik }}</td>
            <td>{{ patient.bpjs }}</td>
            <td>{{ patient.phone_number }}</td>
            <td style="position: relative;">
                <div class="drop-down">
                    <button class="act" @click="menuDataPatient[index] = !menuDataPatient[index]">:</button>
                    <div class="menu" v-if="menuDataPatient[index]">
                        <ul>
                            <li>
                              <button class="button-action" @click="edit(patient); mr = patient.medical_record; editAction = true; menuDataPatient[index] = false">Edit</button>
                            </li>
                            <li>
                              <button class="button-action" @click="handleDeletePatient(patient.medical_record); menuDataPatient[index] = false">Delete</button>
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
