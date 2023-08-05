<div>
    <section class="bg-gambar">
      <div class="py-8 px-4 mx-auto max-w-screen-xl text-center lg:py-16 lg:px-12">
        <div class="flex max-w-full justify-center items-center">
          <div class="flex w-24 justify-center items-center">
            <img src="/images/pens-color.png">
            <img src="/images/hmce-color.png">
          </div>
        </div>  
        <h1 class="mb-4 text-4xl font-extrabold tracking-tight leading-none text-gray-900 md:text-5xl lg:text-6xl dark:text-white">Welcome to CEMIS</h1>
          <p class="mb-8 text-lg font-normal text-gray-500 lg:text-xl sm:px-16 xl:px-48 dark:text-gray-400">Computer Engineering Management and Information System</p>
          
          <button type="button" class="text-center font-medium focus:ring-4 focus:outline-none inline-flex items-center justify-center px-5 py-2.5 text-sm text-white bg-blue-700 hover:bg-blue-900  focus:ring-blue-300 rounded-lg">Login <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" class="h-4 w-4 ml-2" aria-label="chevron down" fill="none" viewBox="0 0 20 20" role="img" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" stroke="currentColor"></path></svg></button>
          <Dropdown >
            <DropdownItem href={$url('/auth/login')}>Mahasiswa</DropdownItem>
            <DropdownItem href={$url('/auth/login-lecture')}>Dosen</DropdownItem>
            <DropdownItem href={$url('')}>Admin</DropdownItem>
          </Dropdown>
         </div>
    </section>
  </div>
  
  <script>
    import { Dropdown, DropdownItem } from "flowbite-svelte";
    import { url } from "@roxi/routify";
  </script>

<style>
    .bg-gambar {
        background: url("../../../public/images/background.png") no-repeat center center fixed;
        height: 100vh;
        background-repeat: no-repeat;
        background-position: center center;
        background-attachment: fixed;
        background-size: cover;
    }
  </style>
  