import React, { useState } from 'react';
import Link from 'next/link';
import useSWR from 'swr';
import axios from 'axios';

interface Product {
  id: string;
  sku: string;
  title: string;
  brand: string;
  condition: string;
  meta_description: string;
  description: string;
  colors?: string[];
  sizes?: string[];
  price: number;
  qty: number;
  warranty_id: number;
  shop: string;
  is_active: boolean;
  publish_date: string;
  created_at: string;
  updated_at: string;
  deleted_at?: string;
}

const ProductList: React.FC = () => {
  const [pageIndex, setPageIndex] = useState<number>(1);

  const productApi = `${process.env.GO_MONGO_ENDPOINT_API}/product?page=${pageIndex}`;
  const { data: products, revalidate } = useSWR(productApi);

  const handleRemoveProduct = (id: string) => {
    axios
      .delete(`${process.env.GO_MONGO_ENDPOINT_API}/product/${id}`)
      .then((res) => {
        const result = confirm('Are you sure to delete this product?');
        if (result) {
          revalidate();
        }
      })
      .catch((err) => console.error(err));
  };

  return (
    <React.Fragment>
      <table>
        <thead>
          <tr>
            <th>SKU</th>
            <th>Title</th>
            <th>Brand</th>
            <th>Condition</th>
            <th>Meta Description</th>
            <th>Description</th>
            <th>Colors</th>
            <th>Sizes</th>
            <th>Price</th>
            <th>Qty</th>
            <th>WarrantyId</th>
            <th>ShopId</th>
            <th>IsActive</th>
            <th>PublishDate</th>
            <th>CreatedAt</th>
            <th>UpdatedAt</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          {products?.data?.map((product: Product) => (
            <tr key={product.id}>
              <td>{product.sku}</td>
              <td>{product.title}</td>
              <td>{product.brand}</td>
              <td>{product.condition}</td>
              <td>{product.meta_description}</td>
              <td>{product.description}</td>
              <td>{product.colors}</td>
              <td>{product.sizes}</td>
              <td>{product.price}</td>
              <td>{product.qty}</td>
              <td>{product.warranty_id}</td>
              <td>{product.shop}</td>
              <td>{product.is_active}</td>
              <td>{product.publish_date}</td>
              <td>{product.created_at}</td>
              <td>{product.updated_at}</td>
              <td>
                <div>
                  <Link href={`/product/${product.id}`}>
                    <a>Edit</a>
                  </Link>
                </div>
                <div>
                  <a onClick={() => handleRemoveProduct(product.id)}>Remove</a>
                </div>
              </td>
            </tr>
          ))}
        </tbody>
      </table>

      <div>
        <button onClick={() => setPageIndex(pageIndex - 1)}>Previous</button>
        <button onClick={() => setPageIndex(pageIndex + 1)}>Next</button>
      </div>
    </React.Fragment>
  );
};

export default ProductList;
