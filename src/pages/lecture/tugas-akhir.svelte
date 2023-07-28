<script>
    import { Input, Label, Helper } from 'flowbite-svelte';
    import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from 'flowbite-svelte';
    import Uploadfile from '../../components/uploadfile.svelte';
    import { formatDate } from '../../utils';
    import { Button, Dropdown, DropdownItem, Chevron } from 'flowbite-svelte'
    import {  Heading, P, A } from 'flowbite-svelte'
  
    const kp = [
      {
        studentname: "Fidiyanti Al Ma'idha",
        title: "Smart Odontogram Berbasis Mobile Menggunakan Deep Learning",
        dosbim: "Riyanto Sigit",
        file: "Download",
      },
      {
        studentname: "Firmansyah",
        title: "Amour Unit .........",
        dosbim: "Dewi Mutiara Sari",
        file: "Download",
      },
    ]
  </script>
<div class="m-5">
    <Heading tag="h2" customSize="text-4xl font-extrabold ">Data Proyek Akhir Mahasiswa</Heading>
    <Table striped={true} divClass="m-5">
        <TableHead theadClass="text-left">
          <TableHeadCell>No</TableHeadCell>
          <TableHeadCell>Nama Mahasiswa</TableHeadCell>
          <TableHeadCell>Judul</TableHeadCell>
          <TableHeadCell>Dosen Pembimbing</TableHeadCell>
          <TableHeadCell>Dokumen PA</TableHeadCell>
        </TableHead>
        <TableBody tableBodyClass="divide-y">
          {#each kp as dokumen, index}
          <TableBodyRow class="py-3">
            <TableBodyCell>{ index + 1}.</TableBodyCell>
            <TableBodyCell>{ dokumen.studentname }</TableBodyCell>
            <TableBodyCell>{ dokumen.title }</TableBodyCell>
            <TableBodyCell>{ dokumen.dosbim}</TableBodyCell>
            <TableBodyCell>{ dokumen.file ? dokumen.file : "-" }</TableBodyCell>
          </TableBodyRow>
          {/each}
        </TableBody>
      </Table>
</div>