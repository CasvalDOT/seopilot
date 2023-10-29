import type { EventDispatcher } from 'svelte';

export async function getEntity(id: number) {
	if (!id) return;

	const [data, err] = await api.view(id);
	if (err) {
		error = err.message;
		return;
	}

	entity = data;
}

export function handleCancel() {
	dispatch('cancel');
}

export async function handleDelete<T>(dispatch: EventDispatcher<any>, entityId: T) {
	const [entity, err] = await api.delete(entityId);
	if (err) {
		return;
	}

	dispatch('delete', entity);
}

export function updateData(dispatch: EventDispatcher<any>, evt: CustomEvent) {
	dispatch('update', evt.detail);
	return evt.detail;
}
