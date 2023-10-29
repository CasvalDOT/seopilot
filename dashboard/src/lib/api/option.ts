import type { Option } from './../types';
import type { ViewAny } from './../types/responses/option';
import API from './main';
import type { ApiError } from './../types/request/main';

const groupPath = '/options';

export default new (class APISite extends API {
	constructor() {
		super();
	}

	async getRoles(): Promise<[Option[], ApiError?]> {
		const [response, error] = await this.request('GET', `${groupPath}/roles`, null);
		if (error) {
			return [[], error];
		}

		const body: ViewAny = await response.json();

		return [body.data, error];
	}

	async getAlerts(): Promise<[Option[], ApiError?]> {
		const [response, error] = await this.request('GET', `${groupPath}/alerts`, null);
		if (error) {
			return [[], error];
		}

		const body: ViewAny = await response.json();

		return [body.data, error];
	}
})();
