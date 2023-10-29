<script lang="ts">
import {createEventDispatcher} from 'svelte';
import api from '$lib/api/alert';
import type { Alert } from '$lib/types';
import Upsert from './upsert.svelte';
import Header from '../detail/header.svelte';

const dispatch = createEventDispatcher();

export let entityId: number = 0
let entity:Alert
let error:string = ""

async function getEntity(id:number) {
  if (!id) return

  const [data, err] = await api.view(id)
  if (err) {
    error = err.message
    return
  }

  entity = data
}

function handleCancel() {
  dispatch('cancel')
}

async function handleDelete() {
  const [entity, err] = await api.delete(entityId)  	
  if (err) {
    return 
  }

  dispatch('delete', entity)
}

function updateData(evt:CustomEvent) {
  entity = evt.detail
  dispatch('update', entity)
}

$: getEntity(entityId)

</script>

{#if !error && entity}
  <section class="p-6">
    <Header title="Alert details" on:delete={handleDelete} />

    <h3 class="h3 mb-2">{entity.name}</h3>
    {#if entity.description}
      <p class="mb-4">{entity.description}</p>
    {/if}
    <span>Updated at: {entity.updated_at}</span>
    <br />
    <small>Created at: {entity.created_at}</small>
    <br />
    <br />

    <Upsert entity={entity} on:success={updateData} on:cancel={handleCancel} />
  </section>
{:else}
  <p>{error}</p>
{/if}
