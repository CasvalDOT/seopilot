import { redirect, type Load } from '@sveltejs/kit';

export const load: Load = async () => {
	redirect(307, '/sites');
};
