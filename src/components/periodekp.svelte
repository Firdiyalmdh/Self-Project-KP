<script>
    let periodekp = [
        { id: 0, text: `Pilih periode KP` },
        { id: 1, text: `Ganjil 2022/2023` },
		{ id: 2, text: `Genap 2022/2023` },
	  ];

    let periode= '';
    
	let selected;
</script>
  
<form on:submit>
    <select class="h-10 font-thin border border-black w-full text-gray-500 text-left focus:outline-1 inline-flex items-center  px-5 py-2.5 bg-white hover:bg-gray-200 rounded-lg"
    bind:value={selected} on:change={() => (periode = '')}>
    {#each periodekp as periode}
        <option value={periode}>
        {periode.text}
        </option>
    {/each}
    </select>
</form>
