export interface Stringable {
	stringify(): string;
}

export class ApiRequest implements Stringable {
	constructor() {}

	stringify(): string {
		return '';
	}
}

export class ApiError {
	key: string = 'Generic.Error';
	message: string = 'Generic Error';
	status: number = 500;

	constructor() {}

	toJSON() {
		return {
			key: this.key,
			message: this.message,
			status: this.status
		};
	}

	setKey(key: string) {
		this.key = key;
	}

	setMessage(message: string) {
		this.message = message;
	}

	setStatus(status: number) {
		this.status = status;
	}
}
