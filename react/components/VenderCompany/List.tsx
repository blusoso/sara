import React, { useState } from 'react';
import useSWR from 'swr';
import Link from 'next/link';
import axios from 'axios';
import ReactPaginate from 'react-paginate';

interface VenderCompany {
  id: string;
  name: string;
  address: string;
  phone_number: string;
  created_at: string;
  updated_at: string;
  deleted_at?: string;
}

const VenderCompanyList: React.FC = () => {
  const [pageIndex, setPageIndex] = useState<number>(1);
  const [pageCount, setPageCount] = useState<number>(1);

  const venderCompanyApi = `${process.env.GO_MONGO_ENDPOINT_API}/vender-company`;
  const { data: venderCompanies, revalidate } = useSWR(venderCompanyApi);

  const handleRemoveVenderCompany = (id: string) => {
    const result = confirm('Are you sure to delete this product?');
    if (result) {
      axios
        .delete(`${process.env.GO_MONGO_ENDPOINT_API}/vender-company/${id}`)
        .then((res) => {
          revalidate();
        })
        .catch((err) => console.error(err));
    }
  };

  const handlePageClick = ({ selected }) => {
    setPageIndex(selected + 1);
    revalidate();
  };

  return (
    <React.Fragment>
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Address</th>
            <th>Phone number</th>
            <th>Created At</th>
            <th>Updated At</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          {venderCompanies?.data?.map((venderCompany: VenderCompany) => (
            <tr key={venderCompany.id}>
              <td>{venderCompany.name}</td>
              <td>{venderCompany.address}</td>
              <td>{venderCompany.phone_number}</td>
              <td>{venderCompany.created_at}</td>
              <td>{venderCompany.updated_at}</td>
              <td>
                <div>
                  <Link href={`/vender-company/${venderCompany.id}`}>
                    <a>Edit</a>
                  </Link>
                </div>
                <div>
                  <a
                    onClick={() => handleRemoveVenderCompany(venderCompany.id)}
                  >
                    Remove
                  </a>
                </div>
              </td>
            </tr>
          ))}
        </tbody>
      </table>

      <div>
        <ReactPaginate
          previousLabel={'previous'}
          nextLabel={'next'}
          breakLabel={'...'}
          breakClassName={'break-me'}
          pageCount={venderCompanies?.total_pages || pageCount}
          pageRangeDisplayed={5}
          onPageChange={handlePageClick}
          containerClassName={'pagination'}
          subContainerClassName={'pages pagination'}
          activeClassName={'active'}
        />
      </div>
    </React.Fragment>
  );
};

export default VenderCompanyList;
