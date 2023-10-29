import type { Site } from './../index';
import type { Stringable } from './main';
import { ApiRequest } from './main';

export class WritableRequest extends ApiRequest implements Stringable {
	entity: Site;

	constructor(entity: Site) {
		super();
		this.entity = entity;
	}

	stringify(): string {
		return JSON.stringify({
			url: this.entity.url,
			alert_id: this.entity.alert ? parseInt(String(this.entity.alert.id), 10) : null
		});
	}
}
