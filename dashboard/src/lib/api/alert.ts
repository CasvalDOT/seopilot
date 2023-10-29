import API from './main';
import type { Alert } from './../types';
import type { Filters } from './../types/filters/alert';
import type { ViewAny } from './../types/responses/alert';
import type { WritableRequest } from './../types/request/alert';
import type { ApiError } from './../types/request/main';

const groupPath = '/alerts';

export default new (class APIAlert extends API {
	constructor() {
		super();
	}

	async viewAny(filters: Filters): Promise<[Alert[], ApiError?]> {
		const queryString = await this.toQueryString(filters);

		const [response, error] = await this.request('GET', `${groupPath}?${queryString}`, null);
		if (error) {
			return [[], error];
		}

		const body: ViewAny = await response.json();

		return [body.data, error];
	}

	async view(id: number): Promise<[Alert, ApiError?]> {
		const [response, error] = await this.request('GET', `${groupPath}/${id}`, null);
		if (error) {
			return [{} as Alert, error];
		}

		const body: Alert = await response.json();

		return [body, error];
	}

	async create(request: WritableRequest): Promise<[Alert, ApiError?]> {
		const [response, error] = await this.request('POST', groupPath, request);
		if (error) {
			return [{} as Alert, error];
		}

		const body: Alert = await response.json();

		return [body, error];
	}

	async update(id: number, request: WritableRequest): Promise<[Alert, ApiError?]> {
		const [response, error] = await this.request('PATCH', `${groupPath}/${id}`, request);
		if (error) {
			return [{} as Alert, error];
		}

		const body: Alert = await response.json();

		return [body, error];
	}

	async delete(id: number): Promise<[Alert, ApiError?]> {
		const [response, error] = await this.request('DELETE', `${groupPath}/${id}`, null);
		if (error) {
			return [{} as Alert, error];
		}

		const body: Alert = await response.json();

		return [body, error];
	}
})();
