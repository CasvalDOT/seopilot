import type { Login } from './../types/responses/auth';
import API from './main';
import type { ApiError } from './../types/request/main';
import type {
	ForgotPasswordRequest,
	LoginRequest,
	ResetPasswordRequest
} from '$lib/types/request/auth';

const groupPath = '/auth';

export default new (class APISite extends API {
	constructor() {
		super();
	}

	async login(request: LoginRequest): Promise<[Login, ApiError?]> {
		const [response, error] = await this.request('POST', groupPath, request);
		if (error) {
			return [{} as Login, error];
		}

		const body: Login = await response.json();

		return [body, undefined];
	}
	async forgotPassword(request: ForgotPasswordRequest): Promise<[null, ApiError?]> {
		const [, error] = await this.request('POST', `${groupPath}/forgot-password`, request);
		if (error) {
			return [null, error];
		}

		return [null, undefined];
	}
	async resetPassword(request: ResetPasswordRequest): Promise<[null, ApiError?]> {
		const [, error] = await this.request('POST', `${groupPath}/reset-password`, request);
		if (error) {
			return [null, error];
		}

		return [null, undefined];
	}
})();
