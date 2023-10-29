<script lang="ts">
import {createEventDispatcher} from 'svelte';
import {Tab,TabGroup} from '@skeletonlabs/skeleton';
import Icon from './../icon.svelte';
import api from '$lib/api/site';
import type { Site } from '$lib/types';
import Upsert from './upsert.svelte';
import Attributes from './attributes.svelte';
import SourceBadge from '../source-badge.svelte';
import Header from '../detail/header.svelte';

const dispatch = createEventDispatcher();

export let entityId: number = 0
let tabSet: number = 0;
let entity:Site
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

function handleVerification() {
  // api.verifySite(entityId)
}

function updateData(evt:CustomEvent) {
  entity = evt.detail
  dispatch('update',entity)
}

$: getEntity(entityId)

</script>

{#if !error && entity}
  <section class="p-6">
    <Header title="Site details" hasDeleteButton={entity.source?.code === 'manual'} on:delete={handleDelete}>
      <button class="btn btn-sm variant-filled-success" on:click={handleVerification}>
        Check 
        &nbsp;
        <Icon icon="solar:check-circle-broken" height="auto" />
      </button>
    </Header>

    <div class="mb-8">
      <h3 class="h3 mb-2">{entity.url}</h3>
      <div class="flex items-center gap-x-4">
        <SourceBadge value={entity.source} />
        <div>
          Status: <span class="badge variant-ghost-surface">{entity.status || 'UNKNOW'}</span>
        </div>
        <div>
          {#if entity.alert?.id}
            <span class="badge variant-ghost-error">
              <Icon height="20" icon="solar:bell-bing-bold" />
              &nbsp;
              {entity.alert.name}
            </span>
          {/if}
        </div>
      </div>
    </div>
    <span>Updated at: {entity.updated_at}</span>
    <br />
    <small>Created at: {entity.created_at}</small>
    <br />
    <br />

    <TabGroup>
      <Tab bind:group={tabSet} name="details" value={0}>
        Details
      </Tab>
      <Tab bind:group={tabSet} name="attributes" value={1}>
        Attributes
      </Tab>

      <svelte:fragment slot="panel">
        {#if tabSet == 0}
          <Upsert entity={entity} on:success={updateData} on:cancel={handleCancel} />
        {/if}
        {#if tabSet == 1}
          <Attributes entity={entity} />
        {/if}
      </svelte:fragment>
    </TabGroup>
  </section>
{:else}
  <p>{error}</p>
{/if}
