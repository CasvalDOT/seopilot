import { writable, get } from 'svelte/store';

export const optionsRolesStore = writable([]);

export const optionsRoles = get(optionsRolesStore);
