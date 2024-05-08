import { Product } from "../models/Product.ts";

interface ICartProductProps {
	product: Product;
	removeProductFromCart: (productId: string) => void;
}

export default function CartProduct(props: ICartProductProps) {
	return (
		<div style={{ border: "1px solid", padding: "10px", display: "block" }}>
			<h2>{props.product.name}</h2>
			<h3>{props.product.description}</h3>
			<button onClick={() => props.removeProductFromCart(props.product.id!)}>Remove from cart</button>
		</div>
	)
}