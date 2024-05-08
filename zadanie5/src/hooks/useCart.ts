import { useState } from "react";
import { CartResponse, CartsService } from "../client";
import { Product } from "../models/Product.ts";

type CartFunction = (productId: string) => void;
type UseCartReturn = [Product[], CartFunction, CartFunction];

export default function useCart(cartId: string): UseCartReturn {
	const [products, setProducts] = useState<Product[]>([])

	const addProductToCart = (productId: string) => {
		CartsService.addProductToCart({cartId: cartId, productId: productId}, ).then((response: CartResponse) => {
			setProducts(response.products as Product[]);
		})
	}

	const removeProductFromCart = (productId: string) => {
		CartsService.deleteProductFromCart({cartId: cartId, productId: productId}).then((response: CartResponse) => {
			setProducts(response.products as Product[]);
		})
	}

	return [ products, addProductToCart, removeProductFromCart ]
}