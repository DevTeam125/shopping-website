import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Home from './feature/home';
import ProductLayout from './feature/product/Layout';
import Products from './feature/product/products';
import Product from './feature/product/product';
import ProductPost from './feature/product/productPost';
import ProductDelete from './feature/product/productDelete';
import useStartup from './feature/home/useStartup';

function App() {
  useStartup();
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/">
          <Route index element={<Home />} />
          <Route path="Product" element={<ProductLayout />}>
            <Route index element={<Products />} />
            <Route path=":productId" element={<Product />} />
            <Route path="create" element={<ProductPost />} />
            <Route path="edit/:productId" element={<ProductPost />} />
            <Route path="delete/:productId" element={<ProductDelete />} />
          </Route>
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
