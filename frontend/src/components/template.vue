<script setup lang="ts">
import InputData from '@/components/Extras/InputData.vue';

</script>

<template>
  <section class="anim-slide" ref="pages">
    <h3 style="margin: 0.5rem;">Data Pegawai</h3>
    <div style="padding-top: 0.5rem; padding-bottom: 1rem;" class="bottom-line">
      <form class="form-data-custom" v-on:submit.prevent="">
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Tambah pegawai</h4>
        <div style="display: grid; grid-template-columns: auto auto auto; padding-left: 1rem;">
          <InputData :props="{id: 'id', name: 'ID pegawai'}">
            <input type="text" id="id" placeholder="ID pemeriksaan" required>
          </InputData>
        </div>
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Action</h4>
        <div>
          <button type="button">Reset</button>
          <button type="submit">Add</button>
        </div>
      </form>
    </div>

    <div>
      <form class="form-data-custom">
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Cari data pegawai</h4>
        <div class="center" style="justify-content: flex-start; align-items: flex-end; padding-left: 1rem;">
          <InputData :props="{id: 'nv', name: 'Nama pegawai'}">
            <input type="text" id="nv" placeholder="Nama pegawai">
          </InputData>
          <InputData :props="{id: 'l', name: 'Limit'}">
            <input type="number" id="l"  placeholder="Limit">
          </InputData>
          <button type="submit">Cari</button>
        </div>
      </form>
    </div>

    <div style="width: 100%; overflow-x: scroll; overflow-y: scroll; height: 20rem; scrollbar-width: thin;" class="bottom-line">
      <table class="table-custom">
        <thead>
          <tr>
            <td>ID pegawai</td>
            <td>Nama</td>
            <td>Jenis kelamin</td>
            <td>Tempat, tanggal lahir</td>
            <td>Status menikah</td>
            <td>Alamat</td>
            <td>NIK</td>
            <td>No BPJS</td>
            <td>No NPWP</td>
            <td>Nomor telepon</td>
            <td>Email</td>
            <td>Action</td>
          </tr>
        </thead>

      </table>
    </div>
  </section>
</template>
