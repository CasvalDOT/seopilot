<script lang="ts">
  import Icon from '@iconify/svelte';
  export let icon: string;
  export let height: string = '30';
</script>

<Icon icon="{icon}" height="{height}" inline={true} style="display: inline"/>
