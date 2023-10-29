import type { Stringable } from './main';
import { ApiRequest } from './main';

export class LoginRequest extends ApiRequest implements Stringable {
	email: string;
	password: string;

	constructor(email: string, password: string) {
		super();
		this.email = email;
		this.password = password;
	}

	stringify(): string {
		return JSON.stringify({
			email: this.email,
			password: this.password
		});
	}
}

export class ForgotPasswordRequest extends ApiRequest implements Stringable {
	email: string;

	constructor(email: string) {
		super();
		this.email = email;
	}

	stringify(): string {
		return JSON.stringify({
			email: this.email
		});
	}
}

export class ResetPasswordRequest extends ApiRequest implements Stringable {
	email: string;
	password: string;
	passwordConfirmation: string;
	token: string;

	constructor(email: string, password: string, passwordConfirmation: string, token: string) {
		super();
		this.email = email;
		this.password = password;
		this.passwordConfirmation = passwordConfirmation;
		this.token = token;
	}

	stringify(): string {
		return JSON.stringify({
			email: this.email,
			password: this.password,
			password_confirmation: this.passwordConfirmation,
			token: this.token
		});
	}
}
