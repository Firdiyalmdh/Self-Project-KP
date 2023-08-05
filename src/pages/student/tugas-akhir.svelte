<script>
import { Toast } from "flowbite-svelte";
import { postData } from "../../lib/request";
import { user } from "../../stores/UserStore";
import { getCurrentDate } from "../../utils";

let judul = ""
let link = ""
let openToast = false
let toastMessage = ""

$:createPengumpulan = async () => {
  postData({
    endpoint: '/pengumpulan',
    payload: {
      jenis: "pa",
      nama: $user.name,
      nomor_pengenal: $user.id,
      berkas: {
          nama_berkas: judul,
          url_berkas: link
      },
      tgl: getCurrentDate()
    },
    onSuccess: () => {
      openToast = true
      toastMessage = "berhasil mengunggah berkas"
      
      judul = ""
      link = ""
      setTimeout(() => {
        openToast = false
        toastMessage = ""
      }, 3000);
    },
    onFailed: (response) => {
      console.log(response);
      alert(response.data.message)
    }
  })
}
</script>

<form class="m-10 w-96" on:submit|preventDefault={createPengumpulan}>   
    <div class="my-5">
    <label>Judul PA</label>
      <input 
        bind:value={judul}
        class="h-10 font-thin border border-black rounded-lg w-full text-gray-400 leading-tight focus:outline-1 outline-blue-300 focus:shadow-outline" 
        id="judul" 
        type="text" 
        placeholder="Masukkan judul PA Anda">
    </div>

  <div class="my-5">
    <label>Unggah Dokumen PA</label>
    <input 
    bind:value={link}
    class="h-10 font-thin border border-black rounded-lg w-full text-gray-400 leading-tight focus:outline-1 outline-blue-300 focus:shadow-outline" 
    id="username" 
    type="text" 
    placeholder="Masukkan link google drive dokumen Anda">
</div>
    <button type="submit" class="text-white bg-blue-500 w-full rounded-lg h-10 shadow-sm shadow-gray-300">Submit</button>
</form>

<Toast color="green" open={openToast} position="top-right">
  <svelte:fragment slot="icon">
    <svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path></svg>
    <span class="sr-only">Check icon</span>
  </svelte:fragment>
  {toastMessage}
</Toast>