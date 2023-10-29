<script lang="ts">
import api from "$lib/api/auth";
import { goto } from "$app/navigation";
import { ResetPasswordRequest } from "$lib/types/request/auth";
	import { page } from "$app/stores";

let error:string = "";
let email:string = "";
let password:string = "";
let passwordConfirmation:string = "";

async function handleSubmit() {
  const token = $page.url.searchParams.get("token") || ''

  const apiRequest = new ResetPasswordRequest(
    email, 
    password, 
    passwordConfirmation, 
    token
  );

  const [, err] = await api.resetPassword(apiRequest);
  if (err) {
    error = err.message
    return
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
        <input bind:value={email} name="email" class="input" type="email" required  placeholder="es. youremail@domain.com" />
      </label>  

      <label class="label">
        <strong>Password</strong>
        <input bind:value={password} name="password" class="input" type="password" required  placeholder="" />
      </label>  

      <label class="label">
        <strong>Password Confirmation</strong>
        <input bind:value={passwordConfirmation} name="password_confirmation" class="input" type="password" required  placeholder="" />
      </label>  

      <nav class="flex items-center justify-between">
        <button class="btn variant-filled" type="submit">Confirm</button>
        <a href="/login" title="Return to login">Return to login</a>
      </nav>

      {#if error}
        <p class="text-error mt-10">{error}</p>
      {/if}
		</form>
	</div>
</div>

