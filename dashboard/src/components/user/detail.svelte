<script lang="ts">
import {createEventDispatcher} from 'svelte';
import {Avatar} from '@skeletonlabs/skeleton';
import api from '$lib/api/user';
import type { User } from '$lib/types';
import Upsert from './upsert.svelte';
import Header from '../detail/header.svelte';

const dispatch = createEventDispatcher();

export let entityId: number = 0
let entity:User
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
    <Header title="User details" on:delete={handleDelete} />

    <div class="flex mb-6 items-center space-x-4">
      <Avatar initials={entity.name} />
      <div>
        <h3 class="h3 mb-2">{entity.name}</h3>
        <h4 class="mb-2">Email: {entity.email}</h4>
      </div>
    </div>
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
