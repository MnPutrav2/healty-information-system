<template>
  <section class="anim-slide" ref="pages">
    <h3 style="margin: 0.5rem;">Data Pemeriksaan</h3>
    <div style="padding-top: 2rem; padding-bottom: 2rem;" class="bottom-line">
      <form class="form-data-custom">
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Buat data pemeriksaan labolatorium</h4>
        <div style="display: grid; grid-template-columns: auto auto auto; padding-left: 1rem;">
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="id">ID pemeriksaan</label>
            </div>
            <input type="text" id="id" placeholder="ID pemeriksaan">
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="nm">Nama pemeriksaan</label>
            </div>
            <input type="text" id="nm" placeholder="Nama pemeriksaan">
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="prd">Biaya perujuk</label>
            </div>
            <input type="number" id="prd" placeholder="Biaya perujuk">
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="pr">Biaya petugas</label>
            </div>
            <input type="number" id="pr" placeholder="Biaya petugas">
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="ma">Manajement</label>
            </div>
            <input type="number" id="ma" placeholder="Biaya manajement">
          </div>
          <div>
            <button type="submit">Create</button>
          </div>
        </div>
      </form>
    </div>

    <!-- Search -->
    <div style="padding-top: 2rem; padding-bottom: 2rem;">
      <form class="form-data-custom">
        <h4 style="margin: 0.5rem; color: var(--font-color-sec);">Data Obat</h4>
        <div class="center" style="justify-content: flex-start; align-items: flex-end; padding-left: 1rem;">
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="sc">Cari</label>
            </div>
            <input type="text" id="sc" placeholder="Nama obat">
          </div>
          <div style="padding: 0.5rem;">
            <div style="margin-bottom: 0.5rem;">
              <label for="lm">Limit</label>
            </div>
            <input type="number" id="lm" placeholder="limit">
          </div>
          <button type="submit">Cari</button>
        </div>
      </form>
    </div>

    <!-- Data -->
    <div style="width: 100%; overflow-x: scroll; overflow-y: scroll; height: 20rem; scrollbar-width: thin;" class="bottom-line">
      <table class="table-custom">
        <thead>
          <tr>
            <td>No</td>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>1</td>
          </tr>
        </tbody>
      </table>
    </div>

  </section>
</template>
