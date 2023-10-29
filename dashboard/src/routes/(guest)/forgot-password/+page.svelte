<script lang="ts">
import { ForgotPasswordRequest } from '$lib/types/request/auth';
import api from '$lib/api/auth';
	import { goto } from '$app/navigation';

let error: string = '';
let email: string = '';

async function handleSubmit() {
  const forgotPasswordRequest = new ForgotPasswordRequest(email);

  const [, err] = await api.forgotPassword(forgotPasswordRequest);
  if (err) {
    error = err.message;
    return;
  } 

  goto('/login');
}
</script>

<div class="container h-full mx-auto flex justify-center items-center">
	<div class="space-y-10 text-center flex flex-col items-center w-full sm:w-1/3">
		<h2 class="h2">Recupero password</h2>

		<form method="POST" class="flex flex-col gap-6 w-full" on:submit|preventDefault={handleSubmit}>
      <label class="label">
        <strong>Email</strong>
        <input name="email" class="input" type="email" required  placeholder="es. youremail@domain.com" />
      </label>  
      <nav class="flex items-center justify-between">
        <button class="btn variant-filled" type="submit">Invia procedura</button>
        <a href="/forgot-password" title="Forgot password?">Torna alla login</a>
      </nav>

      {#if error}
        <p class="text-error mt-10">{error}</p>
      {/if}
		</form>
	</div>
</div>

