<script lang="ts">
import { createEventDispatcher } from 'svelte';
import type { Site } from '$lib/types';
import NoResults from '../no-results.svelte';
import Feedback from '../feedback.svelte';
import Icon from '../icon.svelte';
import SourceBadge from '../source-badge.svelte';

const dispatch = createEventDispatcher();

export let items: Site[] = [];
export let loading = true;
export let error:string = "";
export let columns: string[] = [];

function handleSelection(id: number) {
  dispatch('select', id);  
}
</script>

<div>
  <!-- Print error if any -->
  {#if error}
    <Feedback message={error} />
  {/if}

  {#if loading}
    <p>Loading...</p>
  {:else if items.length > 0}
    <table class="table table-hover">
      <thead>
        <tr>
          {#each columns as column}
            <th>{column}</th>
          {/each}
        </tr>
      </thead>
      <tbody>
        {#each items as item}
          <tr on:click={() => handleSelection(item.id)}>
            <td class="table-cell-fit">
              <strong>{item.url}</strong>
            </td>
            <td class="table-cell-fit">
              {#if item.alert?.id}
                <span class="badge variant-ghost-error">
                  <Icon height="20" icon="solar:bell-bing-bold" />
                  &nbsp;
                  {item.alert?.name || 'UNKNOW'}
                </span>
              {:else}
              -
              {/if}
            </td>
            <td class="table-cell-fit">
              <span class="badge variant-ghost-surface">{item.status || 'UNKNOW'}</span>
            </td>
            <td class="table-cell-fit">
              <SourceBadge value={item.source} />
            </td>
            <td class="table-cell-fit">{item.created_at}</td>
            <td class="table-cell-fit">{item.updated_at}</td>
          </tr>
        {/each}
      </tbody>
      <tfoot>
        <tr>
          <td colspan="100">
            <div class="flex font-bold justify-end">
              Results: {items.length}
            </div>
          </td>
        </tr>
      </tfoot>
    </table>
  {:else}
    <NoResults />
  {/if}
</div>
