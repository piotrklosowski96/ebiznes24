import { Product } from "../models/Product.ts";

interface ProductDetailsProps {
	product: Product;
	addProductToCart: (productId: string) => void;
}

export default function ProductDetails(props: ProductDetailsProps) {

	return (
		<div style={{ border: "1px solid", padding: "10px", display: "block" }}>
			<h2>{props.product.name}</h2>
			<h3>{props.product.description}</h3>
			<button onClick={() => props.addProductToCart(props.product.id!)}>Add to Cart</button>
		</div>
	)
}