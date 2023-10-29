<script lang="ts">
  import type { User ,Option} from '$lib/types';
  import api from '$lib/api/user'
  import apiOptions from '$lib/api/option'
  import Icon from './../icon.svelte';
  import Feedback from './../feedback.svelte';
  import { createEventDispatcher } from 'svelte';
	import { optionsRolesStore } from '$lib/stores/global';
	import {get} from 'svelte/store';
	import { WritableRequest } from '$lib/types/request/user';

  let dispatch = createEventDispatcher();

  let message: string = '';
  let status: string = 'success';
  let availableRoles: Option[] = get(optionsRolesStore) as Option[];
  let roles: number[] = [];

  export let entity:User = {
    id: 0,
    name: '',
    email: '',
    roles: [],
    created_at: '',
    updated_at: '',
  }

  async function update() {
    const request = new WritableRequest(entity,roles)

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
    const request = new WritableRequest(entity,roles)
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

  async function fetchRolesOptions() {
    const [items,err] = await apiOptions.getRoles()
    if (err) {
      message = err.message
      return
    }

    availableRoles = items

    for (const role of entity.roles) {
      roles = [...roles, role.id]
    }
  }

  $: fetchRolesOptions()

</script>
  
<form class="flex flex-col gap-6" on:submit|preventDefault={upsert}>
  <label class="label">
    <strong>Name and surname</strong>
    <input bind:value={entity.name} class="input" type="text" required placeholder="Name & Surname" />
  </label>

  <label class="label">
    <strong>Email</strong>
    <input bind:value={entity.email} class="input" type="email" required  placeholder="es. youremail@domain.com" />
  </label>  

  <div class="space-y-2">
    <strong>Roles</strong>

    {#each availableRoles as role}
      <label class="flex items-center space-x-2">
        <input class="checkbox" 
          type="checkbox" 
          value={role.value}
          bind:group={roles}
          checked={roles.includes(parseInt(role.value))}
        />

        <p>{role.label}</p>
      </label>
    {/each}
  </div>

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

