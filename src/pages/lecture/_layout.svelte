<script>
    import { Sidebar, SidebarGroup, SidebarItem, SidebarWrapper, SidebarDropdownItem, SidebarDropdownWrapper } from 'flowbite-svelte';
    import { url, goto } from '@roxi/routify';
    import { user } from '../../stores/UserStore';
    import { deleteData } from '../../lib/request';

    $:logout = async () => {
        deleteData({
            endpoint: `/logout/${$user.sessionId}`,
            onSuccess: () => {
                user.set(null)
                $goto("/")
            },
            onFailed: () => {}
        })
    }
</script>

<svelte:head>
	<title>CEMIS | Dosen</title>
</svelte:head>
{#if $user !== null && $user.role === "dosen"}
<div class="flex w-full">
    <Sidebar class="min-w-[25%] bg-gradient-to-r from-amber-700 via-blue-00 to-yellow-900 h-screen">
        <div class="w-100 my-5 flex justify-center items-center">
            <div class="flex">
                <img src="/images/pens.png" width="70">
                <img src="/images/hmce.png" width="70">
            </div>
        </div>
        <p class="text-base flex items-center justify-center mb-3 font-bold text-white leading-normal text-left whitespace-normal">{$user.name }</p>
        <SidebarWrapper divClass="bg-transparent">
            <SidebarGroup>
            <SidebarItem label="Home" href={$url('/lecture/home')} aClass="hover:bg-black/25 flex items-center p-2 text-base font-normal text-white rounded-lg">
                <svelte:fragment slot="icon">
                    <svg class="w-6 h-6 text-white dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20">
                        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M3 8v10a1 1 0 0 0 1 1h4v-5a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v5h4a1 1 0 0 0 1-1V8M1 10l9-9 9 9"/>
                    </svg>        
                </svelte:fragment>
            </SidebarItem>
            <SidebarItem label="Permohonan Nomor Surat" href={$url('/lecture/nomor-surat')} aClass="hover:bg-black/25 flex items-center p-2 text-base font-normal text-white rounded-lg">
                <svelte:fragment slot="icon">
                <svg class="w-6 h-6 text-white dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 16 20">
                    <path stroke="currentColor" stroke-linejoin="round" stroke-width="1" d="M6 1v4a1 1 0 0 1-1 1H1m14-4v16a.97.97 0 0 1-.933 1H1.933A.97.97 0 0 1 1 18V5.828a2 2 0 0 1 .586-1.414l2.828-2.828A2 2 0 0 1 5.828 1h8.239A.97.97 0 0 1 15 2Z"/>
                </svg>          </svelte:fragment>
            </SidebarItem>
            <SidebarItem label="Permohonan Tanda Tangan"href={$url('/lecture/ttd')} aClass="hover:bg-black/25 flex items-center p-2 text-base font-normal text-white rounded-lg">
                <svelte:fragment slot="icon">
                <svg class="w-6 h-6 text-white dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 16 20">
                    <path stroke="currentColor" stroke-linejoin="round" stroke-width="1" d="M6 1v4a1 1 0 0 1-1 1H1m14-4v16a.97.97 0 0 1-.933 1H1.933A.97.97 0 0 1 1 18V5.828a2 2 0 0 1 .586-1.414l2.828-2.828A2 2 0 0 1 5.828 1h8.239A.97.97 0 0 1 15 2Z"/>
                </svg>          </svelte:fragment>
            </SidebarItem>
            <SidebarDropdownWrapper label="Berkas Masuk" btnClass="flex items-center p-2 w-full text-base font-normal text-white rounded-lg transition duration-75 group hover:bg-black/25 dark:text-white dark:hover:bg-gray-700" spanClass="grow-0 mx-3">
                <svelte:fragment slot="icon">
                    <svg class="w-6 h-6 text-white dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 18 20">
                        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="m7.708 2.292.706-.706A2 2 0 0 1 9.828 1h6.239A.97.97 0 0 1 17 2v12a.97.97 0 0 1-.933 1H15M6 5v4a1 1 0 0 1-1 1H1m11-4v12a.97.97 0 0 1-.933 1H1.933A.97.97 0 0 1 1 18V9.828a2 2 0 0 1 .586-1.414l2.828-2.828A2 2 0 0 1 5.828 5h5.239A.97.97 0 0 1 12 6Z"/>
                    </svg>          
                </svelte:fragment>
                <SidebarDropdownItem label="Kerja Praktik" href={$url('/lecture/kerja-praktik')} aClass="flex items-center p-2 pl-11 w-full text-base font-normal text-white rounded-lg transition duration-75 group hover:bg-black/25 dark:text-white dark:hover:bg-gray-700"/>
                <SidebarDropdownItem label="Proyek Akhir" href={$url('/lecture/tugas-akhir')} aClass="flex items-center p-2 pl-11 w-full text-base font-normal text-white rounded-lg transition duration-75 group hover:bg-black/25 dark:text-white dark:hover:bg-gray-700"/>
                <SidebarDropdownItem label="PKM" href={$url('/lecture/pkm')} aClass="flex items-center p-2 pl-11 w-full text-base font-normal text-white rounded-lg transition duration-75 group hover:bg-black/25 dark:text-white dark:hover:bg-gray-700"/>
            </SidebarDropdownWrapper>
            <SidebarItem on:click={logout} label="Logout" aClass="hover:bg-black/25 flex items-center p-2 text-base font-normal text-white rounded-lg">
                <svelte:fragment slot="icon">
                <svg class="w-6 h-6 text-white dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10" > 
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M13 5H1m0 0 4 4M1 5l4-4"/>
                </svg>          </svelte:fragment>
            </SidebarItem>
            </SidebarGroup>
        </SidebarWrapper>
    </Sidebar>
    <main class="grow">
        <slot></slot>
    </main>
</div>
{:else}
    {$goto("/auth/login-lecture")}
{/if}