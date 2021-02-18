import React from 'react';
import useSWR from 'swr';
import axios from 'axios';
import { useRouter } from 'next/router';
import VenderCompanyForm from '@/components/VenderCompany/Form';

const ProductCreate: React.FC<{ id: string }> = ({ id }) => {
  const router = useRouter();
  const venderCompanyApi = `${process.env.GO_MONGO_ENDPOINT_API}/vender-company/${id}`;
  const { data: defaultVenderCompany, revalidate } = useSWR(venderCompanyApi);

  const handleSubmit = (newVenderCompany) => {
    axios
      .put(venderCompanyApi, newVenderCompany)
      .then((res) => {
        revalidate();
        router.push('/vender-company');
      })
      .catch((err) => console.error(err));
  };

  return (
    <React.Fragment>
      <h1>Vender Company #{defaultVenderCompany?.data?.name}</h1>
      {defaultVenderCompany && (
        <VenderCompanyForm
          submitVenderCompany={handleSubmit.bind(this)}
          defaultVenderCompany={defaultVenderCompany.data}
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
