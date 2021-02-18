import React from 'react';
import Link from 'next/link';
import VenderCompanyList from '@/components/VenderCompany/List';

const VenderCompany: React.FC = () => {
  return (
    <React.Fragment>
      <main>
        <h1>Vender Company</h1>
        <Link href="/vender-company/create">
          <a>Add</a>
        </Link>

        <section>
          <VenderCompanyList />
        </section>
      </main>
    </React.Fragment>
  );
};

export default VenderCompany;
