<script>
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();

    const handleChange = () => {
        dispatch('dosenSelected', {
            dosen: selected.text
        });
    }

    let listdosen = [
        { id: 0, text: `Pilih nama dosen` },
        { id: 1, text: `Riyanto Sigit, ST., M.Kom., Ph.D` },
		{ id: 2, text: `Heny Yuniarti, S.ST. MT` },
		{ id: 3, text: `Tri Harsono, S.Si., M.Kom., Ph.D` }
	  ];
    
	let selected;
</script>
  
<form on:submit>
    <select class="h-10 font-thin border border-black w-full text-gray-500 text-left focus:outline-1 inline-flex items-center  px-5 py-2.5 bg-white hover:bg-gray-200 rounded-lg"
        bind:value={selected} on:change={handleChange}>
    {#each listdosen as namadosen}
        <option value={namadosen}>
            {namadosen.text}
        </option>
    {/each}
    </select>
</form>
