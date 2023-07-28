<script>
    import Periode from '../../components/periodekp.svelte';
    const permohonan = [
      {
        name: "Surat Keterangan",
        date: "27/07/2023",
        status: "diproses",
        file: null,
      },
      {
        name: "Surat Rekomendasi",
        date: "20/07/2023",
        status: "selesai",
        file: "Download",
      },
      {
        name: "Surat Embuh",
        date: "23/06/2023",
        status: "ditolak",
        file: null,
      },
    ]
  </script>
  
  <form class="m-10 w-96">   
      <div class="my-5">
      <label>Nama Perusahaan</label>
        <input 
          class="h-10 font-thin border border-black rounded-lg w-full text-gray-400 leading-tight focus:outline-1 outline-blue-300 focus:shadow-outline" 
          id="company" 
          type="text" 
          placeholder="Masukkan perusahan yang anda tuju">
      </div>
      <div class="my-5">
        <label>Periode</label>
        <Periode/>
    </div>
  
    <div class="my-5">
      <label>Unggah Bukti Penerimaan</label>
      <input 
      class="h-10 font-thin border border-black rounded-lg w-full text-gray-400 leading-tight focus:outline-1 outline-blue-300 focus:shadow-outline" 
      id="username" 
      type="text" 
      placeholder="Masukkan link google drive dokumen Anda">
  </div>
      <button class="text-white bg-blue-500 w-full rounded-lg h-10 shadow-sm shadow-gray-300">Submit</button>
  </form>