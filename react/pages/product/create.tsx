import React from 'react';
import ProductForm from '@/components/Product/Form';
import useSWR from 'swr';
import axios from 'axios';
import { useRouter } from 'next/router';

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

const ProductCreate: React.FC = () => {
  const router = useRouter();
  const productApi = `${process.env.GO_MONGO_ENDPOINT_API}/product`;
  const { revalidate } = useSWR(productApi);

  const handleSubmit = (product: Product) => {
      product.colors= [product.colors];
      product.sizes= [product.sizes];

      axios
      .post(`${process.env.GO_MONGO_ENDPOINT_API}/product`, product)
      .then((res) => {
        revalidate();
        router.push('/product');
      })
      .catch((err) => console.error(err));
  };

  return (
    <React.Fragment>
      <h1>Product</h1>

      <ProductForm submitProduct={handleSubmit.bind(this)} />
    </React.Fragment>
  );
};

export default ProductCreate;
