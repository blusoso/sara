import React from 'react';
import ProductList from '@/components/Product/List';
import Link from 'next/link';

const Home: React.FC = () => {
  return (
    <React.Fragment>
      <main>
        <h1>Products</h1>
        <Link href="/product/create">
          <a>Add</a>
        </Link>

        <section>
          <ProductList />
        </section>
      </main>
    </React.Fragment>
  );
};

export default Home;
