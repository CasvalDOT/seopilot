import type { Alert } from './../index';
import type { Stringable } from './main';
import { ApiRequest } from './main';

export class WritableRequest extends ApiRequest implements Stringable {
	entity: Alert;

	constructor(entity: Alert) {
		super();
		this.entity = entity;
	}

	stringify(): string {
		return JSON.stringify({
			name: this.entity.name,
			description: this.entity.description
		});
	}
}
