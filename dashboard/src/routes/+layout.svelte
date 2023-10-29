<script lang="ts">
	import { browser } from '$app/environment';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
  import api from '$lib/api/user';
	import '../app.postcss';

  const unProtectedRoutes: string[] = ['/login', '/forgot-password', '/reset-password'];

	async function fetchData() {
	  if(!browser) {
  	  return
	  }
    if (unProtectedRoutes.includes($page.url.pathname)) {
      return
    }

    const [, error] = await api.me();
    if (error) {
      goto('/login');
    }
  }

	$: $page, fetchData();
</script>

<slot />
