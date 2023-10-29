<script lang="ts">
import type { Source } from '$lib/types';
import Icon from '@iconify/svelte';
export let value: Source;

function getIcon(code: string) {
  switch (code) {
    case 'forge':
      return 'logos:forge';
    case 'spinupwp':
      return 'logos:wordpress-icon';
    case 'manual':
      return 'solar:server-bold-duotone';
    default:
      return 'solar:server-bold-duotone';
  }
}

$: sourceIcon = getIcon(value.code);
</script>

<span class="badge bg-white text-black px-4">
  <Icon icon={sourceIcon} height="25" />
  &nbsp;
  {value.name}
</span>
