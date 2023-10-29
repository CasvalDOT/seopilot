<script lang="ts">
import type { User } from '$lib/types';
import type { Filters } from '$lib/types/filters/user';
import List from '../../../components/user/list.svelte';
import Page from '../../../components/page.svelte';
import { page } from '$app/stores';
import api from '$lib/api/user';
import { browser } from '$app/environment';

let entitySelected:number = 0;
let items:User[] = [];
let error:string = "";
let loading = true;


function onCreate(evt:CustomEvent) {
  const entity:User = evt.detail;
  items = [...items, entity];
}

function onUpdate(evt:CustomEvent) {
  const entity:User = evt.detail;
  items = items.map(item => item.id === entity.id ? entity : item);
}

function onDelete(evt:CustomEvent) {
  const entity:User = evt.detail;
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
  entityName="user"
  entityNewIcon="solar:user-plus-broken"
  on:update={onUpdate}
  on:delete={onDelete}
  on:create={onCreate}
  on:close-details={() => entitySelected = 0}
>
    <svelte:fragment slot="table">
      <List 
        columns={['Name','Role','Created at','Updated at']}
        loading={loading}
        error={error}
        items={items}
        on:select={(e) => entitySelected = e.detail}
      />
    </svelte:fragment>
</Page>
