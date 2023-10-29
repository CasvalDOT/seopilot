<script lang="ts">
import { createEventDispatcher } from 'svelte';
import type { User } from '$lib/types';
import NoResults from '../no-results.svelte';
import Feedback from '../feedback.svelte';
	import { Avatar } from '@skeletonlabs/skeleton';

const dispatch = createEventDispatcher();

export let items: User[] = [];
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
            <td>
              <div class="flex items-center gap-x-2">
                <Avatar initials={item.name} width="w-12" />
                <div>
                  <strong>{item.name}</strong>
                  <br />
                  {item.email}
                </div>
              </div>
            </td>
            <td>
              {#each item.roles as role}
                <span class="badge variant-ghost-surface">{role.name}</span>
                &nbsp;
              {/each}
            </td>
            <td>{item.created_at}</td>
            <td>{item.updated_at}</td>
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
