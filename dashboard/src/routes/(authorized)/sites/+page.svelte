<script lang="ts">
import { browser } from '$app/environment';
import { page } from '$app/stores';
import type { Site } from '$lib/types';
import type { Filters } from '$lib/types/filters/site';
import api from '$lib/api/site';
import Page from '../../../components/page.svelte';
import List from '../../../components/site/list.svelte';

let entitySelected:number = 0;
let items:Site[] = [];
let error:string = "";
let loading = true;

function onCreate(evt:CustomEvent) {
  const entity:Site = evt.detail;
  items = [...items, entity];
}

function onUpdate(evt:CustomEvent) {
  const entity:Site = evt.detail;
  items = items.map(item => item.id === entity.id ? entity : item);
}

function onDelete(evt:CustomEvent) {
  const entity:Site = evt.detail;
  items = items.filter(item => item.id !== entity.id);
}

async function fetchData() {
  if(!browser) {
    return
  }
  loading = true
	const search = $page.url.searchParams.get('search');

	const filters: Filters = {
		search: search || ''
	};
  const [data,err] = await api.viewAny(filters);
  items = data || [];
  error = err?.message || "";
  loading = false;
}

$: $page, fetchData();

</script>

<Page
  entitySelected={entitySelected}
  entityName="site"
  entityNewIcon="solar:server-broken"
  on:update={onUpdate}
  on:delete={onDelete}
  on:create={onCreate}
  on:close-details={() => entitySelected = 0}
>
    <svelte:fragment slot="table">
      <List 
        columns={['Url','Alert','Status','Source','Created at','Updated at']}
        loading={loading}
        error={error}
        items={items}
        on:select={(e) => entitySelected = e.detail}
      />
    </svelte:fragment>
</Page>

