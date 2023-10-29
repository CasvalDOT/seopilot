import type { Site } from './../types';
import API from './main';
import type { ViewAny } from './../types/responses/site';
import type { Filters } from './../types/filters/site';
import type { WritableRequest } from './../types/request/site';
import type { ApiError } from './../types/request/main';

const groupPath = '/sites';

export default new (class APISite extends API {
	constructor() {
		super();
	}

	async viewAny(filters: Filters): Promise<[Site[], ApiError?]> {
		const queryString = await this.toQueryString(filters);

		const [response, error] = await this.request('GET', `${groupPath}?${queryString}`, null);
		if (error) {
			return [[], error];
		}

		const body: ViewAny = await response.json();

		return [body.data, error];
	}

	async view(id: number): Promise<[Site, ApiError?]> {
		const [response, error] = await this.request('GET', `${groupPath}/${id}`, null);
		if (error) {
			return [{} as Site, error];
		}

		const body: Site = await response.json();

		return [body, error];
	}

	async create(request: WritableRequest): Promise<[Site, ApiError?]> {
		const [response, error] = await this.request('POST', groupPath, request);
		if (error) {
			return [{} as Site, error];
		}

		const body: Site = await response.json();

		return [body, error];
	}

	async update(id: number, request: WritableRequest): Promise<[Site, ApiError?]> {
		const [response, error] = await this.request('PATCH', `${groupPath}/${id}`, request);
		if (error) {
			return [{} as Site, error];
		}

		const body: Site = await response.json();

		return [body, error];
	}

	async delete(id: number): Promise<[Site, ApiError?]> {
		const [response, error] = await this.request('DELETE', `${groupPath}/${id}`, null);
		if (error) {
			return [{} as Site, error];
		}

		const body: Site = await response.json();

		return [body, error];
	}

	async verify(id: number): Promise<[boolean, ApiError?]> {
		const [, error] = await this.request('GET', `${groupPath}/${id}/verify`, null);
		if (error) {
			return [false, error];
		}

		return [true, undefined];
	}
})();
