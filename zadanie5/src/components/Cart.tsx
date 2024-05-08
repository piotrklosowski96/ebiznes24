import { Product } from "../models/Product.ts";
import CartProduct from "./CartProduct.tsx";

interface CartProps {
	products: Product[];
	removeProductFromCart: (productId: string) => void;
}

export default function Cart(props: CartProps) {
	return (
		<div>
			<h1>Cart </h1>
			{props.products?.map((product: Product) => <CartProduct product={product} removeProductFromCart={props.removeProductFromCart}/>)}
		</div>
	)
}