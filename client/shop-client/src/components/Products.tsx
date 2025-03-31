import { useEffect, useState } from "react";
import { api } from "../api/api";
import { useCart } from "../context/CartContext";

interface ApiProduct {
    ID: number;
    name: string;
    price: number;
  }

interface Product {
  id: number;
  name: string;
  price: number;
}

const Products = () => {
    const [products, setProducts] = useState<Product[]>([]);
    const { addToCart } = useCart();
  
    useEffect(() => {
      api.get<ApiProduct[]>("/products")
        .then((res) => {
          const mappedProducts: Product[] = res.data.map((product) => ({
            id: product.ID,
            name: product.name,
            price: product.price
          }));
          setProducts(mappedProducts);
        })
        .catch((err) => console.error(err));
    }, []);

  return (
    <div>
      <h2>Products</h2>
      {products.map((product) => (
        <div key={product.id}>
          <span>{product.name} - ${product.price}</span>
          <button onClick={() => addToCart({ ...product, quantity: 1 })}>Add to Cart</button>
        </div>
      ))}
    </div>
  );
};

export default Products;