<script lang="ts">
  import Icon from './icon.svelte';
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let value: string = '';
  let timer:number = -1;

  function debounce(v:Event) {
    window.clearTimeout(timer);

    timer = window.setTimeout(() => {
      value = (v.target as HTMLInputElement).value
      dispatch('change', value)
    },750)
  }

</script>

<label class="label">
  <div class="input-group input-group-divider grid-cols-[auto_1fr_auto]">
	  <div class="input-group-shim">
	    <Icon icon="solar:card-search-broken" />
	  </div>
	  <input 
	    type="search" 
	    placeholder="Search..." 
      on:keyup={debounce}
	  />
  </div>
</label>
