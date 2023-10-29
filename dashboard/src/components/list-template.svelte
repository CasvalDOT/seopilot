<script lang="ts">
import { goto } from "$app/navigation";
import { page } from "$app/stores"; 
import Icon from './icon.svelte';
import { SvelteComponent, createEventDispatcher, onMount } from 'svelte';
import Feedback from './feedback.svelte';
import SearchField from './search-field.svelte';
import { getDrawerStore,getModalStore} from '@skeletonlabs/skeleton';
import type { DrawerSettings} from '@skeletonlabs/skeleton';
import NoResults from "./no-results.svelte";

export let entityDrawerId:string = 'entity-details';
export let entityNewModalId:string = 'entity-new';
export let entityName:string = '(entity)';
export let error:string = '';
export let columns:string[] = [];
export let entitySelected:number = 0;
export let count:number = 0;
export let entityNewIcon:string = 'solar:user-plus-broken';
export let entityDrawerWidth:string = 'w-full sm:w-1/3';

let searchField:SvelteComponent;
let dispatch = createEventDispatcher();

const drawerStore = getDrawerStore();
const modalStore = getModalStore();

const drawerSettings: DrawerSettings = {
	id: entityDrawerId,
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
  	onUpdate: onUpdateSuccess,
  	onDelete: onDeleteSuccess
	}
};

function toggleEntityDetails<T>(entity:T) {
  if(!entity) {
    drawerStore.close();
    return;
  }
  drawerStore.open({...drawerSettings,meta: {...drawerSettings.meta,entity}});
}

// TODO: better handler this event
function onCreateSuccess<T>(entity:T) {
  dispatch('create-success', entity);
}

function onUpdateSuccess(event:CustomEvent) {
  dispatch('update-success', event.detail);
}

function onDeleteSuccess(event:CustomEvent) {
  dispatch('delete-success', event.detail);
  drawerStore.close();
}

function openNewEntityModal() {
  modalStore.trigger({
    type: 'component',
    component: entityNewModalId,
	  backdropClasses: 'bg-gradient-to-tr from-indigo-500/50 via-purple-500/50 to-pink-500/50',
	  meta: {
  	  onCreateSuccess: onCreateSuccess,
  	}
  });
}

// Apply search
function applySearch(event: CustomEvent) {
  let query = new URLSearchParams($page.url.searchParams.toString());
  query.set('search', event.detail);
  goto(`?${query.toString()}`);
}

onMount(() => {
  //searchField.focus();
});

$: toggleEntityDetails(entitySelected);

</script>


<section class="p-10">
  <!-- Just a feedback on top when something went wrong -->
  {#if error}
  <Feedback message={error}/>
  {:else} 

  {#if count > 0}
  <!-- The header of the page -->
  <div class="flex justify-between">
    <div class="w-1/4 mb-10">
      <SearchField on:change={applySearch} bind:this={searchField} />
    </div>

    <div>
      <button class="btn variant-ghost" on:click={openNewEntityModal}>
        <Icon icon={entityNewIcon} />
        &nbsp;
        New {entityName}
      </button>
    </div>
  </div>

  <!-- The count of the page -->
  {#if count > 0}
    <p class="font-bold mb-2">Records found: {count}</p>
  {/if}
  <!-- The table of the page -->
  <table class="table table-hover">
		<thead>
			<tr>
        {#each columns as column}
          <th>{column}</th>
        {/each}
			</tr>
		</thead>
		<tbody>
		  <slot />
		</tbody>
	</table>
  {:else}
    <NoResults on:new={openNewEntityModal} />
	{/if}
  {/if}
</section>

