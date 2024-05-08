import { useContext } from "react";
import { ApplicationContext } from "../App.tsx";
import { Product } from "../models/Product.ts";
import ProductDetails from "./ProductDetails.tsx";

interface ProductsProps {
	addProductToCart: (productId: string) => void
}

export default function Products(props: ProductsProps) {
	const { products } = useContext(ApplicationContext)

	return (
		<div className="products">
			{
				products?.map((product: Product) => ProductDetails({
					product,
					addProductToCart: () => props.addProductToCart(product.id!)
				}))
			}
		</div>
	)
}