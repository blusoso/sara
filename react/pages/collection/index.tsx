import React from 'react';
import CollectionList from '@/components/Collection/List';
import Link from 'next/link';

const Home: React.FC = () => {
  return (
    <React.Fragment>
      <main>
        <h1>Collections</h1>
        {/* <Link href="/category/create">
          <a>Add</a>
        </Link> */}

        <section>
          <CollectionList />
        </section>
      </main>
    </React.Fragment>
  );
};

export default Home;
