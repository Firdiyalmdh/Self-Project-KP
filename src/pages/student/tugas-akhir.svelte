<script>
  import ListDosen from '../../components/list-dosen.svelte';
</script>
  
  <form class="m-10 w-96">   
      <div class="my-5">
      <label>Judul PA</label>
        <input 
          class="h-10 font-thin border border-black rounded-lg w-full text-gray-400 leading-tight focus:outline-1 outline-blue-300 focus:shadow-outline" 
          id="company" 
          type="text" 
          placeholder="Masukkan perusahan yang anda tuju">
      </div>
      <div class="my-5">
        <label>Dosen Pembimbing Utama</label>
        <ListDosen/>
    </div>
  
    <div class="my-5">
      <label>Unggah Dokumen Proyek Akhir</label>
      <input 
      class="h-10 font-thin border border-black rounded-lg w-full text-gray-400 leading-tight focus:outline-1 outline-blue-300 focus:shadow-outline" 
      id="username" 
      type="text" 
      placeholder="Masukkan link google drive dokumen Anda">
  </div>
      <button class="text-white bg-blue-500 w-full rounded-lg h-10 shadow-sm shadow-gray-300">Submit</button>
  </form>