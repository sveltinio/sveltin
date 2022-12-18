/**
 ** Sveltin namespace reflects types exported by some of the @sveltinio/[packages].
 ** This file exists to allow using sveltin's features with no lock-in to the sveltinio packages.
 */
export namespace Sveltin {
	export type ResourceContent = {
		resource: string;
		metadata: YAMLFrontmatter;
		html?: string;
	};

	export type ContentMetadata = {
		name: string;
		items: Array<YAMLFrontmatter>;
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
		author?: string;
		headline?: string;
		created_at?: string;
		updated_at?: string;
		cover?: string;
		[key: string]: string | undefined | Array<TocEntry> | boolean;
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

	export type WebSite = {
		name: string;
		baseURL: string;
		language: string;
		title: string;
		slogan?: string;
		description: string;
		seoDescription: string;
		favicon: string;
		logo: string;
		copyright: string;
		keywords: string;
		contactEmail: string;
		socials: Socials;
		webmaster: WebMaster;
	};

	export type Socials = {
		linkedin: string;
		twitter: string;
		github: string;
		facebook: string;
		instagram: string;
		youtube: string;
	};

	export type WebMaster = {
		name: string;
		address: string;
		contactEmail: string;
	};
}
