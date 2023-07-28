<script>
    import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from 'flowbite-svelte';
    import { formatDate } from '../../utils';
    import ListDosen from '../../components/list-dosen.svelte';

    const ttd = [
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
      <label>Nama Dokumen</label>
        <input 
          class="h-10 font-thin border border-black rounded-lg w-full text-gray-400 leading-tight focus:outline-1 outline-blue-300 focus:shadow-outline" 
          id="username" 
          type="text" 
          placeholder="Masukkan nama dokumen">
      </div>
      <div class="my-5">
        <label>Nama Dosen yang dituju</label>
        <ListDosen/>
    </div>
    <div class="my-5">
      <label>Unggah Dokumen</label>
      <input 
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
        <TableHeadCell>Tanggal Pengajuan</TableHeadCell>
        <TableHeadCell>Status</TableHeadCell>
        <TableHeadCell>File</TableHeadCell>
      </TableHead>
      <TableBody tableBodyClass="divide-y">
        {#each ttd as dokumen, index}
        <TableBodyRow class="py-3">
          <TableBodyCell>{ index + 1}.</TableBodyCell>
          <TableBodyCell>{ dokumen.name }</TableBodyCell>
          <TableBodyCell>{ formatDate(dokumen.date) }</TableBodyCell>
          <TableBodyCell>{ dokumen.status }</TableBodyCell>
          <TableBodyCell>{ dokumen.file ? dokumen.file : "-" }</TableBodyCell>
        </TableBodyRow>
        {/each}
      </TableBody>
    </Table>

   