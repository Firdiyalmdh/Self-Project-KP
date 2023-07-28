<script>
    import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from 'flowbite-svelte';
    import {  Heading, P, A } from 'flowbite-svelte'
  
    const kp = [
      {
        studentname: "Fidiyanti Al Ma'idha",
        periode: "Genap 2022/2023",
        company: "PT ABCD",
        file: "Download",
      },
      {
        studentname: "Firmansyah",
        periode: "Genap 2022/2023",
        company: "PT XYZ",
        file: "Download",
      },
    ]
  </script>
<div class="m-5">
    <Heading tag="h2" customSize="text-4xl font-extrabold ">Data Kerja Praktik Mahasiswa</Heading>
    <Table striped={true} divClass="m-5">
        <TableHead theadClass="text-left">
          <TableHeadCell>No</TableHeadCell>
          <TableHeadCell>Nama Mahasiswa</TableHeadCell>
          <TableHeadCell>Periode</TableHeadCell>
          <TableHeadCell>Perusahaan</TableHeadCell>
          <TableHeadCell>Bukti Penerimaan</TableHeadCell>
        </TableHead>
        <TableBody tableBodyClass="divide-y">
          {#each kp as dokumen, index}
          <TableBodyRow class="py-3">
            <TableBodyCell>{ index + 1}.</TableBodyCell>
            <TableBodyCell>{ dokumen.studentname }</TableBodyCell>
            <TableBodyCell>{ dokumen.periode }</TableBodyCell>
            <TableBodyCell>{ dokumen.company}</TableBodyCell>
            <TableBodyCell>{ dokumen.file ? dokumen.file : "-" }</TableBodyCell>
          </TableBodyRow>
          {/each}
        </TableBody>
      </Table>
</div>