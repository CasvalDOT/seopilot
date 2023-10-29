import API from './main';
import type { User } from './../types';
import type { Filters } from './../types/filters/user';
import type { ViewAny } from './../types/responses/user';
import type { WritableRequest } from './../types/request/user';
import type { ApiError } from './../types/request/main';

const groupPath = '/users';

export default new (class APIUser extends API {
	constructor() {
		super();
	}

	async me(): Promise<[User, ApiError?]> {
		const [response, error] = await this.request('GET', `${groupPath}/me`, null);
		if (error) {
			return [{} as User, error];
		}

		const body: User = await response.json();

		return [body, error];
	}

	async viewAny(filters: Filters): Promise<[User[], ApiError?]> {
		const queryString = await this.toQueryString(filters);

		const [response, error] = await this.request('GET', `${groupPath}?${queryString}`, null);
		if (error) {
			return [[], error];
		}

		const body: ViewAny = await response.json();

		return [body.data, error];
	}

	async view(id: number): Promise<[User, ApiError?]> {
		const [response, error] = await this.request('GET', `${groupPath}/${id}`, null);
		if (error) {
			return [{} as User, error];
		}

		const body: User = await response.json();

		return [body, error];
	}

	async create(request: WritableRequest): Promise<[User, ApiError?]> {
		const [response, error] = await this.request('POST', groupPath, request);
		if (error) {
			return [{} as User, error];
		}

		const body: User = await response.json();

		return [body, error];
	}

	async update(id: number, request: WritableRequest): Promise<[User, ApiError?]> {
		const [response, error] = await this.request('PATCH', `${groupPath}/${id}`, request);
		if (error) {
			return [{} as User, error];
		}

		const body: User = await response.json();

		return [body, error];
	}

	async delete(id: number): Promise<[User, ApiError?]> {
		const [response, error] = await this.request('DELETE', `${groupPath}/${id}`, null);
		if (error) {
			return [{} as User, error];
		}

		const body: User = await response.json();

		return [body, error];
	}
})();
