<script>
    export let color = "rose"
    export let title=""
    export let description =""

    const trimmedDesc = `${description.substring(0, 50)}...`
    console.log(color)
</script>
<div>
    <div class="m-3 shadow-md bg-{color}-300 rounded-2xl p-5 w-72 h-72">
        <div class="flex justify-end">
            <button class="rounded-full bg-white shadow">
                <img class="rounded-xl" src="https://icongr.am/clarity/times.svg?size=24&color=currentColor"/>
            </button>
        </div>
        <p id="notes" class="my-4 w-full text-sm h-2/3">{trimmedDesc}</p>
        <div class="flex justify-between">
            <p class="font-semibold">{title}</p>
            <button class="rounded-full bg-white shadow inline-block">
                <img class="rounded-2xl p-2" src="https://icongr.am/clarity/pencil.svg?size=14&color=currentColor"/>
            </button>
        </div>
    </div>
</div>