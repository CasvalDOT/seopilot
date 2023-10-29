export declare type User = {
	name: string;
	email: string;
	id: number;
	roles: Role[];
	created_at: string;
	updated_at: string;
};

export declare type Alert = {
	id: number;
	name: string;
	description: string;
	created_at: string;
	updated_at: string;
};

export declare type Role = {
	id: number;
	name: string;
};

export declare type Option = {
	value: string;
	label: string;
};

export declare type Source = {
	id: string;
	name: string;
	code: string;
};

export declare type SiteAttibute = {
	id: number;
	name: string;
	value: string;
};

export declare type Site = {
	id: number;
	url: string;
	source: Source;
	alert: Alert;
	attributes: SiteAttibute[];
	status: string;
	created_at: string;
	updated_at: string;
};

export declare type Query = {
	[key: string]: string;
};

export type Item = {
	id: number;
	fields: object;
};
