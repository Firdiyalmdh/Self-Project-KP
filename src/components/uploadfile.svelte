<!-- src/components/UploadFile.svelte -->
<script>
    // Import dependencies
    import { createEventDispatcher } from 'svelte';
  
    // Create event dispatcher
    const dispatch = createEventDispatcher();
  
    // Handle file input change
    function handleFileChange(event) {
      const file = event.target.files[0];
      dispatch('fileSelected', file);
    }
  </script>
  
  <div>
    <input type="file" on:change={handleFileChange} />
  </div>
  