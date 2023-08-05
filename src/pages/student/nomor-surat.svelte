<script>
  import { Toast } from 'flowbite-svelte';
  import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from 'flowbite-svelte';
  import { formatDate, getCurrentDate } from '../../utils';
  import { getData, postData } from '../../lib/request';
  import { onMount } from 'svelte';
  import { user } from "../../stores/UserStore";

  let permohonanList = []
  let openToast = false
  let toastMessage = ""
  let title = ""
  let link = ""

  $:getPermohonan = async () => {
    getData({
      endpoint: `/permohonan`,
      params:{
        "nama": $user.name,
        "tipe" : "nomor"
      },
      onSuccess: (response) => {
          permohonanList = response.map(data => ({
            id: data._id,
            name: data.berkas.nama_berkas,
            status: data.status,
            date: data.tgl_masuk,
            url: data.berkas.url_berkas,
            result: data.hasil,
          }))
      },
      onFailed: (response) => {
          console.log(response);
          alert(response.data.message)
      }
    })
  }

  $:createPermohonan = async () => {
    postData({
      endpoint: '/permohonan',
      payload: {
        tipe: "nomor",
        status: "diproses",
        pemohon: {
            nama: $user.name,
            nomor_pengenal: $user.id
        },
        berkas: {
            nama_berkas: title,
            url_berkas: link
        },
        tgl_masuk: getCurrentDate()
      },
      onSuccess: () => {
        openToast = true
        toastMessage = "berhasil membuat permohonan"
        getPermohonan()
        
        title = ""
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

  onMount(async () => {
    getPermohonan()
  })
</script>

<form class="m-10 w-96" on:submit|preventDefault={createPermohonan}>   
  <div class="my-5">
    <label>Nama Dokumen</label>
    <input 
      bind:value={title}
      class="h-10 font-thin border border-black rounded-lg w-full text-gray-400 leading-tight focus:outline-1 outline-blue-300 focus:shadow-outline" 
      type="text" 
      placeholder="Masukkan nama dokumen">
  </div>

    <div class="my-5">
      <label>Unggah Dokumen</label>
      <input 
        bind:value={link}
        class="h-10 font-thin border border-black rounded-lg w-full text-gray-400 leading-tight focus:outline-1 outline-blue-300 focus:shadow-outline"  
        type="text" 
        placeholder="Masukkan link google drive dokumen Anda">
    </div>
    <button type="submit" class="text-white bg-blue-500 w-full rounded-lg h-10 shadow-sm shadow-gray-300">Submit</button>
</form>
  
  <Table striped={true} divClass="m-5">
    <TableHead theadClass="text-left">
      <TableHeadCell>No</TableHeadCell>
      <TableHeadCell>Nama Dokumen</TableHeadCell>
      <TableHeadCell>Tanggal Pengajuan</TableHeadCell>
      <TableHeadCell>Status</TableHeadCell>
      <TableHeadCell>Nomor Surat</TableHeadCell>
      <TableHeadCell>Link</TableHeadCell>
    </TableHead>
    <TableBody tableBodyClass="divide-y">
      {#each permohonanList as dokumen, index}
      <TableBodyRow class="py-3">
        <TableBodyCell>{ index + 1}.</TableBodyCell>
        <TableBodyCell>{ dokumen.name }</TableBodyCell>
        <TableBodyCell>{ formatDate(dokumen.date) }</TableBodyCell>
        <TableBodyCell>{ dokumen.status }</TableBodyCell>
        <TableBodyCell>{ dokumen.result ? dokumen.result : "-" }</TableBodyCell>
        <TableBodyCell>
          <a class="underline" href={dokumen.url} target="_blank" rel="noreferrer">Buka dokumen</a>
        </TableBodyCell>
      </TableBodyRow>
      {/each}
    </TableBody>
  </Table>

  <Toast color="green" open={openToast} position="top-right">
    <svelte:fragment slot="icon">
      <svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path></svg>
      <span class="sr-only">Check icon</span>
    </svelte:fragment>
    {toastMessage}
  </Toast>