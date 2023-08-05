<script>
    import { goto } from "@roxi/routify";
    import { user } from "../../stores/UserStore";
</script>

<style>
    .bg-login {
        background: url("../../../public/images/background.png") no-repeat center center fixed; 
        height: 100vh;
        background-repeat: no-repeat;
        background-position: center center;
        background-attachment: fixed;
    background-size: cover; 
    }
</style>

<svelte:head>
	<title>CEMIS | Login</title>
</svelte:head>
{#if $user === null}
<div class="h-screen flex justify-center items-center bg-login"> 
    <slot></slot>
</div>
{:else if $user.role === "mahasiswa"}
    {$goto("/student/home")}
{:else if $user.role === "dosen"}
    {$goto("/lecture/home")}
{:else if $user.role === "admin"}
    {$goto("/admin/home")}
{/if}