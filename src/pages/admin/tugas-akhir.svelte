<script>
  import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from 'flowbite-svelte';
  import { truncateString } from '../../utils';
  import {  Heading } from 'flowbite-svelte'
  import { getData } from '../../lib/request';
  import { onMount } from 'svelte';

  let pengumpulanList = []

  $:getPengumpulan = async () => {
  getData({
    endpoint: `/pengumpulan`,
    // params:{
    //   jenis : "pa"
    // },
    onSuccess: (response) => {
        pengumpulanList = response.map(data => ({
          id: data._id,
          name: data.nama,
          title: data.berkas.nama_berkas,
          link: data.berkas.url_berkas
        }))
    },
    onFailed: (response) => {
        console.log(response);
        alert(response.data.message)
    }
  })
}

onMount(async () => {
  getPengumpulan()
})
</script>
<div class="m-5">
  <Heading tag="h2" customSize="text-4xl font-extrabold ">Data Proyek Akhir Mahasiswa</Heading>
  <Table striped={true} divClass="m-5">
      <TableHead theadClass="text-left">
        <TableHeadCell>No</TableHeadCell>
        <TableHeadCell>Nama Mahasiswa</TableHeadCell>
        <TableHeadCell>Judul</TableHeadCell>
        <TableHeadCell>Dokumen PA</TableHeadCell>
      </TableHead>
      <TableBody tableBodyClass="divide-y">
        {#each pengumpulanList as data, index}
        <TableBodyRow class="py-3">
          <TableBodyCell>{ index + 1}.</TableBodyCell>
          <TableBodyCell>{ data.name }</TableBodyCell>
          <TableBodyCell>{ truncateString(data.title, 25) }</TableBodyCell>
          <TableBodyCell>
            {#if data.link}
              <a class="underline" href={data.link} target="_blank" rel="noreferrer">Buka dokumen</a>
            {:else}
              <p>-</p>
            {/if}
          </TableBodyCell>
        </TableBodyRow>
        {/each}
      </TableBody>
    </Table>
</div>