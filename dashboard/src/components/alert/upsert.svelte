<script lang="ts">
  import type { Alert } from '$lib/types';
  import Icon from './../icon.svelte';
  import Feedback from './../feedback.svelte';
  import api  from '$lib/api/alert';
  import { createEventDispatcher } from 'svelte';
  import { WritableRequest } from '$lib/types/request/alert';

  let dispatch = createEventDispatcher();

  let message: string = '';
  let status: string = 'error';

  export let entity:Alert = {
    id: 0,
    name: '',
    description: '',
    created_at: '',
    updated_at: '',
  }

  async function update() {
    const request = new WritableRequest(entity)

    const [response, err] = await api.update(entity.id, request) 
    if (err) {
      message = err.message
      status = 'error'
      dispatch('error', message)
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
      status = 'error'
      dispatch('error', message)
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

</script>
  
<form class="flex flex-col gap-6" on:submit|preventDefault={upsert}>
  <label class="label">
    <strong>Name</strong>
    <input bind:value={entity.name} class="input" type="text" required placeholder="A name for the alert..." />
  </label>  

  <label class="label">
    <strong>Description</strong>
    <textarea bind:value={entity.description} class="textarea" placeholder="A description for the alert..." />
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

