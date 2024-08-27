import { BrowserRouter, Routes, Route } from "react-router-dom";
import Products from "./components/Products.tsx";
import Payments from "./components/Payments.tsx";
import LayoutWrapper from "./components/LayoutWrapper.tsx";
import Cart from "./components/Cart.tsx";
import { createContext, useEffect, useRef, useState } from "react";

import { Product as IProduct } from "./models/Product"
import { Cart as ICart } from "./models/Cart";
import {
	CartCreate,
	CartsService,
	CreateCartResponse,
	OpenAPI,
	ProductResponseArray,
	ProductsService
} from "./client";
import useCart from "./hooks/useCart.ts";

OpenAPI.BASE = 'http://localhost:8080/api/v1';

export interface IContext {
	products?: IProduct[];
}

export const ApplicationContext = createContext<IContext>({})

function App() {
	const cartId = useRef('');
	const [products, setProducts] = useState<IProduct[]>([]);
	const [cartProducts, addProductToCart, removeProductFromCart] = useCart(cartId.current);

	useEffect(() => {
		CartsService.createCart({body: {productIds: []} as CartCreate}).then((response: CreateCartResponse) => {
			cartId.current = (response as ICart).id!;
		})

		ProductsService.getProducts().then((response: ProductResponseArray) => {
			setProducts(response.products as IProduct[]);
		})
	}, []);

	return (
		<>
			<ApplicationContext.Provider value={{
				products: products,
			}}>
				<BrowserRouter>
					<Routes>
						<Route path="/" element={<LayoutWrapper/>}>
							<Route path="products" element={<Products addProductToCart={addProductToCart} />}/>
							<Route path="payments" element={<Payments/>}/>
							<Route path="cart" element={<Cart products={cartProducts} removeProductFromCart={removeProductFromCart}/>}/>
						</Route>
					</Routes>
				</BrowserRouter>
			</ApplicationContext.Provider>
		</>
	)
}

export default App
