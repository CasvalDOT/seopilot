<script lang="ts">
  import { page } from '$app/stores';
	import Icon from '../../components/icon.svelte';
	import UserDetail from '../../components/user/detail.svelte';
	import UserNew from '../../components/user/new.svelte';
	import SiteDetail from '../../components/site/detail.svelte';
	import SiteNew from '../../components/site/new.svelte';
	import AlertDetail from '../../components/alert/detail.svelte';
	import AlertNew from '../../components/alert/new.svelte';
	import { getDrawerStore, Modal, Drawer,AppShell, AppRail, AppRailAnchor} from '@skeletonlabs/skeleton';
  import { initializeStores} from '@skeletonlabs/skeleton';
  import { Avatar, type ModalComponent } from '@skeletonlabs/skeleton';

	// Floating UI for Popups
	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';
	import TimeToLive from '../../components/time-to-live.svelte';
	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });

  const modalRegistry: Record<string, ModalComponent> = {
    "user-new": { ref: UserNew },
    "site-new": { ref: SiteNew },
    "alert-new": { ref: AlertNew}
  };

  initializeStores();

	const drawerStore = getDrawerStore();
</script>

<!-- App Shell -->
<Drawer on:backdrop={$drawerStore.meta.onClose}>
  {#if $drawerStore.id === 'site-details'}
    <SiteDetail 
      entityId={$drawerStore.meta.entity} 
      on:cancel={$drawerStore.meta.onCancel} 
      on:update={$drawerStore.meta.onUpdate} 
      on:delete={$drawerStore.meta.onDelete}
    />
  {:else if $drawerStore.id === 'user-details'}
    <UserDetail 
      entityId={$drawerStore.meta.entity} 
      on:cancel={$drawerStore.meta.onCancel} 
      on:update={$drawerStore.meta.onUpdate} 
      on:delete={$drawerStore.meta.onDelete}
    />
  {:else if $drawerStore.id === 'alert-details'}
    <AlertDetail 
      entityId={$drawerStore.meta.entity} 
      on:cancel={$drawerStore.meta.onCancel} 
      on:update={$drawerStore.meta.onUpdate} 
      on:delete={$drawerStore.meta.onDelete}
    />
  {/if}
</Drawer>

<Modal components={modalRegistry} />

<AppShell>
  <div class="fixed z-20 bottom-10 ml-10">
    <TimeToLive />
  </div>
  <svelte:fragment slot="sidebarLeft">
    <AppRail>
      <svelte:fragment slot="lead">
        icon
      </svelte:fragment>

      <AppRailAnchor href="/sites" title="Sites" selected={$page.url.pathname === '/sites'}>
        <svelte:fragment slot="lead">
          <Icon icon="solar:server-broken" />
        </svelte:fragment>
        <span>Sites</span>
      </AppRailAnchor>

      <AppRailAnchor href="/alerts" selected={$page.url.pathname === '/alerts'} >
        <svelte:fragment slot="lead">
          <Icon icon="solar:bell-bing-broken" />
        </svelte:fragment>
        <span>Alerts</span>
      </AppRailAnchor>

      <AppRailAnchor href="/users" selected={$page.url.pathname === '/users'} >
        <svelte:fragment slot="lead">
          <Icon icon="solar:user-circle-broken" />
        </svelte:fragment>
        <span>Users</span>
      </AppRailAnchor>

      <svelte:fragment slot="trail">
        <div class="flex justify-center items-center">
          <Avatar initials="S" width="w-12" />
        </div>
		    <AppRailAnchor selected={$page.url.pathname === '/settings'} href="/settings" title="Settings">
		      <Icon icon="solar:tuning-square-2-broken" />
		    </AppRailAnchor>
	    </svelte:fragment>
    </AppRail>
  </svelte:fragment>

	<!-- Page Route Content -->
	<slot />
</AppShell>
