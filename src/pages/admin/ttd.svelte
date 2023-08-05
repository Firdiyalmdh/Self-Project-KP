<script>
    import { Input, Label, Modal, Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell, Button, Toast } from 'flowbite-svelte';
    import { formatDate } from '../../utils';
    import { getData, putData } from '../../lib/request';
    import { onMount } from 'svelte';
  
    let permohonanList = [];
    let acceptModal = false;
    let rejectModal = false;
    let openToast = false;
    let toastMessage = "";
    let selectedData = {};

    $:getPermohonan = async () =>{
      getData({
            endpoint: '/permohonan',
            params: {
              tipe: "ttd"
            },
            onSuccess: (response) => {
                permohonanList = response.map(data => ({
                  id: data._id,
                  name: data.berkas.nama_berkas,
                  pemohon: data.pemohon.nama,
                  status: data.status,
                  date: data.tgl_masuk,
                  url: data.berkas.url_berkas,
                  result: data.hasil,
                  origin: data,
                }))
            },
            onFailed: (response) => {
                console.log(response);
                alert(response.data.message)
            }
      })
    }

    $:acceptPermohonan = async (event) => {
      const result = event.target[0].value
      putData({
        endpoint: `/permohonan/${selectedData.id}`,
        payload: {
          ...selectedData.origin,
          status: "selesai",
          hasil: result
        },
        onSuccess: () => {
          openToast = true
          toastMessage = "berhasil menyetujui permohonan"
          acceptModal = false
          getPermohonan()
          
          setTimeout(() => {
            openToast = false
            toastMessage = ""
            selectedData = {}
          }, 3000);
        },
        onFailed: (response) => {
          acceptModal = false
          console.log(response);
          alert(response.data.message)
        }
      })
    }

    $:rejectPermohonan = async () => {
      putData({
        endpoint: `/permohonan/${selectedData.id}`,
        payload: {
            ...selectedData.origin,
          status: "ditolak"
        },
        onSuccess: () => {
          openToast = true
          toastMessage = "Berhasil menolak permohonan"
          rejectModal = false
          getPermohonan()
          setTimeout(() => {
            openToast = false
            toastMessage = ""
            selectedData = {}
          }, 3000);
        },
        onFailed: (response) => {
          rejectModal = false
          console.log(response);
          alert(response.data.message)
        }
      })
    }

    onMount(async () => {
      await getPermohonan()
    })
  </script>
  
    <Table striped={true} divClass="m-5">
      <TableHead theadClass="text-left">
        <TableHeadCell>No</TableHeadCell>
        <TableHeadCell>Nama Pemohon</TableHeadCell>
        <TableHeadCell>Nama Dokumen</TableHeadCell>
        <TableHeadCell>Tanggal Pengajuan</TableHeadCell>
        <TableHeadCell>Status</TableHeadCell>
        <TableHeadCell>Action</TableHeadCell>
      </TableHead>
      <TableBody tableBodyClass="divide-y">
        {#each permohonanList as dokumen, index}
        <TableBodyRow class="py-3">
          <TableBodyCell>{ index + 1}.</TableBodyCell>
          <TableBodyCell>{ dokumen.pemohon }</TableBodyCell>
          <TableBodyCell>{ dokumen.name }</TableBodyCell>
          <TableBodyCell>{ formatDate(dokumen.date) }</TableBodyCell>
          <TableBodyCell>{ dokumen.status }</TableBodyCell>
          <TableBodyCell>
            <Button outline={true} color="blue" size="xs">
              <a href={dokumen.url} target="_blank" rel="noreferrer">
                <svg class="w-3 h-3" fill="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><title>eye</title><path fill-rule="evenodd" d="M12,9A3,3 0 0,0 9,12A3,3 0 0,0 12,15A3,3 0 0,0 15,12A3,3 0 0,0 12,9M12,17A5,5 0 0,1 7,12A5,5 0 0,1 12,7A5,5 0 0,1 17,12A5,5 0 0,1 12,17M12,4.5C7,4.5 2.73,7.61 1,12C2.73,16.39 7,19.5 12,19.5C17,19.5 21.27,16.39 23,12C21.27,7.61 17,4.5 12,4.5Z" clip-rule="evenodd" /></svg>
              </a>
            </Button>
            <Button on:click={() => {acceptModal = true; selectedData = dokumen}} outline={true} color="green" size="xs" disabled={dokumen.status === "selesai" || dokumen.status === "ditolak"}>
              <svg class="w-3 h-3" fill="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><title>check-circle</title><path fill-rule="evenodd" d="M12 2C6.5 2 2 6.5 2 12S6.5 22 12 22 22 17.5 22 12 17.5 2 12 2M10 17L5 12L6.41 10.59L10 14.17L17.59 6.58L19 8L10 17Z" clip-rule="evenodd" /></svg>
            </Button>
            <Button on:click={() => {rejectModal = true; selectedData = dokumen}} outline={true} color="red" size="xs" disabled={dokumen.status === "selesai" || dokumen.status === "ditolak"}>
              <svg class="w-3 h-3" fill="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><title>close-circle</title><path fill-rule="evenodd" d="M12,2C17.53,2 22,6.47 22,12C22,17.53 17.53,22 12,22C6.47,22 2,17.53 2,12C2,6.47 6.47,2 12,2M15.59,7L12,10.59L8.41,7L7,8.41L10.59,12L7,15.59L8.41,17L12,13.41L15.59,17L17,15.59L13.41,12L17,8.41L15.59,7Z" clip-rule="evenodd" /></svg>
            </Button>
          </TableBodyCell>
        </TableBodyRow>
        {/each}
      </TableBody>
    </Table>

    <Modal bind:open={acceptModal} size="xs" autoclose={false} class="w-full">
      <form class="flex flex-col space-y-6" on:submit|preventDefault={acceptPermohonan}>
        <h3 class="mb-4 text-xl font-medium text-gray-900 dark:text-white">Masukkan Dokumen Bertanda tangan</h3>
        <Label class="space-y-2">
          <span>Link dokumen yang sudah bertanda tangan (gdrive, atau lainnya)</span>
          <Input type="text" placeholder="Contoh: https://drive.google.com" required />
        </Label>
        <Button type="submit" class="w-full1">Kirim</Button>
      </form>
    </Modal>

    <Modal bind:open={rejectModal} size="xs" autoclose>
      <div class="text-center">
        <svg aria-hidden="true" class="mx-auto mb-4 w-14 h-14 text-gray-400 dark:text-gray-200" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
        <h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">Anda yakin ingin menolak permohonan ini?</h3>
        <Button on:click={rejectPermohonan} color="red" class="mr-2">Ya, tolak permohonan</Button>
        <Button color='alternative'>Tidak, batalkan</Button>
      </div>
    </Modal>

    <Toast color="green" open={openToast} position="top-right">
      <svelte:fragment slot="icon">
        <svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path></svg>
        <span class="sr-only">Check icon</span>
      </svelte:fragment>
      {toastMessage}
    </Toast>