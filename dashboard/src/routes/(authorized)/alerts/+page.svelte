<script lang="ts">
import { browser } from '$app/environment';
import { page } from '$app/stores';
import type { Alert} from '$lib/types';
import type { Filters } from '$lib/types/filters/alert';
import api from '$lib/api/alert';
import Page from '../../../components/page.svelte';
	import List from '../../../components/alert/list.svelte';

let entitySelected:number = 0;
let items:Alert[] = [];
let error:string = "";
let loading = true;

function onCreate(evt:CustomEvent) {
  const entity:Alert = evt.detail;
  items = [...items, entity];
}

function onUpdate(evt:CustomEvent) {
  const entity:Alert = evt.detail;
  items = items.map(item => item.id === entity.id ? entity : item);
}

function onDelete(evt:CustomEvent) {
  const entity:Alert = evt.detail;
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
  entityName="alert"
  entityNewIcon="solar:bell-bing-broken"
  on:update={onUpdate}
  on:delete={onDelete}
  on:create={onCreate}
  on:close-details={() => entitySelected = 0}
>
    <svelte:fragment slot="table">
      <List
        columns={['Name','Created at','Updated at']}
        loading={loading}
        error={error}
        items={items}
        on:select={(e) => entitySelected = e.detail}
      />
    </svelte:fragment>
</Page>

