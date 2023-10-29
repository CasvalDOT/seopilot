<script lang="ts">
import type { Site } from '$lib/types'
import NoResults from './../no-results.svelte'
export let entity: Site = {} as Site
</script>

<div class="flex flex-col gap-4">
  <h6 class="h5">Attributes</h6>
  <p>Attributes collected from the site</p>
  {#if entity.attributes.length > 0}
  <div class="flex flex-col gap-6">
    {#each entity.attributes as attribute}
      <div class="flex items-center gap-4 flex-col">
        <strong class="w-full">{attribute.name}</strong>
        <span class="w-full">{attribute.value}</span>
      </div>
    {/each}
  </div>
  {:else}
  <NoResults variant='sm' withCreate={false} />
  {/if}
</div>
