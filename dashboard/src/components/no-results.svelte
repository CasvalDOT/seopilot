<script lang="ts">
import {createEventDispatcher} from 'svelte';

export let variant: 'sm' | 'lg' = 'lg';
export let withCreate: boolean = true;
export let message: string = 'Sorry, there are no results for this entity';

const dispatch = createEventDispatcher();
function openNewEntityModal() {
  dispatch('new');
}

$: titleClass = variant === 'sm' ? 'h5' : 'h1';
</script>

<div 
  class:p-10={variant !== 'sm'} 
  class:p-4={variant === 'sm'} 
  class:gap-y-10={variant !== 'sm'} 
  class:gap-y-4={variant === 'sm'} 
  class="bg-surface-backdrop-token rounded-xl items-center h-full justify-center flex flex-col"
>
  <h1 class={titleClass}>No results</h1>
  <p>{message}</p>
  {#if withCreate}
    <button class="btn variant-ghost" on:click={openNewEntityModal}>Create new</button>
  {/if}
</div>
