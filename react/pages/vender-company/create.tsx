import React from 'react';
import VenderCompanyForm from '@/components/VenderCompany/Form';
import useSWR from 'swr';
import axios from 'axios';
import { useRouter } from 'next/router';

interface VenderCompany {
    id: string;
    name: string;
    address: string;
    phone_number: string;
    created_at: string;
    updated_at: string;
    deleted_at?: string;
  }

const VenderCompanyCreate: React.FC = () => {
  const router = useRouter();
  const productApi = `${process.env.GO_MONGO_ENDPOINT_API}/vender-company`;
  const { revalidate } = useSWR(productApi);

  const handleSubmit = (venderCompany: VenderCompany) => {
    axios
      .post(
        `${process.env.GO_MONGO_ENDPOINT_API}/vender-company`,
        venderCompany
      )
      .then((res) => {
        revalidate();
        router.push('/vender-company');
      })
      .catch((err) => console.error(err));
  };

  return (
    <React.Fragment>
      <h1>Vender Company</h1>

      <VenderCompanyForm submitVenderCompany={handleSubmit.bind(this)} />
    </React.Fragment>
  );
};

export default VenderCompanyCreate;
