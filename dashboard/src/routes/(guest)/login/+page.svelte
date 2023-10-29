<script lang="ts">
import { LoginRequest } from "$lib/types/request/auth";
import api from "$lib/api/auth";
import { goto } from "$app/navigation";

let form:{email:string, password:string, error:string} = {
  email: '',
  password: '',
  error: ''
};

async function handleSubmit() {
  const loginRequest = new LoginRequest(form.email, form.password);
  const [response,error] = await api.login(loginRequest);
  if(error) {
    form.error = error.message;
    return
  }

  // Save the token
  window.localStorage.setItem('token', response.token);
  window.localStorage.setItem("expiration_date", response.expiration_date);

  goto('/users');

}
</script>

<div class="container h-full mx-auto flex justify-center items-center">
	<div class="space-y-10 text-center flex flex-col items-center w-full sm:w-1/3">
		<h2 class="h2">Welcome to SeoPilot</h2>

		<form method="POST" class="flex flex-col gap-6 w-full" on:submit|preventDefault={handleSubmit}>
      <label class="label">
        <strong>Email</strong>
        <input bind:value={form.email} name="email" class="input" type="email" required  placeholder="es. youremail@domain.com" />
      </label>  
      <label class="label">
        <strong>Password</strong>
        <input bind:value={form.password} name="password" class="input" type="password" required  placeholder="" />
      </label>  
      <nav class="flex items-center justify-between">
        <button class="btn variant-filled" type="submit">Accedi</button>
        <a href="/forgot-password" title="Forgot password?">Forgot password?</a>
      </nav>

      {#if form?.error}
        <p class="text-error mt-10">{form.error}</p>
      {/if}
		</form>
	</div>
</div>

