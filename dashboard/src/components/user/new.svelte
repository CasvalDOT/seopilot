<script lang="ts">
import type { User } from '$lib/types';
import {getModalStore} from '@skeletonlabs/skeleton';
import Upsert from './upsert.svelte';

const modalStore = getModalStore();

let entity:User;

function handleSuccess(evt:CustomEvent) {
  $modalStore[0].meta.onCreate({...evt.detail});
  modalStore.close();
}

</script>

<section class="flex flex-col gap-y-6 p-6 w-modal bg-surface-200-700-token">
  <h2 class="h3">New user</h2>
  <p>Invite your team mate to collaborate on this project</p>


  <Upsert on:success={handleSuccess} entity={entity} />
</section>
