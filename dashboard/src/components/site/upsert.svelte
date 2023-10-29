<script lang="ts">
  import type { Site,Source,Option,Alert } from '$lib/types';
  import Icon from './../icon.svelte';
  import Feedback from './../feedback.svelte';
  import api  from '$lib/api/site';
  import apiOptions from '$lib/api/option';
  import { createEventDispatcher } from 'svelte';
  import { WritableRequest } from '$lib/types/request/site';

  let dispatch = createEventDispatcher();

  let message: string = '';
  let status: string = 'success';
  let availableAlerts: Option[] = [];
  export let entity:Site = {
    id: 0,
    url: '',
    attributes: [],
    source: {} as Source,
    alert: {} as Alert,
    status: '',
    created_at: '',
    updated_at: '',
  }

  async function update() {
    const request = new WritableRequest(entity)

    const [response, err] = await api.update(entity.id, request) 
    if (err) {
      message = err.message
      dispatch('error', message)
      status = 'error'
      return
    }

    message = 'Site updated successfully'
    status = 'success'
    dispatch('success', response)
  }

  function handleCancel() {
    dispatch('cancel')
  }

  async function create() {
    const request = new WritableRequest(entity)
    const [response, err] = await api.create(request);
    if (err) {
      message = err.message
      dispatch('error', message)
      status = 'error'
      return
    }

    message = 'User created successfully'
    status = 'success'
    dispatch('success', response)
  }

  function upsert() {
    if (entity.id) {
      return update()
    } else {
      return create()
    }
  }

  async function fetchAlertOptions() {
    const [alerts,err] = await apiOptions.getAlerts()
    if (err) {
      message = err.message
      return
    }

    availableAlerts = alerts
  }

  $: fetchAlertOptions()

</script>
  
<form class="flex flex-col gap-6" on:submit|preventDefault={upsert}>
  <label class="label">
    <strong>URL</strong>
    <input bind:value={entity.url} class="input" type="text" required placeholder="https://www.google.it" />
  </label>  

  <label class="label">
    <strong>Alert</strong>

    <select bind:value={entity.alert.id} class="select" placeholder="Select alert..." >
      {#each availableAlerts as alert}
        <option 
          value={alert.value}
          selected={String(entity.alert.id) === String(alert.value)}
        >
          {alert.label}
        </option>
      {/each}
    </select>
  </label>  

  <hr class="!border-t-4" />

  <Feedback message={message} status={status} />

  <div class="flex gap-4 justify-between items-center">
    <button type="button" class="w-1/2 btn flex items-center variant-ghost" on:click={handleCancel}>
      <Icon icon="solar:arrow-left-broken" height="auto" />
    </button>

    <button type="submit" class="w-full btn flex items-center variant-filled">
      <span>Confirm</span>
      <Icon icon="solar:arrow-right-broken" height="auto" />
    </button>
  </div>
</form>

