import { BrowserRouter, Routes, Route } from "react-router-dom";
import Products from "./components/Products.tsx";
import Payments from "./components/Payments.tsx";
import LayoutWrapper from "./components/LayoutWrapper.tsx";
import Cart from "./components/Cart.tsx";
import { createContext, useEffect, useState } from "react";
import { Product } from "./services/products/models/Product.ts";


function App() {
  const [products, setProducts] = useState<Product[]>([]);

  useEffect(() => {

  })

  const ApplicationContext = createContext({})

  return (
    <>
      <ApplicationContext.Provider value={ApplicationContext}>
        <BrowserRouter>
          <Routes>
            <Route path="/" element={ <LayoutWrapper /> }>
              <Route path="products" element={ <Products /> } />
              <Route path="payments" element={ <Payments /> } />
              <Route path="cart" element={ <Cart /> } />
            </Route>
          </Routes>
        </BrowserRouter>
      </ApplicationContext.Provider>
    </>
  )
}

export default App
