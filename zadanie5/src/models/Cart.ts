import { ProductResponse } from "../client";

export interface Cart {
	id?: string;
	name?: string;
	description?: string;
	products?: Array<ProductResponse>;
}
