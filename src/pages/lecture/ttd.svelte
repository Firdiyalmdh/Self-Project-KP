<script>
  import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from 'flowbite-svelte';
  import { formatDate, getCurrentDate } from '../../utils';
  import ListDosen from '../../components/list-dosen.svelte';
  import { getData, postData } from '../../lib/request';
  import { user } from "../../stores/UserStore";
  import { onMount } from 'svelte';

  let permohonanList = []
  let openToast = false
  let toastMessage = ""
  let tujuan = ""
  let title =""
  let link =""

  $:getPermohonan = async () => {
  getData({
    endpoint: `/permohonan`,
    params: {
      "nama": $user.name,
      "tipe": "ttd"
    },
    onSuccess: (response) => {
        permohonanList = response.map(data => ({
          id: data._id,
          name: data.berkas.nama_berkas,
          lecture: data.tujuan,
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

$:createPermohonan = async (event) => {
  const name = event.target[0].value
  const url = event.target[1].value
  postData({
    endpoint: '/permohonan',
    payload: {
      tipe: "ttd",
      status: "diproses",
      pemohon: {
          nama: $user.name,
          nomor_pengenal: $user.id
      },
      berkas: {
          nama_berkas: name,
          url_berkas: url
      },
      tgl_masuk: getCurrentDate(),
      tujuan,
    },
    onSuccess: () => {
      openToast = true
      toastMessage = "berhasil membuat permohonan"
      getPermohonan()
      
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
        id="username" 
        type="text" 
        placeholder="Masukkan nama dokumen">
    </div>
    <div class="my-5">
      <label>Nama Dosen yang dituju</label>
      <ListDosen on:dosenSelected={(data) => tujuan = data.detail.dosen}/>
  </div>
  <div class="my-5">
    <label>Unggah Dokumen</label>
    <input 
    bind:value={link}
    class="h-10 font-thin border border-black rounded-lg w-full text-gray-400 leading-tight focus:outline-1 outline-blue-300 focus:shadow-outline" 
    id="username" 
    type="text" 
    placeholder="Masukkan link google drive dokumen Anda">
</div>
    <button class="text-white bg-blue-500 w-full rounded-lg h-10 shadow-sm shadow-gray-300">Submit</button>
</form>
  
  <Table striped={true} divClass="m-5">
    <TableHead theadClass="text-left">
      <TableHeadCell>No</TableHeadCell>
      <TableHeadCell>Nama Dokumen</TableHeadCell>
      <TableHeadCell>Nama Dosen yang Dituju</TableHeadCell>
      <TableHeadCell>Tanggal Pengajuan</TableHeadCell>
      <TableHeadCell>Status</TableHeadCell>
      <TableHeadCell>File</TableHeadCell>
    </TableHead>
    <TableBody tableBodyClass="divide-y">
      {#each permohonanList as dokumen, index}
      <TableBodyRow class="py-3">
        <TableBodyCell>{ index + 1}.</TableBodyCell>
        <TableBodyCell>{ dokumen.name }</TableBodyCell>
        <TableBodyCell>{ dokumen.lecture }</TableBodyCell>
        <TableBodyCell>{ formatDate(dokumen.date) }</TableBodyCell>
        <TableBodyCell>{ dokumen.status }</TableBodyCell>
        <TableBodyCell>{ dokumen.result ? dokumen.result : "-" }</TableBodyCell>
      </TableBodyRow>
      {/each}
    </TableBody>
  </Table>

 