<script lang="ts">
import { createEventDispatcher } from 'svelte';
import { getModalStore, type ModalSettings} from '@skeletonlabs/skeleton';
import Icon from './../icon.svelte';

const modalStore = getModalStore();

export let hasDeleteButton: boolean = true
export let title: string = ''

const dispatch = createEventDispatcher()

const modal: ModalSettings = {
	type: 'confirm',
	title: 'Please Confirm',
	body: 'Are you sure you wish to proceed?',
	response: async (r: boolean) => {
	  if(!r) return

    dispatch('delete')
	},
};

function handleDelete() {
  modalStore.trigger(modal)
}

</script>

<div class="flex justify-between items-center mb-10">
  <h2 class="h6">{title}</h2>
  <nav>
    {#if hasDeleteButton }
      <button class="btn btn-sm variant-filled-error" on:click={handleDelete}>
        Delete
        &nbsp;
        <Icon icon="solar:trash-bin-minimalistic-broken" height="auto" />
      </button>
    {/if}

    <slot />
  </nav>
</div>
