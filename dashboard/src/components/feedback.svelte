<script lang="ts">
  import Icon from './icon.svelte';
  export let message: string = '';
  export let status: string = '';
</script>

{#if message}
  <aside 
    class="alert" 
    class:variant-ghost-error={status === 'error'} 
    class:variant-ghost-warning={status === 'warning'} 
    class:variant-ghost-success={status === 'success'}
  >
    <div>
      <Icon icon="solar:shield-warning-broken" height="40" />
    </div>
    <div class="alert-message font-bold">
        {status}
        <p>{message}</p>
    </div>
  </aside>
{/if}
