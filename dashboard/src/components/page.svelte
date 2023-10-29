<script lang="ts">
import { goto } from "$app/navigation";
import { page } from "$app/stores"; 
import { getDrawerStore,getModalStore, type DrawerSettings} from '@skeletonlabs/skeleton';
import { SvelteComponent, createEventDispatcher } from 'svelte';
import Icon from './icon.svelte';
import SearchField from './search-field.svelte';
import { browser } from "$app/environment";

export let entityName:string = '(entity)';
export let entitySelected:number = 0;
export let entityNewIcon:string = 'solar:user-plus-broken';
export let entityDrawerWidth:string = 'w-full sm:w-1/3';
let searchField:SvelteComponent;
let filters: Record<string,string> = {}

const drawerStore = getDrawerStore();
const dispatch = createEventDispatcher();
const modalStore = getModalStore();

function removeFilter(key:string) {
  filters = {...filters, [key]: ''};
}

function applySearch(event: CustomEvent) {
  filters.search = event.detail;
}

function onCreate<T>(entity:T) {
  dispatch('create', entity);
}

function onUpdate(event:CustomEvent) {
  dispatch('update', event.detail);
}

function onDelete(event:CustomEvent) {
  dispatch('delete', event.detail);
  drawerStore.close();
}

function applyAllFilters(filters: Record<string,string>) {
  if(!browser) {
    return
  }

  let query = new URLSearchParams("");
  const keys = Object.keys(filters);
  for(const key of keys) {
    query.set(key, filters[key]);
  }
  
  goto(`?${query.toString()}`);
}

$: applyAllFilters(filters);

const drawerSettings: DrawerSettings = {
	position: 'right',
	bgBackdrop: 'bg-gradient-to-tr from-indigo-500/50 via-purple-500/50 to-pink-500/50',
	width: entityDrawerWidth,
	padding: 'p-4',
	rounded: 'rounded-xl',
	meta: {
  	onClose: () => {
  	  dispatch('close-details');
  	},
  	onCancel: () => {
  	  dispatch('close-details');
  	},
  	onUpdate,
  	onDelete,
	}
};

function openNewEntityModal() {
  modalStore.trigger({
    type: 'component',
    component: entityNewModalId,
	  backdropClasses: 'bg-gradient-to-tr from-indigo-500/50 via-purple-500/50 to-pink-500/50',
	  meta: {
	    onCreate
  	}
  });
}

function toggleEntityDetails<T>(entity:T) {
  if(!entity) {
    drawerStore.close();
    return;
  }
  drawerStore.open({
    ...drawerSettings,
    id: entityDrawerId,
    meta: {...drawerSettings.meta,entity}
  });
}

function retrieveFilters() {
  const newFilters: Record<string,string> = {}

  for(const entry of $page.url.searchParams.entries()) {
    const [key, value] = entry
    newFilters[key] = value;
  }

  filters = newFilters
}

$: toggleEntityDetails(entitySelected);
$: entityDrawerId = `${entityName}-details`;
$: entityNewModalId = `${entityName}-new`;
$: retrieveFilters()
</script>

<div class="p-10 pb-20">
  <!-- The header of the page -->
  <div class="flex justify-between">
    <div class="w-1/4 mb-10">
      <SearchField on:change={applySearch} bind:this={searchField} />
    </div>

    <div>
      <button class="btn variant-ghost" on:click={openNewEntityModal}>
        <Icon icon={entityNewIcon} height="20" />
        &nbsp;
        New {entityName}
      </button>
    </div>
  </div>

  <!-- Filters -->
  <div class="py-4">
    {#each Object.keys(filters) as key }
      {#if filters[key]}
        <button class="badge py-2 px-4 flex gap-x-4 variant-ghost-surface" on:click={() => removeFilter(key)}>
          {key}: {filters[key]}
          <Icon icon="solar:close-circle-broken" height="24" />
        </button>
      {/if}
    {/each}
  </div>

  <!-- Table -->
  <slot name="table" />
  <slot />
</div>
