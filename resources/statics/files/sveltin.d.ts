/**
 ** Sveltin namespace reflects types exported by some of the @sveltinio/[packages].
 ** This file exists to allow using sveltin's features with no lock-in to the sveltinio packages.
 */

import type { ComponentType, SvelteComponentTyped } from 'svelte';

export namespace Sveltin {
	export type MdSveXFile = {
		metadata: Record<PropertyKey, string>;
	};

	export type MdSveXResolver = () => Promise<MdsvexFile>;

	export type MdSvexContentComponent = Promise<
		ComponentType<SvelteComponentTyped>
	>;

	export type ResourceContent = {
		component?: MdSvexContentComponent;
		metadata: YAMLFrontmatter;
	};

	export type ContentMetadata = {
		name: string;
		items?: Array<YAMLFrontmatter> | null;
	};

	export type TocEntry = {
		id: string;
		depth: number;
		value: string;
		children?: Array<TocEntry>;
	};

	export type YAMLFrontmatter = {
		title: string;
		slug: string;
		draft: boolean;
		headings?: Array<TocEntry>;
		keywords?: Array<string>;
		author?: string;
		headline?: string;
		created_at?: string;
		updated_at?: string;
		readingTime?: Record<PropertyKey, string>;
		cover?: string;
		misc?: DynamicObject;
	};

	export type DynamicObject = {
		[key: string]: string | number | object | [];
	};

	export type MenuItem = {
		identifier: string;
		name: string;
		url: string;
		weight: number;
		external?: boolean;
		children?: Array<MenuItem>;
	};

	export type Address = {
		city?: string;
		state?: string;
		postalCode?: string;
		streetAddress?: string;
	};

	export type Contact = {
		name?: string;
		jobTitle?: string;
		email?: string;
		telephone?: string;
		url?: string;
		address?: Address | string;
	};

	export type Person = Contact;

	export type Organization = Contact;

	export type WebSite = {
		name: string;
		baseURL: string;
		language: string;
		title: string;
		description: string;
		favicon?: string;
		logo?: string;
		copyright?: string;
		keywords?: Array<string>;
		contactEmail?: string;
		socials?: Socials;
		creator?: Person | Organization;
	};

	export type Socials = {
		[key: string]: string;
	};
}
