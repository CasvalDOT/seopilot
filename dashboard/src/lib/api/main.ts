import type { ApiRequest, ApiError } from './../types/request/main';
import { PUBLIC_API_BASE_URL } from '$env/static/public';

export default class API {
	constructor() {}

	async request(
		method: string,
		endpoint: string,
		body: null | ApiRequest
	): Promise<[Response, ApiError?]> {
		const url = PUBLIC_API_BASE_URL + endpoint;

		const token = window.localStorage.getItem('token') || '';
		try {
			const response = await fetch(url, {
				method,
				headers: {
					'Content-Type': 'application/json',
					Authorization: `Bearer ${token}`
				},
				body: body ? body.stringify() : null
			});

			if (!response.ok) {
				const error: ApiError = await response.json();

				return [response, error];
			}

			return [response, undefined];
		} catch (error) {
			console.error(error);
			return [
				new Response(),
				{ message: 'An error occurred, check the logs', status: 500 } as ApiError
			];
		}
	}

	async toQueryString(query: Record<string, string>): Promise<string> {
		const data: URLSearchParams = new URLSearchParams();
		for (const [key, value] of Object.entries(query)) {
			data.append(key, value);
		}
		return data.toString();
	}

	async executeMultiple<T>(calls: Promise<T>[]) {
		const results = await Promise.all(calls);
		return results;
	}
}
