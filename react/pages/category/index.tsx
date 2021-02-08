import React from 'react';
import CategoryList from '@/components/Category/List';
import Link from 'next/link';

const Home: React.FC = () => {
  return (
    <React.Fragment>
      <main>
        <h1>Categories</h1>
        {/* <Link href="/category/create">
          <a>Add</a>
        </Link> */}

        <section>
          <CategoryList />
        </section>
      </main>
    </React.Fragment>
  );
};

export default Home;
