import type { User } from './../index';
import type { Stringable } from './main';
import { ApiRequest } from './main';

export class WritableRequest extends ApiRequest implements Stringable {
	entity: User;
	roles: number[];

	constructor(entity: User, roles: number[]) {
		super();
		this.entity = entity;
		this.roles = roles;
	}

	stringify(): string {
		return JSON.stringify({
			name: this.entity.name,
			email: this.entity.email,
			roles: (this.roles || []).map((role) => parseInt(role.toString(), 10))
		});
	}
}
