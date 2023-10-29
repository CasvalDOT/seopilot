import { writable, get } from 'svelte/store';

export const userStore = writable();
export const tokenStore = writable('');
export const user = get(userStore);
export const token = get(tokenStore);
