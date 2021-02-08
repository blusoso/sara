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

const ProductCreate: React.FC<{ id: string }> = ({ id }) => {
  const router = useRouter();
  const productApi = `${process.env.GO_MONGO_ENDPOINT_API}/product/${id}`;
  const { data: defaultProduct, revalidate } = useSWR(productApi);

  const handleSubmit = (newProduct: Product) => {
    newProduct.colors= [newProduct.colors];
    newProduct.sizes= [newProduct.sizes];

    axios
      .put(productApi, newProduct)
      .then((res) => {
        revalidate();
        router.push('/product');
      })
      .catch((err) => console.error(err));
  };

  return (
    <React.Fragment>
      <h1>Product #{defaultProduct?.sku}</h1>
      {defaultProduct && (
        <ProductForm
          submitProduct={handleSubmit.bind(this)}
          defaultProduct={defaultProduct}
        />
      )}
    </React.Fragment>
  );
};

export const getServerSideProps = async ({ query }) => {
  return {
    props: { id: query.id },
  };
};

export default ProductCreate;
